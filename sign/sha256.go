package sign

import (
	"crypto/hmac"
	"crypto/sha256"
)

func SHA256(plain, key []byte) []byte {
	algor := hmac.New(sha256.New, key)
	algor.Write(plain)
	return algor.Sum(nil)
}
