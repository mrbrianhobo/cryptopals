package set1

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestFixedXOR(t *testing.T) {
	gomega := NewWithT(t)

	const (
		hexString1 = "1c0111001f010100061a024b53535009181c"
		hexString2 = "686974207468652062756c6c277320657965"
		expected   = "746865206b696420646f6e277420706c6179"
	)

	actual := FixedXOR(hexString1, hexString2)

	gomega.Expect(actual).To(Equal(expected))
}
