package helper

import (
	"crypto/sha256"
	"encoding/hex"
)

func Hash(pin string) string {
	hash := sha256.New()
	hash.Write([]byte(pin))
	return hex.EncodeToString(hash.Sum(nil))
}
