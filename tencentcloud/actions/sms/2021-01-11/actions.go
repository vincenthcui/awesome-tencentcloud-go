package sms

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	DescribeSmsTemplateList = tencentcloud.Action{
		Service: "sms",
		Version: "2021-01-11",
		Action:  "DescribeSmsTemplateList",
	}
)
