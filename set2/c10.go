package set2

import (
	"crypto/aes"
	"crypto/cipher"
	"log"
)

func EncryptAESCBC(plaintext, key, iv []byte) []byte {
	validateInputs(plaintext, key, iv)

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	encrypted := make([]byte, len(plaintext))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(encrypted, plaintext)

	return encrypted
}

func DecryptAESCBC(ciphertext, key, iv []byte) []byte {
	validateInputs(ciphertext, key, iv)

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	decrypted := make([]byte, len(ciphertext))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(decrypted, ciphertext)

	return decrypted
}

func validateInputs(text, key, iv []byte) {
	if len(iv) != aes.BlockSize {
		log.Fatal("iv must be same lenght as the block size")
	}

	if len(text) < aes.BlockSize {
		log.Fatal("plaintext/ciphertext too short")
	}

	if len(text)%aes.BlockSize != 0 {
		log.Fatal("plaintext/ciphertext is not a multiple of the block size")
	}
}
