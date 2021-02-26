package wss

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	UploadCert = tencentcloud.Action{
		Service: "wss",
		Version: "2018-04-26",
		Action:  "UploadCert",
	}
	DescribeCertList = tencentcloud.Action{
		Service: "wss",
		Version: "2018-04-26",
		Action:  "DescribeCertList",
	}
	DeleteCert = tencentcloud.Action{
		Service: "wss",
		Version: "2018-04-26",
		Action:  "DeleteCert",
	}
)
