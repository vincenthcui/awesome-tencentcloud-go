package soe

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	TransmitOralProcessWithInit = tencentcloud.Action{
		Service: "soe",
		Version: "2018-07-24",
		Action:  "TransmitOralProcessWithInit",
	}
	TransmitOralProcess = tencentcloud.Action{
		Service: "soe",
		Version: "2018-07-24",
		Action:  "TransmitOralProcess",
	}
	KeywordEvaluate = tencentcloud.Action{
		Service: "soe",
		Version: "2018-07-24",
		Action:  "KeywordEvaluate",
	}
	InitOralProcess = tencentcloud.Action{
		Service: "soe",
		Version: "2018-07-24",
		Action:  "InitOralProcess",
	}
)
