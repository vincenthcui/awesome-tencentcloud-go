package emr

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	TerminateTasks = tencentcloud.Action{
		Service: "emr",
		Version: "2019-01-03",
		Action:  "TerminateTasks",
	}
	TerminateInstance = tencentcloud.Action{
		Service: "emr",
		Version: "2019-01-03",
		Action:  "TerminateInstance",
	}
	ScaleOutInstance = tencentcloud.Action{
		Service: "emr",
		Version: "2019-01-03",
		Action:  "ScaleOutInstance",
	}
	RunJobFlow = tencentcloud.Action{
		Service: "emr",
		Version: "2019-01-03",
		Action:  "RunJobFlow",
	}
	InquiryPriceUpdateInstance = tencentcloud.Action{
		Service: "emr",
		Version: "2019-01-03",
		Action:  "InquiryPriceUpdateInstance",
	}
	InquiryPriceScaleOutInstance = tencentcloud.Action{
		Service: "emr",
		Version: "2019-01-03",
		Action:  "InquiryPriceScaleOutInstance",
	}
	InquiryPriceRenewInstance = tencentcloud.Action{
		Service: "emr",
		Version: "2019-01-03",
		Action:  "InquiryPriceRenewInstance",
	}
	InquiryPriceCreateInstance = tencentcloud.Action{
		Service: "emr",
		Version: "2019-01-03",
		Action:  "InquiryPriceCreateInstance",
	}
	DescribeJobFlow = tencentcloud.Action{
		Service: "emr",
		Version: "2019-01-03",
		Action:  "DescribeJobFlow",
	}
	DescribeInstances = tencentcloud.Action{
		Service: "emr",
		Version: "2019-01-03",
		Action:  "DescribeInstances",
	}
	DescribeClusterNodes = tencentcloud.Action{
		Service: "emr",
		Version: "2019-01-03",
		Action:  "DescribeClusterNodes",
	}
	CreateInstance = tencentcloud.Action{
		Service: "emr",
		Version: "2019-01-03",
		Action:  "CreateInstance",
	}
)
