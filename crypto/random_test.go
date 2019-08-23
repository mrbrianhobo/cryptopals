package crypto

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestRandom(t *testing.T) {
	RegisterTestingT(t)

	t.Log(RandomInt(0, 20))
	t.Log(RandomBytes(5))
	t.Log(RandomByte())
	t.Log(RandomKey())
}
