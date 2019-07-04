package set1

import (
	"log"

	"github.com/mrbrianhobo/cryptopals/utils"
)

func FixedXOR(hexString1 string, hexString2 string) string {
	src1 := utils.HexToBytes(hexString1)
	src2 := utils.HexToBytes(hexString2)

	if len(src1) != len(src2) {
		log.Fatal("buffers are not equal length")
	}

	out := make([]byte, len(src1))

	for i, b := range src1 {
		out[i] = b ^ src2[i]
	}

	return utils.BytesToHex(out)
}
