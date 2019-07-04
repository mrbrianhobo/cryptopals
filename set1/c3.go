package set1

import (
	"math"

	"github.com/mrbrianhobo/cryptopals/utils"
)

func DecryptSingleByteXORCipher(ciphertext []byte) byte {
	m := utils.NewMonogramStats()
	max, char := -math.MaxFloat64, byte(0)
	for i := 0; i < 256; i++ {
		current := string(singleByteXOR(ciphertext, byte(i)))
		currentLogScore := m.Score(current)
		if currentLogScore > max {
			max = currentLogScore
			char = byte(i)
		}
	}

	return char
}

func singleByteXOR(src []byte, char byte) []byte {
	out := make([]byte, len(src))
	for i := range src {
		out[i] = src[i] ^ char
	}
	return out
}

func DecryptXORCipherWithKey(ciphertext []byte, key byte) string {
	return string(singleByteXOR(ciphertext, byte(key)))
}
