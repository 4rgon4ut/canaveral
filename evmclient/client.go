package evmclient

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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
	c.Account.Signer.GasLimit = uint64(1e18)      // in units
	c.Account.Signer.GasPrice = gasPrice
	return nil
}

func (c *Client) GetDefaultTxOptions() (*bind.TransactOpts, error) {
	err := c.SetupTxOptions(0, 0)
	if err != nil {
		return nil, err
	}
	return c.Account.Signer, nil
}
