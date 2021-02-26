package habo

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	StartAnalyse = tencentcloud.Action{
		Service: "habo",
		Version: "2018-12-03",
		Action:  "StartAnalyse",
	}
	DescribeStatus = tencentcloud.Action{
		Service: "habo",
		Version: "2018-12-03",
		Action:  "DescribeStatus",
	}
)
