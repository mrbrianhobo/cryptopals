package set1

import (
	"log"
	"math"

	"encoding/hex"

	"github.com/mrbrianhobo/cryptopals/utils"
)

const (
	alphabet = "QWERTYUIOPASDFGHJKLZXCVBNM"
)

func DecryptSingleByteXORCipher(src string) byte {
	decodedHex, err := hex.DecodeString(src)
	if err != nil {
		log.Fatal(err)
	}

	m := utils.NewMonogramStats()
	max, char := -math.MaxFloat64, byte(0)
	for _, c := range alphabet {
		current := string(singleByteXOR(decodedHex, byte(c)))
		currentLogScore := m.Score(current)
		if currentLogScore > max {
			max = currentLogScore
			char = byte(c)
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

func DecryptXORCipherWithKey(src string, key byte) string {
	decodedHex, err := hex.DecodeString(src)
	if err != nil {
		log.Fatal(err)
	}

	return string(singleByteXOR(decodedHex, byte(key)))
}
