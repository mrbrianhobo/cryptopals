package set1

import (
	"log"

	"encoding/hex"
)

func fixedXOR(src1 string, src2 string) string {
	if len(src1) != len(src2) {
		log.Fatal("buffers are not equal length")
	}

	decodedHex1, err := hex.DecodeString(src1)
	if err != nil {
		log.Fatal(err)
	}

	decodedHex2, err := hex.DecodeString(src2)
	if err != nil {
		log.Fatal(err)
	}

	out := make([]byte, len(decodedHex1))

	for i, b := range decodedHex1 {
		out[i] = b ^ decodedHex2[i]
	}

	return hex.EncodeToString(out)
}
