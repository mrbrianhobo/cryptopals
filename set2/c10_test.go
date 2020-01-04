package set2

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mrbrianhobo/cryptopals/utils"
)

func TestDecryptAESCBC(t *testing.T) {

	var (
		key       = []byte("YELLOW SUBMARINE")
		iv        = make([]byte, 16)
		plaintext = utils.GetRawText("../expected/10_out.txt")
	)

	ciphertext := utils.GetBase64Ciphertext("../challenge-data/10.txt")
	decrypted := DecryptAESCBC(ciphertext, key, iv)
	encrypted := EncryptAESCBC(decrypted, key, iv)

	assert.Equal(t, plaintext, string(decrypted))
	assert.Equal(t, ciphertext, encrypted)
}
