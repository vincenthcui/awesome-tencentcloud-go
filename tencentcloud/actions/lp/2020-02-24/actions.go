package lp

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	QueryLoginProtection = tencentcloud.Action{
		Service: "lp",
		Version: "2020-02-24",
		Action:  "QueryLoginProtection",
	}
)
