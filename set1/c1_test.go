package set1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCovertHexToBase64(t *testing.T) {

	const (
		hexString = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
		expected  = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	)

	actual := ConvertHexToBase64(hexString)

	assert.Equal(t, expected, actual)
}
