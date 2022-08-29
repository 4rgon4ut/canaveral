# Welcome to Canaveral!
Canaveral is a cli tool written in Go for compiling, deploying and making calls to solidity smart contracts on Evmos blockchain 
(potentially on every EVM-compatible chain).

- [Pre-Requisites](#pre-requisites)
- [Installation and Setup](#installation-and-setup)
- [Short Summary](#short-summary)
- [Evmos Node](#evmos-node)
  - [Configuration](#configuration)
  - [Running the Node](#running-the-node)
- [ERC20 Smart Contract](#erc20-smart-contract)
  - [Compilation](#compilation)
- [Deployment Using Go-Ethereum](#deployment-using-go-ethereum)
- [Further Scope](#further-scope)

## Pre-Requisites
The following software has to be installed on your machine in order to use the 
latest version of Evmos (currently v5.0.0):
- [Go v1.18+](https://go.dev/)
- [Solc](https://docs.soliditylang.org/en/v0.8.15/installing-solidity.html#macos-packages)
- [Abigen](https://goethereumbook.org/smart-contract-compile/)
- [Golangci-lint](https://golangci-lint.run/usage/install/#local-installation)
- [Ginkgo](https://onsi.github.io/ginkgo/#installing-ginkgo)

## Dependencies and Makefile
  `solc`, `abigen`, and `golangci-lint` can be installed along with other dependencies by:
  ```
  make deps
  ```
  Full list of available make commands:
  ```
  Makefile available targets:
  * help            Display this help screen.
  * contract-tools  Install ethereum specific tools (solc, solhing, abigen, ...)
  * deps            Install all necessary dependencies
  * lint            Runs linting tool
  * test            Runs all tests
  * install         Install canaveral cli to go binaries dir
  * clean           Clean generated artifacts
  ```
## Short Summary
When you have installed the required software, configured and ran a local Evmos
node, you can clone this repository, and install the OpenZeppelin contracts using 
`npm install`. 

When executing 
```shell
 $ ./init.sh
```
an instance of the Maltcoin ERC20 token contract is deployed
to the running localnet, the contents of transaction receipt printed, as well as 
a simple transfer of Maltcoin tokens between two accounts executed.

## Installation
### 1. Firstly we need to configure Evmos local node.
There are several ways to do that. You can find more detailed information about every way in [docs](https://docs.evmos.org/developers/localnet/single_node.html)
I personally prefer using docker container:
```sh
docker run -it  -p 26657:26657 -p 26656:26656  -p 8545:8545  tharsishq/evmos:v8.0.0 bash  
```
Configuration either be done [manually](https://docs.evmos.org/validators/quickstart/run_node.html#manual-deployment) 
or using the `init.sh` shell script, that is contained in the 
[Evmos GitHub repository](https://github.com/evmos/evmos).


### 2. Running the node
You can start your configured Evmos node using `evmosd start` and should see blocks 
being produced. <br>
Now it's possible to interact with the node through the `evmosd` CLI.


### Compilation
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
