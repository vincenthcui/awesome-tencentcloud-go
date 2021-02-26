package cloudaudit

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	UpdateAudit = tencentcloud.Action{
		Service: "cloudaudit",
		Version: "2019-03-19",
		Action:  "UpdateAudit",
	}
	StopLogging = tencentcloud.Action{
		Service: "cloudaudit",
		Version: "2019-03-19",
		Action:  "StopLogging",
	}
	StartLogging = tencentcloud.Action{
		Service: "cloudaudit",
		Version: "2019-03-19",
		Action:  "StartLogging",
	}
	LookUpEvents = tencentcloud.Action{
		Service: "cloudaudit",
		Version: "2019-03-19",
		Action:  "LookUpEvents",
	}
	ListKeyAliasByRegion = tencentcloud.Action{
		Service: "cloudaudit",
		Version: "2019-03-19",
		Action:  "ListKeyAliasByRegion",
	}
	ListCosEnableRegion = tencentcloud.Action{
		Service: "cloudaudit",
		Version: "2019-03-19",
		Action:  "ListCosEnableRegion",
	}
	ListCmqEnableRegion = tencentcloud.Action{
		Service: "cloudaudit",
		Version: "2019-03-19",
		Action:  "ListCmqEnableRegion",
	}
	ListAudits = tencentcloud.Action{
		Service: "cloudaudit",
		Version: "2019-03-19",
		Action:  "ListAudits",
	}
	InquireAuditCredit = tencentcloud.Action{
		Service: "cloudaudit",
		Version: "2019-03-19",
		Action:  "InquireAuditCredit",
	}
	GetAttributeKey = tencentcloud.Action{
		Service: "cloudaudit",
		Version: "2019-03-19",
		Action:  "GetAttributeKey",
	}
	DescribeAudit = tencentcloud.Action{
		Service: "cloudaudit",
		Version: "2019-03-19",
		Action:  "DescribeAudit",
	}
	DeleteAudit = tencentcloud.Action{
		Service: "cloudaudit",
		Version: "2019-03-19",
		Action:  "DeleteAudit",
	}
	CreateAudit = tencentcloud.Action{
		Service: "cloudaudit",
		Version: "2019-03-19",
		Action:  "CreateAudit",
	}
)
