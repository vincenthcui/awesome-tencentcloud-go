package ims

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	ImageRecognition = tencentcloud.Action{
		Service: "ims",
		Version: "2020-12-29",
		Action:  "ImageRecognition",
	}
	ImageModeration = tencentcloud.Action{
		Service: "ims",
		Version: "2020-12-29",
		Action:  "ImageModeration",
	}
)
