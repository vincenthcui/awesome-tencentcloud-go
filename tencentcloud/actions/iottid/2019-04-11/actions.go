package iottid

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	VerifyChipBurnInfo = tencentcloud.Action{
		Service: "iottid",
		Version: "2019-04-11",
		Action:  "VerifyChipBurnInfo",
	}
	UploadDeviceUniqueCode = tencentcloud.Action{
		Service: "iottid",
		Version: "2019-04-11",
		Action:  "UploadDeviceUniqueCode",
	}
	DownloadTids = tencentcloud.Action{
		Service: "iottid",
		Version: "2019-04-11",
		Action:  "DownloadTids",
	}
	DescribePermission = tencentcloud.Action{
		Service: "iottid",
		Version: "2019-04-11",
		Action:  "DescribePermission",
	}
	DescribeAvailableLibCount = tencentcloud.Action{
		Service: "iottid",
		Version: "2019-04-11",
		Action:  "DescribeAvailableLibCount",
	}
	DeliverTids = tencentcloud.Action{
		Service: "iottid",
		Version: "2019-04-11",
		Action:  "DeliverTids",
	}
	DeliverTidNotify = tencentcloud.Action{
		Service: "iottid",
		Version: "2019-04-11",
		Action:  "DeliverTidNotify",
	}
	BurnTidNotify = tencentcloud.Action{
		Service: "iottid",
		Version: "2019-04-11",
		Action:  "BurnTidNotify",
	}
	AuthTestTid = tencentcloud.Action{
		Service: "iottid",
		Version: "2019-04-11",
		Action:  "AuthTestTid",
	}
)
