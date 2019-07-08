package set1

import (
	"io/ioutil"
	"log"
	"math"

	"github.com/mrbrianhobo/cryptopals/utils"
)

func GetRawText(filename string) string {
	rawBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return string(rawBytes)
}

func GetCiphertext(filename string) []byte {
	base64Bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	base64String := string(base64Bytes)
	return utils.Base64ToBytes(base64String)
}

func DecryptRepeatingKeyXORCipher(ciphertext []byte) string {
	keysize := findKeysize(ciphertext)
	blocks := blockify(ciphertext, keysize)
	transposed := transposeBlocks(blocks, keysize)

	key := make([]byte, keysize)
	for i, block := range transposed {
		currentKey := DecryptSingleByteXORCipher(block)
		key[i] = currentKey
	}

	return string(key)
}

func DecryptRepeatingKeyXORCipherWithKey(ciphertext []byte, key string) string {
	keyBytes := []byte(key)
	out := make([]byte, len(ciphertext))

	for i, char := range ciphertext {
		out[i] = char ^ keyBytes[i%len(key)]
	}

	return string(out)
}

func blockify(ciphertext []byte, keysize int) [][]byte {
	blocks := make([][]byte, 0)
	for i := 0; i < len(ciphertext); i += keysize {
		end := i + keysize
		if end > len(ciphertext) {
			end = len(ciphertext)
		}
		block := ciphertext[i:end]
		blocks = append(blocks, block)
	}
	return blocks
}

func transposeBlocks(blocks [][]byte, keysize int) [][]byte {
	transposed := make([][]byte, keysize)
	for i := 0; i < len(blocks); i++ {
		for j := range blocks[i] {
			transposed[j] = append(transposed[j], blocks[i][j])
		}
	}
	return transposed
}

func findKeysize(ciphertext []byte) int {
	minHammingDist := math.MaxFloat64
	minKeysize := 0
	for keysize := 2; keysize <= 40; keysize++ {
		blocks := blockify(ciphertext, keysize)
		totalDistance := 0
		for i := 0; i < len(blocks)-2; i++ {
			block1, block2 := string(blocks[i]), string(blocks[i+1])
			totalDistance += HammingDistance(block1, block2)
		}
		normalized := float64(totalDistance) / float64(len(blocks)) / float64(keysize)
		if normalized < minHammingDist {
			minHammingDist = normalized
			minKeysize = keysize
		}
	}
	return minKeysize
}

func HammingDistance(str1 string, str2 string) int {
	bitstring1 := []byte(str1)
	bitstring2 := []byte(str2)

	if len(bitstring1) != len(bitstring2) {
		log.Fatal("strings are not equal length")
	}

	dist := 0
	for i := range bitstring1 {
		b := bitstring1[i] ^ bitstring2[i]
		dist += countBits(b)
	}

	return dist
}

func countBits(b byte) int {
	bits := 0
	for b != byte(0) {
		b &= b - 1
		bits += 1
	}
	return bits
}
