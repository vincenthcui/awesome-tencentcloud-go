package cms

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	TextModeration = tencentcloud.Action{
		Service: "cms",
		Version: "2019-03-21",
		Action:  "TextModeration",
	}
	ManualReview = tencentcloud.Action{
		Service: "cms",
		Version: "2019-03-21",
		Action:  "ManualReview",
	}
	ImageModeration = tencentcloud.Action{
		Service: "cms",
		Version: "2019-03-21",
		Action:  "ImageModeration",
	}
	DescribeTextSample = tencentcloud.Action{
		Service: "cms",
		Version: "2019-03-21",
		Action:  "DescribeTextSample",
	}
	DescribeFileSample = tencentcloud.Action{
		Service: "cms",
		Version: "2019-03-21",
		Action:  "DescribeFileSample",
	}
	DeleteTextSample = tencentcloud.Action{
		Service: "cms",
		Version: "2019-03-21",
		Action:  "DeleteTextSample",
	}
	DeleteFileSample = tencentcloud.Action{
		Service: "cms",
		Version: "2019-03-21",
		Action:  "DeleteFileSample",
	}
	CreateTextSample = tencentcloud.Action{
		Service: "cms",
		Version: "2019-03-21",
		Action:  "CreateTextSample",
	}
	CreateFileSample = tencentcloud.Action{
		Service: "cms",
		Version: "2019-03-21",
		Action:  "CreateFileSample",
	}
)
