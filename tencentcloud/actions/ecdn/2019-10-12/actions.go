package ecdn

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	UpdateDomainConfig = tencentcloud.Action{
		Service: "ecdn",
		Version: "2019-10-12",
		Action:  "UpdateDomainConfig",
	}
	StopEcdnDomain = tencentcloud.Action{
		Service: "ecdn",
		Version: "2019-10-12",
		Action:  "StopEcdnDomain",
	}
	StartEcdnDomain = tencentcloud.Action{
		Service: "ecdn",
		Version: "2019-10-12",
		Action:  "StartEcdnDomain",
	}
	PurgeUrlsCache = tencentcloud.Action{
		Service: "ecdn",
		Version: "2019-10-12",
		Action:  "PurgeUrlsCache",
	}
	PurgePathCache = tencentcloud.Action{
		Service: "ecdn",
		Version: "2019-10-12",
		Action:  "PurgePathCache",
	}
	DescribePurgeTasks = tencentcloud.Action{
		Service: "ecdn",
		Version: "2019-10-12",
		Action:  "DescribePurgeTasks",
	}
	DescribePurgeQuota = tencentcloud.Action{
		Service: "ecdn",
		Version: "2019-10-12",
		Action:  "DescribePurgeQuota",
	}
	DescribeIpStatus = tencentcloud.Action{
		Service: "ecdn",
		Version: "2019-10-12",
		Action:  "DescribeIpStatus",
	}
	DescribeEcdnStatistics = tencentcloud.Action{
		Service: "ecdn",
		Version: "2019-10-12",
		Action:  "DescribeEcdnStatistics",
	}
	DescribeEcdnDomainStatistics = tencentcloud.Action{
		Service: "ecdn",
		Version: "2019-10-12",
		Action:  "DescribeEcdnDomainStatistics",
	}
	DescribeEcdnDomainLogs = tencentcloud.Action{
		Service: "ecdn",
		Version: "2019-10-12",
		Action:  "DescribeEcdnDomainLogs",
	}
	DescribeDomainsConfig = tencentcloud.Action{
		Service: "ecdn",
		Version: "2019-10-12",
		Action:  "DescribeDomainsConfig",
	}
	DescribeDomains = tencentcloud.Action{
		Service: "ecdn",
		Version: "2019-10-12",
		Action:  "DescribeDomains",
	}
	DeleteEcdnDomain = tencentcloud.Action{
		Service: "ecdn",
		Version: "2019-10-12",
		Action:  "DeleteEcdnDomain",
	}
	AddEcdnDomain = tencentcloud.Action{
		Service: "ecdn",
		Version: "2019-10-12",
		Action:  "AddEcdnDomain",
	}
)
