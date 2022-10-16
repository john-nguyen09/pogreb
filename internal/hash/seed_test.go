package hash

import (
	"testing"

	"github.com/john-nguyen09/pogreb/internal/assert"
)

func TestRandSeed(t *testing.T) {
	_, err := RandSeed()
	assert.Nil(t, err)
}
