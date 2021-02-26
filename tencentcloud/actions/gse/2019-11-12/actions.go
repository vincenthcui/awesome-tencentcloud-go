package gse

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	UpdateRuntimeConfiguration = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "UpdateRuntimeConfiguration",
	}
	UpdateGameServerSessionQueue = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "UpdateGameServerSessionQueue",
	}
	UpdateGameServerSession = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "UpdateGameServerSession",
	}
	UpdateFleetPortSettings = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "UpdateFleetPortSettings",
	}
	UpdateFleetName = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "UpdateFleetName",
	}
	UpdateFleetCapacity = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "UpdateFleetCapacity",
	}
	UpdateFleetAttributes = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "UpdateFleetAttributes",
	}
	UpdateBucketCORSOpt = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "UpdateBucketCORSOpt",
	}
	UpdateBucketAccelerateOpt = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "UpdateBucketAccelerateOpt",
	}
	UpdateAsset = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "UpdateAsset",
	}
	UpdateAlias = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "UpdateAlias",
	}
	StopGameServerSessionPlacement = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "StopGameServerSessionPlacement",
	}
	StopFleetActions = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "StopFleetActions",
	}
	StartGameServerSessionPlacement = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "StartGameServerSessionPlacement",
	}
	StartFleetActions = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "StartFleetActions",
	}
	SetServerWeight = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "SetServerWeight",
	}
	SetServerReserved = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "SetServerReserved",
	}
	SearchGameServerSessions = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "SearchGameServerSessions",
	}
	ResolveAlias = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "ResolveAlias",
	}
	PutTimerScalingPolicy = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "PutTimerScalingPolicy",
	}
	PutScalingPolicy = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "PutScalingPolicy",
	}
	ListFleets = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "ListFleets",
	}
	ListAliases = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "ListAliases",
	}
	JoinGameServerSessionBatch = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "JoinGameServerSessionBatch",
	}
	JoinGameServerSession = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "JoinGameServerSession",
	}
	GetUploadFederationToken = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "GetUploadFederationToken",
	}
	GetUploadCredentials = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "GetUploadCredentials",
	}
	GetInstanceAccess = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "GetInstanceAccess",
	}
	GetGameServerSessionLogUrl = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "GetGameServerSessionLogUrl",
	}
	DetachCcnInstances = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "DetachCcnInstances",
	}
	DescribeUserQuotas = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "DescribeUserQuotas",
	}
	DescribeUserQuota = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "DescribeUserQuota",
	}
	DescribeTimerScalingPolicies = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "DescribeTimerScalingPolicies",
	}
	DescribeScalingPolicies = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "DescribeScalingPolicies",
	}
	DescribeRuntimeConfiguration = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "DescribeRuntimeConfiguration",
	}
	DescribePlayerSessions = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "DescribePlayerSessions",
	}
	DescribeInstancesExtend = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "DescribeInstancesExtend",
	}
	DescribeInstances = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "DescribeInstances",
	}
	DescribeInstanceTypes = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "DescribeInstanceTypes",
	}
	DescribeInstanceLimit = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "DescribeInstanceLimit",
	}
	DescribeGameServerSessions = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "DescribeGameServerSessions",
	}
	DescribeGameServerSessionQueues = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "DescribeGameServerSessionQueues",
	}
	DescribeGameServerSessionPlacement = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "DescribeGameServerSessionPlacement",
	}
	DescribeGameServerSessionDetails = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "DescribeGameServerSessionDetails",
	}
	DescribeFleetUtilization = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "DescribeFleetUtilization",
	}
	DescribeFleetStatisticSummary = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "DescribeFleetStatisticSummary",
	}
	DescribeFleetStatisticFlows = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "DescribeFleetStatisticFlows",
	}
	DescribeFleetStatisticDetails = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "DescribeFleetStatisticDetails",
	}
	DescribeFleetPortSettings = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "DescribeFleetPortSettings",
	}
	DescribeFleetEvents = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "DescribeFleetEvents",
	}
	DescribeFleetCapacity = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "DescribeFleetCapacity",
	}
	DescribeFleetAttributes = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "DescribeFleetAttributes",
	}
	DescribeCcnInstances = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "DescribeCcnInstances",
	}
	DescribeAssets = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "DescribeAssets",
	}
	DescribeAssetSystems = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "DescribeAssetSystems",
	}
	DescribeAsset = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "DescribeAsset",
	}
	DescribeAlias = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "DescribeAlias",
	}
	DeleteTimerScalingPolicy = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "DeleteTimerScalingPolicy",
	}
	DeleteScalingPolicy = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "DeleteScalingPolicy",
	}
	DeleteGameServerSessionQueue = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "DeleteGameServerSessionQueue",
	}
	DeleteFleet = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "DeleteFleet",
	}
	DeleteAsset = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "DeleteAsset",
	}
	DeleteAlias = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "DeleteAlias",
	}
	CreateGameServerSessionQueue = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "CreateGameServerSessionQueue",
	}
	CreateGameServerSession = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "CreateGameServerSession",
	}
	CreateFleet = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "CreateFleet",
	}
	CreateAssetWithImage = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "CreateAssetWithImage",
	}
	CreateAsset = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "CreateAsset",
	}
	CreateAlias = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "CreateAlias",
	}
	CopyFleet = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "CopyFleet",
	}
	AttachCcnInstances = tencentcloud.Action{
		Service: "gse",
		Version: "2019-11-12",
		Action:  "AttachCcnInstances",
	}
)
