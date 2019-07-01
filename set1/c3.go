package set1

import (
	"log"
	"math"

	"encoding/hex"

	"github.com/mrbrianhobo/cryptopals/utils"
)

func DecryptSingleByteXORCipher(ciphertext string) byte {
	decodedHex, err := hex.DecodeString(ciphertext)
	if err != nil {
		log.Fatal(err)
	}

	m := utils.NewMonogramStats()
	max, char := -math.MaxFloat64, byte(0)
	for i := 0; i < 256; i++ {
		current := string(singleByteXOR(decodedHex, byte(i)))
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

func DecryptXORCipherWithKey(ciphertext string, key byte) string {
	decodedHex, err := hex.DecodeString(ciphertext)
	if err != nil {
		log.Fatal(err)
	}

	return string(singleByteXOR(decodedHex, byte(key)))
}
