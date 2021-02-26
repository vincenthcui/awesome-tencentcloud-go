package tms

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	TextModeration = tencentcloud.Action{
		Service: "tms",
		Version: "2020-07-13",
		Action:  "TextModeration",
	}
	DescribeTextStat = tencentcloud.Action{
		Service: "tms",
		Version: "2020-07-13",
		Action:  "DescribeTextStat",
	}
	DescribeTextLib = tencentcloud.Action{
		Service: "tms",
		Version: "2020-07-13",
		Action:  "DescribeTextLib",
	}
	AccountTipoffAccess = tencentcloud.Action{
		Service: "tms",
		Version: "2020-07-13",
		Action:  "AccountTipoffAccess",
	}
)
