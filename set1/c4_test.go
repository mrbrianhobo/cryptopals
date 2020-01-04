package set1

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mrbrianhobo/cryptopals/utils"
)

func TestDetectSingleByteXOR(t *testing.T) {

	const (
		hexString = "7b5a4215415d544115415d5015455447414c155c46155f4058455c5b523f"
		key       = byte('5')
		expected  = "Now that the party is jumping\n"
	)

	rawtext := utils.GetRawText("../challenge-data/4.txt")
	detected := DetectSingleByteXOR(rawtext)
	ciphertext := utils.HexToBytes(detected)
	actualKey := DecryptSingleByteXORCipher(ciphertext)
	decrypted := DecryptXORCipherWithKey(ciphertext, actualKey)

	assert.Equal(t, hexString, detected)
	assert.Equal(t, key, actualKey)
	assert.Equal(t, expected, decrypted)
}
