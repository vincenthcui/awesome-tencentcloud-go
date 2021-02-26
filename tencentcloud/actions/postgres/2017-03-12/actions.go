package postgres

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	UpgradeDBInstance = tencentcloud.Action{
		Service: "postgres",
		Version: "2017-03-12",
		Action:  "UpgradeDBInstance",
	}
	SetAutoRenewFlag = tencentcloud.Action{
		Service: "postgres",
		Version: "2017-03-12",
		Action:  "SetAutoRenewFlag",
	}
	RestartDBInstance = tencentcloud.Action{
		Service: "postgres",
		Version: "2017-03-12",
		Action:  "RestartDBInstance",
	}
	ResetAccountPassword = tencentcloud.Action{
		Service: "postgres",
		Version: "2017-03-12",
		Action:  "ResetAccountPassword",
	}
	RenewInstance = tencentcloud.Action{
		Service: "postgres",
		Version: "2017-03-12",
		Action:  "RenewInstance",
	}
	RemoveDBInstanceFromReadOnlyGroup = tencentcloud.Action{
		Service: "postgres",
		Version: "2017-03-12",
		Action:  "RemoveDBInstanceFromReadOnlyGroup",
	}
	RebalanceReadOnlyGroup = tencentcloud.Action{
		Service: "postgres",
		Version: "2017-03-12",
		Action:  "RebalanceReadOnlyGroup",
	}
	OpenServerlessDBExtranetAccess = tencentcloud.Action{
		Service: "postgres",
		Version: "2017-03-12",
		Action:  "OpenServerlessDBExtranetAccess",
	}
	OpenDBExtranetAccess = tencentcloud.Action{
		Service: "postgres",
		Version: "2017-03-12",
		Action:  "OpenDBExtranetAccess",
	}
	ModifyReadOnlyGroupConfig = tencentcloud.Action{
		Service: "postgres",
		Version: "2017-03-12",
		Action:  "ModifyReadOnlyGroupConfig",
	}
	ModifyDBInstancesProject = tencentcloud.Action{
		Service: "postgres",
		Version: "2017-03-12",
		Action:  "ModifyDBInstancesProject",
	}
	ModifyDBInstanceReadOnlyGroup = tencentcloud.Action{
		Service: "postgres",
		Version: "2017-03-12",
		Action:  "ModifyDBInstanceReadOnlyGroup",
	}
	ModifyDBInstanceName = tencentcloud.Action{
		Service: "postgres",
		Version: "2017-03-12",
		Action:  "ModifyDBInstanceName",
	}
	ModifyAccountRemark = tencentcloud.Action{
		Service: "postgres",
		Version: "2017-03-12",
		Action:  "ModifyAccountRemark",
	}
	InquiryPriceUpgradeDBInstance = tencentcloud.Action{
		Service: "postgres",
		Version: "2017-03-12",
		Action:  "InquiryPriceUpgradeDBInstance",
	}
	InquiryPriceRenewDBInstance = tencentcloud.Action{
		Service: "postgres",
		Version: "2017-03-12",
		Action:  "InquiryPriceRenewDBInstance",
	}
	InquiryPriceCreateDBInstances = tencentcloud.Action{
		Service: "postgres",
		Version: "2017-03-12",
		Action:  "InquiryPriceCreateDBInstances",
	}
	InitDBInstances = tencentcloud.Action{
		Service: "postgres",
		Version: "2017-03-12",
		Action:  "InitDBInstances",
	}
	DestroyDBInstance = tencentcloud.Action{
		Service: "postgres",
		Version: "2017-03-12",
		Action:  "DestroyDBInstance",
	}
	DescribeZones = tencentcloud.Action{
		Service: "postgres",
		Version: "2017-03-12",
		Action:  "DescribeZones",
	}
	DescribeServerlessDBInstances = tencentcloud.Action{
		Service: "postgres",
		Version: "2017-03-12",
		Action:  "DescribeServerlessDBInstances",
	}
	DescribeRegions = tencentcloud.Action{
		Service: "postgres",
		Version: "2017-03-12",
		Action:  "DescribeRegions",
	}
	DescribeReadOnlyGroups = tencentcloud.Action{
		Service: "postgres",
		Version: "2017-03-12",
		Action:  "DescribeReadOnlyGroups",
	}
	DescribeProductConfig = tencentcloud.Action{
		Service: "postgres",
		Version: "2017-03-12",
		Action:  "DescribeProductConfig",
	}
	DescribeOrders = tencentcloud.Action{
		Service: "postgres",
		Version: "2017-03-12",
		Action:  "DescribeOrders",
	}
	DescribeDatabases = tencentcloud.Action{
		Service: "postgres",
		Version: "2017-03-12",
		Action:  "DescribeDatabases",
	}
	DescribeDBXlogs = tencentcloud.Action{
		Service: "postgres",
		Version: "2017-03-12",
		Action:  "DescribeDBXlogs",
	}
	DescribeDBSlowlogs = tencentcloud.Action{
		Service: "postgres",
		Version: "2017-03-12",
		Action:  "DescribeDBSlowlogs",
	}
	DescribeDBInstances = tencentcloud.Action{
		Service: "postgres",
		Version: "2017-03-12",
		Action:  "DescribeDBInstances",
	}
	DescribeDBInstanceAttribute = tencentcloud.Action{
		Service: "postgres",
		Version: "2017-03-12",
		Action:  "DescribeDBInstanceAttribute",
	}
	DescribeDBErrlogs = tencentcloud.Action{
		Service: "postgres",
		Version: "2017-03-12",
		Action:  "DescribeDBErrlogs",
	}
	DescribeDBBackups = tencentcloud.Action{
		Service: "postgres",
		Version: "2017-03-12",
		Action:  "DescribeDBBackups",
	}
	DescribeAccounts = tencentcloud.Action{
		Service: "postgres",
		Version: "2017-03-12",
		Action:  "DescribeAccounts",
	}
	DeleteServerlessDBInstance = tencentcloud.Action{
		Service: "postgres",
		Version: "2017-03-12",
		Action:  "DeleteServerlessDBInstance",
	}
	DeleteReadOnlyGroup = tencentcloud.Action{
		Service: "postgres",
		Version: "2017-03-12",
		Action:  "DeleteReadOnlyGroup",
	}
	CreateServerlessDBInstance = tencentcloud.Action{
		Service: "postgres",
		Version: "2017-03-12",
		Action:  "CreateServerlessDBInstance",
	}
	CreateReadOnlyGroup = tencentcloud.Action{
		Service: "postgres",
		Version: "2017-03-12",
		Action:  "CreateReadOnlyGroup",
	}
	CreateReadOnlyDBInstance = tencentcloud.Action{
		Service: "postgres",
		Version: "2017-03-12",
		Action:  "CreateReadOnlyDBInstance",
	}
	CreateDBInstances = tencentcloud.Action{
		Service: "postgres",
		Version: "2017-03-12",
		Action:  "CreateDBInstances",
	}
	CloseServerlessDBExtranetAccess = tencentcloud.Action{
		Service: "postgres",
		Version: "2017-03-12",
		Action:  "CloseServerlessDBExtranetAccess",
	}
	CloseDBExtranetAccess = tencentcloud.Action{
		Service: "postgres",
		Version: "2017-03-12",
		Action:  "CloseDBExtranetAccess",
	}
	AddDBInstanceToReadOnlyGroup = tencentcloud.Action{
		Service: "postgres",
		Version: "2017-03-12",
		Action:  "AddDBInstanceToReadOnlyGroup",
	}
)
