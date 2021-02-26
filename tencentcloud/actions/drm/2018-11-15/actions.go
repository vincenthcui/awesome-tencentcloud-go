package drm

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	StartEncryption = tencentcloud.Action{
		Service: "drm",
		Version: "2018-11-15",
		Action:  "StartEncryption",
	}
	ModifyFairPlayPem = tencentcloud.Action{
		Service: "drm",
		Version: "2018-11-15",
		Action:  "ModifyFairPlayPem",
	}
	DescribeKeys = tencentcloud.Action{
		Service: "drm",
		Version: "2018-11-15",
		Action:  "DescribeKeys",
	}
	DescribeFairPlayPem = tencentcloud.Action{
		Service: "drm",
		Version: "2018-11-15",
		Action:  "DescribeFairPlayPem",
	}
	DescribeAllKeys = tencentcloud.Action{
		Service: "drm",
		Version: "2018-11-15",
		Action:  "DescribeAllKeys",
	}
	DeleteFairPlayPem = tencentcloud.Action{
		Service: "drm",
		Version: "2018-11-15",
		Action:  "DeleteFairPlayPem",
	}
	CreateLicense = tencentcloud.Action{
		Service: "drm",
		Version: "2018-11-15",
		Action:  "CreateLicense",
	}
	CreateEncryptKeys = tencentcloud.Action{
		Service: "drm",
		Version: "2018-11-15",
		Action:  "CreateEncryptKeys",
	}
	AddFairPlayPem = tencentcloud.Action{
		Service: "drm",
		Version: "2018-11-15",
		Action:  "AddFairPlayPem",
	}
)
