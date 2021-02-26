package tbp

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	TextReset = tencentcloud.Action{
		Service: "tbp",
		Version: "2019-03-11",
		Action:  "TextReset",
	}
	TextProcess = tencentcloud.Action{
		Service: "tbp",
		Version: "2019-03-11",
		Action:  "TextProcess",
	}
	Reset = tencentcloud.Action{
		Service: "tbp",
		Version: "2019-03-11",
		Action:  "Reset",
	}
	CreateBot = tencentcloud.Action{
		Service: "tbp",
		Version: "2019-03-11",
		Action:  "CreateBot",
	}
)
