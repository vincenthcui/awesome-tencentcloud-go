package tbp

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	TextReset = tencentcloud.Action{
		Service: "tbp",
		Version: "2019-06-27",
		Action:  "TextReset",
	}
	TextProcess = tencentcloud.Action{
		Service: "tbp",
		Version: "2019-06-27",
		Action:  "TextProcess",
	}
)
