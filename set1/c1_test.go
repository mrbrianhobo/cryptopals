package set1

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestCovertHexToBase64(t *testing.T) {
	gomega := NewWithT(t)

	const (
		hexString = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
		expected  = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	)

	actual := ConvertHexToBase64(hexString)

	gomega.Expect(actual).To(Equal(expected))
}
