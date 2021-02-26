package ic

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	SendSms = tencentcloud.Action{
		Service: "ic",
		Version: "2019-03-07",
		Action:  "SendSms",
	}
	SendMultiSms = tencentcloud.Action{
		Service: "ic",
		Version: "2019-03-07",
		Action:  "SendMultiSms",
	}
	RenewCards = tencentcloud.Action{
		Service: "ic",
		Version: "2019-03-07",
		Action:  "RenewCards",
	}
	DescribeCards = tencentcloud.Action{
		Service: "ic",
		Version: "2019-03-07",
		Action:  "DescribeCards",
	}
	DescribeCard = tencentcloud.Action{
		Service: "ic",
		Version: "2019-03-07",
		Action:  "DescribeCard",
	}
	DescribeApp = tencentcloud.Action{
		Service: "ic",
		Version: "2019-03-07",
		Action:  "DescribeApp",
	}
)
