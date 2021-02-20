package tencentcloud

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func sha256crypt(plain, key []byte) []byte {
	algor := hmac.New(sha256.New, key)
	algor.Write(plain) // nolint:errcheck
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
	secret := sha256crypt([]byte(date), []byte(version3+key))
	secret = sha256crypt([]byte(service), secret)
	secret = sha256crypt([]byte(tc3Request), secret)
	return hex.EncodeToString(sha256crypt([]byte(plain), secret))
}
