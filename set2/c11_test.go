package set2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryptionOracle(t *testing.T) {
	var (
		plaintext = "Detect the block cipher mode the function is using each time. You should end up with a piece of code that, pointed at a block box that might be encrypting ECB or CBC, tells you which one is happening."
	)

	encryption := EncryptionOracle(plaintext)
	mode := DetectMode(encryption.ciphertext)

	assert.Equal(t, encryption.mode, mode)
}
