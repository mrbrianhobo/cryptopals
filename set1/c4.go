package set1

import (
	"math"
	"strings"

	"github.com/mrbrianhobo/cryptopals/utils"
)

func DetectSingleByteXOR(rawtext string) string {
	m := utils.NewMonogramStats()
	max, hexString := -math.MaxFloat64, ""

	for _, line := range strings.Split(rawtext, "\n") {
		ciphertext := utils.HexToBytes(line)
		key := DecryptSingleByteXORCipher(ciphertext)
		decrypted := DecryptXORCipherWithKey(ciphertext, key)
		currentLogScore := m.Score(decrypted)
		if currentLogScore > max {
			max = currentLogScore
			hexString = line
		}
	}

	return hexString
}
