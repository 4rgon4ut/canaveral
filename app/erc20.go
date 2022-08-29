package app

import (
	"fmt"
	"math/big"

	"github.com/canaveral/bindings"
	"github.com/canaveral/config"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Minimal ERC20 functionality interface.
// Should match every bind prodused from IERC20-compatibe contract.
type ERC20Minimal interface {
	BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error)
	Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error)
	Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error)
	Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)
}

// Returns ERC20 bind instance stored in App, which implements ERC20Minimal interface.
// If no instance intialized returns error.
//
// Also, function used as a guard from nil pointer dereference in erc20 specific calls
// which addressing ERC20Example instance.
func (a *App) getERC20Instance() (ERC20Minimal, error) {
	if a.ERC20Instance == nil {
		return nil, fmt.Errorf("no deployed erc20 instance available")
	}
	return a.ERC20Instance, nil
}

// Deploy predifined ERC20Example bind object and rewrite config example contract address
// to keep address of deployed contract persistent between application reboots.
func (a *App) DeployExampleERC20() error {
	addr, tx, instance, err := bindings.Deployer.DeployExampleERC20(
		bindings.Deployer{},
		a.EVMClient.Account.Signer, a.EVMClient,
		"ExampleToken",
		"EXT",
	)
	if err != nil {
		return fmt.Errorf("deployment error: %w", err)
	}
	a.UpdateERC20Instance(instance)
	fmt.Printf(
		"%s seccessfully deployed. \nTX hash: %s \nContract address: %s\n\n",
		"ExampleERC20",
		tx.Hash().Hex(),
		addr.Hex(),
	)
	err = config.AddExampleContractAddr(addr.Hex())
	if err != nil {
		return fmt.Errorf("can't add deployed example address to config: %w", err)
	}
	return nil
}

func (a *App) BalanceOf(address string) error {
	instance, err := a.getERC20Instance()
	if err != nil {
		return err
	}
	addr := common.HexToAddress(address)
	if err := a.EVMClient.SetupTxOptions(0, 0); err != nil {
		return err
	}
	balance, err := instance.BalanceOf(nil, addr)
	if err != nil {
		return err
	}
	fmt.Printf("Address: %s\nBalance: %s\n\n", addr.Hex(), balance)
	return nil
}

func (a *App) Transfer(to string, amount string) error {
	instance, err := a.getERC20Instance()
	if err != nil {
		return err
	}
	addr := common.HexToAddress(to)
	n := new(big.Int)
	n, ok := n.SetString(amount, 10)
	if !ok {
		return fmt.Errorf("can't convert anount to *big.Int")
	}
	if err := a.EVMClient.SetupTxOptions(0, 0); err != nil {
		return err
	}
	tx, err := instance.Transfer(a.EVMClient.Account.Signer, addr, n)
	if err != nil {
		return err
	}
	fmt.Printf("Success! TX hash: %s\n\n", tx.Hash().Hex())
	return nil
}

func (a *App) Mint(to string, amount string) error {
	instance, err := a.getERC20Instance()
	if err != nil {
		return err
	}
	addr := common.HexToAddress(to)
	n := new(big.Int)
	n, ok := n.SetString(amount, 10)
	if !ok {
		return fmt.Errorf("can't convert anount to *big.Int")
	}
	if err := a.EVMClient.SetupTxOptions(0, 0); err != nil {
		return err
	}
	tx, err := instance.Mint(a.EVMClient.Account.Signer, addr, n)
	if err != nil {
		return err
	}
	fmt.Printf("Success! TX hash: %s\n\n", tx.Hash().Hex())
	return nil
}

func (a *App) Burn(amount string) error {
	instance, err := a.getERC20Instance()
	if err != nil {
		return err
	}
	n := new(big.Int)
	n, ok := n.SetString(amount, 10)
	if !ok {
		return fmt.Errorf("can't convert anount to *big.Int")
	}
	if err := a.EVMClient.SetupTxOptions(0, 0); err != nil {
		return err
	}
	tx, err := instance.Burn(a.EVMClient.Account.Signer, n)
	if err != nil {
		return err
	}
	fmt.Printf("Success! TX hash: %s\n\n", tx.Hash().Hex())
	return nil
}
