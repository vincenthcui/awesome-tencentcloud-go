package tencentcloud

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud/actions"
	"github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud/sign"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
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

type ClientOpts struct {
	Region    string
	Language  string
	SecretID  string
	SecretKey string
}

func NewClient(opts ClientOpts) Client {
	if opts.Language == "" {
		opts.Language = "zh-CN"
	}
	if opts.Region == "" {
		opts.Region = "ap-guangzhou"
	}

	return &client{secretID: opts.SecretID, secretKey: opts.SecretKey, region: opts.Region, language: opts.Language, client: &http.Client{}}
}

type Client interface {
	Send(ctx context.Context, action actions.Action, request, response interface{}) error
}

type client struct {
	client *http.Client

	secretID  string
	secretKey string
	region    string
	language  string
}

func (c *client) Send(ctx context.Context, action actions.Action, request interface{}, response interface{}) error {
	now := time.Now()
	timestamp := strconv.FormatInt(now.Unix(), 10)
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
	signature := sign.Sign(string2sign, c.secretKey, action.Service(), date)
	authorization := sign.Authorize(algorithmSHA256, c.secretID, scope, signedHeaderFields, signature)
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

	output, _ := httputil.DumpRequest(httpRequest, true)
	fmt.Println(string(output))

	httpResponse, err := c.client.Do(httpRequest)
	if err != nil {
		return err
	}
	defer httpResponse.Body.Close()

	output, _ = httputil.DumpResponse(httpResponse, true)
	fmt.Println(string(output))
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
