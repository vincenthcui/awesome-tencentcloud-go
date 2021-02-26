package ims

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	ImageModeration = tencentcloud.Action{
		Service: "ims",
		Version: "2020-07-13",
		Action:  "ImageModeration",
	}
	DescribeImsList = tencentcloud.Action{
		Service: "ims",
		Version: "2020-07-13",
		Action:  "DescribeImsList",
	}
	DescribeImageStat = tencentcloud.Action{
		Service: "ims",
		Version: "2020-07-13",
		Action:  "DescribeImageStat",
	}
)
