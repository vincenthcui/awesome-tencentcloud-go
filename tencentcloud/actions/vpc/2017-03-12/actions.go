package vpc

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	WithdrawNotifyRoutes = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "WithdrawNotifyRoutes",
	}
	UnassignPrivateIpAddresses = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "UnassignPrivateIpAddresses",
	}
	UnassignIpv6SubnetCidrBlock = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "UnassignIpv6SubnetCidrBlock",
	}
	UnassignIpv6CidrBlock = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "UnassignIpv6CidrBlock",
	}
	UnassignIpv6Addresses = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "UnassignIpv6Addresses",
	}
	TransformAddress = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "TransformAddress",
	}
	SetCcnRegionBandwidthLimits = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "SetCcnRegionBandwidthLimits",
	}
	ResetVpnGatewayInternetMaxBandwidth = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ResetVpnGatewayInternetMaxBandwidth",
	}
	ResetVpnConnection = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ResetVpnConnection",
	}
	ResetRoutes = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ResetRoutes",
	}
	ResetNatGatewayConnection = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ResetNatGatewayConnection",
	}
	ResetAttachCcnInstances = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ResetAttachCcnInstances",
	}
	ReplaceSecurityGroupPolicy = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ReplaceSecurityGroupPolicy",
	}
	ReplaceRoutes = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ReplaceRoutes",
	}
	ReplaceRouteTableAssociation = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ReplaceRouteTableAssociation",
	}
	ReplaceDirectConnectGatewayCcnRoutes = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ReplaceDirectConnectGatewayCcnRoutes",
	}
	RenewVpnGateway = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "RenewVpnGateway",
	}
	RenewAddresses = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "RenewAddresses",
	}
	RemoveIp6Rules = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "RemoveIp6Rules",
	}
	RemoveBandwidthPackageResources = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "RemoveBandwidthPackageResources",
	}
	ReleaseIp6AddressesBandwidth = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ReleaseIp6AddressesBandwidth",
	}
	ReleaseAddresses = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ReleaseAddresses",
	}
	RejectAttachCcnInstances = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "RejectAttachCcnInstances",
	}
	NotifyRoutes = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "NotifyRoutes",
	}
	ModifyVpnGatewayCcnRoutes = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ModifyVpnGatewayCcnRoutes",
	}
	ModifyVpnGatewayAttribute = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ModifyVpnGatewayAttribute",
	}
	ModifyVpnConnectionAttribute = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ModifyVpnConnectionAttribute",
	}
	ModifyVpcAttribute = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ModifyVpcAttribute",
	}
	ModifySubnetAttribute = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ModifySubnetAttribute",
	}
	ModifyServiceTemplateGroupAttribute = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ModifyServiceTemplateGroupAttribute",
	}
	ModifyServiceTemplateAttribute = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ModifyServiceTemplateAttribute",
	}
	ModifySecurityGroupPolicies = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ModifySecurityGroupPolicies",
	}
	ModifySecurityGroupAttribute = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ModifySecurityGroupAttribute",
	}
	ModifyRouteTableAttribute = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ModifyRouteTableAttribute",
	}
	ModifyPrivateIpAddressesAttribute = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ModifyPrivateIpAddressesAttribute",
	}
	ModifyNetworkInterfaceAttribute = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ModifyNetworkInterfaceAttribute",
	}
	ModifyNetworkAclEntries = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ModifyNetworkAclEntries",
	}
	ModifyNetworkAclAttribute = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ModifyNetworkAclAttribute",
	}
	ModifyNetDetect = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ModifyNetDetect",
	}
	ModifyNatGatewaySourceIpTranslationNatRule = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ModifyNatGatewaySourceIpTranslationNatRule",
	}
	ModifyNatGatewayDestinationIpPortTranslationNatRule = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ModifyNatGatewayDestinationIpPortTranslationNatRule",
	}
	ModifyNatGatewayAttribute = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ModifyNatGatewayAttribute",
	}
	ModifyIpv6AddressesAttribute = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ModifyIpv6AddressesAttribute",
	}
	ModifyIp6Translator = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ModifyIp6Translator",
	}
	ModifyIp6Rule = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ModifyIp6Rule",
	}
	ModifyIp6AddressesBandwidth = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ModifyIp6AddressesBandwidth",
	}
	ModifyHaVipAttribute = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ModifyHaVipAttribute",
	}
	ModifyGatewayFlowQos = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ModifyGatewayFlowQos",
	}
	ModifyFlowLogAttribute = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ModifyFlowLogAttribute",
	}
	ModifyDirectConnectGatewayAttribute = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ModifyDirectConnectGatewayAttribute",
	}
	ModifyDhcpIpAttribute = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ModifyDhcpIpAttribute",
	}
	ModifyCustomerGatewayAttribute = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ModifyCustomerGatewayAttribute",
	}
	ModifyCcnRegionBandwidthLimitsType = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ModifyCcnRegionBandwidthLimitsType",
	}
	ModifyCcnAttribute = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ModifyCcnAttribute",
	}
	ModifyBandwidthPackageAttribute = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ModifyBandwidthPackageAttribute",
	}
	ModifyAssistantCidr = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ModifyAssistantCidr",
	}
	ModifyAddressesBandwidth = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ModifyAddressesBandwidth",
	}
	ModifyAddressTemplateGroupAttribute = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ModifyAddressTemplateGroupAttribute",
	}
	ModifyAddressTemplateAttribute = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ModifyAddressTemplateAttribute",
	}
	ModifyAddressInternetChargeType = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ModifyAddressInternetChargeType",
	}
	ModifyAddressAttribute = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "ModifyAddressAttribute",
	}
	MigratePrivateIpAddress = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "MigratePrivateIpAddress",
	}
	MigrateNetworkInterface = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "MigrateNetworkInterface",
	}
	InquiryPriceResetVpnGatewayInternetMaxBandwidth = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "InquiryPriceResetVpnGatewayInternetMaxBandwidth",
	}
	InquiryPriceRenewVpnGateway = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "InquiryPriceRenewVpnGateway",
	}
	InquiryPriceCreateVpnGateway = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "InquiryPriceCreateVpnGateway",
	}
	InquirePriceCreateDirectConnectGateway = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "InquirePriceCreateDirectConnectGateway",
	}
	HaVipDisassociateAddressIp = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "HaVipDisassociateAddressIp",
	}
	HaVipAssociateAddressIp = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "HaVipAssociateAddressIp",
	}
	GetCcnRegionBandwidthLimits = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "GetCcnRegionBandwidthLimits",
	}
	EnableRoutes = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "EnableRoutes",
	}
	EnableGatewayFlowMonitor = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "EnableGatewayFlowMonitor",
	}
	EnableCcnRoutes = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "EnableCcnRoutes",
	}
	DownloadCustomerGatewayConfiguration = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DownloadCustomerGatewayConfiguration",
	}
	DisassociateNetworkInterfaceSecurityGroups = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DisassociateNetworkInterfaceSecurityGroups",
	}
	DisassociateNetworkAclSubnets = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DisassociateNetworkAclSubnets",
	}
	DisassociateNatGatewayAddress = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DisassociateNatGatewayAddress",
	}
	DisassociateDirectConnectGatewayNatGateway = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DisassociateDirectConnectGatewayNatGateway",
	}
	DisassociateDhcpIpWithAddressIp = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DisassociateDhcpIpWithAddressIp",
	}
	DisassociateAddress = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DisassociateAddress",
	}
	DisableRoutes = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DisableRoutes",
	}
	DisableGatewayFlowMonitor = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DisableGatewayFlowMonitor",
	}
	DisableCcnRoutes = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DisableCcnRoutes",
	}
	DetachNetworkInterface = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DetachNetworkInterface",
	}
	DetachClassicLinkVpc = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DetachClassicLinkVpc",
	}
	DetachCcnInstances = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DetachCcnInstances",
	}
	DescribeVpnGateways = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeVpnGateways",
	}
	DescribeVpnGatewayCcnRoutes = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeVpnGatewayCcnRoutes",
	}
	DescribeVpnConnections = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeVpnConnections",
	}
	DescribeVpcs = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeVpcs",
	}
	DescribeVpcResourceDashboard = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeVpcResourceDashboard",
	}
	DescribeVpcPrivateIpAddresses = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeVpcPrivateIpAddresses",
	}
	DescribeVpcLimits = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeVpcLimits",
	}
	DescribeVpcIpv6Addresses = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeVpcIpv6Addresses",
	}
	DescribeVpcInstances = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeVpcInstances",
	}
	DescribeTemplateLimits = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeTemplateLimits",
	}
	DescribeTaskResult = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeTaskResult",
	}
	DescribeSubnets = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeSubnets",
	}
	DescribeServiceTemplates = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeServiceTemplates",
	}
	DescribeServiceTemplateGroups = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeServiceTemplateGroups",
	}
	DescribeSecurityGroups = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeSecurityGroups",
	}
	DescribeSecurityGroupReferences = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeSecurityGroupReferences",
	}
	DescribeSecurityGroupPolicies = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeSecurityGroupPolicies",
	}
	DescribeSecurityGroupLimits = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeSecurityGroupLimits",
	}
	DescribeSecurityGroupAssociationStatistics = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeSecurityGroupAssociationStatistics",
	}
	DescribeRouteTables = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeRouteTables",
	}
	DescribeRouteConflicts = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeRouteConflicts",
	}
	DescribeProductQuota = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeProductQuota",
	}
	DescribeNetworkInterfaces = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeNetworkInterfaces",
	}
	DescribeNetworkInterfaceLimit = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeNetworkInterfaceLimit",
	}
	DescribeNetworkAcls = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeNetworkAcls",
	}
	DescribeNetDetects = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeNetDetects",
	}
	DescribeNetDetectStates = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeNetDetectStates",
	}
	DescribeNatGateways = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeNatGateways",
	}
	DescribeNatGatewaySourceIpTranslationNatRules = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeNatGatewaySourceIpTranslationNatRules",
	}
	DescribeNatGatewayDestinationIpPortTranslationNatRules = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeNatGatewayDestinationIpPortTranslationNatRules",
	}
	DescribeIpGeolocationInfos = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeIpGeolocationInfos",
	}
	DescribeIpGeolocationDatabaseUrl = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeIpGeolocationDatabaseUrl",
	}
	DescribeIp6Translators = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeIp6Translators",
	}
	DescribeIp6TranslatorQuota = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeIp6TranslatorQuota",
	}
	DescribeIp6Addresses = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeIp6Addresses",
	}
	DescribeHaVips = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeHaVips",
	}
	DescribeGatewayFlowQos = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeGatewayFlowQos",
	}
	DescribeGatewayFlowMonitorDetail = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeGatewayFlowMonitorDetail",
	}
	DescribeFlowLogs = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeFlowLogs",
	}
	DescribeFlowLog = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeFlowLog",
	}
	DescribeDirectConnectGateways = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeDirectConnectGateways",
	}
	DescribeDirectConnectGatewayCcnRoutes = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeDirectConnectGatewayCcnRoutes",
	}
	DescribeDhcpIps = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeDhcpIps",
	}
	DescribeCustomerGateways = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeCustomerGateways",
	}
	DescribeCustomerGatewayVendors = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeCustomerGatewayVendors",
	}
	DescribeCrossBorderCompliance = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeCrossBorderCompliance",
	}
	DescribeClassicLinkInstances = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeClassicLinkInstances",
	}
	DescribeCcns = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeCcns",
	}
	DescribeCcnRoutes = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeCcnRoutes",
	}
	DescribeCcnRegionBandwidthLimits = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeCcnRegionBandwidthLimits",
	}
	DescribeCcnAttachedInstances = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeCcnAttachedInstances",
	}
	DescribeBandwidthPackages = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeBandwidthPackages",
	}
	DescribeBandwidthPackageResources = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeBandwidthPackageResources",
	}
	DescribeBandwidthPackageQuota = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeBandwidthPackageQuota",
	}
	DescribeBandwidthPackageBillUsage = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeBandwidthPackageBillUsage",
	}
	DescribeAssistantCidr = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeAssistantCidr",
	}
	DescribeAddresses = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeAddresses",
	}
	DescribeAddressTemplates = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeAddressTemplates",
	}
	DescribeAddressTemplateGroups = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeAddressTemplateGroups",
	}
	DescribeAddressQuota = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeAddressQuota",
	}
	DescribeAccountAttributes = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DescribeAccountAttributes",
	}
	DeleteVpnGateway = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DeleteVpnGateway",
	}
	DeleteVpnConnection = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DeleteVpnConnection",
	}
	DeleteVpc = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DeleteVpc",
	}
	DeleteSubnet = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DeleteSubnet",
	}
	DeleteServiceTemplateGroup = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DeleteServiceTemplateGroup",
	}
	DeleteServiceTemplate = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DeleteServiceTemplate",
	}
	DeleteSecurityGroupPolicies = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DeleteSecurityGroupPolicies",
	}
	DeleteSecurityGroup = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DeleteSecurityGroup",
	}
	DeleteRoutes = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DeleteRoutes",
	}
	DeleteRouteTable = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DeleteRouteTable",
	}
	DeleteNetworkInterface = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DeleteNetworkInterface",
	}
	DeleteNetworkAcl = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DeleteNetworkAcl",
	}
	DeleteNetDetect = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DeleteNetDetect",
	}
	DeleteNatGatewaySourceIpTranslationNatRule = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DeleteNatGatewaySourceIpTranslationNatRule",
	}
	DeleteNatGatewayDestinationIpPortTranslationNatRule = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DeleteNatGatewayDestinationIpPortTranslationNatRule",
	}
	DeleteNatGateway = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DeleteNatGateway",
	}
	DeleteIp6Translators = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DeleteIp6Translators",
	}
	DeleteHaVip = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DeleteHaVip",
	}
	DeleteFlowLog = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DeleteFlowLog",
	}
	DeleteDirectConnectGatewayCcnRoutes = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DeleteDirectConnectGatewayCcnRoutes",
	}
	DeleteDirectConnectGateway = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DeleteDirectConnectGateway",
	}
	DeleteDhcpIp = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DeleteDhcpIp",
	}
	DeleteCustomerGateway = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DeleteCustomerGateway",
	}
	DeleteCcn = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DeleteCcn",
	}
	DeleteBandwidthPackage = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DeleteBandwidthPackage",
	}
	DeleteAssistantCidr = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DeleteAssistantCidr",
	}
	DeleteAddressTemplateGroup = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DeleteAddressTemplateGroup",
	}
	DeleteAddressTemplate = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "DeleteAddressTemplate",
	}
	CreateVpnGateway = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "CreateVpnGateway",
	}
	CreateVpnConnection = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "CreateVpnConnection",
	}
	CreateVpc = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "CreateVpc",
	}
	CreateSubnets = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "CreateSubnets",
	}
	CreateSubnet = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "CreateSubnet",
	}
	CreateServiceTemplateGroup = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "CreateServiceTemplateGroup",
	}
	CreateServiceTemplate = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "CreateServiceTemplate",
	}
	CreateSecurityGroupWithPolicies = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "CreateSecurityGroupWithPolicies",
	}
	CreateSecurityGroupPolicies = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "CreateSecurityGroupPolicies",
	}
	CreateSecurityGroup = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "CreateSecurityGroup",
	}
	CreateRoutes = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "CreateRoutes",
	}
	CreateRouteTable = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "CreateRouteTable",
	}
	CreateNetworkInterface = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "CreateNetworkInterface",
	}
	CreateNetworkAcl = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "CreateNetworkAcl",
	}
	CreateNetDetect = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "CreateNetDetect",
	}
	CreateNatGatewaySourceIpTranslationNatRule = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "CreateNatGatewaySourceIpTranslationNatRule",
	}
	CreateNatGatewayDestinationIpPortTranslationNatRule = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "CreateNatGatewayDestinationIpPortTranslationNatRule",
	}
	CreateNatGateway = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "CreateNatGateway",
	}
	CreateIp6Translators = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "CreateIp6Translators",
	}
	CreateHaVip = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "CreateHaVip",
	}
	CreateFlowLog = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "CreateFlowLog",
	}
	CreateDirectConnectGatewayCcnRoutes = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "CreateDirectConnectGatewayCcnRoutes",
	}
	CreateDirectConnectGateway = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "CreateDirectConnectGateway",
	}
	CreateDhcpIp = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "CreateDhcpIp",
	}
	CreateDefaultVpc = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "CreateDefaultVpc",
	}
	CreateDefaultSecurityGroup = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "CreateDefaultSecurityGroup",
	}
	CreateCustomerGateway = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "CreateCustomerGateway",
	}
	CreateCcn = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "CreateCcn",
	}
	CreateBandwidthPackage = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "CreateBandwidthPackage",
	}
	CreateAssistantCidr = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "CreateAssistantCidr",
	}
	CreateAndAttachNetworkInterface = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "CreateAndAttachNetworkInterface",
	}
	CreateAddressTemplateGroup = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "CreateAddressTemplateGroup",
	}
	CreateAddressTemplate = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "CreateAddressTemplate",
	}
	CloneSecurityGroup = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "CloneSecurityGroup",
	}
	CheckNetDetectState = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "CheckNetDetectState",
	}
	CheckDefaultSubnet = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "CheckDefaultSubnet",
	}
	CheckAssistantCidr = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "CheckAssistantCidr",
	}
	AuditCrossBorderCompliance = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "AuditCrossBorderCompliance",
	}
	AttachNetworkInterface = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "AttachNetworkInterface",
	}
	AttachClassicLinkVpc = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "AttachClassicLinkVpc",
	}
	AttachCcnInstances = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "AttachCcnInstances",
	}
	AssociateNetworkInterfaceSecurityGroups = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "AssociateNetworkInterfaceSecurityGroups",
	}
	AssociateNetworkAclSubnets = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "AssociateNetworkAclSubnets",
	}
	AssociateNatGatewayAddress = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "AssociateNatGatewayAddress",
	}
	AssociateDirectConnectGatewayNatGateway = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "AssociateDirectConnectGatewayNatGateway",
	}
	AssociateDhcpIpWithAddressIp = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "AssociateDhcpIpWithAddressIp",
	}
	AssociateAddress = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "AssociateAddress",
	}
	AssignPrivateIpAddresses = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "AssignPrivateIpAddresses",
	}
	AssignIpv6SubnetCidrBlock = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "AssignIpv6SubnetCidrBlock",
	}
	AssignIpv6CidrBlock = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "AssignIpv6CidrBlock",
	}
	AssignIpv6Addresses = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "AssignIpv6Addresses",
	}
	AllocateIp6AddressesBandwidth = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "AllocateIp6AddressesBandwidth",
	}
	AllocateAddresses = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "AllocateAddresses",
	}
	AddIp6Rules = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "AddIp6Rules",
	}
	AddBandwidthPackageResources = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "AddBandwidthPackageResources",
	}
	AcceptAttachCcnInstances = tencentcloud.Action{
		Service: "vpc",
		Version: "2017-03-12",
		Action:  "AcceptAttachCcnInstances",
	}
)
