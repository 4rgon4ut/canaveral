package app

import (
	"fmt"
)

func (a *App) getABIPath(name string) string {
	return fmt.Sprintf("%s/%s.abi", a.abiDir, name)
}

func (a *App) getBinPath(name string) string {
	return fmt.Sprintf("%s/%s.bin", a.binDir, name)
}

func (a *App) getContractPath(name string) string {
	return fmt.Sprintf("%s/%s.sol", a.contractsDir, name)
}
