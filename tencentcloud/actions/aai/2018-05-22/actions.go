package aai

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	TextToVoice = tencentcloud.Action{
		Service: "aai",
		Version: "2018-05-22",
		Action:  "TextToVoice",
	}
	SimultaneousInterpreting = tencentcloud.Action{
		Service: "aai",
		Version: "2018-05-22",
		Action:  "SimultaneousInterpreting",
	}
	SentenceRecognition = tencentcloud.Action{
		Service: "aai",
		Version: "2018-05-22",
		Action:  "SentenceRecognition",
	}
	Chat = tencentcloud.Action{
		Service: "aai",
		Version: "2018-05-22",
		Action:  "Chat",
	}
)
