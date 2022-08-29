package evmclient

import (
	"context"
	"fmt"
	"math/big"
	"reflect"

	"github.com/canaveral/bindings"
	"github.com/canaveral/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	*ethclient.Client

	Account *Account
}

// Creates new Client constructed from go-ethereum ethclient and Account object.
// If no rpcPort specified, uses :8545 by default.
func New(rpcAddr string, rpcPort string, privateKey string) (*Client, error) {
	if rpcPort == "" {
		rpcPort = "8545" // default
	}
	client, err := ethclient.Dial(rpcAddr + ":" + rpcPort)
	if err != nil {
		return nil, err
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	account, err := NewAccount(privateKey, chainID)
	if err != nil {
		return nil, fmt.Errorf("account creation error: %w", err)
	}
	return &Client{
		Client:  client,
		Account: account,
	}, nil
}

// Setup signers transaction options.
// Sets nonce, msg.value in wei, gas limit for call and gas price.
// If gasLimit 0 --> estimate. Gas price is suggested.
func (c *Client) SetupTxOptions(msgValue int64, gasLimit int) error {
	nonce, err := c.PendingNonceAt(context.Background(), c.Account.Address)
	if err != nil {
		return err
	}
	gasPrice, err := c.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}
	c.Account.Signer.Nonce = big.NewInt(int64(nonce))
	c.Account.Signer.Value = big.NewInt(msgValue) // in wei
	c.Account.Signer.GasLimit = uint64(gasLimit)  // in units
	c.Account.Signer.GasPrice = gasPrice
	return nil
}

// Find deploy function in specific contract *.go bind file by provided contracts name.
// Perform reflect.Call of that function with provided params.
func (c *Client) DeployContract(name string, args []string) (addr common.Address, tx *types.Transaction, err error) {
	method, err := getGeployMethodByContractName(name)
	if err != nil {
		return
	}
	// len(args)+2 here because of default first two args for every bind deploy function:
	// Deploy<contract-name>(auth *bind.TransactOpts, backend bind.ContractBackend, ..., ...)
	if len(args)+2 != (method.Type().NumIn()) {
		err = fmt.Errorf("number of params out of index")
		return
	}
	if err = c.SetupTxOptions(0, 0); err != nil {
		return
	}
	input, err := c.prepareCallInput(method.Type(), args)
	if err != nil {
		err = fmt.Errorf("call input preparation error: %w", err)
		return
	}
	res := method.Call(input)
	return castToDeployResults(res)
}

// Find deploy function by contract name.
// Returns specific contract bind Deploy<contract-name>( ) function value.
func getGeployMethodByContractName(name string) (f reflect.Value, err error) {
	// deployer is a special type which used as receiver is every generated deploy function
	// to perform .MethodByName( ) lookups in its scope
	v := reflect.ValueOf(bindings.Deployer{})
	f = v.MethodByName("Deploy" + name)
	if !f.IsValid() || f.Kind() != reflect.Func {
		err = fmt.Errorf("no such contract bind")
		return
	}
	return
}

// Make reflect.Value slice from provided arguments.
// This type([ ]reflect.Value) of input is specific for reflect.Value.Call( ).
//
// Also perform arguments casting into resulting function arguments types:
//
// (porvided args) [ ]string --> [ ]interface{ } (casted to Depoy<contract-name>( ) input types)
// --> [ ]reflect.Value (to make reflect.Call( ) of deploy function)
func (c *Client) prepareCallInput(callType reflect.Type, args []string) (input []reflect.Value, err error) {
	// cast inputs for resulting function call types
	assertedArgs, err := utils.CastInputs(callType, args)
	if err != nil {
		err = fmt.Errorf("input type assertion error: %w", err)
		return
	}
	input = make([]reflect.Value, len(assertedArgs)+2)
	input[0] = reflect.ValueOf(c.Account.Signer)
	input[1] = reflect.ValueOf(c)
	for k, param := range assertedArgs {
		input[k+2] = reflect.ValueOf(param)
	}
	return
}

// Casts called thru reflect.Call( ) contract deploy function output to corresponding types.
func castToDeployResults(res []reflect.Value) (addr common.Address, tx *types.Transaction, err error) {
	var ok bool
	// cast and check error value first
	if e := res[3].Interface(); e != nil {
		err, ok = e.(error)
		if !ok {
			err = fmt.Errorf("error type assertion failed")
			return
		}
		return
	}
	addr, ok = res[0].Interface().(common.Address)
	if !ok {
		err = fmt.Errorf("address type assertion failed")
		return
	}
	tx, ok = res[1].Interface().(*types.Transaction)
	if !ok {
		err = fmt.Errorf("transaction type assertion failed")
		return
	}
	return
}
