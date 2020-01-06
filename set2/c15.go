package set2

import (
	"errors"
)

func ValidatePKCS7(plaintext string) (string, error) {
	bytes := []byte(plaintext)

	last := bytes[len(bytes)-1]

	if int(last) > len(bytes) || last == byte(0) {
		return "", errors.New("Invalid padding")
	}

	for i := 0; i < int(last); i++ {
		if bytes[len(bytes)-i-1] != last {
			// return error
			return "", errors.New("Invalid padding")
		}
	}

	return string(bytes[0 : len(bytes)-int(last)]), nil
}
