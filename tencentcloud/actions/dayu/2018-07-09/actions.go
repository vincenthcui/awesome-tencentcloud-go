package dayu

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	ModifyResourceRenewFlag = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "ModifyResourceRenewFlag",
	}
	ModifyResBindDDoSPolicy = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "ModifyResBindDDoSPolicy",
	}
	ModifyNewL4Rule = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "ModifyNewL4Rule",
	}
	ModifyNewDomainRules = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "ModifyNewDomainRules",
	}
	ModifyNetReturnSwitch = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "ModifyNetReturnSwitch",
	}
	ModifyL7Rules = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "ModifyL7Rules",
	}
	ModifyL4Rules = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "ModifyL4Rules",
	}
	ModifyL4KeepTime = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "ModifyL4KeepTime",
	}
	ModifyL4Health = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "ModifyL4Health",
	}
	ModifyElasticLimit = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "ModifyElasticLimit",
	}
	ModifyDDoSWaterKey = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "ModifyDDoSWaterKey",
	}
	ModifyDDoSThreshold = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "ModifyDDoSThreshold",
	}
	ModifyDDoSSwitch = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "ModifyDDoSSwitch",
	}
	ModifyDDoSPolicyName = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "ModifyDDoSPolicyName",
	}
	ModifyDDoSPolicyCase = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "ModifyDDoSPolicyCase",
	}
	ModifyDDoSPolicy = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "ModifyDDoSPolicy",
	}
	ModifyDDoSLevel = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "ModifyDDoSLevel",
	}
	ModifyDDoSDefendStatus = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "ModifyDDoSDefendStatus",
	}
	ModifyDDoSAlarmThreshold = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "ModifyDDoSAlarmThreshold",
	}
	ModifyDDoSAIStatus = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "ModifyDDoSAIStatus",
	}
	ModifyCCUrlAllow = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "ModifyCCUrlAllow",
	}
	ModifyCCThreshold = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "ModifyCCThreshold",
	}
	ModifyCCSelfDefinePolicy = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "ModifyCCSelfDefinePolicy",
	}
	ModifyCCPolicySwitch = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "ModifyCCPolicySwitch",
	}
	ModifyCCLevel = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "ModifyCCLevel",
	}
	ModifyCCIpAllowDeny = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "ModifyCCIpAllowDeny",
	}
	ModifyCCHostProtection = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "ModifyCCHostProtection",
	}
	ModifyCCFrequencyRulesStatus = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "ModifyCCFrequencyRulesStatus",
	}
	ModifyCCFrequencyRules = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "ModifyCCFrequencyRules",
	}
	ModifyCCAlarmThreshold = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "ModifyCCAlarmThreshold",
	}
	DescribleRegionCount = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribleRegionCount",
	}
	DescribleNewL7Rules = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribleNewL7Rules",
	}
	DescribleL7Rules = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribleL7Rules",
	}
	DescribleL4Rules = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribleL4Rules",
	}
	DescribeUnBlockStatis = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribeUnBlockStatis",
	}
	DescribeTransmitStatis = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribeTransmitStatis",
	}
	DescribeSourceIpSegment = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribeSourceIpSegment",
	}
	DescribeSecIndex = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribeSecIndex",
	}
	DescribeSchedulingDomainList = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribeSchedulingDomainList",
	}
	DescribeRuleSets = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribeRuleSets",
	}
	DescribeResourceList = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribeResourceList",
	}
	DescribeResIpList = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribeResIpList",
	}
	DescribePolicyCase = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribePolicyCase",
	}
	DescribePcap = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribePcap",
	}
	DescribePackIndex = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribePackIndex",
	}
	DescribeNewL7RulesErrHealth = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribeNewL7RulesErrHealth",
	}
	DescribeNewL4RulesErrHealth = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribeNewL4RulesErrHealth",
	}
	DescribeNewL4Rules = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribeNewL4Rules",
	}
	DescribeL7HealthConfig = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribeL7HealthConfig",
	}
	DescribeL4RulesErrHealth = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribeL4RulesErrHealth",
	}
	DescribeL4HealthConfig = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribeL4HealthConfig",
	}
	DescribeIpUnBlockList = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribeIpUnBlockList",
	}
	DescribeIpBlockList = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribeIpBlockList",
	}
	DescribeInsurePacks = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribeInsurePacks",
	}
	DescribeIPProductInfo = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribeIPProductInfo",
	}
	DescribeDDoSUsedStatis = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribeDDoSUsedStatis",
	}
	DescribeDDoSTrend = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribeDDoSTrend",
	}
	DescribeDDoSPolicy = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribeDDoSPolicy",
	}
	DescribeDDoSNetTrend = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribeDDoSNetTrend",
	}
	DescribeDDoSNetIpLog = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribeDDoSNetIpLog",
	}
	DescribeDDoSNetEvList = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribeDDoSNetEvList",
	}
	DescribeDDoSNetEvInfo = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribeDDoSNetEvInfo",
	}
	DescribeDDoSNetCount = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribeDDoSNetCount",
	}
	DescribeDDoSIpLog = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribeDDoSIpLog",
	}
	DescribeDDoSEvList = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribeDDoSEvList",
	}
	DescribeDDoSEvInfo = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribeDDoSEvInfo",
	}
	DescribeDDoSDefendStatus = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribeDDoSDefendStatus",
	}
	DescribeDDoSCount = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribeDDoSCount",
	}
	DescribeDDoSAttackSource = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribeDDoSAttackSource",
	}
	DescribeDDoSAttackIPRegionMap = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribeDDoSAttackIPRegionMap",
	}
	DescribeDDoSAlarmThreshold = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribeDDoSAlarmThreshold",
	}
	DescribeCCUrlAllow = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribeCCUrlAllow",
	}
	DescribeCCTrend = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribeCCTrend",
	}
	DescribeCCSelfDefinePolicy = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribeCCSelfDefinePolicy",
	}
	DescribeCCIpAllowDeny = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribeCCIpAllowDeny",
	}
	DescribeCCFrequencyRules = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribeCCFrequencyRules",
	}
	DescribeCCEvList = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribeCCEvList",
	}
	DescribeCCAlarmThreshold = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribeCCAlarmThreshold",
	}
	DescribeBizTrend = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribeBizTrend",
	}
	DescribeBasicDeviceThreshold = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribeBasicDeviceThreshold",
	}
	DescribeBasicCCThreshold = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribeBasicCCThreshold",
	}
	DescribeBaradData = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribeBaradData",
	}
	DescribeBGPIPL7RuleMaxCnt = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribeBGPIPL7RuleMaxCnt",
	}
	DescribeActionLog = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DescribeActionLog",
	}
	DeleteNewL7Rules = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DeleteNewL7Rules",
	}
	DeleteNewL4Rules = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DeleteNewL4Rules",
	}
	DeleteL7Rules = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DeleteL7Rules",
	}
	DeleteL4Rules = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DeleteL4Rules",
	}
	DeleteDDoSPolicyCase = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DeleteDDoSPolicyCase",
	}
	DeleteDDoSPolicy = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DeleteDDoSPolicy",
	}
	DeleteCCSelfDefinePolicy = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DeleteCCSelfDefinePolicy",
	}
	DeleteCCFrequencyRules = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "DeleteCCFrequencyRules",
	}
	CreateUnblockIp = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "CreateUnblockIp",
	}
	CreateNewL7RulesUpload = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "CreateNewL7RulesUpload",
	}
	CreateNewL7Rules = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "CreateNewL7Rules",
	}
	CreateNewL4Rules = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "CreateNewL4Rules",
	}
	CreateNetReturn = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "CreateNetReturn",
	}
	CreateL7RulesUpload = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "CreateL7RulesUpload",
	}
	CreateL7Rules = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "CreateL7Rules",
	}
	CreateL7RuleCert = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "CreateL7RuleCert",
	}
	CreateL7HealthConfig = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "CreateL7HealthConfig",
	}
	CreateL7CCRule = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "CreateL7CCRule",
	}
	CreateL4Rules = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "CreateL4Rules",
	}
	CreateL4HealthConfig = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "CreateL4HealthConfig",
	}
	CreateInstanceName = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "CreateInstanceName",
	}
	CreateDDoSPolicyCase = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "CreateDDoSPolicyCase",
	}
	CreateDDoSPolicy = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "CreateDDoSPolicy",
	}
	CreateCCSelfDefinePolicy = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "CreateCCSelfDefinePolicy",
	}
	CreateCCFrequencyRules = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "CreateCCFrequencyRules",
	}
	CreateBoundIP = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "CreateBoundIP",
	}
	CreateBasicDDoSAlarmThreshold = tencentcloud.Action{
		Service: "dayu",
		Version: "2018-07-09",
		Action:  "CreateBasicDDoSAlarmThreshold",
	}
)
