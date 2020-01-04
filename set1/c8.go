package set1

import (
	"crypto/aes"
	"strings"

	"github.com/mrbrianhobo/cryptopals/utils"
)

func DetectAESECB(rawtext string) string {
	max, hexString := 0, ""

	for _, line := range strings.Split(rawtext, "\n") {
		ciphertext := utils.HexToBytes(line)
		score := countRepeatingBlocks(ciphertext)
		if score > max {
			max = score
			hexString = line
		}
	}

	return hexString
}

func countRepeatingBlocks(ciphertext []byte) int {
	blocks := make(map[string]int)
	for len(ciphertext) > 0 {
		block := string(ciphertext[:aes.BlockSize])
		if val, ok := blocks[block]; ok {
			blocks[block] = val + 1
		} else {
			blocks[block] = 0
		}
		ciphertext = ciphertext[aes.BlockSize:]
	}

	score := 0
	for _, val := range blocks {
		score += val
	}

	return score
}
