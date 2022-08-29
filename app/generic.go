package app

import (
	"fmt"
	"os/exec"

	"github.com/canaveral/utils"
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
		fmt.Sprintf("%s/%s.sol", a.contractsDir, name), // path and name of solidity file
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
		fmt.Sprintf("%s/%s.sol", a.contractsDir, name),
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
		fmt.Sprintf("%s/%s.bin", a.binDir, name), // path to compiled binaries
		"--abi",
		fmt.Sprintf("%s/%s.abi", a.abiDir, name), // path to compiled abi
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

// Generic deploy smart contracts on specified EVM-compatible chain.
func (a *App) Deploy(fileName string, args []string) error {
	name := utils.RemoveExtension(fileName)
	addr, tx, err := a.EVMClient.DeployContract(
		name,
		args,
	)
	if err != nil {
		return fmt.Errorf("deployment error: %w", err)
	}
	fmt.Printf(
		"%s seccessfully deployed. \nTX hash: %s \nContract address: %s\n\n",
		name,
		tx.Hash().Hex(),
		addr.Hex(),
	)
	return nil
}
