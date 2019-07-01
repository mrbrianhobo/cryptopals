package set1

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestDetectSingleByteXOR(t *testing.T) {
	gomega := NewWithT(t)

	const (
		ciphertext = "7b5a4215415d544115415d5015455447414c155c46155f4058455c5b523f"
		key        = byte('5')
		expected   = "Now that the party is jumping\n"
	)

	actualCiphertext := DetectSingleByteXOR("../challenge-data/4.txt")
	actualKey := DecryptSingleByteXORCipher(actualCiphertext)
	decrypted := DecryptXORCipherWithKey(actualCiphertext, actualKey)

	gomega.Expect(actualCiphertext).To(Equal(ciphertext))
	gomega.Expect(actualKey).To(Equal(key))
	gomega.Expect(decrypted).To(Equal(expected))
}
