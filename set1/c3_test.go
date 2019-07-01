package set1

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestSingleByteXORCipher(t *testing.T) {
	gomega := NewWithT(t)

	const (
		ciphertext = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
		key        = byte('X')
		expected   = "Cooking MC's like a pound of bacon"
	)

	actualKey := DecryptSingleByteXORCipher(ciphertext)
	decrypted := DecryptXORCipherWithKey(ciphertext, actualKey)

	gomega.Expect(actualKey).To(Equal(key))
	gomega.Expect(decrypted).To(Equal(expected))
}
