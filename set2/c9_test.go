package set2

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestPKCS7Padding(t *testing.T) {
	RegisterTestingT(t)

	const (
		plaintext = "YELLOW SUBMARINE"
		expected  = "YELLOW SUBMARINE\x04\x04\x04\x04"
	)

	actual := PKCS7Padding(plaintext, 20)

	Expect(actual).To(Equal(expected))
}
