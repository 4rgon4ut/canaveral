package utils

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// Takes path to the *.abi file, read it and returns parsed ABI interface object
func GetABIObject(abiPath string) (*abi.ABI, error) {
	b, err := os.ReadFile(abiPath)
	if err != nil {
		return nil, err
	}
	evmABI, err := abi.JSON(strings.NewReader(string(b)))
	if err != nil {
		return nil, err
	}
	return &evmABI, nil
}

// TODO: deprecate?
func GetBytecode(binPath string) ([]byte, error) {
	b, err := os.ReadFile(binPath)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Cast types of args([]string) to input types of provided function type.
// Ignores first two inputs of provided function(starts with index 2) and try to cast
// inputs respectively:
//
// funcType(_, _, type1, type2, ...) --> args[0].(type1), args[1].(type2), ...
//
// Function supports only several types(any <uintX>, <string>), plus ETH <common.Address> and <[32]byte>
// and is not safe to buffer overflow
//
// TODO: add checks for resulting types overflow
// TODO: spread supported types number
func CastInputs(abiArgs abi.Arguments, args []string) ([]interface{}, error) {
	assertedArgs := make([]interface{}, len(args))
	// +2 to args index here to match arguments properly
	// Deploy<ContractName>(auth *bind.TransactOpts, backend bind.ContractBackend, ... ) always have first two arguments
	// here me casting only contract constructor inputs
	for i, arg := range abiArgs {
		stringArg := args[i]
		// string representation of argument to find a match
		switch arg.Type.String() {
		case "uint8":
			u64, err := strconv.ParseUint(stringArg, 10, 32)
			if err != nil {
				return nil, err
			}
			assertedArgs[i] = uint8(u64)
		case "uint16":
			u64, err := strconv.ParseUint(stringArg, 10, 32)
			if err != nil {
				return nil, err
			}
			assertedArgs[i] = uint16(u64)
		case "uint32":
			u64, err := strconv.ParseUint(stringArg, 10, 32)
			if err != nil {
				return nil, err
			}
			assertedArgs[i] = uint32(u64)
		case "uint64":
			u64, err := strconv.ParseUint(stringArg, 10, 32)
			if err != nil {
				return nil, err
			}
			assertedArgs[i] = u64
		case "uint128":
			fallthrough // EVM u128 & u256 both represented as big.Int in Go
		case "uint256":
			n := new(big.Int)
			n, ok := n.SetString(stringArg, 10)
			if !ok {
				return nil, fmt.Errorf("can't convert %s to big.Int", stringArg)
			}
			assertedArgs[i] = n
		case "string":
			assertedArgs[i] = stringArg
		case "address":
			assertedArgs[i] = common.HexToAddress(stringArg)
		case "bytes32":
			var param [32]uint8
			copy(param[:], stringArg)
			assertedArgs[i] = param
		}
	}
	return assertedArgs, nil
}
