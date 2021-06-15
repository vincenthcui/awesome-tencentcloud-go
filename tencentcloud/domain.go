package tencentcloud

import "fmt"

const (
	defaultBasicDomain = "tencentcloudapi.com"
)

func (c Client) domain(service string) string {
	return fmt.Sprintf("%s.%s", service, c.basicDomain)
}
