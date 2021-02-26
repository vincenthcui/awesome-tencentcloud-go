package rkp

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	QueryDevAndRisk = tencentcloud.Action{
		Service: "rkp",
		Version: "2019-12-09",
		Action:  "QueryDevAndRisk",
	}
	GetToken = tencentcloud.Action{
		Service: "rkp",
		Version: "2019-12-09",
		Action:  "GetToken",
	}
	GetOpenId = tencentcloud.Action{
		Service: "rkp",
		Version: "2019-12-09",
		Action:  "GetOpenId",
	}
)
