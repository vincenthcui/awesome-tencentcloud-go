package sslpod

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	ResolveDomain = tencentcloud.Action{
		Service: "sslpod",
		Version: "2019-06-05",
		Action:  "ResolveDomain",
	}
	RefreshDomain = tencentcloud.Action{
		Service: "sslpod",
		Version: "2019-06-05",
		Action:  "RefreshDomain",
	}
	ModifyDomainTags = tencentcloud.Action{
		Service: "sslpod",
		Version: "2019-06-05",
		Action:  "ModifyDomainTags",
	}
	DescribeNoticeInfo = tencentcloud.Action{
		Service: "sslpod",
		Version: "2019-06-05",
		Action:  "DescribeNoticeInfo",
	}
	DescribeDomains = tencentcloud.Action{
		Service: "sslpod",
		Version: "2019-06-05",
		Action:  "DescribeDomains",
	}
	DescribeDomainTags = tencentcloud.Action{
		Service: "sslpod",
		Version: "2019-06-05",
		Action:  "DescribeDomainTags",
	}
	DescribeDomainCerts = tencentcloud.Action{
		Service: "sslpod",
		Version: "2019-06-05",
		Action:  "DescribeDomainCerts",
	}
	DescribeDashboard = tencentcloud.Action{
		Service: "sslpod",
		Version: "2019-06-05",
		Action:  "DescribeDashboard",
	}
	DeleteDomain = tencentcloud.Action{
		Service: "sslpod",
		Version: "2019-06-05",
		Action:  "DeleteDomain",
	}
	CreateDomain = tencentcloud.Action{
		Service: "sslpod",
		Version: "2019-06-05",
		Action:  "CreateDomain",
	}
)
