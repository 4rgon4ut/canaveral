package evmclient_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestEvmclient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Evmclient Suite")
}
