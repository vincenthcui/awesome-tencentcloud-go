package tke

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	UpgradeClusterInstances = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "UpgradeClusterInstances",
	}
	UpdateEKSCluster = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "UpdateEKSCluster",
	}
	UpdateClusterVersion = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "UpdateClusterVersion",
	}
	SyncPrometheusTemplate = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "SyncPrometheusTemplate",
	}
	RemoveNodeFromNodePool = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "RemoveNodeFromNodePool",
	}
	ModifyPrometheusTemplate = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "ModifyPrometheusTemplate",
	}
	ModifyNodePoolDesiredCapacityAboutAsg = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "ModifyNodePoolDesiredCapacityAboutAsg",
	}
	ModifyClusterNodePool = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "ModifyClusterNodePool",
	}
	ModifyClusterEndpointSP = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "ModifyClusterEndpointSP",
	}
	ModifyClusterAttribute = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "ModifyClusterAttribute",
	}
	ModifyClusterAsGroupOptionAttribute = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "ModifyClusterAsGroupOptionAttribute",
	}
	ModifyClusterAsGroupAttribute = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "ModifyClusterAsGroupAttribute",
	}
	GetUpgradeInstanceProgress = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "GetUpgradeInstanceProgress",
	}
	DescribeRouteTableConflicts = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "DescribeRouteTableConflicts",
	}
	DescribeRegions = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "DescribeRegions",
	}
	DescribePrometheusTemplates = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "DescribePrometheusTemplates",
	}
	DescribePrometheusTemplateSync = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "DescribePrometheusTemplateSync",
	}
	DescribePrometheusTargets = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "DescribePrometheusTargets",
	}
	DescribePrometheusOverviews = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "DescribePrometheusOverviews",
	}
	DescribePrometheusAlertRule = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "DescribePrometheusAlertRule",
	}
	DescribePrometheusAlertHistory = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "DescribePrometheusAlertHistory",
	}
	DescribePrometheusAgents = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "DescribePrometheusAgents",
	}
	DescribePrometheusAgentInstances = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "DescribePrometheusAgentInstances",
	}
	DescribeImages = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "DescribeImages",
	}
	DescribeExistedInstances = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "DescribeExistedInstances",
	}
	DescribeEKSClusters = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "DescribeEKSClusters",
	}
	DescribeEKSClusterCredential = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "DescribeEKSClusterCredential",
	}
	DescribeClusters = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "DescribeClusters",
	}
	DescribeClusterSecurity = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "DescribeClusterSecurity",
	}
	DescribeClusterRoutes = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "DescribeClusterRoutes",
	}
	DescribeClusterRouteTables = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "DescribeClusterRouteTables",
	}
	DescribeClusterNodePools = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "DescribeClusterNodePools",
	}
	DescribeClusterNodePoolDetail = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "DescribeClusterNodePoolDetail",
	}
	DescribeClusterKubeconfig = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "DescribeClusterKubeconfig",
	}
	DescribeClusterInstances = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "DescribeClusterInstances",
	}
	DescribeClusterEndpointVipStatus = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "DescribeClusterEndpointVipStatus",
	}
	DescribeClusterEndpointStatus = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "DescribeClusterEndpointStatus",
	}
	DescribeClusterAsGroups = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "DescribeClusterAsGroups",
	}
	DescribeClusterAsGroupOption = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "DescribeClusterAsGroupOption",
	}
	DescribeAvailableClusterVersion = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "DescribeAvailableClusterVersion",
	}
	DeletePrometheusTemplateSync = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "DeletePrometheusTemplateSync",
	}
	DeletePrometheusTemplate = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "DeletePrometheusTemplate",
	}
	DeleteEKSCluster = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "DeleteEKSCluster",
	}
	DeleteClusterRouteTable = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "DeleteClusterRouteTable",
	}
	DeleteClusterRoute = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "DeleteClusterRoute",
	}
	DeleteClusterNodePool = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "DeleteClusterNodePool",
	}
	DeleteClusterInstances = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "DeleteClusterInstances",
	}
	DeleteClusterEndpointVip = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "DeleteClusterEndpointVip",
	}
	DeleteClusterEndpoint = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "DeleteClusterEndpoint",
	}
	DeleteClusterAsGroups = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "DeleteClusterAsGroups",
	}
	DeleteCluster = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "DeleteCluster",
	}
	CreatePrometheusTemplate = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "CreatePrometheusTemplate",
	}
	CreatePrometheusDashboard = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "CreatePrometheusDashboard",
	}
	CreateEKSCluster = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "CreateEKSCluster",
	}
	CreateClusterRouteTable = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "CreateClusterRouteTable",
	}
	CreateClusterRoute = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "CreateClusterRoute",
	}
	CreateClusterNodePoolFromExistingAsg = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "CreateClusterNodePoolFromExistingAsg",
	}
	CreateClusterNodePool = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "CreateClusterNodePool",
	}
	CreateClusterInstances = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "CreateClusterInstances",
	}
	CreateClusterEndpointVip = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "CreateClusterEndpointVip",
	}
	CreateClusterEndpoint = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "CreateClusterEndpoint",
	}
	CreateClusterAsGroup = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "CreateClusterAsGroup",
	}
	CreateCluster = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "CreateCluster",
	}
	CheckInstancesUpgradeAble = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "CheckInstancesUpgradeAble",
	}
	AddNodeToNodePool = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "AddNodeToNodePool",
	}
	AddExistedInstances = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "AddExistedInstances",
	}
	AcquireClusterAdminRole = tencentcloud.Action{
		Service: "tke",
		Version: "2018-05-25",
		Action:  "AcquireClusterAdminRole",
	}
)
