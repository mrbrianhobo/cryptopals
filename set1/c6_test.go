package set1

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestHammingDistance(t *testing.T) {
	gomega := NewWithT(t)

	const (
		str1     = "this is a test"
		str2     = "wokka wokka!!!"
		expected = 37
	)

	actual := HammingDistance(str1, str2)

	gomega.Expect(actual).To(Equal(expected))
}

func TestDecryptRepeatingKeyXORCipher(t *testing.T) {
	gomega := NewWithT(t)

	var (
		keysize  = 29
		key      = "Terminator X: Bring the noise"
		expected = GetRawText("../expected/6_out.txt")
	)

	ciphertext := GetCiphertext("../challenge-data/6.txt")
	actualKey := DecryptRepeatingKeyXORCipher(ciphertext)
	decrypted := DecryptRepeatingKeyXORCipherWithKey(ciphertext, actualKey)

	gomega.Expect(len(actualKey)).To(Equal(keysize))
	gomega.Expect(actualKey).To(Equal(key))
	gomega.Expect(decrypted).To(Equal(expected))
}
