package utils_test

import (
	"github.com/canaveral/utils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Format", func() {

	Context("RemoveExtension", func() {
		fileName := "BlahBlah"
		fileNameWithDot := "Blah.Blah"
		extension := ".sol"
		fileNameExtended := fileName + extension
		fileNameWithDotExtended := fileNameWithDot + extension
		It("should remove extension from file name", func() {
			Expect(utils.RemoveExtension(fileNameExtended)).To(Equal(fileName))
		})
		It("should ignore dots in filename and revome extension only", func() {
			Expect(utils.RemoveExtension(fileNameWithDotExtended)).To(Equal(fileNameWithDot))
		})
	})
})
