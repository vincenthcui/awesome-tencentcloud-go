package tsf

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	UpdateRepository = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "UpdateRepository",
	}
	UpdateHealthCheckSettings = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "UpdateHealthCheckSettings",
	}
	UpdateGatewayApi = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "UpdateGatewayApi",
	}
	UpdateApiRateLimitRules = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "UpdateApiRateLimitRules",
	}
	UpdateApiRateLimitRule = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "UpdateApiRateLimitRule",
	}
	UpdateApiGroup = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "UpdateApiGroup",
	}
	UnbindApiGroup = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "UnbindApiGroup",
	}
	TerminateTaskFlowBatch = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "TerminateTaskFlowBatch",
	}
	StopTaskExecute = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "StopTaskExecute",
	}
	StopTaskBatch = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "StopTaskBatch",
	}
	StopGroup = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "StopGroup",
	}
	StopContainerGroup = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "StopContainerGroup",
	}
	StartGroup = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "StartGroup",
	}
	StartContainerGroup = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "StartContainerGroup",
	}
	ShrinkInstances = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "ShrinkInstances",
	}
	ShrinkGroup = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "ShrinkGroup",
	}
	RollbackConfig = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "RollbackConfig",
	}
	RevocationPublicConfig = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "RevocationPublicConfig",
	}
	RevocationConfig = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "RevocationConfig",
	}
	RemoveInstances = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "RemoveInstances",
	}
	ReleasePublicConfig = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "ReleasePublicConfig",
	}
	ReleaseConfig = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "ReleaseConfig",
	}
	ReleaseApiGroup = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "ReleaseApiGroup",
	}
	RedoTaskFlowBatch = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "RedoTaskFlowBatch",
	}
	RedoTaskExecute = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "RedoTaskExecute",
	}
	RedoTaskBatch = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "RedoTaskBatch",
	}
	RedoTask = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "RedoTask",
	}
	ModifyUploadInfo = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "ModifyUploadInfo",
	}
	ModifyTask = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "ModifyTask",
	}
	ModifyPathRewrite = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "ModifyPathRewrite",
	}
	ModifyMicroservice = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "ModifyMicroservice",
	}
	ModifyLaneRule = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "ModifyLaneRule",
	}
	ModifyLane = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "ModifyLane",
	}
	ModifyContainerReplicas = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "ModifyContainerReplicas",
	}
	ModifyContainerGroup = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "ModifyContainerGroup",
	}
	ExpandGroup = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "ExpandGroup",
	}
	ExecuteTaskFlow = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "ExecuteTaskFlow",
	}
	ExecuteTask = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "ExecuteTask",
	}
	EnableTaskFlow = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "EnableTaskFlow",
	}
	EnableTask = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "EnableTask",
	}
	DraftApiGroup = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DraftApiGroup",
	}
	DisableTaskFlow = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DisableTaskFlow",
	}
	DisableTask = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DisableTask",
	}
	DescribeUploadInfo = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribeUploadInfo",
	}
	DescribeTaskLastStatus = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribeTaskLastStatus",
	}
	DescribeSimpleNamespaces = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribeSimpleNamespaces",
	}
	DescribeSimpleGroups = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribeSimpleGroups",
	}
	DescribeSimpleClusters = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribeSimpleClusters",
	}
	DescribeSimpleApplications = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribeSimpleApplications",
	}
	DescribeServerlessGroups = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribeServerlessGroups",
	}
	DescribeServerlessGroup = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribeServerlessGroup",
	}
	DescribeRepository = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribeRepository",
	}
	DescribeRepositories = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribeRepositories",
	}
	DescribeReleasedConfig = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribeReleasedConfig",
	}
	DescribePublicConfigs = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribePublicConfigs",
	}
	DescribePublicConfigSummary = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribePublicConfigSummary",
	}
	DescribePublicConfigReleases = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribePublicConfigReleases",
	}
	DescribePublicConfigReleaseLogs = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribePublicConfigReleaseLogs",
	}
	DescribePublicConfig = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribePublicConfig",
	}
	DescribePodInstances = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribePodInstances",
	}
	DescribePkgs = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribePkgs",
	}
	DescribePathRewrites = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribePathRewrites",
	}
	DescribePathRewrite = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribePathRewrite",
	}
	DescribeMsApiList = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribeMsApiList",
	}
	DescribeMicroservices = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribeMicroservices",
	}
	DescribeMicroservice = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribeMicroservice",
	}
	DescribeLanes = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribeLanes",
	}
	DescribeLaneRules = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribeLaneRules",
	}
	DescribeImageTags = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribeImageTags",
	}
	DescribeImageRepository = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribeImageRepository",
	}
	DescribeGroups = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribeGroups",
	}
	DescribeGroupUseDetail = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribeGroupUseDetail",
	}
	DescribeGroupInstances = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribeGroupInstances",
	}
	DescribeGroupGateways = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribeGroupGateways",
	}
	DescribeGroupBindedGateways = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribeGroupBindedGateways",
	}
	DescribeGroup = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribeGroup",
	}
	DescribeGatewayMonitorOverview = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribeGatewayMonitorOverview",
	}
	DescribeGatewayAllGroupApis = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribeGatewayAllGroupApis",
	}
	DescribeFlowLastBatchState = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribeFlowLastBatchState",
	}
	DescribeDownloadInfo = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribeDownloadInfo",
	}
	DescribeCreateGatewayApiStatus = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribeCreateGatewayApiStatus",
	}
	DescribeContainerGroups = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribeContainerGroups",
	}
	DescribeContainerGroupDetail = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribeContainerGroupDetail",
	}
	DescribeConfigs = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribeConfigs",
	}
	DescribeConfigSummary = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribeConfigSummary",
	}
	DescribeConfigReleases = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribeConfigReleases",
	}
	DescribeConfigReleaseLogs = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribeConfigReleaseLogs",
	}
	DescribeConfig = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribeConfig",
	}
	DescribeClusterInstances = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribeClusterInstances",
	}
	DescribeBasicResourceUsage = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribeBasicResourceUsage",
	}
	DescribeApplications = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribeApplications",
	}
	DescribeApplicationAttribute = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribeApplicationAttribute",
	}
	DescribeApplication = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribeApplication",
	}
	DescribeApiVersions = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribeApiVersions",
	}
	DescribeApiUseDetail = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribeApiUseDetail",
	}
	DescribeApiRateLimitRules = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribeApiRateLimitRules",
	}
	DescribeApiGroups = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribeApiGroups",
	}
	DescribeApiGroup = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribeApiGroup",
	}
	DescribeApiDetail = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DescribeApiDetail",
	}
	DeployServerlessGroup = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DeployServerlessGroup",
	}
	DeployGroup = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DeployGroup",
	}
	DeployContainerGroup = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DeployContainerGroup",
	}
	DeleteTask = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DeleteTask",
	}
	DeleteServerlessGroup = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DeleteServerlessGroup",
	}
	DeleteRepository = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DeleteRepository",
	}
	DeletePublicConfig = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DeletePublicConfig",
	}
	DeletePkgs = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DeletePkgs",
	}
	DeletePathRewrites = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DeletePathRewrites",
	}
	DeleteNamespace = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DeleteNamespace",
	}
	DeleteMicroservice = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DeleteMicroservice",
	}
	DeleteLane = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DeleteLane",
	}
	DeleteImageTags = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DeleteImageTags",
	}
	DeleteGroup = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DeleteGroup",
	}
	DeleteContainerGroup = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DeleteContainerGroup",
	}
	DeleteConfig = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DeleteConfig",
	}
	DeleteApplication = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DeleteApplication",
	}
	DeleteApiGroup = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "DeleteApiGroup",
	}
	CreateTaskFlow = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "CreateTaskFlow",
	}
	CreateTask = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "CreateTask",
	}
	CreateServerlessGroup = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "CreateServerlessGroup",
	}
	CreateRepository = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "CreateRepository",
	}
	CreatePublicConfig = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "CreatePublicConfig",
	}
	CreatePathRewrites = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "CreatePathRewrites",
	}
	CreateNamespace = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "CreateNamespace",
	}
	CreateMicroservice = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "CreateMicroservice",
	}
	CreateLaneRule = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "CreateLaneRule",
	}
	CreateLane = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "CreateLane",
	}
	CreateGroup = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "CreateGroup",
	}
	CreateGatewayApi = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "CreateGatewayApi",
	}
	CreateContainGroup = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "CreateContainGroup",
	}
	CreateConfig = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "CreateConfig",
	}
	CreateCluster = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "CreateCluster",
	}
	CreateApplication = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "CreateApplication",
	}
	CreateApiRateLimitRule = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "CreateApiRateLimitRule",
	}
	CreateApiGroup = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "CreateApiGroup",
	}
	CreateAllGatewayApiAsync = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "CreateAllGatewayApiAsync",
	}
	ContinueRunFailedTaskBatch = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "ContinueRunFailedTaskBatch",
	}
	ChangeApiUsableStatus = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "ChangeApiUsableStatus",
	}
	BindApiGroup = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "BindApiGroup",
	}
	AddInstances = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "AddInstances",
	}
	AddClusterInstances = tencentcloud.Action{
		Service: "tsf",
		Version: "2018-03-26",
		Action:  "AddClusterInstances",
	}
)
