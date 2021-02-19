package tencentcloud

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
)

const (
	defaultMethod = http.MethodPost
	defaultURI    = "/"
	defaultQuery  = ""
	schemaHttps   = "https"
	authorizeTpl  = "%s Credential=%s/%s, SignedHeaders=%s, Signature=%s"

	clientVersion = "awesome-tencentcloud-go@v1"
	dateLayout    = "2006-01-02" // ref: package time

	headerHost            = "Host"
	headerContentType     = "Content-Type"
	headerAuthorization   = "Authorization"
	headerTCRequestClient = "X-TC-RequestClient"
	headerTCAction        = "X-TC-Action"
	headerTCVersion       = "X-TC-Version"
	headerTCTimestamp     = "X-TC-Timestamp"
	headerTCLanguage      = "X-TC-Language"
	headerTCRegion        = "X-TC-Region"

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
	client HttpClient

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
	u := url.URL{Scheme: schemaHttps, Host: domainName(action.Service()), Path: defaultURI, RawQuery: defaultQuery}
	httpRequest, err := http.NewRequest(defaultMethod, u.String(), bytes.NewReader(body))
	if err != nil {
		return err
	}

	now := time.Now()
	headers := map[string]string{
		headerHost:            domainName(action.Service()),
		headerContentType:     contentTypeJson,
		headerTCRequestClient: clientVersion,
		headerTCAction:        action.Action(),
		headerTCVersion:       action.Version(),
		headerTCTimestamp:     strconv.FormatInt(now.Unix(), 10),
		headerTCLanguage:      c.language,
		headerTCRegion:        c.region,
	}
	headers[headerAuthorization] = c.authorize(action, headers, body, now)
	httpRequest.Header = toHttpHeader(headers)

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
	scope := fmt.Sprintf("%s/%s/%s", date, action.Service(), scopeTC3Request)
	signedHeaders, signedHeadersVal := signHeaders(headers, headerContentType, headerHost)

	payload := sha256hex(body)
	payload = joinLines(c.httpMethod, c.httpURI, c.httpQuery, signedHeadersVal, signedHeaders, payload)
	payload = sha256hex([]byte(payload))
	payload = joinLines(c.algorithm, timestamp, scope, payload)
	payload = sign(payload, c.secretKey, action.Service(), date)
	payload = fmt.Sprintf(authorizeTpl, c.algorithm, c.secretID, scope, signedHeaders, payload)
	return payload
}
