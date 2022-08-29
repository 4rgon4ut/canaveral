# Welcome to Canaveral!
Canaveral is a cli tool written in Go for compiling, deploying and making calls to solidity smart contracts on Evmos blockchain 
(potentially on every EVM-compatible chain). 

Another words canaveral is a launchsite for smart contracts.

- [Pre-Requisites](#pre-requisites)
- [Dependencies and Makefile](#dependencies-and-makefile)
- [Installation](#installation)
- [CLI](#cli)
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
As I mentioned earlier Canaveral is a cli tool, so all interaction goes by specific commands.

### Base
To see full list of commands run:
```sh
canaveral help
```

Also, you can run:
```sh
canaveral example
```
to see special demonstration script which compile, deploy and interact with `ExampleERC20` contract.

## Methods
Canaveral have couple generic methods and some predifined for interaction with contracts which implementing specific interface(ERC20Minimal).

### Generic
Canaveral allows user to compile and deploy every smart contract which is valid and can be compiled by `solc`. 

So, we can declare that
```sh
canaveral compile [contract-name]

canaveral deploy [contract-name] arg1, arg2, ...
```
are generic methods. 

But, for now, there is a restriction:
deploying contract constructor arguments must be simple. This is beacuse casting CMD inputs(of types string) to specific for solidity types in elegant and  eficient way is complicated design problem. 
I'm planning to spread supported types during the time. 

List of supported constructor types:
```
|    Go          |    Solidity    |
|                |                |
|----------------|----------------|
|    uint8       |    uint8       | 
|----------------|----------------|
|    uint16      |    uint16      | 
|----------------|----------------|
|    uint32      |    uint32      | 
|----------------|----------------|
|    uint64      |    uint64      | 
|----------------|----------------|
|    big.Int     |    uint128     | 
|----------------|----------------|
|    big.Int     |    uint256     | 
|----------------|----------------|
|     string     |    string      | 
|----------------|----------------|
|common.Address* |    address     | 
|----------------|----------------|
|    [32]byte    |    bytes32     | 
|----------------|----------------|


*
common.Address is a specific go-ethereum type

```

### Predifined ERC20 methods
Application also have secific for ERC20 implementation methods to query balances, mint, burn, and transfer tokens.
Every deployed ERC20-compatible contract can be called by this methods, because application uses special ERC20Minimal interface for ERC20 contract bind(generated by abigen *.go representation of a contract) file:

```go
type ERC20Minimal interface {
	
  BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error)
	
  Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error)
	
  Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error)
	
  Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)
}
```

## Workflow

### Testing
Poject contains BDD tests based on [ginkgo](https://onsi.github.io/ginkgo/) tool. 
Run tests:
```
make test
```
### Linting
Gollangci-lint used with `.golangci.yml` config obtained from [evmos](https://github.com/evmos/evmos) repository with little changes.

Run linter:
```
make lint
```

CI pipepline contain `build`, `lint` and `test` tasks.

[Conventional commits semantic](https://www.conventionalcommits.org/en/v1.0.0/) used in most of commits. 

## Further scope

Additional things, that may be done for the further development of this basic repository:

1. As every poduct created with time pressure this project have two common problems:
    - Tech debt consisting of non-reusable and not eficient code.
    - Lack of tests.
   So, this both points is a fields for improvement.
   
2. Would be nice to implement fully generic way of interaction with every function of smart contract.

3. Workflow and CI\CD improvements. For example, one of first improvements may be packaging application to docker container.
