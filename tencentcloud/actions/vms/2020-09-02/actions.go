package vms

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	SendTtsVoice = tencentcloud.Action{
		Service: "vms",
		Version: "2020-09-02",
		Action:  "SendTtsVoice",
	}
	SendCodeVoice = tencentcloud.Action{
		Service: "vms",
		Version: "2020-09-02",
		Action:  "SendCodeVoice",
	}
)
