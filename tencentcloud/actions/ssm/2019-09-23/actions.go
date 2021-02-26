package ssm

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	UpdateSecret = tencentcloud.Action{
		Service: "ssm",
		Version: "2019-09-23",
		Action:  "UpdateSecret",
	}
	UpdateDescription = tencentcloud.Action{
		Service: "ssm",
		Version: "2019-09-23",
		Action:  "UpdateDescription",
	}
	RestoreSecret = tencentcloud.Action{
		Service: "ssm",
		Version: "2019-09-23",
		Action:  "RestoreSecret",
	}
	PutSecretValue = tencentcloud.Action{
		Service: "ssm",
		Version: "2019-09-23",
		Action:  "PutSecretValue",
	}
	ListSecrets = tencentcloud.Action{
		Service: "ssm",
		Version: "2019-09-23",
		Action:  "ListSecrets",
	}
	ListSecretVersionIds = tencentcloud.Action{
		Service: "ssm",
		Version: "2019-09-23",
		Action:  "ListSecretVersionIds",
	}
	GetServiceStatus = tencentcloud.Action{
		Service: "ssm",
		Version: "2019-09-23",
		Action:  "GetServiceStatus",
	}
	GetSecretValue = tencentcloud.Action{
		Service: "ssm",
		Version: "2019-09-23",
		Action:  "GetSecretValue",
	}
	GetRegions = tencentcloud.Action{
		Service: "ssm",
		Version: "2019-09-23",
		Action:  "GetRegions",
	}
	EnableSecret = tencentcloud.Action{
		Service: "ssm",
		Version: "2019-09-23",
		Action:  "EnableSecret",
	}
	DisableSecret = tencentcloud.Action{
		Service: "ssm",
		Version: "2019-09-23",
		Action:  "DisableSecret",
	}
	DescribeSecret = tencentcloud.Action{
		Service: "ssm",
		Version: "2019-09-23",
		Action:  "DescribeSecret",
	}
	DeleteSecretVersion = tencentcloud.Action{
		Service: "ssm",
		Version: "2019-09-23",
		Action:  "DeleteSecretVersion",
	}
	DeleteSecret = tencentcloud.Action{
		Service: "ssm",
		Version: "2019-09-23",
		Action:  "DeleteSecret",
	}
	CreateSecret = tencentcloud.Action{
		Service: "ssm",
		Version: "2019-09-23",
		Action:  "CreateSecret",
	}
)
