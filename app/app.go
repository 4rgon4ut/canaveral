package app

import (
	"fmt"

	"github.com/canaveral/bindings"
	"github.com/canaveral/config"
	evmclient "github.com/canaveral/evmclient"
	"github.com/canaveral/registry"
	"github.com/ethereum/go-ethereum/common"
)

// App is the core application object.
// It binds other project components together to perform core functions on hight abstrction level.
type App struct {
	EVMClient *evmclient.Client

	Registry *registry.Registry
	// ERC20 contract bind which implements ERC20Minimal interface
	ERC20Instance ERC20Minimal

	// pack of dirs for proper files generation and management
	binDir       string
	abiDir       string
	bindsDir     string
	contractsDir string
}

// Creates new App instance utilizing config fields.
// Creates EVM RPC client and tries to load ExampleERC20 instance is address provided by config
// otherwise ERC20Instance is <nil>.
func New(cfg *config.Config) (*App, error) {
	client, err := evmclient.New(cfg.RPCAddr, cfg.RPCPort, cfg.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("evm client initialization error: %w", err)
	}
	r, err := registry.New("canaveral.db", 0600)
	if err != nil {
		return nil, fmt.Errorf("db open error: %w", err)
	}
	a := &App{
		EVMClient: client,
		Registry:  r,

		binDir:       cfg.BinDir,
		abiDir:       cfg.ABIDir,
		bindsDir:     cfg.BindsDir,
		contractsDir: cfg.ContractsDir,
	}
	// checks if config were updated with new deployed example contrct address.
	// If so, construct bind instance
	if cfg.ERC20ExampleAddress != "" {
		var instance *bindings.ExampleERC20
		instance, _ = bindings.NewExampleERC20(common.HexToAddress(cfg.ERC20ExampleAddress), client)
		a.UpdateERC20Instance(instance)
	}
	return a, nil
}

// Update ERC20 instance field with new bind
func (a *App) UpdateERC20Instance(instance ERC20Minimal) {
	a.ERC20Instance = instance
}
