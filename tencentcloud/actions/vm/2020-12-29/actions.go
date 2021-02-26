package vm

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	DescribeTasks = tencentcloud.Action{
		Service: "vm",
		Version: "2020-12-29",
		Action:  "DescribeTasks",
	}
	DescribeTaskDetail = tencentcloud.Action{
		Service: "vm",
		Version: "2020-12-29",
		Action:  "DescribeTaskDetail",
	}
	CreateVideoModerationTask = tencentcloud.Action{
		Service: "vm",
		Version: "2020-12-29",
		Action:  "CreateVideoModerationTask",
	}
	CancelTask = tencentcloud.Action{
		Service: "vm",
		Version: "2020-12-29",
		Action:  "CancelTask",
	}
)
