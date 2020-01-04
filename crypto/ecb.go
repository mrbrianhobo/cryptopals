package crypto

import (
	"crypto/aes"
	"log"
)

func ECBEncrypt(plaintext, key []byte) []byte {
	plaintext = Pad(plaintext, aes.BlockSize)

	validateInputsECB(plaintext, key)

	block, _ := aes.NewCipher(key)
	encrypted := make([]byte, 0)

	for len(plaintext) > 0 {
		dst := make([]byte, block.BlockSize())
		block.Encrypt(dst, plaintext)
		encrypted = append(encrypted, dst...)
		plaintext = plaintext[block.BlockSize():]
	}

	return encrypted
}

func ECBDecrypt(ciphertext, key []byte) []byte {
	validateInputsECB(ciphertext, key)

	block, _ := aes.NewCipher(key)
	decrypted := make([]byte, 0)

	for len(ciphertext) > 0 {
		dst := make([]byte, block.BlockSize())
		block.Decrypt(dst, ciphertext)
		decrypted = append(decrypted, dst...)
		ciphertext = ciphertext[block.BlockSize():]
	}

	return decrypted
}

func validateInputsECB(text, key []byte) {
	if len(text) < aes.BlockSize {
		log.Fatal("plaintext/ciphertext too short")
	}

	if len(text)%aes.BlockSize != 0 {
		log.Fatal("plaintext/ciphertext is not a multiple of the block size")
	}
}
