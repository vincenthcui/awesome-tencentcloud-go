package af

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	QueryAntiFraud = tencentcloud.Action{
		Service: "af",
		Version: "2020-02-26",
		Action:  "QueryAntiFraud",
	}
)
