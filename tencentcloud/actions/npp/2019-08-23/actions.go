package npp

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	GetVirtualNum = tencentcloud.Action{
		Service: "npp",
		Version: "2019-08-23",
		Action:  "GetVirtualNum",
	}
	Get400Cdr = tencentcloud.Action{
		Service: "npp",
		Version: "2019-08-23",
		Action:  "Get400Cdr",
	}
	DescribeCallerDisplayList = tencentcloud.Action{
		Service: "npp",
		Version: "2019-08-23",
		Action:  "DescribeCallerDisplayList",
	}
	DescribeCallBackStatus = tencentcloud.Action{
		Service: "npp",
		Version: "2019-08-23",
		Action:  "DescribeCallBackStatus",
	}
	DescribeCallBackCdr = tencentcloud.Action{
		Service: "npp",
		Version: "2019-08-23",
		Action:  "DescribeCallBackCdr",
	}
	DeleteCallBack = tencentcloud.Action{
		Service: "npp",
		Version: "2019-08-23",
		Action:  "DeleteCallBack",
	}
	DelVirtualNum = tencentcloud.Action{
		Service: "npp",
		Version: "2019-08-23",
		Action:  "DelVirtualNum",
	}
	CreateCallBack = tencentcloud.Action{
		Service: "npp",
		Version: "2019-08-23",
		Action:  "CreateCallBack",
	}
)
