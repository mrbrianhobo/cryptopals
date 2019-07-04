package set1

import (
	"bufio"
	"log"
	"math"
	"os"

	"github.com/mrbrianhobo/cryptopals/utils"
)

func DetectSingleByteXOR(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	m := utils.NewMonogramStats()
	max, hexString := -math.MaxFloat64, ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
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
