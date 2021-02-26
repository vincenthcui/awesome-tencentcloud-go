package ecm

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	TerminateInstances = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "TerminateInstances",
	}
	StopInstances = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "StopInstances",
	}
	StartInstances = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "StartInstances",
	}
	SetSecurityGroupForLoadbalancers = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "SetSecurityGroupForLoadbalancers",
	}
	SetLoadBalancerSecurityGroups = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "SetLoadBalancerSecurityGroups",
	}
	RunInstances = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "RunInstances",
	}
	ResetRoutes = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "ResetRoutes",
	}
	ResetInstancesPassword = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "ResetInstancesPassword",
	}
	ResetInstancesMaxBandwidth = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "ResetInstancesMaxBandwidth",
	}
	ResetInstances = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "ResetInstances",
	}
	ReplaceSecurityGroupPolicy = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "ReplaceSecurityGroupPolicy",
	}
	ReplaceRoutes = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "ReplaceRoutes",
	}
	ReplaceRouteTableAssociation = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "ReplaceRouteTableAssociation",
	}
	RemovePrivateIpAddresses = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "RemovePrivateIpAddresses",
	}
	ReleaseIpv6Addresses = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "ReleaseIpv6Addresses",
	}
	ReleaseAddresses = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "ReleaseAddresses",
	}
	RebootInstances = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "RebootInstances",
	}
	ModifyVpcAttribute = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "ModifyVpcAttribute",
	}
	ModifyTargetWeight = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "ModifyTargetWeight",
	}
	ModifyTargetPort = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "ModifyTargetPort",
	}
	ModifySubnetAttribute = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "ModifySubnetAttribute",
	}
	ModifySecurityGroupPolicies = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "ModifySecurityGroupPolicies",
	}
	ModifySecurityGroupAttribute = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "ModifySecurityGroupAttribute",
	}
	ModifyRouteTableAttribute = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "ModifyRouteTableAttribute",
	}
	ModifyPrivateIpAddressesAttribute = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "ModifyPrivateIpAddressesAttribute",
	}
	ModifyModuleSecurityGroups = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "ModifyModuleSecurityGroups",
	}
	ModifyModuleNetwork = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "ModifyModuleNetwork",
	}
	ModifyModuleName = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "ModifyModuleName",
	}
	ModifyModuleIpDirect = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "ModifyModuleIpDirect",
	}
	ModifyModuleImage = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "ModifyModuleImage",
	}
	ModifyModuleDisableWanIp = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "ModifyModuleDisableWanIp",
	}
	ModifyModuleConfig = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "ModifyModuleConfig",
	}
	ModifyLoadBalancerAttributes = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "ModifyLoadBalancerAttributes",
	}
	ModifyListener = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "ModifyListener",
	}
	ModifyIpv6AddressesAttribute = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "ModifyIpv6AddressesAttribute",
	}
	ModifyInstancesAttribute = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "ModifyInstancesAttribute",
	}
	ModifyImageAttribute = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "ModifyImageAttribute",
	}
	ModifyHaVipAttribute = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "ModifyHaVipAttribute",
	}
	ModifyDefaultSubnet = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "ModifyDefaultSubnet",
	}
	ModifyAddressesBandwidth = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "ModifyAddressesBandwidth",
	}
	ModifyAddressAttribute = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "ModifyAddressAttribute",
	}
	MigratePrivateIpAddress = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "MigratePrivateIpAddress",
	}
	MigrateNetworkInterface = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "MigrateNetworkInterface",
	}
	ImportImage = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "ImportImage",
	}
	ImportCustomImage = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "ImportCustomImage",
	}
	EnableRoutes = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "EnableRoutes",
	}
	DisassociateSecurityGroups = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DisassociateSecurityGroups",
	}
	DisassociateAddress = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DisassociateAddress",
	}
	DisableRoutes = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DisableRoutes",
	}
	DetachNetworkInterface = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DetachNetworkInterface",
	}
	DescribeVpcs = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DescribeVpcs",
	}
	DescribeTaskStatus = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DescribeTaskStatus",
	}
	DescribeTaskResult = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DescribeTaskResult",
	}
	DescribeTargets = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DescribeTargets",
	}
	DescribeTargetHealth = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DescribeTargetHealth",
	}
	DescribeSubnets = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DescribeSubnets",
	}
	DescribeSecurityGroups = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DescribeSecurityGroups",
	}
	DescribeSecurityGroupPolicies = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DescribeSecurityGroupPolicies",
	}
	DescribeSecurityGroupLimits = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DescribeSecurityGroupLimits",
	}
	DescribeSecurityGroupAssociationStatistics = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DescribeSecurityGroupAssociationStatistics",
	}
	DescribeRouteTables = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DescribeRouteTables",
	}
	DescribeRouteConflicts = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DescribeRouteConflicts",
	}
	DescribePeakNetworkOverview = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DescribePeakNetworkOverview",
	}
	DescribePeakBaseOverview = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DescribePeakBaseOverview",
	}
	DescribeNode = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DescribeNode",
	}
	DescribeNetworkInterfaces = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DescribeNetworkInterfaces",
	}
	DescribeModuleDetail = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DescribeModuleDetail",
	}
	DescribeModule = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DescribeModule",
	}
	DescribeLoadBalancers = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DescribeLoadBalancers",
	}
	DescribeLoadBalanceTaskStatus = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DescribeLoadBalanceTaskStatus",
	}
	DescribeListeners = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DescribeListeners",
	}
	DescribeInstancesDeniedActions = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DescribeInstancesDeniedActions",
	}
	DescribeInstances = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DescribeInstances",
	}
	DescribeInstanceVncUrl = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DescribeInstanceVncUrl",
	}
	DescribeInstanceTypeConfig = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DescribeInstanceTypeConfig",
	}
	DescribeImportImageOs = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DescribeImportImageOs",
	}
	DescribeImage = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DescribeImage",
	}
	DescribeHaVips = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DescribeHaVips",
	}
	DescribeDefaultSubnet = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DescribeDefaultSubnet",
	}
	DescribeCustomImageTask = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DescribeCustomImageTask",
	}
	DescribeConfig = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DescribeConfig",
	}
	DescribeBaseOverview = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DescribeBaseOverview",
	}
	DescribeAddresses = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DescribeAddresses",
	}
	DescribeAddressQuota = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DescribeAddressQuota",
	}
	DeleteVpc = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DeleteVpc",
	}
	DeleteSubnet = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DeleteSubnet",
	}
	DeleteSecurityGroupPolicies = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DeleteSecurityGroupPolicies",
	}
	DeleteSecurityGroup = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DeleteSecurityGroup",
	}
	DeleteRoutes = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DeleteRoutes",
	}
	DeleteRouteTable = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DeleteRouteTable",
	}
	DeleteNetworkInterface = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DeleteNetworkInterface",
	}
	DeleteModule = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DeleteModule",
	}
	DeleteLoadBalancerListeners = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DeleteLoadBalancerListeners",
	}
	DeleteLoadBalancer = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DeleteLoadBalancer",
	}
	DeleteListener = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DeleteListener",
	}
	DeleteImage = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DeleteImage",
	}
	DeleteHaVip = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "DeleteHaVip",
	}
	CreateVpc = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "CreateVpc",
	}
	CreateSubnet = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "CreateSubnet",
	}
	CreateSecurityGroupPolicies = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "CreateSecurityGroupPolicies",
	}
	CreateSecurityGroup = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "CreateSecurityGroup",
	}
	CreateRoutes = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "CreateRoutes",
	}
	CreateRouteTable = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "CreateRouteTable",
	}
	CreateNetworkInterface = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "CreateNetworkInterface",
	}
	CreateModule = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "CreateModule",
	}
	CreateLoadBalancer = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "CreateLoadBalancer",
	}
	CreateListener = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "CreateListener",
	}
	CreateImage = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "CreateImage",
	}
	CreateHaVip = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "CreateHaVip",
	}
	BatchRegisterTargets = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "BatchRegisterTargets",
	}
	BatchModifyTargetWeight = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "BatchModifyTargetWeight",
	}
	BatchDeregisterTargets = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "BatchDeregisterTargets",
	}
	AttachNetworkInterface = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "AttachNetworkInterface",
	}
	AssociateSecurityGroups = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "AssociateSecurityGroups",
	}
	AssociateAddress = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "AssociateAddress",
	}
	AssignPrivateIpAddresses = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "AssignPrivateIpAddresses",
	}
	AssignIpv6Addresses = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "AssignIpv6Addresses",
	}
	AllocateAddresses = tencentcloud.Action{
		Service: "ecm",
		Version: "2019-07-19",
		Action:  "AllocateAddresses",
	}
)
