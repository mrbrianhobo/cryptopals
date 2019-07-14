package set1

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/mrbrianhobo/cryptopals/utils"
)

func TestDecryptAESECB(t *testing.T) {
	RegisterTestingT(t)

	var (
		expected = utils.GetRawText("../expected/7_out.txt")
	)

	ciphertext := utils.GetBase64Ciphertext("../challenge-data/7.txt")
	decrypted := DecryptAESECB(ciphertext)

	Expect(decrypted).To(Equal(expected))
}
