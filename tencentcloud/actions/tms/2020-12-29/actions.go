package tms

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	TextModeration = tencentcloud.Action{
		Service: "tms",
		Version: "2020-12-29",
		Action:  "TextModeration",
	}
)
