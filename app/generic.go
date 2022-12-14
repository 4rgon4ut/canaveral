package app

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/canaveral/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// Uses solc (solidity compiller) to compile *.sol contract by provided name into *.abi and *.bin files.
// After success compilation uses abigen tool to create *.go bind file from compilled objects.
//
// Writes steps info to stdout and patches generated Go contract binding to be accesible thrue reflect calls.
//
//nolint:gosec,lll
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
	return nil
}

// Generic contract deploy function.
// Loads contract ABI & bin representations by provided name and configured paths.
// Deploys contracts on pre-configured network.
//
// Returns contract address and tx hash.
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
	txOpts, err := a.EVMClient.GetDefaultTxOptions()
	if err != nil {
		return "", "", err
	}
	_, tx, _, err := bind.DeployContract(
		txOpts,
		*contractABI,
		common.FromHex(string(bytecode)),
		a.EVMClient,
		input...,
	)
	if err != nil {
		return "", "", err
	}
	addr, err := bind.WaitDeployed(context.Background(), a.EVMClient, tx)
	if err != nil {
		return "", "", err
	}
	err = a.Registry.AddContract(name, addr)
	if err != nil {
		return "", "", fmt.Errorf("registry error: %w", err)
	}
	return addr.Hex(), tx.Hash().Hex(), nil
}

// Call make calls to contract functions.
// Searches for method in contract ABI by provided method name.
//
// If method is view/pure perform contract.Call and prints results.
//
// If method is changing the state perform transaction.
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
	txOpts, err := a.EVMClient.GetDefaultTxOptions()
	if err != nil {
		return err
	}
	input, err := utils.CastInputs(abiMethod.Inputs, args)
	if err != nil {
		return fmt.Errorf("inputs casting error: %w", err)
	}
	if abiMethod.IsConstant() {
		res := &[]interface{}{}
		err := instance.Call(nil, res, method, input...)
		if err != nil {
			return fmt.Errorf("call error: %w", err)
		}
		fmt.Println(*res...)
	} else {
		tx, err := instance.Transact(txOpts, method, input...)
		if err != nil {
			return fmt.Errorf("transact error: %w", err)
		}
		_, err = bind.WaitMined(context.Background(), a.EVMClient, tx)
		if err != nil {
			return fmt.Errorf("tx not mined: %w", err)
		}
		fmt.Printf("tx mined: %x", tx.Hash())
	}
	return nil
}
