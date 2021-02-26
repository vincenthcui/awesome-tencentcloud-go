package apcas

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	UploadId = tencentcloud.Action{
		Service: "apcas",
		Version: "2020-11-27",
		Action:  "UploadId",
	}
	QueryGeneralStat = tencentcloud.Action{
		Service: "apcas",
		Version: "2020-11-27",
		Action:  "QueryGeneralStat",
	}
	QueryCallStat = tencentcloud.Action{
		Service: "apcas",
		Version: "2020-11-27",
		Action:  "QueryCallStat",
	}
	QueryCallDetails = tencentcloud.Action{
		Service: "apcas",
		Version: "2020-11-27",
		Action:  "QueryCallDetails",
	}
	PredictRating = tencentcloud.Action{
		Service: "apcas",
		Version: "2020-11-27",
		Action:  "PredictRating",
	}
	GetTaskList = tencentcloud.Action{
		Service: "apcas",
		Version: "2020-11-27",
		Action:  "GetTaskList",
	}
	GetTaskDetail = tencentcloud.Action{
		Service: "apcas",
		Version: "2020-11-27",
		Action:  "GetTaskDetail",
	}
)
