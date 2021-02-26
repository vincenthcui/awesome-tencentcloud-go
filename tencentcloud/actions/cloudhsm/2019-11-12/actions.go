package cloudhsm

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	ModifyVsmAttributes = tencentcloud.Action{
		Service: "cloudhsm",
		Version: "2019-11-12",
		Action:  "ModifyVsmAttributes",
	}
	InquiryPriceBuyVsm = tencentcloud.Action{
		Service: "cloudhsm",
		Version: "2019-11-12",
		Action:  "InquiryPriceBuyVsm",
	}
	DescribeVsms = tencentcloud.Action{
		Service: "cloudhsm",
		Version: "2019-11-12",
		Action:  "DescribeVsms",
	}
	DescribeVsmAttributes = tencentcloud.Action{
		Service: "cloudhsm",
		Version: "2019-11-12",
		Action:  "DescribeVsmAttributes",
	}
	DescribeVpc = tencentcloud.Action{
		Service: "cloudhsm",
		Version: "2019-11-12",
		Action:  "DescribeVpc",
	}
	DescribeUsgRule = tencentcloud.Action{
		Service: "cloudhsm",
		Version: "2019-11-12",
		Action:  "DescribeUsgRule",
	}
	DescribeUsg = tencentcloud.Action{
		Service: "cloudhsm",
		Version: "2019-11-12",
		Action:  "DescribeUsg",
	}
	DescribeSubnet = tencentcloud.Action{
		Service: "cloudhsm",
		Version: "2019-11-12",
		Action:  "DescribeSubnet",
	}
	DescribeHSMByVpcId = tencentcloud.Action{
		Service: "cloudhsm",
		Version: "2019-11-12",
		Action:  "DescribeHSMByVpcId",
	}
	DescribeHSMBySubnetId = tencentcloud.Action{
		Service: "cloudhsm",
		Version: "2019-11-12",
		Action:  "DescribeHSMBySubnetId",
	}
)
