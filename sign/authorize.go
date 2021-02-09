package sign

import "fmt"

const (
	authorizeFormat = "%s Credential=%s/%s, SignedHeaders=%s, Signature=%s"
)

func Authorize(algorithm, secretID, scope, signedHeaders, signature string) string {
	return fmt.Sprintf(authorizeFormat, algorithm, secretID, scope, signedHeaders, signature)
}
