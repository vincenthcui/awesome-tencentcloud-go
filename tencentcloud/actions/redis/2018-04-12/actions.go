package redis

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	UpgradeInstanceVersion = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "UpgradeInstanceVersion",
	}
	UpgradeInstance = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "UpgradeInstance",
	}
	SwitchInstanceVip = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "SwitchInstanceVip",
	}
	StartupInstance = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "StartupInstance",
	}
	RestoreInstance = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "RestoreInstance",
	}
	ResetPassword = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "ResetPassword",
	}
	RenewInstance = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "RenewInstance",
	}
	ModifyNetworkConfig = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "ModifyNetworkConfig",
	}
	ModifyMaintenanceWindow = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "ModifyMaintenanceWindow",
	}
	ModifyInstanceParams = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "ModifyInstanceParams",
	}
	ModifyInstanceAccount = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "ModifyInstanceAccount",
	}
	ModifyInstance = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "ModifyInstance",
	}
	ModifyDBInstanceSecurityGroups = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "ModifyDBInstanceSecurityGroups",
	}
	ModifyConnectionConfig = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "ModifyConnectionConfig",
	}
	ModifyAutoBackupConfig = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "ModifyAutoBackupConfig",
	}
	ModfiyInstancePassword = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "ModfiyInstancePassword",
	}
	ManualBackupInstance = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "ManualBackupInstance",
	}
	InquiryPriceUpgradeInstance = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "InquiryPriceUpgradeInstance",
	}
	InquiryPriceRenewInstance = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "InquiryPriceRenewInstance",
	}
	InquiryPriceCreateInstance = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "InquiryPriceCreateInstance",
	}
	EnableReplicaReadonly = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "EnableReplicaReadonly",
	}
	DisassociateSecurityGroups = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "DisassociateSecurityGroups",
	}
	DisableReplicaReadonly = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "DisableReplicaReadonly",
	}
	DestroyPrepaidInstance = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "DestroyPrepaidInstance",
	}
	DestroyPostpaidInstance = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "DestroyPostpaidInstance",
	}
	DescribeTendisSlowLog = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "DescribeTendisSlowLog",
	}
	DescribeTaskList = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "DescribeTaskList",
	}
	DescribeTaskInfo = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "DescribeTaskInfo",
	}
	DescribeSlowLog = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "DescribeSlowLog",
	}
	DescribeProxySlowLog = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "DescribeProxySlowLog",
	}
	DescribeProjectSecurityGroups = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "DescribeProjectSecurityGroups",
	}
	DescribeProjectSecurityGroup = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "DescribeProjectSecurityGroup",
	}
	DescribeProductInfo = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "DescribeProductInfo",
	}
	DescribeMaintenanceWindow = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "DescribeMaintenanceWindow",
	}
	DescribeInstances = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "DescribeInstances",
	}
	DescribeInstanceZoneInfo = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "DescribeInstanceZoneInfo",
	}
	DescribeInstanceShards = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "DescribeInstanceShards",
	}
	DescribeInstanceSecurityGroup = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "DescribeInstanceSecurityGroup",
	}
	DescribeInstanceParams = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "DescribeInstanceParams",
	}
	DescribeInstanceParamRecords = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "DescribeInstanceParamRecords",
	}
	DescribeInstanceNodeInfo = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "DescribeInstanceNodeInfo",
	}
	DescribeInstanceMonitorTopNCmdTook = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "DescribeInstanceMonitorTopNCmdTook",
	}
	DescribeInstanceMonitorTopNCmd = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "DescribeInstanceMonitorTopNCmd",
	}
	DescribeInstanceMonitorTookDist = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "DescribeInstanceMonitorTookDist",
	}
	DescribeInstanceMonitorSIP = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "DescribeInstanceMonitorSIP",
	}
	DescribeInstanceMonitorHotKey = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "DescribeInstanceMonitorHotKey",
	}
	DescribeInstanceMonitorBigKeyTypeDist = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "DescribeInstanceMonitorBigKeyTypeDist",
	}
	DescribeInstanceMonitorBigKeySizeDist = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "DescribeInstanceMonitorBigKeySizeDist",
	}
	DescribeInstanceMonitorBigKey = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "DescribeInstanceMonitorBigKey",
	}
	DescribeInstanceDealDetail = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "DescribeInstanceDealDetail",
	}
	DescribeInstanceDTSInfo = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "DescribeInstanceDTSInfo",
	}
	DescribeInstanceBackups = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "DescribeInstanceBackups",
	}
	DescribeInstanceAccount = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "DescribeInstanceAccount",
	}
	DescribeDBSecurityGroups = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "DescribeDBSecurityGroups",
	}
	DescribeCommonDBInstances = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "DescribeCommonDBInstances",
	}
	DescribeBackupUrl = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "DescribeBackupUrl",
	}
	DescribeAutoBackupConfig = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "DescribeAutoBackupConfig",
	}
	DeleteInstanceAccount = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "DeleteInstanceAccount",
	}
	CreateInstances = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "CreateInstances",
	}
	CreateInstanceAccount = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "CreateInstanceAccount",
	}
	ClearInstance = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "ClearInstance",
	}
	CleanUpInstance = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "CleanUpInstance",
	}
	AssociateSecurityGroups = tencentcloud.Action{
		Service: "redis",
		Version: "2018-04-12",
		Action:  "AssociateSecurityGroups",
	}
)
