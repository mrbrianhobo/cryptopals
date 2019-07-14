package set2

func PKCS7Padding(plaintext string, blockLength int) string {
	bytes := []byte(plaintext)
	n := blockLength - len(plaintext)%blockLength

	padding := make([]byte, n)
	for i := range padding {
		padding[i] = byte(n)
	}

	bytes = append(bytes, padding...)
	return string(bytes)
}
