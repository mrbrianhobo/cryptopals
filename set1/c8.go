package set1

import (
	"bufio"
	"crypto/aes"
	"log"
	"os"

	"github.com/mrbrianhobo/cryptopals/utils"
)

func DetectAESECB(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	max, hexString := 0, ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
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
