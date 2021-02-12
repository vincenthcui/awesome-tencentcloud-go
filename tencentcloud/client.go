package tencentcloud

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	"github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud/sign"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	defaultDomain        = "tencentcloudapi.com"
	defaultRequestClient = "awesome-tencentcloud-go"
	defaultMethod        = http.MethodPost
	defaultURI           = "/"
	defaultQuery         = ""
	schemaHttps          = "https"
	authorizeTpl         = "%s Credential=%s, SignedHeaders=%s, Signature=%s"

	dateLayout = "2006-01-02" // ref: package time

	headerHost          = "FullDomainName"
	headerContentType   = "Content-Type"
	headerAuthorization = "authorization"
	headerTCAction      = "X-TC-Action"
	headerTCVersion     = "X-TC-Version"
	headerTCTimestamp   = "X-TC-Timestamp"
	headerTCLanguage    = "X-TC-Language"
	headerTCRegion      = "X-TC-Region"
	headerRequestClient = "X-TC-RequestClient"

	contentTypeJson = "application/json"
	algorithmSHA256 = "TC3-HMAC-SHA256"
	scopeTC3Request = "tc3_request"
	languageZhCN    = "zh-CN"
)

// NewClient create Client with defaults, mutate client with Option
func NewClient(opts ...Option) *Client {
	cli := &Client{
		client:    &http.Client{},
		language:  languageZhCN,
		region:    regions.Guangzhou,
		algorithm: algorithmSHA256,

		httpMethod: defaultMethod,
		httpURI:    defaultURI,
		httpQuery:  defaultQuery,

		interceptors: []Interceptor{
			OnNetworkFailure,
			OnRateLimitExceeded,
		},
	}
	for idx := range opts {
		opts[idx](cli)
	}
	return cli
}

type Client struct {
	client *http.Client

	secretID  string
	secretKey string
	region    string
	language  string
	algorithm string

	httpMethod string
	httpURI    string
	httpQuery  string

	interceptors []Interceptor
}

// Send a common api request
func (c *Client) Send(ctx context.Context, action Action, request, response interface{}) error {
	injectClientToken(request)
	itv := time.Duration(0)
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(itv):
			err := c.send(action, request, response)
			for _, interceptor := range c.interceptors {
				err = interceptor(ctx, action, request, response, err)
			}
			if err, ok := err.(NeedRetry); ok {
				itv = err.In
				continue
			}
			return err
		}
	}
}

func (c *Client) send(action Action, request interface{}, response interface{}) error {
	body, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("tencentcloud: marshal request failed: %+v", err)
	}
	u := url.URL{Scheme: schemaHttps, Host: action.FullDomainName(defaultDomain), Path: defaultURI, RawQuery: defaultQuery}
	httpRequest, err := http.NewRequest(defaultMethod, u.String(), bytes.NewReader(body))
	if err != nil {
		return err
	}

	now := time.Now()
	headers := map[string]string{
		headerHost:        action.FullDomainName(defaultDomain),
		headerContentType: contentTypeJson,

		headerTCAction:      action.Action(),
		headerTCVersion:     action.Version(),
		headerTCTimestamp:   strconv.FormatInt(now.Unix(), 10),
		headerRequestClient: defaultRequestClient,
		headerTCLanguage:    c.language,
		headerTCRegion:      c.region,
	}
	headers[headerAuthorization] = c.authorize(action, headers, body, now)
	for k, v := range headers {
		httpRequest.Header[k] = []string{v}
	}

	httpResponse, err := c.client.Do(httpRequest)
	if err != nil {
		return err
	}
	defer httpResponse.Body.Close()

	byts, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return err
	}

	if err = maybeFailed(byts); err != nil {
		return err
	}

	return json.Unmarshal(byts, response)
}

func (c *Client) authorize(action Action, headers map[string]string, body []byte, now time.Time) string {
	date := now.UTC().Format(dateLayout)
	timestamp := strconv.FormatInt(now.Unix(), 10)
	scope := fmt.Sprintf("%s/%s/%s/%s", c.secretID, date, action.Service(), scopeTC3Request)
	signedHeaders, signedHeadersVal := sign.SignedHeaders(headers).PickOut(headerContentType, headerHost)

	payload := sign.SHA256Hex(body)
	payload = joinLines(c.httpMethod, c.httpURI, c.httpQuery, signedHeadersVal, signedHeaders, payload)
	payload = sign.SHA256Hex([]byte(payload))
	payload = joinLines(c.algorithm, timestamp, scope, payload)
	payload = sign.Sign(payload, c.secretKey, action.Service(), date)
	payload = fmt.Sprintf(authorizeTpl, c.algorithm, scope, signedHeaders, payload)
	return payload
}
