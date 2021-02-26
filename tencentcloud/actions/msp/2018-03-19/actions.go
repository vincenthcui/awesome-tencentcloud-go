package msp

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	RegisterMigrationTask = tencentcloud.Action{
		Service: "msp",
		Version: "2018-03-19",
		Action:  "RegisterMigrationTask",
	}
	ModifyMigrationTaskStatus = tencentcloud.Action{
		Service: "msp",
		Version: "2018-03-19",
		Action:  "ModifyMigrationTaskStatus",
	}
	ModifyMigrationTaskBelongToProject = tencentcloud.Action{
		Service: "msp",
		Version: "2018-03-19",
		Action:  "ModifyMigrationTaskBelongToProject",
	}
	ListMigrationTask = tencentcloud.Action{
		Service: "msp",
		Version: "2018-03-19",
		Action:  "ListMigrationTask",
	}
	ListMigrationProject = tencentcloud.Action{
		Service: "msp",
		Version: "2018-03-19",
		Action:  "ListMigrationProject",
	}
	DescribeMigrationTask = tencentcloud.Action{
		Service: "msp",
		Version: "2018-03-19",
		Action:  "DescribeMigrationTask",
	}
	DeregisterMigrationTask = tencentcloud.Action{
		Service: "msp",
		Version: "2018-03-19",
		Action:  "DeregisterMigrationTask",
	}
)
