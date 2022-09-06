// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ExampleERC20

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ExampleERC20MetaData contains all meta data concerning the ExampleERC20 contract.
var ExampleERC20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001e9238038062001e92833981810160405281019062000037919062000252565b818181600390816200004a919062000522565b5080600490816200005c919062000522565b5050506200006f620000b760201b60201c565b600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505062000609565b600033905090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6200012882620000dd565b810181811067ffffffffffffffff821117156200014a5762000149620000ee565b5b80604052505050565b60006200015f620000bf565b90506200016d82826200011d565b919050565b600067ffffffffffffffff82111562000190576200018f620000ee565b5b6200019b82620000dd565b9050602081019050919050565b60005b83811015620001c8578082015181840152602081019050620001ab565b60008484015250505050565b6000620001eb620001e58462000172565b62000153565b9050828152602081018484840111156200020a5762000209620000d8565b5b62000217848285620001a8565b509392505050565b600082601f830112620002375762000236620000d3565b5b815162000249848260208601620001d4565b91505092915050565b600080604083850312156200026c576200026b620000c9565b5b600083015167ffffffffffffffff8111156200028d576200028c620000ce565b5b6200029b858286016200021f565b925050602083015167ffffffffffffffff811115620002bf57620002be620000ce565b5b620002cd858286016200021f565b9150509250929050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200032a57607f821691505b60208210810362000340576200033f620002e2565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620003aa7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826200036b565b620003b686836200036b565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b600062000403620003fd620003f784620003ce565b620003d8565b620003ce565b9050919050565b6000819050919050565b6200041f83620003e2565b620004376200042e826200040a565b84845462000378565b825550505050565b600090565b6200044e6200043f565b6200045b81848462000414565b505050565b5b8181101562000483576200047760008262000444565b60018101905062000461565b5050565b601f821115620004d2576200049c8162000346565b620004a7846200035b565b81016020851015620004b7578190505b620004cf620004c6856200035b565b83018262000460565b50505b505050565b600082821c905092915050565b6000620004f760001984600802620004d7565b1980831691505092915050565b6000620005128383620004e4565b9150826002028217905092915050565b6200052d82620002d7565b67ffffffffffffffff811115620005495762000548620000ee565b5b62000555825462000311565b6200056282828562000487565b600060209050601f8311600181146200059a576000841562000585578287015190505b62000591858262000504565b86555062000601565b601f198416620005aa8662000346565b60005b82811015620005d457848901518255600182019150602085019450602081019050620005ad565b86831015620005f45784890151620005f0601f891682620004e4565b8355505b6001600288020188555050505b505050505050565b61187980620006196000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c806342966c681161008c57806395d89b411161006657806395d89b411461025f578063a457c2d71461027d578063a9059cbb146102ad578063dd62ed3e146102dd576100ea565b806342966c68146101f557806370a08231146102115780638da5cb5b14610241576100ea565b806323b872dd116100c857806323b872dd1461015b578063313ce5671461018b57806339509351146101a957806340c10f19146101d9576100ea565b806306fdde03146100ef578063095ea7b31461010d57806318160ddd1461013d575b600080fd5b6100f761030d565b6040516101049190610f6f565b60405180910390f35b6101276004803603810190610122919061102a565b61039f565b6040516101349190611085565b60405180910390f35b6101456103c2565b60405161015291906110af565b60405180910390f35b610175600480360381019061017091906110ca565b6103cc565b6040516101829190611085565b60405180910390f35b6101936103fb565b6040516101a09190611139565b60405180910390f35b6101c360048036038101906101be919061102a565b610404565b6040516101d09190611085565b60405180910390f35b6101f360048036038101906101ee919061102a565b61043b565b005b61020f600480360381019061020a9190611154565b6104aa565b005b61022b60048036038101906102269190611181565b6104be565b60405161023891906110af565b60405180910390f35b610249610506565b60405161025691906111bd565b60405180910390f35b61026761052c565b6040516102749190610f6f565b60405180910390f35b6102976004803603810190610292919061102a565b6105be565b6040516102a49190611085565b60405180910390f35b6102c760048036038101906102c2919061102a565b610635565b6040516102d49190611085565b60405180910390f35b6102f760048036038101906102f291906111d8565b610658565b60405161030491906110af565b60405180910390f35b60606003805461031c90611247565b80601f016020809104026020016040519081016040528092919081815260200182805461034890611247565b80156103955780601f1061036a57610100808354040283529160200191610395565b820191906000526020600020905b81548152906001019060200180831161037857829003601f168201915b5050505050905090565b6000806103aa6106df565b90506103b78185856106e7565b600191505092915050565b6000600254905090565b6000806103d76106df565b90506103e48582856108b0565b6103ef85858561093c565b60019150509392505050565b60006012905090565b60008061040f6106df565b90506104308185856104218589610658565b61042b91906112a7565b6106e7565b600191505092915050565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1661047c6106df565b73ffffffffffffffffffffffffffffffffffffffff161461049c57600080fd5b6104a68282610bb2565b5050565b6104bb6104b56106df565b82610d08565b50565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60606004805461053b90611247565b80601f016020809104026020016040519081016040528092919081815260200182805461056790611247565b80156105b45780601f10610589576101008083540402835291602001916105b4565b820191906000526020600020905b81548152906001019060200180831161059757829003601f168201915b5050505050905090565b6000806105c96106df565b905060006105d78286610658565b90508381101561061c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106139061134d565b60405180910390fd5b61062982868684036106e7565b60019250505092915050565b6000806106406106df565b905061064d81858561093c565b600191505092915050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610756576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161074d906113df565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036107c5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107bc90611471565b60405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040516108a391906110af565b60405180910390a3505050565b60006108bc8484610658565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146109365781811015610928576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161091f906114dd565b60405180910390fd5b61093584848484036106e7565b5b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036109ab576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109a29061156f565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610a1a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a1190611601565b60405180910390fd5b610a25838383610ed5565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015610aab576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610aa290611693565b60405180910390fd5b8181036000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610b9991906110af565b60405180910390a3610bac848484610eda565b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610c21576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c18906116ff565b60405180910390fd5b610c2d60008383610ed5565b8060026000828254610c3f91906112a7565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610cf091906110af565b60405180910390a3610d0460008383610eda565b5050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610d77576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d6e90611791565b60405180910390fd5b610d8382600083610ed5565b60008060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015610e09576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e0090611823565b60405180910390fd5b8181036000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600260008282540392505081905550600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610ebc91906110af565b60405180910390a3610ed083600084610eda565b505050565b505050565b505050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610f19578082015181840152602081019050610efe565b60008484015250505050565b6000601f19601f8301169050919050565b6000610f4182610edf565b610f4b8185610eea565b9350610f5b818560208601610efb565b610f6481610f25565b840191505092915050565b60006020820190508181036000830152610f898184610f36565b905092915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610fc182610f96565b9050919050565b610fd181610fb6565b8114610fdc57600080fd5b50565b600081359050610fee81610fc8565b92915050565b6000819050919050565b61100781610ff4565b811461101257600080fd5b50565b60008135905061102481610ffe565b92915050565b6000806040838503121561104157611040610f91565b5b600061104f85828601610fdf565b925050602061106085828601611015565b9150509250929050565b60008115159050919050565b61107f8161106a565b82525050565b600060208201905061109a6000830184611076565b92915050565b6110a981610ff4565b82525050565b60006020820190506110c460008301846110a0565b92915050565b6000806000606084860312156110e3576110e2610f91565b5b60006110f186828701610fdf565b935050602061110286828701610fdf565b925050604061111386828701611015565b9150509250925092565b600060ff82169050919050565b6111338161111d565b82525050565b600060208201905061114e600083018461112a565b92915050565b60006020828403121561116a57611169610f91565b5b600061117884828501611015565b91505092915050565b60006020828403121561119757611196610f91565b5b60006111a584828501610fdf565b91505092915050565b6111b781610fb6565b82525050565b60006020820190506111d260008301846111ae565b92915050565b600080604083850312156111ef576111ee610f91565b5b60006111fd85828601610fdf565b925050602061120e85828601610fdf565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061125f57607f821691505b60208210810361127257611271611218565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006112b282610ff4565b91506112bd83610ff4565b92508282019050808211156112d5576112d4611278565b5b92915050565b7f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760008201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b6000611337602583610eea565b9150611342826112db565b604082019050919050565b600060208201905081810360008301526113668161132a565b9050919050565b7f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b60006113c9602483610eea565b91506113d48261136d565b604082019050919050565b600060208201905081810360008301526113f8816113bc565b9050919050565b7f45524332303a20617070726f766520746f20746865207a65726f20616464726560008201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b600061145b602283610eea565b9150611466826113ff565b604082019050919050565b6000602082019050818103600083015261148a8161144e565b9050919050565b7f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000600082015250565b60006114c7601d83610eea565b91506114d282611491565b602082019050919050565b600060208201905081810360008301526114f6816114ba565b9050919050565b7f45524332303a207472616e736665722066726f6d20746865207a65726f20616460008201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b6000611559602583610eea565b9150611564826114fd565b604082019050919050565b600060208201905081810360008301526115888161154c565b9050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f206164647260008201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b60006115eb602383610eea565b91506115f68261158f565b604082019050919050565b6000602082019050818103600083015261161a816115de565b9050919050565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206260008201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b600061167d602683610eea565b915061168882611621565b604082019050919050565b600060208201905081810360008301526116ac81611670565b9050919050565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b60006116e9601f83610eea565b91506116f4826116b3565b602082019050919050565b60006020820190508181036000830152611718816116dc565b9050919050565b7f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360008201527f7300000000000000000000000000000000000000000000000000000000000000602082015250565b600061177b602183610eea565b91506117868261171f565b604082019050919050565b600060208201905081810360008301526117aa8161176e565b9050919050565b7f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60008201527f6365000000000000000000000000000000000000000000000000000000000000602082015250565b600061180d602283610eea565b9150611818826117b1565b604082019050919050565b6000602082019050818103600083015261183c81611800565b905091905056fea264697066735822122009de4305a459b515ad371864c86a416701e7346e2d8199808af06e94cea6c43d64736f6c63430008100033",
}

// ExampleERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use ExampleERC20MetaData.ABI instead.
var ExampleERC20ABI = ExampleERC20MetaData.ABI

// ExampleERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ExampleERC20MetaData.Bin instead.
var ExampleERC20Bin = ExampleERC20MetaData.Bin

// DeployExampleERC20 deploys a new Ethereum contract, binding an instance of ExampleERC20 to it.
func DeployExampleERC20(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string) (common.Address, *types.Transaction, *ExampleERC20, error) {
	parsed, err := ExampleERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExampleERC20Bin), backend, name_, symbol_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExampleERC20{ExampleERC20Caller: ExampleERC20Caller{contract: contract}, ExampleERC20Transactor: ExampleERC20Transactor{contract: contract}, ExampleERC20Filterer: ExampleERC20Filterer{contract: contract}}, nil
}

// ExampleERC20 is an auto generated Go binding around an Ethereum contract.
type ExampleERC20 struct {
	ExampleERC20Caller     // Read-only binding to the contract
	ExampleERC20Transactor // Write-only binding to the contract
	ExampleERC20Filterer   // Log filterer for contract events
}

// ExampleERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ExampleERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ExampleERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExampleERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExampleERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExampleERC20Session struct {
	Contract     *ExampleERC20     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExampleERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExampleERC20CallerSession struct {
	Contract *ExampleERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ExampleERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExampleERC20TransactorSession struct {
	Contract     *ExampleERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ExampleERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ExampleERC20Raw struct {
	Contract *ExampleERC20 // Generic contract binding to access the raw methods on
}

// ExampleERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExampleERC20CallerRaw struct {
	Contract *ExampleERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ExampleERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExampleERC20TransactorRaw struct {
	Contract *ExampleERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewExampleERC20 creates a new instance of ExampleERC20, bound to a specific deployed contract.
func NewExampleERC20(address common.Address, backend bind.ContractBackend) (*ExampleERC20, error) {
	contract, err := bindExampleERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExampleERC20{ExampleERC20Caller: ExampleERC20Caller{contract: contract}, ExampleERC20Transactor: ExampleERC20Transactor{contract: contract}, ExampleERC20Filterer: ExampleERC20Filterer{contract: contract}}, nil
}

// NewExampleERC20Caller creates a new read-only instance of ExampleERC20, bound to a specific deployed contract.
func NewExampleERC20Caller(address common.Address, caller bind.ContractCaller) (*ExampleERC20Caller, error) {
	contract, err := bindExampleERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExampleERC20Caller{contract: contract}, nil
}

// NewExampleERC20Transactor creates a new write-only instance of ExampleERC20, bound to a specific deployed contract.
func NewExampleERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ExampleERC20Transactor, error) {
	contract, err := bindExampleERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExampleERC20Transactor{contract: contract}, nil
}

// NewExampleERC20Filterer creates a new log filterer instance of ExampleERC20, bound to a specific deployed contract.
func NewExampleERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ExampleERC20Filterer, error) {
	contract, err := bindExampleERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExampleERC20Filterer{contract: contract}, nil
}

// bindExampleERC20 binds a generic wrapper to an already deployed contract.
func bindExampleERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExampleERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExampleERC20 *ExampleERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExampleERC20.Contract.ExampleERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExampleERC20 *ExampleERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExampleERC20.Contract.ExampleERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExampleERC20 *ExampleERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExampleERC20.Contract.ExampleERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExampleERC20 *ExampleERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExampleERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExampleERC20 *ExampleERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExampleERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExampleERC20 *ExampleERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExampleERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ExampleERC20 *ExampleERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ExampleERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ExampleERC20 *ExampleERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ExampleERC20.Contract.Allowance(&_ExampleERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ExampleERC20 *ExampleERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ExampleERC20.Contract.Allowance(&_ExampleERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ExampleERC20 *ExampleERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ExampleERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ExampleERC20 *ExampleERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _ExampleERC20.Contract.BalanceOf(&_ExampleERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ExampleERC20 *ExampleERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ExampleERC20.Contract.BalanceOf(&_ExampleERC20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ExampleERC20 *ExampleERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ExampleERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ExampleERC20 *ExampleERC20Session) Decimals() (uint8, error) {
	return _ExampleERC20.Contract.Decimals(&_ExampleERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ExampleERC20 *ExampleERC20CallerSession) Decimals() (uint8, error) {
	return _ExampleERC20.Contract.Decimals(&_ExampleERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ExampleERC20 *ExampleERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ExampleERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ExampleERC20 *ExampleERC20Session) Name() (string, error) {
	return _ExampleERC20.Contract.Name(&_ExampleERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ExampleERC20 *ExampleERC20CallerSession) Name() (string, error) {
	return _ExampleERC20.Contract.Name(&_ExampleERC20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ExampleERC20 *ExampleERC20Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExampleERC20.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ExampleERC20 *ExampleERC20Session) Owner() (common.Address, error) {
	return _ExampleERC20.Contract.Owner(&_ExampleERC20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ExampleERC20 *ExampleERC20CallerSession) Owner() (common.Address, error) {
	return _ExampleERC20.Contract.Owner(&_ExampleERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ExampleERC20 *ExampleERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ExampleERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ExampleERC20 *ExampleERC20Session) Symbol() (string, error) {
	return _ExampleERC20.Contract.Symbol(&_ExampleERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ExampleERC20 *ExampleERC20CallerSession) Symbol() (string, error) {
	return _ExampleERC20.Contract.Symbol(&_ExampleERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ExampleERC20 *ExampleERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ExampleERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ExampleERC20 *ExampleERC20Session) TotalSupply() (*big.Int, error) {
	return _ExampleERC20.Contract.TotalSupply(&_ExampleERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ExampleERC20 *ExampleERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ExampleERC20.Contract.TotalSupply(&_ExampleERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ExampleERC20 *ExampleERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ExampleERC20 *ExampleERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.Contract.Approve(&_ExampleERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ExampleERC20 *ExampleERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.Contract.Approve(&_ExampleERC20.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ExampleERC20 *ExampleERC20Transactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ExampleERC20 *ExampleERC20Session) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.Contract.Burn(&_ExampleERC20.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ExampleERC20 *ExampleERC20TransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.Contract.Burn(&_ExampleERC20.TransactOpts, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ExampleERC20 *ExampleERC20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ExampleERC20 *ExampleERC20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.Contract.DecreaseAllowance(&_ExampleERC20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ExampleERC20 *ExampleERC20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.Contract.DecreaseAllowance(&_ExampleERC20.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ExampleERC20 *ExampleERC20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ExampleERC20 *ExampleERC20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.Contract.IncreaseAllowance(&_ExampleERC20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ExampleERC20 *ExampleERC20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.Contract.IncreaseAllowance(&_ExampleERC20.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_ExampleERC20 *ExampleERC20Transactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_ExampleERC20 *ExampleERC20Session) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.Contract.Mint(&_ExampleERC20.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_ExampleERC20 *ExampleERC20TransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.Contract.Mint(&_ExampleERC20.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ExampleERC20 *ExampleERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ExampleERC20 *ExampleERC20Session) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.Contract.Transfer(&_ExampleERC20.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ExampleERC20 *ExampleERC20TransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.Contract.Transfer(&_ExampleERC20.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ExampleERC20 *ExampleERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ExampleERC20 *ExampleERC20Session) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.Contract.TransferFrom(&_ExampleERC20.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ExampleERC20 *ExampleERC20TransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExampleERC20.Contract.TransferFrom(&_ExampleERC20.TransactOpts, from, to, amount)
}

// ExampleERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ExampleERC20 contract.
type ExampleERC20ApprovalIterator struct {
	Event *ExampleERC20Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExampleERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExampleERC20Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExampleERC20Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExampleERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExampleERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExampleERC20Approval represents a Approval event raised by the ExampleERC20 contract.
type ExampleERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ExampleERC20 *ExampleERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ExampleERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ExampleERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ExampleERC20ApprovalIterator{contract: _ExampleERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ExampleERC20 *ExampleERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ExampleERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ExampleERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExampleERC20Approval)
				if err := _ExampleERC20.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ExampleERC20 *ExampleERC20Filterer) ParseApproval(log types.Log) (*ExampleERC20Approval, error) {
	event := new(ExampleERC20Approval)
	if err := _ExampleERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExampleERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ExampleERC20 contract.
type ExampleERC20TransferIterator struct {
	Event *ExampleERC20Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExampleERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExampleERC20Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExampleERC20Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExampleERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExampleERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExampleERC20Transfer represents a Transfer event raised by the ExampleERC20 contract.
type ExampleERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ExampleERC20 *ExampleERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ExampleERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ExampleERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ExampleERC20TransferIterator{contract: _ExampleERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ExampleERC20 *ExampleERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ExampleERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ExampleERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExampleERC20Transfer)
				if err := _ExampleERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ExampleERC20 *ExampleERC20Filterer) ParseTransfer(log types.Log) (*ExampleERC20Transfer, error) {
	event := new(ExampleERC20Transfer)
	if err := _ExampleERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
