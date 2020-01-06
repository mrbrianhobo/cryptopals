package set2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatePKCS7(t *testing.T) {
	const (
		plaintext = "ICE ICE BABY\x04\x04\x04\x04"
		expected  = "ICE ICE BABY"
		invalid1  = "ICE ICE BABY\x05\x05\x05\x05"
		invalid2  = "ICE ICE BABY\x01\x02\x03\x04"
	)

	stripped, err := ValidatePKCS7(plaintext)
	assert.NoError(t, err)
	assert.Equal(t, expected, stripped)

	_, err = ValidatePKCS7(invalid1)
	assert.Error(t, err)

	_, err = ValidatePKCS7(invalid2)
	assert.Error(t, err)
}
