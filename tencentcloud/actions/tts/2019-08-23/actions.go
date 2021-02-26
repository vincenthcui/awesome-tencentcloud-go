package tts

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	TextToVoice = tencentcloud.Action{
		Service: "tts",
		Version: "2019-08-23",
		Action:  "TextToVoice",
	}
)
