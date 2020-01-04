package set1

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mrbrianhobo/cryptopals/utils"
)

func TestHammingDistance(t *testing.T) {

	const (
		str1     = "this is a test"
		str2     = "wokka wokka!!!"
		expected = 37
	)

	actual := HammingDistance(str1, str2)

	assert.Equal(t, expected, actual)
}

func TestDecryptRepeatingKeyXORCipher(t *testing.T) {

	var (
		keysize  = 29
		key      = "Terminator X: Bring the noise"
		expected = utils.GetRawText("../expected/6_out.txt")
	)

	ciphertext := utils.GetBase64Ciphertext("../challenge-data/6.txt")
	actualKey := DecryptRepeatingKeyXORCipher(ciphertext)
	decrypted := DecryptRepeatingKeyXORCipherWithKey(ciphertext, actualKey)

	assert.Len(t, actualKey, keysize)
	assert.Equal(t, key, actualKey)
	assert.Equal(t, expected, decrypted)
}
