package set2

import (
	"crypto/aes"

	"github.com/mrbrianhobo/cryptopals/crypto"
)

type EncryptionObject struct {
	ciphertext []byte
	mode       string
}

func EncryptionOracle(plaintext string) EncryptionObject {
	bytes := []byte(plaintext)
	key := crypto.RandomKey()
	prefix := crypto.RandomBytes(crypto.RandomInt(5, 10))
	suffix := crypto.RandomBytes(crypto.RandomInt(5, 10))

	bytes = append(bytes, suffix...)
	bytes = append(prefix, bytes...)

	bit := crypto.RandomInt(0, 1)
	if bit == 0 {
		// encrypt under ECB
		return EncryptionObject{
			ciphertext: crypto.ECBEncrypt(bytes, key),
			mode:       "ECB",
		}
	} else {
		// encrypt under CBC (with random iv)
		return EncryptionObject{
			ciphertext: crypto.CBCEncrypt(bytes, key, crypto.RandomKey()),
			mode:       "CBC",
		}
	}
}

func detectECB(ciphertext []byte) bool {
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

	if score > 1 {
		return true
	}

	return false
}

func DetectMode(ciphertext []byte) string {
	if detectECB(ciphertext) {
		return "ECB"
	} else {
		return "CBC"
	}
}
