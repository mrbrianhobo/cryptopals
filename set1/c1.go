package set1

import (
	"log"

	"encoding/base64"
	"encoding/hex"
)

func convertHexToBase64(src string) string {
	decodedHex, err := hex.DecodeString(src)
	if err != nil {
		log.Fatal(err)
	}

	base64String := base64.StdEncoding.EncodeToString(decodedHex)
	return base64String
}
