package mariadb

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	UpgradeDBInstance = tencentcloud.Action{
		Service: "mariadb",
		Version: "2017-03-12",
		Action:  "UpgradeDBInstance",
	}
	RestartDBInstances = tencentcloud.Action{
		Service: "mariadb",
		Version: "2017-03-12",
		Action:  "RestartDBInstances",
	}
	ResetAccountPassword = tencentcloud.Action{
		Service: "mariadb",
		Version: "2017-03-12",
		Action:  "ResetAccountPassword",
	}
	RenewDBInstance = tencentcloud.Action{
		Service: "mariadb",
		Version: "2017-03-12",
		Action:  "RenewDBInstance",
	}
	OpenDBExtranetAccess = tencentcloud.Action{
		Service: "mariadb",
		Version: "2017-03-12",
		Action:  "OpenDBExtranetAccess",
	}
	ModifyLogFileRetentionPeriod = tencentcloud.Action{
		Service: "mariadb",
		Version: "2017-03-12",
		Action:  "ModifyLogFileRetentionPeriod",
	}
	ModifyDBParameters = tencentcloud.Action{
		Service: "mariadb",
		Version: "2017-03-12",
		Action:  "ModifyDBParameters",
	}
	ModifyDBInstancesProject = tencentcloud.Action{
		Service: "mariadb",
		Version: "2017-03-12",
		Action:  "ModifyDBInstancesProject",
	}
	ModifyDBInstanceSecurityGroups = tencentcloud.Action{
		Service: "mariadb",
		Version: "2017-03-12",
		Action:  "ModifyDBInstanceSecurityGroups",
	}
	ModifyDBInstanceName = tencentcloud.Action{
		Service: "mariadb",
		Version: "2017-03-12",
		Action:  "ModifyDBInstanceName",
	}
	ModifyBackupTime = tencentcloud.Action{
		Service: "mariadb",
		Version: "2017-03-12",
		Action:  "ModifyBackupTime",
	}
	ModifyAccountDescription = tencentcloud.Action{
		Service: "mariadb",
		Version: "2017-03-12",
		Action:  "ModifyAccountDescription",
	}
	InitDBInstances = tencentcloud.Action{
		Service: "mariadb",
		Version: "2017-03-12",
		Action:  "InitDBInstances",
	}
	GrantAccountPrivileges = tencentcloud.Action{
		Service: "mariadb",
		Version: "2017-03-12",
		Action:  "GrantAccountPrivileges",
	}
	FlushBinlog = tencentcloud.Action{
		Service: "mariadb",
		Version: "2017-03-12",
		Action:  "FlushBinlog",
	}
	DisassociateSecurityGroups = tencentcloud.Action{
		Service: "mariadb",
		Version: "2017-03-12",
		Action:  "DisassociateSecurityGroups",
	}
	DescribeUpgradePrice = tencentcloud.Action{
		Service: "mariadb",
		Version: "2017-03-12",
		Action:  "DescribeUpgradePrice",
	}
	DescribeSqlLogs = tencentcloud.Action{
		Service: "mariadb",
		Version: "2017-03-12",
		Action:  "DescribeSqlLogs",
	}
	DescribeSaleInfo = tencentcloud.Action{
		Service: "mariadb",
		Version: "2017-03-12",
		Action:  "DescribeSaleInfo",
	}
	DescribeRenewalPrice = tencentcloud.Action{
		Service: "mariadb",
		Version: "2017-03-12",
		Action:  "DescribeRenewalPrice",
	}
	DescribeProjectSecurityGroups = tencentcloud.Action{
		Service: "mariadb",
		Version: "2017-03-12",
		Action:  "DescribeProjectSecurityGroups",
	}
	DescribePrice = tencentcloud.Action{
		Service: "mariadb",
		Version: "2017-03-12",
		Action:  "DescribePrice",
	}
	DescribeOrders = tencentcloud.Action{
		Service: "mariadb",
		Version: "2017-03-12",
		Action:  "DescribeOrders",
	}
	DescribeLogFileRetentionPeriod = tencentcloud.Action{
		Service: "mariadb",
		Version: "2017-03-12",
		Action:  "DescribeLogFileRetentionPeriod",
	}
	DescribeFlow = tencentcloud.Action{
		Service: "mariadb",
		Version: "2017-03-12",
		Action:  "DescribeFlow",
	}
	DescribeDatabases = tencentcloud.Action{
		Service: "mariadb",
		Version: "2017-03-12",
		Action:  "DescribeDatabases",
	}
	DescribeDBSlowLogs = tencentcloud.Action{
		Service: "mariadb",
		Version: "2017-03-12",
		Action:  "DescribeDBSlowLogs",
	}
	DescribeDBSecurityGroups = tencentcloud.Action{
		Service: "mariadb",
		Version: "2017-03-12",
		Action:  "DescribeDBSecurityGroups",
	}
	DescribeDBResourceUsageDetails = tencentcloud.Action{
		Service: "mariadb",
		Version: "2017-03-12",
		Action:  "DescribeDBResourceUsageDetails",
	}
	DescribeDBResourceUsage = tencentcloud.Action{
		Service: "mariadb",
		Version: "2017-03-12",
		Action:  "DescribeDBResourceUsage",
	}
	DescribeDBPerformanceDetails = tencentcloud.Action{
		Service: "mariadb",
		Version: "2017-03-12",
		Action:  "DescribeDBPerformanceDetails",
	}
	DescribeDBPerformance = tencentcloud.Action{
		Service: "mariadb",
		Version: "2017-03-12",
		Action:  "DescribeDBPerformance",
	}
	DescribeDBParameters = tencentcloud.Action{
		Service: "mariadb",
		Version: "2017-03-12",
		Action:  "DescribeDBParameters",
	}
	DescribeDBLogFiles = tencentcloud.Action{
		Service: "mariadb",
		Version: "2017-03-12",
		Action:  "DescribeDBLogFiles",
	}
	DescribeDBInstances = tencentcloud.Action{
		Service: "mariadb",
		Version: "2017-03-12",
		Action:  "DescribeDBInstances",
	}
	DescribeDBInstanceSpecs = tencentcloud.Action{
		Service: "mariadb",
		Version: "2017-03-12",
		Action:  "DescribeDBInstanceSpecs",
	}
	DescribeBackupTime = tencentcloud.Action{
		Service: "mariadb",
		Version: "2017-03-12",
		Action:  "DescribeBackupTime",
	}
	DescribeAccounts = tencentcloud.Action{
		Service: "mariadb",
		Version: "2017-03-12",
		Action:  "DescribeAccounts",
	}
	DescribeAccountPrivileges = tencentcloud.Action{
		Service: "mariadb",
		Version: "2017-03-12",
		Action:  "DescribeAccountPrivileges",
	}
	DeleteAccount = tencentcloud.Action{
		Service: "mariadb",
		Version: "2017-03-12",
		Action:  "DeleteAccount",
	}
	CreateTmpInstances = tencentcloud.Action{
		Service: "mariadb",
		Version: "2017-03-12",
		Action:  "CreateTmpInstances",
	}
	CreateDBInstance = tencentcloud.Action{
		Service: "mariadb",
		Version: "2017-03-12",
		Action:  "CreateDBInstance",
	}
	CreateAccount = tencentcloud.Action{
		Service: "mariadb",
		Version: "2017-03-12",
		Action:  "CreateAccount",
	}
	CopyAccountPrivileges = tencentcloud.Action{
		Service: "mariadb",
		Version: "2017-03-12",
		Action:  "CopyAccountPrivileges",
	}
	CloseDBExtranetAccess = tencentcloud.Action{
		Service: "mariadb",
		Version: "2017-03-12",
		Action:  "CloseDBExtranetAccess",
	}
	CloneAccount = tencentcloud.Action{
		Service: "mariadb",
		Version: "2017-03-12",
		Action:  "CloneAccount",
	}
	AssociateSecurityGroups = tencentcloud.Action{
		Service: "mariadb",
		Version: "2017-03-12",
		Action:  "AssociateSecurityGroups",
	}
)
