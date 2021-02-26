package tcex

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	InvokeService = tencentcloud.Action{
		Service: "tcex",
		Version: "2020-07-27",
		Action:  "InvokeService",
	}
	DescribeInvocationResult = tencentcloud.Action{
		Service: "tcex",
		Version: "2020-07-27",
		Action:  "DescribeInvocationResult",
	}
)
