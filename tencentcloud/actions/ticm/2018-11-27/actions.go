package ticm

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	VideoModeration = tencentcloud.Action{
		Service: "ticm",
		Version: "2018-11-27",
		Action:  "VideoModeration",
	}
	ImageModeration = tencentcloud.Action{
		Service: "ticm",
		Version: "2018-11-27",
		Action:  "ImageModeration",
	}
	DescribeVideoTask = tencentcloud.Action{
		Service: "ticm",
		Version: "2018-11-27",
		Action:  "DescribeVideoTask",
	}
)
