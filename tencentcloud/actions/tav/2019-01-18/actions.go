package tav

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	ScanFileHash = tencentcloud.Action{
		Service: "tav",
		Version: "2019-01-18",
		Action:  "ScanFileHash",
	}
	ScanFile = tencentcloud.Action{
		Service: "tav",
		Version: "2019-01-18",
		Action:  "ScanFile",
	}
	GetScanResult = tencentcloud.Action{
		Service: "tav",
		Version: "2019-01-18",
		Action:  "GetScanResult",
	}
	GetLocalEngine = tencentcloud.Action{
		Service: "tav",
		Version: "2019-01-18",
		Action:  "GetLocalEngine",
	}
)
