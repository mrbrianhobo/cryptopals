package set2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPKCS7Padding(t *testing.T) {

	const (
		plaintext = "YELLOW SUBMARINE"
		expected  = "YELLOW SUBMARINE\x04\x04\x04\x04"
	)

	actual := PKCS7Padding(plaintext, 20)

	assert.Equal(t, expected, actual)
}
