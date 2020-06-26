package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// GetHash returns sha256 hash of provided string
func GetHash(v string) string {
	h := sha256.New()
	h.Write([]byte(v))
	return hex.EncodeToString(h.Sum(nil))
}
