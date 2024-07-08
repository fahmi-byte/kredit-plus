package helper

import (
	"crypto/rand"
	"encoding/base64"
	"io"
)

func GenerateAPIKey(length int) (string, error) {
	key := make([]byte, length)

	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(key), nil
}
