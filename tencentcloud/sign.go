package tencentcloud

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

func sha256hex(plain []byte) string {
	b := sha256.Sum256(plain)
	return hex.EncodeToString(b[:])
}

const (
	version3   = "TC3"
	tc3Request = "tc3_request"
)

func sign(plain, key, service, date string) string {
	secret := SHA256([]byte(date), []byte(version3+key))
	secret = SHA256([]byte(service), secret)
	secret = SHA256([]byte(tc3Request), secret)
	return hex.EncodeToString(SHA256([]byte(plain), secret))
}
