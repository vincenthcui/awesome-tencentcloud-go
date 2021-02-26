package ams

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	DescribeTasks = tencentcloud.Action{
		Service: "ams",
		Version: "2020-12-29",
		Action:  "DescribeTasks",
	}
	DescribeTaskDetail = tencentcloud.Action{
		Service: "ams",
		Version: "2020-12-29",
		Action:  "DescribeTaskDetail",
	}
	CreateAudioModerationTask = tencentcloud.Action{
		Service: "ams",
		Version: "2020-12-29",
		Action:  "CreateAudioModerationTask",
	}
	CancelTask = tencentcloud.Action{
		Service: "ams",
		Version: "2020-12-29",
		Action:  "CancelTask",
	}
)
