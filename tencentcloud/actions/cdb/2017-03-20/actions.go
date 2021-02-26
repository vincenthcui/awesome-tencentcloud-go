package cdb

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	VerifyRootAccount = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "VerifyRootAccount",
	}
	UpgradeDBInstanceEngineVersion = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "UpgradeDBInstanceEngineVersion",
	}
	UpgradeDBInstance = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "UpgradeDBInstance",
	}
	SwitchForUpgrade = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "SwitchForUpgrade",
	}
	SwitchDrInstanceToMaster = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "SwitchDrInstanceToMaster",
	}
	SwitchDBInstanceMasterSlave = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "SwitchDBInstanceMasterSlave",
	}
	StopRollback = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "StopRollback",
	}
	StopDelayReplication = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "StopDelayReplication",
	}
	StopDBImportJob = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "StopDBImportJob",
	}
	StartDelayReplication = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "StartDelayReplication",
	}
	StartBatchRollback = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "StartBatchRollback",
	}
	RestartDBInstances = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "RestartDBInstances",
	}
	RenewDBInstance = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "RenewDBInstance",
	}
	ReleaseIsolatedDBInstances = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "ReleaseIsolatedDBInstances",
	}
	OpenWanService = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "OpenWanService",
	}
	OpenDBInstanceGTID = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "OpenDBInstanceGTID",
	}
	OfflineIsolatedInstances = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "OfflineIsolatedInstances",
	}
	ModifyTimeWindow = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "ModifyTimeWindow",
	}
	ModifyRoType = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "ModifyRoType",
	}
	ModifyRoReplicationDelay = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "ModifyRoReplicationDelay",
	}
	ModifyRoGroupInfo = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "ModifyRoGroupInfo",
	}
	ModifyParamTemplate = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "ModifyParamTemplate",
	}
	ModifyNameOrDescByDpId = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "ModifyNameOrDescByDpId",
	}
	ModifyInstanceTag = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "ModifyInstanceTag",
	}
	ModifyInstanceParam = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "ModifyInstanceParam",
	}
	ModifyDBInstanceVipVport = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "ModifyDBInstanceVipVport",
	}
	ModifyDBInstanceSecurityGroups = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "ModifyDBInstanceSecurityGroups",
	}
	ModifyDBInstanceProject = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "ModifyDBInstanceProject",
	}
	ModifyDBInstanceName = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "ModifyDBInstanceName",
	}
	ModifyBackupConfig = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "ModifyBackupConfig",
	}
	ModifyAutoRenewFlag = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "ModifyAutoRenewFlag",
	}
	ModifyAuditRule = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "ModifyAuditRule",
	}
	ModifyAuditConfig = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "ModifyAuditConfig",
	}
	ModifyAccountPrivileges = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "ModifyAccountPrivileges",
	}
	ModifyAccountPassword = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "ModifyAccountPassword",
	}
	ModifyAccountDescription = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "ModifyAccountDescription",
	}
	IsolateDBInstance = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "IsolateDBInstance",
	}
	InquiryPriceUpgradeInstances = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "InquiryPriceUpgradeInstances",
	}
	InitDBInstances = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "InitDBInstances",
	}
	DisassociateSecurityGroups = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DisassociateSecurityGroups",
	}
	DescribeUploadedFiles = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeUploadedFiles",
	}
	DescribeTimeWindow = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeTimeWindow",
	}
	DescribeTasks = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeTasks",
	}
	DescribeTagsOfInstanceIds = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeTagsOfInstanceIds",
	}
	DescribeTables = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeTables",
	}
	DescribeSupportedPrivileges = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeSupportedPrivileges",
	}
	DescribeSlowLogs = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeSlowLogs",
	}
	DescribeSlowLogData = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeSlowLogData",
	}
	DescribeRollbackTaskDetail = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeRollbackTaskDetail",
	}
	DescribeRollbackRangeTime = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeRollbackRangeTime",
	}
	DescribeRoMinScale = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeRoMinScale",
	}
	DescribeRoGroups = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeRoGroups",
	}
	DescribeProjectSecurityGroups = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeProjectSecurityGroups",
	}
	DescribeParamTemplates = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeParamTemplates",
	}
	DescribeParamTemplateInfo = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeParamTemplateInfo",
	}
	DescribeInstanceParams = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeInstanceParams",
	}
	DescribeInstanceParamRecords = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeInstanceParamRecords",
	}
	DescribeErrorLogData = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeErrorLogData",
	}
	DescribeDeviceMonitorInfo = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeDeviceMonitorInfo",
	}
	DescribeDeployGroupList = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeDeployGroupList",
	}
	DescribeDefaultParams = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeDefaultParams",
	}
	DescribeDatabases = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeDatabases",
	}
	DescribeDataBackupOverview = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeDataBackupOverview",
	}
	DescribeDBZoneConfig = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeDBZoneConfig",
	}
	DescribeDBSwitchRecords = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeDBSwitchRecords",
	}
	DescribeDBSecurityGroups = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeDBSecurityGroups",
	}
	DescribeDBPrice = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeDBPrice",
	}
	DescribeDBInstances = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeDBInstances",
	}
	DescribeDBInstanceRebootTime = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeDBInstanceRebootTime",
	}
	DescribeDBInstanceInfo = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeDBInstanceInfo",
	}
	DescribeDBInstanceGTID = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeDBInstanceGTID",
	}
	DescribeDBInstanceConfig = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeDBInstanceConfig",
	}
	DescribeDBInstanceCharset = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeDBInstanceCharset",
	}
	DescribeDBImportRecords = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeDBImportRecords",
	}
	DescribeCloneList = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeCloneList",
	}
	DescribeBinlogs = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeBinlogs",
	}
	DescribeBinlogBackupOverview = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeBinlogBackupOverview",
	}
	DescribeBackups = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeBackups",
	}
	DescribeBackupTables = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeBackupTables",
	}
	DescribeBackupSummaries = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeBackupSummaries",
	}
	DescribeBackupOverview = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeBackupOverview",
	}
	DescribeBackupDatabases = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeBackupDatabases",
	}
	DescribeBackupConfig = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeBackupConfig",
	}
	DescribeAuditRules = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeAuditRules",
	}
	DescribeAuditPolicies = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeAuditPolicies",
	}
	DescribeAuditLogFiles = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeAuditLogFiles",
	}
	DescribeAuditConfig = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeAuditConfig",
	}
	DescribeAsyncRequestInfo = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeAsyncRequestInfo",
	}
	DescribeAccounts = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeAccounts",
	}
	DescribeAccountPrivileges = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DescribeAccountPrivileges",
	}
	DeleteTimeWindow = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DeleteTimeWindow",
	}
	DeleteParamTemplate = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DeleteParamTemplate",
	}
	DeleteDeployGroups = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DeleteDeployGroups",
	}
	DeleteBackup = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DeleteBackup",
	}
	DeleteAuditRule = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DeleteAuditRule",
	}
	DeleteAuditPolicy = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DeleteAuditPolicy",
	}
	DeleteAuditLogFile = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DeleteAuditLogFile",
	}
	DeleteAccounts = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "DeleteAccounts",
	}
	CreateRoInstanceIp = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "CreateRoInstanceIp",
	}
	CreateParamTemplate = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "CreateParamTemplate",
	}
	CreateDeployGroup = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "CreateDeployGroup",
	}
	CreateDBInstanceHour = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "CreateDBInstanceHour",
	}
	CreateDBInstance = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "CreateDBInstance",
	}
	CreateDBImportJob = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "CreateDBImportJob",
	}
	CreateCloneInstance = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "CreateCloneInstance",
	}
	CreateBackup = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "CreateBackup",
	}
	CreateAuditRule = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "CreateAuditRule",
	}
	CreateAuditPolicy = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "CreateAuditPolicy",
	}
	CreateAuditLogFile = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "CreateAuditLogFile",
	}
	CreateAccounts = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "CreateAccounts",
	}
	CloseWanService = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "CloseWanService",
	}
	BalanceRoGroupLoad = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "BalanceRoGroupLoad",
	}
	AssociateSecurityGroups = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "AssociateSecurityGroups",
	}
	AddTimeWindow = tencentcloud.Action{
		Service: "cdb",
		Version: "2017-03-20",
		Action:  "AddTimeWindow",
	}
)
