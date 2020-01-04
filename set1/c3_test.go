package set1

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mrbrianhobo/cryptopals/utils"
)

func TestSingleByteXORCipher(t *testing.T) {

	const (
		hexString = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
		key       = byte('X')
		expected  = "Cooking MC's like a pound of bacon"
	)

	ciphertext := utils.HexToBytes(hexString)
	actualKey := DecryptSingleByteXORCipher(ciphertext)
	decrypted := DecryptXORCipherWithKey(ciphertext, actualKey)

	assert.Equal(t, key, actualKey)
	assert.Equal(t, expected, decrypted)
}
