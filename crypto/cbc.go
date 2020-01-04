package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"log"
)

func CBCEncrypt(plaintext, key, iv []byte) []byte {
	plaintext = Pad(plaintext, aes.BlockSize)

	validateInputsCBC(plaintext, key, iv)

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	encrypted := make([]byte, len(plaintext))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(encrypted, plaintext)

	return encrypted
}

func CBCDecrypt(ciphertext, key, iv []byte) []byte {
	validateInputsCBC(ciphertext, key, iv)

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	decrypted := make([]byte, len(ciphertext))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(decrypted, ciphertext)

	return decrypted
}

func validateInputsCBC(text, key, iv []byte) {
	if len(iv) != aes.BlockSize {
		log.Fatal("iv must be same length as the block size")
	}

	if len(text) < aes.BlockSize {
		log.Fatal("plaintext/ciphertext too short")
	}

	if len(text)%aes.BlockSize != 0 {
		log.Fatal("plaintext/ciphertext is not a multiple of the block size")
	}
}
