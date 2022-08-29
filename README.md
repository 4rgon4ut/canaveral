# Welcome to Canaveral!
Canaveral is a cli tool written in Go for compiling, deploying and making calls to solidity smart contracts on Evmos blockchain 
(potentially on every EVM-compatible chain). 

Another words canaveral is a launchsite for smart contracts.

- [Pre-Requisites](#pre-requisites)
- [Dependencies and Makefile](#dependencies-and-makefile)
- [Short Summary](#short-summary)
- [Evmos Node](#evmos-node)
  - [Configuration](#configuration)
  - [Running the Node](#running-the-node)
- [ERC20 Smart Contract](#erc20-smart-contract)
  - [Compilation](#compilation)
- [Deployment Using Go-Ethereum](#deployment-using-go-ethereum)
- [Further Scope](#further-scope)

## Pre-Requisites
- [Go v1.18+](https://go.dev/)
- [Solc](https://docs.soliditylang.org/en/v0.8.15/installing-solidity.html#macos-packages)
- [Abigen](https://goethereumbook.org/smart-contract-compile/)
- [Golangci-lint](https://golangci-lint.run/usage/install/#local-installation)
- [Ginkgo](https://onsi.github.io/ginkgo/#installing-ginkgo)

## Dependencies and Makefile
  `solc`, `abigen`, `golangci-lint` and `ginkgo` can be installed along with other dependencies by:
  ```
  make deps
  ```
  To get full list of available make commands:
  ```
  make help
  ```

## Installation
### 1. Firstly we need to configure Evmos local node.
There are several ways to do that. You can find more detailed information about every way in [docs](https://docs.evmos.org/developers/localnet/single_node.html).

I personally prefer using docker container:
```sh
docker run -it  -p 26657:26657 -p 26656:26656  -p 8545:8545  tharsishq/evmos:v8.0.0 bash  
```
Configuration either be done [manually](https://docs.evmos.org/validators/quickstart/run_node.html#manual-deployment) 
or using the `start-node.sh` script from 
[canaveral](https://github.com/4rgon4ut/canaveral/blob/develop/start-node.sh) repo.

To use `start-node.sh` from evmos docker container:
```
docker cp path/to/start-node.sh <container_id>:/target/dir
```
And than just run:
```
sh ./start-node.sh
```

Or you can run commands from `start-node.sh` one by one manually:
```sh
# env vars
KEY="mykey"
CHAINID="evmos_9000-4"
MONIKER="localtestnet"
KEYRING="test"
LOGLEVEL="info"

# clean evmos directory
rm -rf ~/.evmosd/

# init genesis files
evmosd init test --chain-id=$CHAINID

# generate keypair
evmosd keys add $KEY --keyring-backend=$KEYRING

# predefine validator balance and stake
evmosd add-genesis-account $KEY 1000000000000000000000aphoton,10000000000000000000aevmos,1000000000000000000000stake --chain-id=$CHAINID
evmosd gentx $KEY 1000000000000000000000stake --chain-id=$CHAINID --moniker=$MONIKER --keyring-backend=$KEYRING

evmosd collect-gentxs
evmosd validate-genesis

# start node
evmosd start  --log_level $LOGLEVEL  --json-rpc.api eth,txpool,personal,net,debug,web3
```
### 2. Add PrivateKey to Canaveral Config
To get available key pairs run:
```sh
 $ evmosd keys list
```


Then print private key to terminal stdout:
```sh
evmosd keys unsafe-export-eth-key $KEYNAME --keyring-backend test


693F03A42E6F377D2305CB036EAE9BACCC09B230041CC786252A3BD5C34ED0FA
```
> :warning: **THIS IS UNSAFE AND ONLY CAN BE USED FOR LOCAL ENVIROMENT!!!**


### 3. Install CLI binary to your local go/bin
```
make install
```


## CLI 
In order to deploy the smart contract using go, it first must be compiled using
the Solidity compiler. We create the `.abi` as well as `.bin` files, which

## Deployment using go-ethereum

```shell
 $ evmosd keys list
```

This will print the list of keys with additional information like
key name, address and public key.



```shell
 $ evmosd keys unsafe-export-eth-key $KEYNAME --keyring-backend test
693F03A42E6F377D2305CB036EAE9BACCC09B230041CC786252A3BD5C34ED0FA
```


-

## Testing


## Further scope

Additional things, that may be done for the further development of this basic repository:

- 
