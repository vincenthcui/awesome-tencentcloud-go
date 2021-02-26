package cvm

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	TerminateInstances = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "TerminateInstances",
	}
	SyncImages = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "SyncImages",
	}
	StopInstances = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "StopInstances",
	}
	StartInstances = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "StartInstances",
	}
	RunInstances = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "RunInstances",
	}
	ResizeInstanceDisks = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "ResizeInstanceDisks",
	}
	ResetInstancesType = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "ResetInstancesType",
	}
	ResetInstancesPassword = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "ResetInstancesPassword",
	}
	ResetInstancesInternetMaxBandwidth = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "ResetInstancesInternetMaxBandwidth",
	}
	ResetInstance = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "ResetInstance",
	}
	RenewInstances = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "RenewInstances",
	}
	RenewHosts = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "RenewHosts",
	}
	RebootInstances = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "RebootInstances",
	}
	PurchaseReservedInstancesOffering = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "PurchaseReservedInstancesOffering",
	}
	ModifyKeyPairAttribute = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "ModifyKeyPairAttribute",
	}
	ModifyInstancesVpcAttribute = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "ModifyInstancesVpcAttribute",
	}
	ModifyInstancesRenewFlag = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "ModifyInstancesRenewFlag",
	}
	ModifyInstancesProject = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "ModifyInstancesProject",
	}
	ModifyInstancesChargeType = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "ModifyInstancesChargeType",
	}
	ModifyInstancesAttribute = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "ModifyInstancesAttribute",
	}
	ModifyImageSharePermission = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "ModifyImageSharePermission",
	}
	ModifyImageAttribute = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "ModifyImageAttribute",
	}
	ModifyHostsAttribute = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "ModifyHostsAttribute",
	}
	ModifyDisasterRecoverGroupAttribute = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "ModifyDisasterRecoverGroupAttribute",
	}
	InquiryPriceRunInstances = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "InquiryPriceRunInstances",
	}
	InquiryPriceResizeInstanceDisks = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "InquiryPriceResizeInstanceDisks",
	}
	InquiryPriceResetInstancesType = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "InquiryPriceResetInstancesType",
	}
	InquiryPriceResetInstancesInternetMaxBandwidth = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "InquiryPriceResetInstancesInternetMaxBandwidth",
	}
	InquiryPriceResetInstance = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "InquiryPriceResetInstance",
	}
	InquiryPriceRenewInstances = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "InquiryPriceRenewInstances",
	}
	InquiryPriceModifyInstancesChargeType = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "InquiryPriceModifyInstancesChargeType",
	}
	InquirePricePurchaseReservedInstancesOffering = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "InquirePricePurchaseReservedInstancesOffering",
	}
	ImportKeyPair = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "ImportKeyPair",
	}
	ImportImage = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "ImportImage",
	}
	DisassociateSecurityGroups = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "DisassociateSecurityGroups",
	}
	DisassociateInstancesKeyPairs = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "DisassociateInstancesKeyPairs",
	}
	DescribeZones = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "DescribeZones",
	}
	DescribeZoneInstanceConfigInfos = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "DescribeZoneInstanceConfigInfos",
	}
	DescribeReservedInstancesOfferings = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "DescribeReservedInstancesOfferings",
	}
	DescribeReservedInstancesConfigInfos = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "DescribeReservedInstancesConfigInfos",
	}
	DescribeReservedInstances = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "DescribeReservedInstances",
	}
	DescribeRegions = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "DescribeRegions",
	}
	DescribeKeyPairs = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "DescribeKeyPairs",
	}
	DescribeInternetChargeTypeConfigs = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "DescribeInternetChargeTypeConfigs",
	}
	DescribeInstancesStatus = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "DescribeInstancesStatus",
	}
	DescribeInstancesOperationLimit = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "DescribeInstancesOperationLimit",
	}
	DescribeInstances = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "DescribeInstances",
	}
	DescribeInstanceVncUrl = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "DescribeInstanceVncUrl",
	}
	DescribeInstanceTypeConfigs = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "DescribeInstanceTypeConfigs",
	}
	DescribeInstanceInternetBandwidthConfigs = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "DescribeInstanceInternetBandwidthConfigs",
	}
	DescribeInstanceFamilyConfigs = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "DescribeInstanceFamilyConfigs",
	}
	DescribeImportImageOs = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "DescribeImportImageOs",
	}
	DescribeImages = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "DescribeImages",
	}
	DescribeImageSharePermission = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "DescribeImageSharePermission",
	}
	DescribeImageQuota = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "DescribeImageQuota",
	}
	DescribeHosts = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "DescribeHosts",
	}
	DescribeDisasterRecoverGroups = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "DescribeDisasterRecoverGroups",
	}
	DescribeDisasterRecoverGroupQuota = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "DescribeDisasterRecoverGroupQuota",
	}
	DeleteKeyPairs = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "DeleteKeyPairs",
	}
	DeleteImages = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "DeleteImages",
	}
	DeleteDisasterRecoverGroups = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "DeleteDisasterRecoverGroups",
	}
	CreateKeyPair = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "CreateKeyPair",
	}
	CreateImage = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "CreateImage",
	}
	CreateDisasterRecoverGroup = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "CreateDisasterRecoverGroup",
	}
	AssociateSecurityGroups = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "AssociateSecurityGroups",
	}
	AssociateInstancesKeyPairs = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "AssociateInstancesKeyPairs",
	}
	AllocateHosts = tencentcloud.Action{
		Service: "cvm",
		Version: "2017-03-12",
		Action:  "AllocateHosts",
	}
)
