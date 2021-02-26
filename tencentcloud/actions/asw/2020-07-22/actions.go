package asw

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	StartExecution = tencentcloud.Action{
		Service: "asw",
		Version: "2020-07-22",
		Action:  "StartExecution",
	}
	ModifyFlowService = tencentcloud.Action{
		Service: "asw",
		Version: "2020-07-22",
		Action:  "ModifyFlowService",
	}
	DescribeFlowServices = tencentcloud.Action{
		Service: "asw",
		Version: "2020-07-22",
		Action:  "DescribeFlowServices",
	}
	DescribeFlowServiceDetail = tencentcloud.Action{
		Service: "asw",
		Version: "2020-07-22",
		Action:  "DescribeFlowServiceDetail",
	}
	DescribeExecutions = tencentcloud.Action{
		Service: "asw",
		Version: "2020-07-22",
		Action:  "DescribeExecutions",
	}
	DescribeExecution = tencentcloud.Action{
		Service: "asw",
		Version: "2020-07-22",
		Action:  "DescribeExecution",
	}
	CreateFlowService = tencentcloud.Action{
		Service: "asw",
		Version: "2020-07-22",
		Action:  "CreateFlowService",
	}
)
