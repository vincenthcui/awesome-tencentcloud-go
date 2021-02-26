package ecc

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	EHOCR = tencentcloud.Action{
		Service: "ecc",
		Version: "2018-12-13",
		Action:  "EHOCR",
	}
	ECC = tencentcloud.Action{
		Service: "ecc",
		Version: "2018-12-13",
		Action:  "ECC",
	}
	DescribeTask = tencentcloud.Action{
		Service: "ecc",
		Version: "2018-12-13",
		Action:  "DescribeTask",
	}
	CorrectMultiImage = tencentcloud.Action{
		Service: "ecc",
		Version: "2018-12-13",
		Action:  "CorrectMultiImage",
	}
)
