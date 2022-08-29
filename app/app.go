package app

import (
	"fmt"

	"github.com/canaveral/bindings"
	"github.com/canaveral/config"
	evmclient "github.com/canaveral/evmclient"
	"github.com/ethereum/go-ethereum/common"
)

type App struct {
	EVMClient *evmclient.Client

	ERC20Instance ERC20Minimal

	binDir       string
	abiDir       string
	bindsDir     string
	contractsDir string
}

const ERC20ExampleAddressVariableName = "ERC20_EXAMPLE_ADDRESS"

func New(cfg *config.Config) (*App, error) {
	client, err := evmclient.New(cfg.RPCAddr, cfg.RPCPort, cfg.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("evm client initialization error: %w", err)
	}

	a := &App{
		EVMClient: client,

		binDir:       cfg.BinDir,
		abiDir:       cfg.ABIDir,
		bindsDir:     cfg.BindsDir,
		contractsDir: cfg.ContractsDir,
	}
	if cfg.ERC20ExampleAddress != "" {
		var instance *bindings.ExampleERC20
		instance, _ = bindings.NewExampleERC20(common.HexToAddress(cfg.ERC20ExampleAddress), client)
		a.UpdateERC20Instance(instance)
	}
	return a, nil
}

func (a *App) UpdateERC20Instance(instance ERC20Minimal) {
	a.ERC20Instance = instance
}
