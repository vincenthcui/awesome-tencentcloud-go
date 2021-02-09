package sign

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func SHA256(plain, key []byte) []byte {
	algor := hmac.New(sha256.New, key)
	algor.Write(plain)
	return algor.Sum(nil)
}

func SHA256Hex(plain []byte) string {
	b := sha256.Sum256(plain)
	return hex.EncodeToString(b[:])
}
