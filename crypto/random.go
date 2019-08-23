package crypto

import (
	"crypto/aes"
	"crypto/rand"
	"log"
	"math/big"
)

func RandomInt(min, max int) int {
	bg := big.NewInt(int64(max - min + 1))
	n, _ := rand.Int(rand.Reader, bg)

	return int(n.Int64()) + min
}

func RandomBytes(n int) []byte {
	validateInputs(n)

	bytes := make([]byte, n)
	rand.Read(bytes)
	return bytes
}

func RandomByte() byte {
	return RandomBytes(1)[0]
}

func RandomKey() []byte {
	return RandomBytes(aes.BlockSize)
}

func validateInputs(n int) {
	if n < 0 {
		log.Fatal("cannot generate negative length byte array")
	}
}
