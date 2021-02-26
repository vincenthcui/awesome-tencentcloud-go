package tkgdq

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	DescribeTriple = tencentcloud.Action{
		Service: "tkgdq",
		Version: "2019-04-11",
		Action:  "DescribeTriple",
	}
	DescribeRelation = tencentcloud.Action{
		Service: "tkgdq",
		Version: "2019-04-11",
		Action:  "DescribeRelation",
	}
	DescribeEntity = tencentcloud.Action{
		Service: "tkgdq",
		Version: "2019-04-11",
		Action:  "DescribeEntity",
	}
)
