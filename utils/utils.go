package utils

import (
	"encoding/base64"
	"encoding/hex"
	"log"
)

func HexToBytes(hexString string) []byte {
	ciphertext, err := hex.DecodeString(hexString)
	if err != nil {
		log.Fatal(err)
	}

	return ciphertext
}

func Base64ToBytes(base64String string) []byte {
	ciphertext, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		log.Fatal(err)
	}

	return ciphertext
}

func BytesToHex(src []byte) string {
	return hex.EncodeToString(src)
}

func BytesToBase64(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}
