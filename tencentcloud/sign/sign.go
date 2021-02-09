package sign

import "encoding/hex"

const (
	version3   = "TC3"
	tc3Request = "tc3_request"
)

func Sign(plain, key, service, date string) string {
	secret := SHA256([]byte(date), []byte(version3+key))
	secret = SHA256([]byte(service), secret)
	secret = SHA256([]byte(tc3Request), secret)
	return hex.EncodeToString(SHA256([]byte(plain), secret))
}
