package vm

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	DescribeVideoStat = tencentcloud.Action{
		Service: "vm",
		Version: "2020-07-09",
		Action:  "DescribeVideoStat",
	}
	DescribeTaskDetail = tencentcloud.Action{
		Service: "vm",
		Version: "2020-07-09",
		Action:  "DescribeTaskDetail",
	}
	CreateVideoModerationTask = tencentcloud.Action{
		Service: "vm",
		Version: "2020-07-09",
		Action:  "CreateVideoModerationTask",
	}
	CreateBizConfig = tencentcloud.Action{
		Service: "vm",
		Version: "2020-07-09",
		Action:  "CreateBizConfig",
	}
	CancelTask = tencentcloud.Action{
		Service: "vm",
		Version: "2020-07-09",
		Action:  "CancelTask",
	}
)
