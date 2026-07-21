package uuid

import (
	"bytes"
	cryptorand "crypto/rand"
	"encoding/hex"
	"strings"
	"testing"

	"github.com/nobl9/govy/internal/assert"
)

func TestGenerateUUID(t *testing.T) {
	randomBytes, err := hex.DecodeString("919108f752d133205bacf847db4148a8")
	assert.Require(t, assert.NoError(t, err))
	originalReader := cryptorand.Reader
	cryptorand.Reader = bytes.NewReader(randomBytes)
	t.Cleanup(func() { cryptorand.Reader = originalReader })

	id := GenerateUUID()
	assert.Equal(t, "919108f7-52d1-4320-9bac-f847db4148a8", id)

	assert.Require(t, assert.Len(t, id, 36))
	assert.Equal(t, byte('-'), id[8])
	assert.Equal(t, byte('-'), id[13])
	assert.Equal(t, byte('-'), id[18])
	assert.Equal(t, byte('-'), id[23])
	assert.Equal(t, strings.ToLower(id), id)

	decoded, err := hex.DecodeString(strings.ReplaceAll(id, "-", ""))
	assert.Require(t, assert.NoError(t, err))
	assert.Require(t, assert.Len(t, decoded, 16))
	assert.Equal(t, byte(4), decoded[6]>>4)
	assert.Equal(t, byte(2), decoded[8]>>6)
}
