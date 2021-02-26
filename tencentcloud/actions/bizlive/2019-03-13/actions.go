package bizlive

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	StopGame = tencentcloud.Action{
		Service: "bizlive",
		Version: "2019-03-13",
		Action:  "StopGame",
	}
	RegisterIM = tencentcloud.Action{
		Service: "bizlive",
		Version: "2019-03-13",
		Action:  "RegisterIM",
	}
	ForbidLiveStream = tencentcloud.Action{
		Service: "bizlive",
		Version: "2019-03-13",
		Action:  "ForbidLiveStream",
	}
	DescribeWorkers = tencentcloud.Action{
		Service: "bizlive",
		Version: "2019-03-13",
		Action:  "DescribeWorkers",
	}
	DescribeStreamPlayInfoList = tencentcloud.Action{
		Service: "bizlive",
		Version: "2019-03-13",
		Action:  "DescribeStreamPlayInfoList",
	}
	CreateSession = tencentcloud.Action{
		Service: "bizlive",
		Version: "2019-03-13",
		Action:  "CreateSession",
	}
)
