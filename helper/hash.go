package helper

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

func Hash(pin string) string {
	hash := sha256.New()
	hash.Write([]byte(pin))
	return hex.EncodeToString(hash.Sum(nil))
}

func GenerateMD5Hash(merchantCode, merchantOrderId, paymentAmount, apiKey string) string {
	// Gabungkan string
	data := merchantCode + merchantOrderId + paymentAmount + apiKey

	// Buat hash MD5
	hash := md5.Sum([]byte(data))

	// Encode hash ke dalam format heksadesimal
	hashString := hex.EncodeToString(hash[:])

	return hashString
}
