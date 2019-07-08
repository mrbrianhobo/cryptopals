package utils

import (
	"encoding/base64"
	"encoding/hex"
	"io/ioutil"
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

func GetRawText(filename string) string {
	rawBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return string(rawBytes)
}

func GetBase64Ciphertext(filename string) []byte {
	base64Bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	base64String := string(base64Bytes)
	return Base64ToBytes(base64String)
}
