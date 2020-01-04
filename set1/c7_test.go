package set1

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mrbrianhobo/cryptopals/utils"
)

func TestDecryptAESECB(t *testing.T) {

	var (
		expected = utils.GetRawText("../expected/7_out.txt")
		key      = "YELLOW SUBMARINE"
	)

	ciphertext := utils.GetBase64Ciphertext("../challenge-data/7.txt")
	decrypted := DecryptAESECB(ciphertext, key)

	assert.Equal(t, expected, decrypted)
}
