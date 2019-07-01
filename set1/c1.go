package set1

import (
	"log"

	"encoding/base64"
	"encoding/hex"
)

func ConvertHexToBase64(src string) string {
	decodedHex, err := hex.DecodeString(src)
	if err != nil {
		log.Fatal(err)
	}

	base64String := base64.StdEncoding.EncodeToString(decodedHex)
	return base64String
}
