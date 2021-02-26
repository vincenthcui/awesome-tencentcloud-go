package sqlserver

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	UpgradeDBInstance = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "UpgradeDBInstance",
	}
	TerminateDBInstance = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "TerminateDBInstance",
	}
	StopMigration = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "StopMigration",
	}
	StartMigrationCheck = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "StartMigrationCheck",
	}
	StartIncrementalMigration = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "StartIncrementalMigration",
	}
	StartBackupMigration = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "StartBackupMigration",
	}
	RunMigration = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "RunMigration",
	}
	RollbackInstance = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "RollbackInstance",
	}
	RestoreInstance = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "RestoreInstance",
	}
	RestartDBInstance = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "RestartDBInstance",
	}
	ResetAccountPassword = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "ResetAccountPassword",
	}
	RenewPostpaidDBInstance = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "RenewPostpaidDBInstance",
	}
	RenewDBInstance = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "RenewDBInstance",
	}
	RemoveBackups = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "RemoveBackups",
	}
	RecycleReadOnlyGroup = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "RecycleReadOnlyGroup",
	}
	RecycleDBInstance = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "RecycleDBInstance",
	}
	QueryMigrationCheckProcess = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "QueryMigrationCheckProcess",
	}
	ModifyReadOnlyGroupDetails = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "ModifyReadOnlyGroupDetails",
	}
	ModifyPublishSubscribeName = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "ModifyPublishSubscribeName",
	}
	ModifyMigration = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "ModifyMigration",
	}
	ModifyMaintenanceSpan = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "ModifyMaintenanceSpan",
	}
	ModifyIncrementalMigration = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "ModifyIncrementalMigration",
	}
	ModifyDBRemark = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "ModifyDBRemark",
	}
	ModifyDBName = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "ModifyDBName",
	}
	ModifyDBInstanceSecurityGroups = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "ModifyDBInstanceSecurityGroups",
	}
	ModifyDBInstanceRenewFlag = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "ModifyDBInstanceRenewFlag",
	}
	ModifyDBInstanceProject = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "ModifyDBInstanceProject",
	}
	ModifyDBInstanceNetwork = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "ModifyDBInstanceNetwork",
	}
	ModifyDBInstanceName = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "ModifyDBInstanceName",
	}
	ModifyBackupStrategy = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "ModifyBackupStrategy",
	}
	ModifyBackupName = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "ModifyBackupName",
	}
	ModifyBackupMigration = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "ModifyBackupMigration",
	}
	ModifyAccountRemark = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "ModifyAccountRemark",
	}
	ModifyAccountPrivilege = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "ModifyAccountPrivilege",
	}
	InquiryPriceUpgradeDBInstance = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "InquiryPriceUpgradeDBInstance",
	}
	InquiryPriceRenewDBInstance = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "InquiryPriceRenewDBInstance",
	}
	InquiryPriceCreateDBInstances = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "InquiryPriceCreateDBInstances",
	}
	DisassociateSecurityGroups = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "DisassociateSecurityGroups",
	}
	DescribeZones = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "DescribeZones",
	}
	DescribeUploadIncrementalInfo = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "DescribeUploadIncrementalInfo",
	}
	DescribeUploadBackupInfo = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "DescribeUploadBackupInfo",
	}
	DescribeSlowlogs = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "DescribeSlowlogs",
	}
	DescribeRollbackTime = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "DescribeRollbackTime",
	}
	DescribeRegions = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "DescribeRegions",
	}
	DescribeReadOnlyGroupList = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "DescribeReadOnlyGroupList",
	}
	DescribeReadOnlyGroupDetails = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "DescribeReadOnlyGroupDetails",
	}
	DescribeReadOnlyGroupByReadOnlyInstance = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "DescribeReadOnlyGroupByReadOnlyInstance",
	}
	DescribePublishSubscribe = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "DescribePublishSubscribe",
	}
	DescribeProjectSecurityGroups = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "DescribeProjectSecurityGroups",
	}
	DescribeProductConfig = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "DescribeProductConfig",
	}
	DescribeOrders = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "DescribeOrders",
	}
	DescribeMigrations = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "DescribeMigrations",
	}
	DescribeMigrationDetail = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "DescribeMigrationDetail",
	}
	DescribeMigrationDatabases = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "DescribeMigrationDatabases",
	}
	DescribeMaintenanceSpan = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "DescribeMaintenanceSpan",
	}
	DescribeIncrementalMigration = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "DescribeIncrementalMigration",
	}
	DescribeFlowStatus = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "DescribeFlowStatus",
	}
	DescribeDBs = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "DescribeDBs",
	}
	DescribeDBSecurityGroups = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "DescribeDBSecurityGroups",
	}
	DescribeDBInstances = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "DescribeDBInstances",
	}
	DescribeCrossRegionZone = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "DescribeCrossRegionZone",
	}
	DescribeBackups = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "DescribeBackups",
	}
	DescribeBackupUploadSize = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "DescribeBackupUploadSize",
	}
	DescribeBackupMigration = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "DescribeBackupMigration",
	}
	DescribeBackupCommand = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "DescribeBackupCommand",
	}
	DescribeBackupByFlowId = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "DescribeBackupByFlowId",
	}
	DescribeAccounts = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "DescribeAccounts",
	}
	DeletePublishSubscribe = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "DeletePublishSubscribe",
	}
	DeleteMigration = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "DeleteMigration",
	}
	DeleteIncrementalMigration = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "DeleteIncrementalMigration",
	}
	DeleteDBInstance = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "DeleteDBInstance",
	}
	DeleteDB = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "DeleteDB",
	}
	DeleteBackupMigration = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "DeleteBackupMigration",
	}
	DeleteAccount = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "DeleteAccount",
	}
	CreateReadOnlyDBInstances = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "CreateReadOnlyDBInstances",
	}
	CreatePublishSubscribe = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "CreatePublishSubscribe",
	}
	CreateMigration = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "CreateMigration",
	}
	CreateIncrementalMigration = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "CreateIncrementalMigration",
	}
	CreateDBInstances = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "CreateDBInstances",
	}
	CreateDB = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "CreateDB",
	}
	CreateBasicDBInstances = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "CreateBasicDBInstances",
	}
	CreateBackupMigration = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "CreateBackupMigration",
	}
	CreateBackup = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "CreateBackup",
	}
	CreateAccount = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "CreateAccount",
	}
	CompleteMigration = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "CompleteMigration",
	}
	CompleteExpansion = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "CompleteExpansion",
	}
	CloneDB = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "CloneDB",
	}
	AssociateSecurityGroups = tencentcloud.Action{
		Service: "sqlserver",
		Version: "2018-03-28",
		Action:  "AssociateSecurityGroups",
	}
)
