package set2

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/mrbrianhobo/cryptopals/utils"
)

func TestDecryptAESCBC(t *testing.T) {
	RegisterTestingT(t)

	var (
		key       = []byte("YELLOW SUBMARINE")
		iv        = make([]byte, 16)
		plaintext = utils.GetRawText("../expected/10_out.txt")
	)

	ciphertext := utils.GetBase64Ciphertext("../challenge-data/10.txt")
	decrypted := DecryptAESCBC(ciphertext, key, iv)
	encrypted := EncryptAESCBC(decrypted, key, iv)

	Expect(string(decrypted)).To(Equal(plaintext))
	Expect(encrypted).To(Equal(ciphertext))
}
