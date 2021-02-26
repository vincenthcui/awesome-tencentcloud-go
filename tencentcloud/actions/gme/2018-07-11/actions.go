package gme

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	VoiceFilter = tencentcloud.Action{
		Service: "gme",
		Version: "2018-07-11",
		Action:  "VoiceFilter",
	}
	ScanVoice = tencentcloud.Action{
		Service: "gme",
		Version: "2018-07-11",
		Action:  "ScanVoice",
	}
	ModifyAppStatus = tencentcloud.Action{
		Service: "gme",
		Version: "2018-07-11",
		Action:  "ModifyAppStatus",
	}
	DescribeUserInAndOutTime = tencentcloud.Action{
		Service: "gme",
		Version: "2018-07-11",
		Action:  "DescribeUserInAndOutTime",
	}
	DescribeScanResultList = tencentcloud.Action{
		Service: "gme",
		Version: "2018-07-11",
		Action:  "DescribeScanResultList",
	}
	DescribeFilterResultList = tencentcloud.Action{
		Service: "gme",
		Version: "2018-07-11",
		Action:  "DescribeFilterResultList",
	}
	DescribeFilterResult = tencentcloud.Action{
		Service: "gme",
		Version: "2018-07-11",
		Action:  "DescribeFilterResult",
	}
	DescribeApplicationData = tencentcloud.Action{
		Service: "gme",
		Version: "2018-07-11",
		Action:  "DescribeApplicationData",
	}
	DescribeAppStatistics = tencentcloud.Action{
		Service: "gme",
		Version: "2018-07-11",
		Action:  "DescribeAppStatistics",
	}
	CreateApp = tencentcloud.Action{
		Service: "gme",
		Version: "2018-07-11",
		Action:  "CreateApp",
	}
)
