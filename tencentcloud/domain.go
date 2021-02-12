package tencentcloud

import "fmt"

const (
	defaultDomain = "tencentcloudapi.com"
)

func domainName(service string) string {
	return fmt.Sprintf("%s.%s", service, defaultDomain)
}
