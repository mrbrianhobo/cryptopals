package set1

import (
	"testing"

	"github.com/mrbrianhobo/cryptopals/utils"
	. "github.com/onsi/gomega"
)

func TestDetectSingleByteXOR(t *testing.T) {
	RegisterTestingT(t)

	const (
		hexString = "7b5a4215415d544115415d5015455447414c155c46155f4058455c5b523f"
		key       = byte('5')
		expected  = "Now that the party is jumping\n"
	)

	actualHex := DetectSingleByteXOR("../challenge-data/4.txt")
	ciphertext := utils.HexToBytes(actualHex)
	actualKey := DecryptSingleByteXORCipher(ciphertext)
	decrypted := DecryptXORCipherWithKey(ciphertext, actualKey)

	Expect(actualHex).To(Equal(hexString))
	Expect(actualKey).To(Equal(key))
	Expect(decrypted).To(Equal(expected))
}
