package awesome_tencentcloud_go

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/vincenthcui/awesome-tencentcloud-go/actions"
	"github.com/vincenthcui/awesome-tencentcloud-go/sign"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	defaultDomain        = "tencentcloudapi.com"
	defaultRequestClient = "awesome-tencentcloud-go"
	defaultMethod        = http.MethodPost
	defaultURI           = "/"
	defaultQuery         = ""
	schemaHttps          = "https"

	lineSep    = "\n"
	dateLayout = "2006-01-02" // ref: package time

	headerHost          = "Host"
	headerContentType   = "Content-Type"
	headerTCAction      = "X-TC-Action"
	headerTCVersion     = "X-TC-Version"
	headerTCTimestamp   = "X-TC-Timestamp"
	headerTCLanguage    = "X-TC-Language"
	headerTCRegion      = "X-TC-Region"
	headerRequestClient = "X-TC-RequestClient"

	contentTypeJson = "application/json"
	algorithmSHA256 = "TC3-HMAC-SHA256"
	scopeTC3Request = "tc3_request"
)

type client struct {
	client *http.Client

	region   string
	language string
}

func (c *client) Send(ctx context.Context, action actions.Action, request interface{}, response interface{}) error {
	now := time.Now()
	timestamp := "TODO"
	headers := map[string]string{
		headerHost:        action.Host(defaultDomain),
		headerContentType: contentTypeJson,

		headerTCAction:      action.Action(),
		headerTCVersion:     action.Version(),
		headerTCTimestamp:   timestamp,
		headerRequestClient: defaultRequestClient,
		headerTCLanguage:    c.language,
		headerTCRegion:      c.region,
	}
	//if c.credential.Token != "" {
	//	headers["X-TC-Token"] = c.credential.Token
	//}
	signedHeaderFields, signedHeaders := sign.SignedHeaders(headers).PickOut(headerContentType, headerHost)

	requestPayload, err := json.Marshal(request)
	if err != nil {
		return err
	}
	hashedRequestPayload := sign.SHA256Hex(requestPayload)

	canonicalRequest := strings.Join([]string{
		defaultMethod,
		defaultURI,
		defaultQuery,
		signedHeaders,
		signedHeaderFields,
		hashedRequestPayload,
	}, lineSep)
	hashedCanonicalRequest := sign.SHA256Hex([]byte(canonicalRequest))

	date := now.UTC().Format(dateLayout)
	scope := fmt.Sprintf("%s/%s/%s", date, action.Service(), scopeTC3Request)

	string2sign := strings.Join([]string{algorithmSHA256, timestamp, scope, hashedCanonicalRequest}, lineSep)
	signature := sign.Sign(string2sign, "SecretKey", action.Service(), date)
	authorization := sign.Authorize(algorithmSHA256, "SecretID", scopeTC3Request, signedHeaderFields, signature)
	headers["Authorization"] = authorization

	u := url.URL{
		Scheme:   schemaHttps,
		Host:     action.Host(defaultDomain),
		Path:     defaultURI,
		RawQuery: defaultQuery,
	}
	httpRequest, err := http.NewRequest(defaultMethod, u.String(), bytes.NewReader(requestPayload))
	if err != nil {
		return err
	}

	for k, v := range headers {
		httpRequest.Header[k] = []string{v}
	}

	httpResponse, err := c.client.Do(httpRequest)
	if err != nil {
		return err
	}
	byts, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(byts, response)
	//if err != nil {
	//	return errors.NewTencentCloudSDKError("ClientError.NetworkError", msg, "")
	//}
	//err = tchttp.ParseFromHttpResponse(httpResponse, response)
}
