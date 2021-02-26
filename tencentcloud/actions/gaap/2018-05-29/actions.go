package gaap

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	SetAuthentication = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "SetAuthentication",
	}
	RemoveRealServers = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "RemoveRealServers",
	}
	OpenSecurityPolicy = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "OpenSecurityPolicy",
	}
	OpenProxyGroup = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "OpenProxyGroup",
	}
	OpenProxies = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "OpenProxies",
	}
	ModifyUDPListenerAttribute = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "ModifyUDPListenerAttribute",
	}
	ModifyTCPListenerAttribute = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "ModifyTCPListenerAttribute",
	}
	ModifySecurityRule = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "ModifySecurityRule",
	}
	ModifyRuleAttribute = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "ModifyRuleAttribute",
	}
	ModifyRealServerName = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "ModifyRealServerName",
	}
	ModifyProxyGroupAttribute = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "ModifyProxyGroupAttribute",
	}
	ModifyProxyConfiguration = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "ModifyProxyConfiguration",
	}
	ModifyProxiesProject = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "ModifyProxiesProject",
	}
	ModifyProxiesAttribute = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "ModifyProxiesAttribute",
	}
	ModifyHTTPSListenerAttribute = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "ModifyHTTPSListenerAttribute",
	}
	ModifyHTTPListenerAttribute = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "ModifyHTTPListenerAttribute",
	}
	ModifyGroupDomainConfig = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "ModifyGroupDomainConfig",
	}
	ModifyDomain = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "ModifyDomain",
	}
	ModifyCertificateAttributes = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "ModifyCertificateAttributes",
	}
	ModifyCertificate = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "ModifyCertificate",
	}
	InquiryPriceCreateProxy = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "InquiryPriceCreateProxy",
	}
	DestroyProxies = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "DestroyProxies",
	}
	DescribeUDPListeners = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "DescribeUDPListeners",
	}
	DescribeTCPListeners = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "DescribeTCPListeners",
	}
	DescribeSecurityRules = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "DescribeSecurityRules",
	}
	DescribeSecurityPolicyDetail = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "DescribeSecurityPolicyDetail",
	}
	DescribeRulesByRuleIds = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "DescribeRulesByRuleIds",
	}
	DescribeRules = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "DescribeRules",
	}
	DescribeRuleRealServers = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "DescribeRuleRealServers",
	}
	DescribeResourcesByTag = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "DescribeResourcesByTag",
	}
	DescribeRegionAndPrice = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "DescribeRegionAndPrice",
	}
	DescribeRealServersStatus = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "DescribeRealServersStatus",
	}
	DescribeRealServers = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "DescribeRealServers",
	}
	DescribeRealServerStatistics = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "DescribeRealServerStatistics",
	}
	DescribeProxyStatistics = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "DescribeProxyStatistics",
	}
	DescribeProxyGroupStatistics = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "DescribeProxyGroupStatistics",
	}
	DescribeProxyGroupList = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "DescribeProxyGroupList",
	}
	DescribeProxyGroupDetails = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "DescribeProxyGroupDetails",
	}
	DescribeProxyDetail = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "DescribeProxyDetail",
	}
	DescribeProxyAndStatisticsListeners = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "DescribeProxyAndStatisticsListeners",
	}
	DescribeProxiesStatus = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "DescribeProxiesStatus",
	}
	DescribeProxies = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "DescribeProxies",
	}
	DescribeListenerStatistics = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "DescribeListenerStatistics",
	}
	DescribeListenerRealServers = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "DescribeListenerRealServers",
	}
	DescribeHTTPSListeners = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "DescribeHTTPSListeners",
	}
	DescribeHTTPListeners = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "DescribeHTTPListeners",
	}
	DescribeGroupDomainConfig = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "DescribeGroupDomainConfig",
	}
	DescribeGroupAndStatisticsProxy = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "DescribeGroupAndStatisticsProxy",
	}
	DescribeDomainErrorPageInfoByIds = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "DescribeDomainErrorPageInfoByIds",
	}
	DescribeDomainErrorPageInfo = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "DescribeDomainErrorPageInfo",
	}
	DescribeDestRegions = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "DescribeDestRegions",
	}
	DescribeCountryAreaMapping = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "DescribeCountryAreaMapping",
	}
	DescribeCertificates = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "DescribeCertificates",
	}
	DescribeCertificateDetail = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "DescribeCertificateDetail",
	}
	DescribeAccessRegionsByDestRegion = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "DescribeAccessRegionsByDestRegion",
	}
	DescribeAccessRegions = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "DescribeAccessRegions",
	}
	DeleteSecurityRules = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "DeleteSecurityRules",
	}
	DeleteSecurityPolicy = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "DeleteSecurityPolicy",
	}
	DeleteRule = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "DeleteRule",
	}
	DeleteProxyGroup = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "DeleteProxyGroup",
	}
	DeleteListeners = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "DeleteListeners",
	}
	DeleteDomainErrorPageInfo = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "DeleteDomainErrorPageInfo",
	}
	DeleteDomain = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "DeleteDomain",
	}
	DeleteCertificate = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "DeleteCertificate",
	}
	CreateUDPListeners = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "CreateUDPListeners",
	}
	CreateTCPListeners = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "CreateTCPListeners",
	}
	CreateSecurityRules = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "CreateSecurityRules",
	}
	CreateSecurityPolicy = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "CreateSecurityPolicy",
	}
	CreateRule = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "CreateRule",
	}
	CreateProxyGroupDomain = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "CreateProxyGroupDomain",
	}
	CreateProxyGroup = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "CreateProxyGroup",
	}
	CreateProxy = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "CreateProxy",
	}
	CreateHTTPSListener = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "CreateHTTPSListener",
	}
	CreateHTTPListener = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "CreateHTTPListener",
	}
	CreateDomainErrorPageInfo = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "CreateDomainErrorPageInfo",
	}
	CreateDomain = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "CreateDomain",
	}
	CreateCertificate = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "CreateCertificate",
	}
	CloseSecurityPolicy = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "CloseSecurityPolicy",
	}
	CloseProxyGroup = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "CloseProxyGroup",
	}
	CloseProxies = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "CloseProxies",
	}
	CheckProxyCreate = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "CheckProxyCreate",
	}
	BindRuleRealServers = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "BindRuleRealServers",
	}
	BindListenerRealServers = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "BindListenerRealServers",
	}
	AddRealServers = tencentcloud.Action{
		Service: "gaap",
		Version: "2018-05-29",
		Action:  "AddRealServers",
	}
)
