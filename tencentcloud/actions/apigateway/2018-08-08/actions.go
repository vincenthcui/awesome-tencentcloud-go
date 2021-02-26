package apigateway

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	UpdateService = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "UpdateService",
	}
	UpdateApiKey = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "UpdateApiKey",
	}
	UnReleaseService = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "UnReleaseService",
	}
	UnBindSubDomain = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "UnBindSubDomain",
	}
	UnBindSecretIds = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "UnBindSecretIds",
	}
	UnBindIPStrategy = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "UnBindIPStrategy",
	}
	UnBindEnvironment = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "UnBindEnvironment",
	}
	ReleaseService = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "ReleaseService",
	}
	ModifyUsagePlan = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "ModifyUsagePlan",
	}
	ModifySubDomain = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "ModifySubDomain",
	}
	ModifyServiceEnvironmentStrategy = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "ModifyServiceEnvironmentStrategy",
	}
	ModifyService = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "ModifyService",
	}
	ModifyIPStrategy = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "ModifyIPStrategy",
	}
	ModifyApiIncrement = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "ModifyApiIncrement",
	}
	ModifyApiEnvironmentStrategy = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "ModifyApiEnvironmentStrategy",
	}
	ModifyApi = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "ModifyApi",
	}
	GenerateApiDocument = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "GenerateApiDocument",
	}
	EnableApiKey = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "EnableApiKey",
	}
	DisableApiKey = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "DisableApiKey",
	}
	DescribeUsagePlansStatus = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "DescribeUsagePlansStatus",
	}
	DescribeUsagePlanSecretIds = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "DescribeUsagePlanSecretIds",
	}
	DescribeUsagePlanEnvironments = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "DescribeUsagePlanEnvironments",
	}
	DescribeUsagePlan = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "DescribeUsagePlan",
	}
	DescribeServicesStatus = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "DescribeServicesStatus",
	}
	DescribeServiceUsagePlan = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "DescribeServiceUsagePlan",
	}
	DescribeServiceSubDomains = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "DescribeServiceSubDomains",
	}
	DescribeServiceSubDomainMappings = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "DescribeServiceSubDomainMappings",
	}
	DescribeServiceReleaseVersion = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "DescribeServiceReleaseVersion",
	}
	DescribeServiceEnvironmentStrategy = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "DescribeServiceEnvironmentStrategy",
	}
	DescribeServiceEnvironmentReleaseHistory = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "DescribeServiceEnvironmentReleaseHistory",
	}
	DescribeServiceEnvironmentList = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "DescribeServiceEnvironmentList",
	}
	DescribeService = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "DescribeService",
	}
	DescribePlugins = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "DescribePlugins",
	}
	DescribeLogSearch = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "DescribeLogSearch",
	}
	DescribeIPStrategysStatus = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "DescribeIPStrategysStatus",
	}
	DescribeIPStrategyApisStatus = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "DescribeIPStrategyApisStatus",
	}
	DescribeIPStrategy = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "DescribeIPStrategy",
	}
	DescribeApisStatus = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "DescribeApisStatus",
	}
	DescribeApiUsagePlan = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "DescribeApiUsagePlan",
	}
	DescribeApiKeysStatus = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "DescribeApiKeysStatus",
	}
	DescribeApiKey = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "DescribeApiKey",
	}
	DescribeApiEnvironmentStrategy = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "DescribeApiEnvironmentStrategy",
	}
	DescribeApi = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "DescribeApi",
	}
	DemoteServiceUsagePlan = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "DemoteServiceUsagePlan",
	}
	DeleteUsagePlan = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "DeleteUsagePlan",
	}
	DeleteServiceSubDomainMapping = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "DeleteServiceSubDomainMapping",
	}
	DeleteService = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "DeleteService",
	}
	DeleteIPStrategy = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "DeleteIPStrategy",
	}
	DeleteApiKey = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "DeleteApiKey",
	}
	DeleteApi = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "DeleteApi",
	}
	CreateUsagePlan = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "CreateUsagePlan",
	}
	CreateService = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "CreateService",
	}
	CreateIPStrategy = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "CreateIPStrategy",
	}
	CreateApiKey = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "CreateApiKey",
	}
	CreateApi = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "CreateApi",
	}
	BindSubDomain = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "BindSubDomain",
	}
	BindSecretIds = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "BindSecretIds",
	}
	BindIPStrategy = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "BindIPStrategy",
	}
	BindEnvironment = tencentcloud.Action{
		Service: "apigateway",
		Version: "2018-08-08",
		Action:  "BindEnvironment",
	}
)
