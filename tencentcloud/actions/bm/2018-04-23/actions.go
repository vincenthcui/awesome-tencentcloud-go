package bm

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	UnbindPsaTag = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "UnbindPsaTag",
	}
	StartDevices = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "StartDevices",
	}
	ShutdownDevices = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "ShutdownDevices",
	}
	SetOutBandVpnAuthPassword = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "SetOutBandVpnAuthPassword",
	}
	RunUserCmd = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "RunUserCmd",
	}
	ReturnDevices = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "ReturnDevices",
	}
	ResetDevicePassword = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "ResetDevicePassword",
	}
	RepairTaskControl = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "RepairTaskControl",
	}
	ReloadDeviceOs = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "ReloadDeviceOs",
	}
	RecoverDevices = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "RecoverDevices",
	}
	RebootDevices = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "RebootDevices",
	}
	OfflineDevices = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "OfflineDevices",
	}
	ModifyUserCmd = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "ModifyUserCmd",
	}
	ModifyPsaRegulation = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "ModifyPsaRegulation",
	}
	ModifyPayModePre2Post = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "ModifyPayModePre2Post",
	}
	ModifyLanIp = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "ModifyLanIp",
	}
	ModifyDeviceAutoRenewFlag = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "ModifyDeviceAutoRenewFlag",
	}
	ModifyDeviceAliases = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "ModifyDeviceAliases",
	}
	ModifyCustomImageAttribute = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "ModifyCustomImageAttribute",
	}
	DetachCamRole = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "DetachCamRole",
	}
	DescribeUserCmds = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "DescribeUserCmds",
	}
	DescribeUserCmdTasks = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "DescribeUserCmdTasks",
	}
	DescribeUserCmdTaskInfo = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "DescribeUserCmdTaskInfo",
	}
	DescribeTaskOperationLog = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "DescribeTaskOperationLog",
	}
	DescribeTaskInfo = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "DescribeTaskInfo",
	}
	DescribeRepairTaskConstant = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "DescribeRepairTaskConstant",
	}
	DescribeRegions = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "DescribeRegions",
	}
	DescribePsaRegulations = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "DescribePsaRegulations",
	}
	DescribeOsInfo = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "DescribeOsInfo",
	}
	DescribeOperationResult = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "DescribeOperationResult",
	}
	DescribeHostedDeviceOutBandInfo = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "DescribeHostedDeviceOutBandInfo",
	}
	DescribeHardwareSpecification = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "DescribeHardwareSpecification",
	}
	DescribeDevices = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "DescribeDevices",
	}
	DescribeDevicePriceInfo = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "DescribeDevicePriceInfo",
	}
	DescribeDevicePosition = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "DescribeDevicePosition",
	}
	DescribeDevicePartition = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "DescribeDevicePartition",
	}
	DescribeDeviceOperationLog = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "DescribeDeviceOperationLog",
	}
	DescribeDeviceInventory = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "DescribeDeviceInventory",
	}
	DescribeDeviceHardwareInfo = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "DescribeDeviceHardwareInfo",
	}
	DescribeDeviceClassPartition = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "DescribeDeviceClassPartition",
	}
	DescribeDeviceClass = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "DescribeDeviceClass",
	}
	DescribeCustomImages = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "DescribeCustomImages",
	}
	DescribeCustomImageProcess = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "DescribeCustomImageProcess",
	}
	DeleteUserCmds = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "DeleteUserCmds",
	}
	DeletePsaRegulation = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "DeletePsaRegulation",
	}
	DeleteCustomImages = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "DeleteCustomImages",
	}
	CreateUserCmd = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "CreateUserCmd",
	}
	CreateSpotDevice = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "CreateSpotDevice",
	}
	CreatePsaRegulation = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "CreatePsaRegulation",
	}
	CreateCustomImage = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "CreateCustomImage",
	}
	BuyDevices = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "BuyDevices",
	}
	BindPsaTag = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "BindPsaTag",
	}
	AttachCamRole = tencentcloud.Action{
		Service: "bm",
		Version: "2018-04-23",
		Action:  "AttachCamRole",
	}
)
