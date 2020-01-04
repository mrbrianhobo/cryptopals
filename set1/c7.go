package set1

import (
	"crypto/aes"
)

func DecryptAESECB(ciphertext []byte, key []byte) string {
	block, _ := aes.NewCipher(key)
	decrypted := make([]byte, 0)

	for len(ciphertext) > 0 {
		dst := make([]byte, block.BlockSize())
		block.Decrypt(dst, ciphertext)
		decrypted = append(decrypted, dst...)
		ciphertext = ciphertext[block.BlockSize():]
	}

	return string(decrypted)
}
