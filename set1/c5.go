package set1

import (
	"github.com/mrbrianhobo/cryptopals/utils"
)

func RepeatingKeyXOR(plaintext string, key string) string {
	out := make([]byte, len(plaintext))

	for i, char := range plaintext {
		out[i] = byte(char) ^ byte(key[i%len(key)])
	}

	return utils.BytesToHex(out)
}
