package crypto

func Pad(plaintext []byte, blocksize int) []byte {
	n := blocksize - len(plaintext)%blocksize

	padding := make([]byte, n)
	for i := range padding {
		padding[i] = byte(n)
	}

	plaintext = append(plaintext, padding...)
	return plaintext
}
