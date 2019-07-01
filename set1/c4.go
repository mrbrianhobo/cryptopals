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
	max, ciphertext := -math.MaxFloat64, ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		key := DecryptSingleByteXORCipher(line)
		decrypted := DecryptXORCipherWithKey(line, key)
		currentLogScore := m.Score(decrypted)
		if currentLogScore > max {
			max = currentLogScore
			ciphertext = line
		}
	}

	return ciphertext
}
