package evmclient_test

import (
	"math/big"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/canaveral/evmclient"
)

var _ = Describe("Account", func() {
	It("should create new account instance", func() {
		account, err := evmclient.NewAccount(
			"41221710A88C7284D5F5BDE53311CDFA22B20F2724F3F4DA60FE3F567669DF56",
			big.NewInt(9000),
		)
		Expect(err).To(BeNil())
		Expect(account).NotTo(BeNil())
		Expect(account.Address.Hex()).ToNot(Equal(""))
		Expect(account.Signer.From).To(Equal(account.Address))
		// TODO: add signing and recover
	})
	It("should return error with invalid private key", func() {
		account, err := evmclient.NewAccount(
			"1234",
			big.NewInt(9000),
		)
		Expect(err).NotTo(BeNil())
		Expect(account).To(BeNil())
	})
})
