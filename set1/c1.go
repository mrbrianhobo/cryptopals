package set1

import (
	"github.com/mrbrianhobo/cryptopals/utils"
)

func ConvertHexToBase64(hexString string) string {
	src := utils.HexToBytes(hexString)
	return utils.BytesToBase64(src)
}
