package gs

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	TrylockWorker = tencentcloud.Action{
		Service: "gs",
		Version: "2019-11-18",
		Action:  "TrylockWorker",
	}
	SwitchGameArchive = tencentcloud.Action{
		Service: "gs",
		Version: "2019-11-18",
		Action:  "SwitchGameArchive",
	}
	StopGame = tencentcloud.Action{
		Service: "gs",
		Version: "2019-11-18",
		Action:  "StopGame",
	}
	SaveGameArchive = tencentcloud.Action{
		Service: "gs",
		Version: "2019-11-18",
		Action:  "SaveGameArchive",
	}
	DescribeInstancesCount = tencentcloud.Action{
		Service: "gs",
		Version: "2019-11-18",
		Action:  "DescribeInstancesCount",
	}
	CreateSession = tencentcloud.Action{
		Service: "gs",
		Version: "2019-11-18",
		Action:  "CreateSession",
	}
)
