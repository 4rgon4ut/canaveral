package app

import (
	"fmt"
	"os/exec"

	"github.com/canaveral/utils"
)

//nolint:gosec,lll
func (a *App) Compile(fileName string) error {
	name := utils.RemoveExtension(fileName)
	if fileName == "" {
		return fmt.Errorf(".sol file name not specified")
	}
	_, err := exec.Command(
		"solc",
		"--overwrite",
		"--abi",
		fmt.Sprintf("%s/%s.sol", a.contractsDir, name),
		"-o",
		a.abiDir,
	).Output()
	if err != nil {
		return fmt.Errorf("compilation execution error: %w", err)
	}

	_, err = exec.Command(
		"solc",
		"--overwrite",
		"--bin",
		fmt.Sprintf("contracts/%s.sol", name),
		"-o",
		a.binDir,
	).Output()
	if err != nil {
		return fmt.Errorf("compilation execution error: %w", err)
	}

	_, err = exec.Command(
		"abigen",
		"--bin",
		fmt.Sprintf("%s/%s.bin", a.binDir, name),
		"--abi",
		fmt.Sprintf("%s/%s.abi", a.abiDir, name),
		"--pkg",
		name,
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
	if err := utils.PatchBind(fmt.Sprintf("%s/%s.go", a.bindsDir, name)); err != nil {
		return fmt.Errorf("patch binding error: %w", err)
	}
	return nil
}

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
