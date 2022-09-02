package app

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/canaveral/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

//nolint:gosec,lll
// Uses solc (solidity compiller) to compile *.sol contract by provided name into *.abi and *.bin files.
// After success compilation uses abigen tool to create *.go bind file from compilled objects.
//
// Writes steps info to stdout and patches generated Go contract binding to be accesible thrue reflect calls.
func (a *App) Compile(fileName string) error {
	name := utils.RemoveExtension(fileName)
	if fileName == "" {
		return fmt.Errorf(".sol file name not specified")
	}
	_, err := exec.Command(
		"solc",
		"--overwrite", // overwrite existing abi files
		"--abi",
		a.getContractPath(name), // path and name of solidity file
		"-o",
		a.abiDir, // output directory
	).Output()
	if err != nil {
		return fmt.Errorf("compilation execution error: %w", err)
	}
	_, err = exec.Command(
		"solc",
		"--overwrite",
		"--bin",
		a.getContractPath(name),
		"-o",
		a.binDir,
	).Output()
	if err != nil {
		return fmt.Errorf("compilation execution error: %w", err)
	}
	// abigen creates Go representation of smart contract(aka bind)
	_, err = exec.Command(
		"abigen",
		"--bin",
		a.getBinPath(name), // path to compiled binaries
		"--abi",
		a.getABIPath(name), // path to compiled abi
		"--pkg",
		name, // contract name for function names generation
		"--out",
		fmt.Sprintf("%s/%s.go", a.bindsDir, name),
	).Output()
	if err != nil {
		return fmt.Errorf("binding execution error: %w", err)
	}
	fmt.Printf(
		"Compilation and binding runs successful. Resulting files can be found in directories: \n\n(*.go files) --> '%s/' \n(*.abi files) --> '%s/' \n(*.bin files) --> '%s/'\n\n",
		a.bindsDir,
		a.abiDir,
		a.binDir,
	)
	// patching new bind necessary for ../bindings package names consistence and reflect calls to deploy functions
	if err := utils.PatchBind(fmt.Sprintf("%s/%s.go", a.bindsDir, name)); err != nil {
		return fmt.Errorf("patch binding error: %w", err)
	}
	return nil
}

func (a *App) Deploy(name string, args []string) (string, string, error) {
	name = utils.RemoveExtension(name)
	contractABI, err := utils.GetABIObject(a.getABIPath(name))
	if err != nil {
		return "", "", err
	}
	bytecode, err := utils.GetBytecode(a.getBinPath(name))
	if err != nil {
		return "", "", err
	}
	input, err := utils.CastInputs(contractABI.Constructor.Inputs, args)
	if err != nil {
		return "", "", err
	}
	err = a.EVMClient.SetupTxOptions(0, 0)
	if err != nil {
		return "", "", err
	}
	addr, tx, _, err := bind.DeployContract(a.EVMClient.Account.Signer, *contractABI, bytecode, a.EVMClient, input...)
	if err != nil {
		return "", "", err
	}
	err = a.Registry.AddContract(name, addr)
	if err != nil {
		return "", "", fmt.Errorf("registry error: %w", err)
	}
	return addr.Hex(), tx.Hash().Hex(), nil
}

func (a *App) Call(name string, method string, args []string) error {
	name = utils.RemoveExtension(name)
	contractABI, err := utils.GetABIObject(a.getABIPath(name))
	if err != nil {
		return err
	}
	abiMethod, exist := utils.GetMethodByName(*contractABI, method)
	if !exist {
		return fmt.Errorf("no such method")
	}
	addr, err := a.Registry.GetAddress(name)
	if err != nil {
		return err
	}
	instance := bind.NewBoundContract(common.HexToAddress(addr), *contractABI, a.EVMClient, a.EVMClient, a.EVMClient)
	err = a.EVMClient.SetupTxOptions(0, 0)
	if err != nil {
		return err
	}
	input, err := utils.CastInputs(abiMethod.Inputs, args)
	if abiMethod.IsConstant() {
		res := &[]interface{}{}
		err := instance.Call(nil, res, method, input...)
		if err != nil {
			return fmt.Errorf("call error: %w", err)
		}
		fmt.Println(res)
	} else {
		if err != nil {
			return fmt.Errorf("inputs casting error: %w", err)
		}
		tx, err := instance.Transact(a.EVMClient.Account.Signer, method, input...)
		if err != nil {
			return fmt.Errorf("transact error: %w", err)
		}
		_, err = bind.WaitMined(context.Background(), a.EVMClient, tx)
		if err != nil {
			return fmt.Errorf("tx not mined: %w", err)
		}
		fmt.Println("tx mined: %s", tx.Hash().Hex())
	}
	return nil
}
