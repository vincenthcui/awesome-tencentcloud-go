package cdn

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	VerifyDomainRecord = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "VerifyDomainRecord",
	}
	UpdateScdnDomain = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "UpdateScdnDomain",
	}
	UpdatePayType = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "UpdatePayType",
	}
	UpdateImageConfig = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "UpdateImageConfig",
	}
	UpdateDomainConfig = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "UpdateDomainConfig",
	}
	StopScdnDomain = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "StopScdnDomain",
	}
	StopCdnDomain = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "StopCdnDomain",
	}
	StartScdnDomain = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "StartScdnDomain",
	}
	StartCdnDomain = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "StartCdnDomain",
	}
	SearchClsLog = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "SearchClsLog",
	}
	PushUrlsCache = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "PushUrlsCache",
	}
	PurgeUrlsCache = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "PurgeUrlsCache",
	}
	PurgePathCache = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "PurgePathCache",
	}
	ManageClsTopicDomains = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "ManageClsTopicDomains",
	}
	ListTopData = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "ListTopData",
	}
	ListScdnLogTasks = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "ListScdnLogTasks",
	}
	ListScdnDomains = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "ListScdnDomains",
	}
	ListDiagnoseReport = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "ListDiagnoseReport",
	}
	ListClsTopicDomains = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "ListClsTopicDomains",
	}
	ListClsLogTopics = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "ListClsLogTopics",
	}
	GetDisableRecords = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "GetDisableRecords",
	}
	EnableClsLogTopic = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "EnableClsLogTopic",
	}
	EnableCaches = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "EnableCaches",
	}
	DuplicateDomainConfig = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "DuplicateDomainConfig",
	}
	DisableClsLogTopic = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "DisableClsLogTopic",
	}
	DisableCaches = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "DisableCaches",
	}
	DescribeUrlViolations = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "DescribeUrlViolations",
	}
	DescribeTrafficPackages = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "DescribeTrafficPackages",
	}
	DescribeScdnTopData = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "DescribeScdnTopData",
	}
	DescribeScdnConfig = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "DescribeScdnConfig",
	}
	DescribeReportData = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "DescribeReportData",
	}
	DescribePushTasks = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "DescribePushTasks",
	}
	DescribePushQuota = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "DescribePushQuota",
	}
	DescribePurgeTasks = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "DescribePurgeTasks",
	}
	DescribePurgeQuota = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "DescribePurgeQuota",
	}
	DescribePayType = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "DescribePayType",
	}
	DescribeOriginData = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "DescribeOriginData",
	}
	DescribeMapInfo = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "DescribeMapInfo",
	}
	DescribeIpVisit = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "DescribeIpVisit",
	}
	DescribeIpStatus = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "DescribeIpStatus",
	}
	DescribeImageConfig = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "DescribeImageConfig",
	}
	DescribeDomainsConfig = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "DescribeDomainsConfig",
	}
	DescribeDomains = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "DescribeDomains",
	}
	DescribeDistrictIspData = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "DescribeDistrictIspData",
	}
	DescribeDiagnoseReport = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "DescribeDiagnoseReport",
	}
	DescribeCertDomains = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "DescribeCertDomains",
	}
	DescribeCdnOriginIp = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "DescribeCdnOriginIp",
	}
	DescribeCdnIp = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "DescribeCdnIp",
	}
	DescribeCdnDomainLogs = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "DescribeCdnDomainLogs",
	}
	DescribeCdnData = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "DescribeCdnData",
	}
	DescribeBillingData = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "DescribeBillingData",
	}
	DeleteScdnDomain = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "DeleteScdnDomain",
	}
	DeleteClsLogTopic = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "DeleteClsLogTopic",
	}
	DeleteCdnDomain = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "DeleteCdnDomain",
	}
	CreateVerifyRecord = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "CreateVerifyRecord",
	}
	CreateScdnLogTask = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "CreateScdnLogTask",
	}
	CreateEdgePackTask = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "CreateEdgePackTask",
	}
	CreateDiagnoseUrl = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "CreateDiagnoseUrl",
	}
	CreateClsLogTopic = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "CreateClsLogTopic",
	}
	AddCdnDomain = tencentcloud.Action{
		Service: "cdn",
		Version: "2018-06-06",
		Action:  "AddCdnDomain",
	}
)
