package dcdb

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	UpgradeDCDBInstance = tencentcloud.Action{
		Service: "dcdb",
		Version: "2018-04-11",
		Action:  "UpgradeDCDBInstance",
	}
	ResetAccountPassword = tencentcloud.Action{
		Service: "dcdb",
		Version: "2018-04-11",
		Action:  "ResetAccountPassword",
	}
	RenewDCDBInstance = tencentcloud.Action{
		Service: "dcdb",
		Version: "2018-04-11",
		Action:  "RenewDCDBInstance",
	}
	OpenDBExtranetAccess = tencentcloud.Action{
		Service: "dcdb",
		Version: "2018-04-11",
		Action:  "OpenDBExtranetAccess",
	}
	ModifyDBSyncMode = tencentcloud.Action{
		Service: "dcdb",
		Version: "2018-04-11",
		Action:  "ModifyDBSyncMode",
	}
	ModifyDBParameters = tencentcloud.Action{
		Service: "dcdb",
		Version: "2018-04-11",
		Action:  "ModifyDBParameters",
	}
	ModifyDBInstancesProject = tencentcloud.Action{
		Service: "dcdb",
		Version: "2018-04-11",
		Action:  "ModifyDBInstancesProject",
	}
	ModifyDBInstanceSecurityGroups = tencentcloud.Action{
		Service: "dcdb",
		Version: "2018-04-11",
		Action:  "ModifyDBInstanceSecurityGroups",
	}
	ModifyAccountDescription = tencentcloud.Action{
		Service: "dcdb",
		Version: "2018-04-11",
		Action:  "ModifyAccountDescription",
	}
	InitDCDBInstances = tencentcloud.Action{
		Service: "dcdb",
		Version: "2018-04-11",
		Action:  "InitDCDBInstances",
	}
	GrantAccountPrivileges = tencentcloud.Action{
		Service: "dcdb",
		Version: "2018-04-11",
		Action:  "GrantAccountPrivileges",
	}
	FlushBinlog = tencentcloud.Action{
		Service: "dcdb",
		Version: "2018-04-11",
		Action:  "FlushBinlog",
	}
	DisassociateSecurityGroups = tencentcloud.Action{
		Service: "dcdb",
		Version: "2018-04-11",
		Action:  "DisassociateSecurityGroups",
	}
	DescribeUserTasks = tencentcloud.Action{
		Service: "dcdb",
		Version: "2018-04-11",
		Action:  "DescribeUserTasks",
	}
	DescribeSqlLogs = tencentcloud.Action{
		Service: "dcdb",
		Version: "2018-04-11",
		Action:  "DescribeSqlLogs",
	}
	DescribeShardSpec = tencentcloud.Action{
		Service: "dcdb",
		Version: "2018-04-11",
		Action:  "DescribeShardSpec",
	}
	DescribeProjects = tencentcloud.Action{
		Service: "dcdb",
		Version: "2018-04-11",
		Action:  "DescribeProjects",
	}
	DescribeProjectSecurityGroups = tencentcloud.Action{
		Service: "dcdb",
		Version: "2018-04-11",
		Action:  "DescribeProjectSecurityGroups",
	}
	DescribeOrders = tencentcloud.Action{
		Service: "dcdb",
		Version: "2018-04-11",
		Action:  "DescribeOrders",
	}
	DescribeDatabases = tencentcloud.Action{
		Service: "dcdb",
		Version: "2018-04-11",
		Action:  "DescribeDatabases",
	}
	DescribeDatabaseTable = tencentcloud.Action{
		Service: "dcdb",
		Version: "2018-04-11",
		Action:  "DescribeDatabaseTable",
	}
	DescribeDatabaseObjects = tencentcloud.Action{
		Service: "dcdb",
		Version: "2018-04-11",
		Action:  "DescribeDatabaseObjects",
	}
	DescribeDCDBUpgradePrice = tencentcloud.Action{
		Service: "dcdb",
		Version: "2018-04-11",
		Action:  "DescribeDCDBUpgradePrice",
	}
	DescribeDCDBShards = tencentcloud.Action{
		Service: "dcdb",
		Version: "2018-04-11",
		Action:  "DescribeDCDBShards",
	}
	DescribeDCDBSaleInfo = tencentcloud.Action{
		Service: "dcdb",
		Version: "2018-04-11",
		Action:  "DescribeDCDBSaleInfo",
	}
	DescribeDCDBRenewalPrice = tencentcloud.Action{
		Service: "dcdb",
		Version: "2018-04-11",
		Action:  "DescribeDCDBRenewalPrice",
	}
	DescribeDCDBPrice = tencentcloud.Action{
		Service: "dcdb",
		Version: "2018-04-11",
		Action:  "DescribeDCDBPrice",
	}
	DescribeDCDBInstances = tencentcloud.Action{
		Service: "dcdb",
		Version: "2018-04-11",
		Action:  "DescribeDCDBInstances",
	}
	DescribeDBSyncMode = tencentcloud.Action{
		Service: "dcdb",
		Version: "2018-04-11",
		Action:  "DescribeDBSyncMode",
	}
	DescribeDBSecurityGroups = tencentcloud.Action{
		Service: "dcdb",
		Version: "2018-04-11",
		Action:  "DescribeDBSecurityGroups",
	}
	DescribeDBParameters = tencentcloud.Action{
		Service: "dcdb",
		Version: "2018-04-11",
		Action:  "DescribeDBParameters",
	}
	DescribeDBLogFiles = tencentcloud.Action{
		Service: "dcdb",
		Version: "2018-04-11",
		Action:  "DescribeDBLogFiles",
	}
	DescribeAccounts = tencentcloud.Action{
		Service: "dcdb",
		Version: "2018-04-11",
		Action:  "DescribeAccounts",
	}
	DescribeAccountPrivileges = tencentcloud.Action{
		Service: "dcdb",
		Version: "2018-04-11",
		Action:  "DescribeAccountPrivileges",
	}
	DeleteAccount = tencentcloud.Action{
		Service: "dcdb",
		Version: "2018-04-11",
		Action:  "DeleteAccount",
	}
	CreateDCDBInstance = tencentcloud.Action{
		Service: "dcdb",
		Version: "2018-04-11",
		Action:  "CreateDCDBInstance",
	}
	CreateAccount = tencentcloud.Action{
		Service: "dcdb",
		Version: "2018-04-11",
		Action:  "CreateAccount",
	}
	CopyAccountPrivileges = tencentcloud.Action{
		Service: "dcdb",
		Version: "2018-04-11",
		Action:  "CopyAccountPrivileges",
	}
	CloseDBExtranetAccess = tencentcloud.Action{
		Service: "dcdb",
		Version: "2018-04-11",
		Action:  "CloseDBExtranetAccess",
	}
	CloneAccount = tencentcloud.Action{
		Service: "dcdb",
		Version: "2018-04-11",
		Action:  "CloneAccount",
	}
	AssociateSecurityGroups = tencentcloud.Action{
		Service: "dcdb",
		Version: "2018-04-11",
		Action:  "AssociateSecurityGroups",
	}
)
