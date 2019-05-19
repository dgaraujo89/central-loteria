package commons

import (
	"crypto"
	"encoding/hex"
)

// Encrypt a string
func Encrypt(data string) string {
	hash := crypto.SHA512.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}
