package app

import (
	"fmt"

	"github.com/canaveral/config"
	evmclient "github.com/canaveral/evmclient"
	"github.com/canaveral/registry"
)

//nolint:gofumpt
const FileMode = 0600

// App is the core application object.
// It binds other project components together to perform core functions on hight abstrction level.
type App struct {
	EVMClient *evmclient.Client
	Registry  *registry.Registry

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
	r, err := registry.New("canaveral.db", FileMode)
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
	return a, nil
}
