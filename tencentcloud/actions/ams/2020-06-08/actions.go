package ams

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	DescribeTaskDetail = tencentcloud.Action{
		Service: "ams",
		Version: "2020-06-08",
		Action:  "DescribeTaskDetail",
	}
	DescribeBizConfig = tencentcloud.Action{
		Service: "ams",
		Version: "2020-06-08",
		Action:  "DescribeBizConfig",
	}
	DescribeAudioStat = tencentcloud.Action{
		Service: "ams",
		Version: "2020-06-08",
		Action:  "DescribeAudioStat",
	}
	DescribeAmsList = tencentcloud.Action{
		Service: "ams",
		Version: "2020-06-08",
		Action:  "DescribeAmsList",
	}
	CreateBizConfig = tencentcloud.Action{
		Service: "ams",
		Version: "2020-06-08",
		Action:  "CreateBizConfig",
	}
	CreateAudioModerationTask = tencentcloud.Action{
		Service: "ams",
		Version: "2020-06-08",
		Action:  "CreateAudioModerationTask",
	}
	CancelTask = tencentcloud.Action{
		Service: "ams",
		Version: "2020-06-08",
		Action:  "CancelTask",
	}
)
