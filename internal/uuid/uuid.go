package uuid

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateUUID() string {
	var randomBytes [16]byte
	_, _ = rand.Read(randomBytes[:])

	// Set version (4) and variant bits
	randomBytes[6] = (randomBytes[6] & 0x0f) | 0x40
	randomBytes[8] = (randomBytes[8] & 0x3f) | 0x80

	var id [36]byte
	hex.Encode(id[0:8], randomBytes[0:4])
	id[8] = '-'
	hex.Encode(id[9:13], randomBytes[4:6])
	id[13] = '-'
	hex.Encode(id[14:18], randomBytes[6:8])
	id[18] = '-'
	hex.Encode(id[19:23], randomBytes[8:10])
	id[23] = '-'
	hex.Encode(id[24:36], randomBytes[10:16])
	return string(id[:])
}
