package ds

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	SignContractByKeyword = tencentcloud.Action{
		Service: "ds",
		Version: "2018-05-23",
		Action:  "SignContractByKeyword",
	}
	SignContractByCoordinate = tencentcloud.Action{
		Service: "ds",
		Version: "2018-05-23",
		Action:  "SignContractByCoordinate",
	}
	SendVcode = tencentcloud.Action{
		Service: "ds",
		Version: "2018-05-23",
		Action:  "SendVcode",
	}
	DownloadContract = tencentcloud.Action{
		Service: "ds",
		Version: "2018-05-23",
		Action:  "DownloadContract",
	}
	DescribeTaskStatus = tencentcloud.Action{
		Service: "ds",
		Version: "2018-05-23",
		Action:  "DescribeTaskStatus",
	}
	DeleteSeal = tencentcloud.Action{
		Service: "ds",
		Version: "2018-05-23",
		Action:  "DeleteSeal",
	}
	DeleteAccount = tencentcloud.Action{
		Service: "ds",
		Version: "2018-05-23",
		Action:  "DeleteAccount",
	}
	CreateSeal = tencentcloud.Action{
		Service: "ds",
		Version: "2018-05-23",
		Action:  "CreateSeal",
	}
	CreatePersonalAccount = tencentcloud.Action{
		Service: "ds",
		Version: "2018-05-23",
		Action:  "CreatePersonalAccount",
	}
	CreateEnterpriseAccount = tencentcloud.Action{
		Service: "ds",
		Version: "2018-05-23",
		Action:  "CreateEnterpriseAccount",
	}
	CreateContractByUpload = tencentcloud.Action{
		Service: "ds",
		Version: "2018-05-23",
		Action:  "CreateContractByUpload",
	}
	CheckVcode = tencentcloud.Action{
		Service: "ds",
		Version: "2018-05-23",
		Action:  "CheckVcode",
	}
)
