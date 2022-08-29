package utils_test

import (
	"reflect"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/canaveral/utils"
)

var _ = Describe("Abispecfic", func() {
	Context("GetABIObject", func() {
		abiPath := "../artifacts/abi/Test.abi"
		It("should load and return abi type object without error", func() {
			object, err := utils.GetABIObject(abiPath)
			expectedType := new(abi.ABI)
			Expect(err).To(BeNil())
			Expect(object).To(BeAssignableToTypeOf(expectedType))
		})
		It("should return error on invalid .abi file path", func() {
			object, err := utils.GetABIObject("dwefwefwefewwfe")
			Expect(err).Error()
			Expect(err).ToNot(BeNil())
			Expect(object).To(BeNil())
		})
	})

	Context("CastInputs", func() {
		// _, _, here to shift arguments index proerly
		// Deploy<ContractName>(auth *bind.TransactOpts, backend bind.ContractBackend, ... ) always have first two arguments
		function := func(_, _, u uint16, word string, byteArray [32]byte, etherAddr common.Address) {}
		functionType := reflect.TypeOf(function)
		var b [32]byte
		// slice with correct types to match after cast
		correctlyTypedSlice := []interface{}{uint16(666), "Hola!", b, common.HexToAddress("0xA0FA3CaCDc397F49C4D19221Bcef495b283D2422")}
		argsToCast := []string{"666", "Hola!", "string to bytes", "0xA0FA3CaCDc397F49C4D19221Bcef495b283D2422"}

		It("should cast slice of strings to provided function type input params types", func() {
			inputs, err := utils.CastInputs(functionType, argsToCast)
			Expect(err).To(BeNil())
			Expect(len(inputs)).To(Equal(len(argsToCast)))
			// types matching
			for i, arg := range inputs {
				Expect(arg).To(BeAssignableToTypeOf(correctlyTypedSlice[i]))
			}
		})
		It("should return error if unable to cast", func() {
			// make first arg of expected type uint16 unconvertable(string)
			argsToCast[0] = "csadcds"
			inputs, err := utils.CastInputs(functionType, argsToCast)
			Expect(err).ToNot(BeNil())
			Expect(inputs).To(BeNil())
		})
	})
})
