package set1

import (
	"encoding/hex"
)

func RepeatingKeyXOR(plaintext string, key string) string {
	out := make([]byte, len(plaintext))

	for i, char := range plaintext {
		out[i] = byte(char) ^ byte(key[i%len(key)])
	}

	return hex.EncodeToString(out)
}
