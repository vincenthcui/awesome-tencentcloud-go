package sts

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	QueryApiKey = tencentcloud.Action{
		Service: "sts",
		Version: "2018-08-13",
		Action:  "QueryApiKey",
	}
	GetFederationToken = tencentcloud.Action{
		Service: "sts",
		Version: "2018-08-13",
		Action:  "GetFederationToken",
	}
	AssumeRoleWithSAML = tencentcloud.Action{
		Service: "sts",
		Version: "2018-08-13",
		Action:  "AssumeRoleWithSAML",
	}
	AssumeRole = tencentcloud.Action{
		Service: "sts",
		Version: "2018-08-13",
		Action:  "AssumeRole",
	}
)
