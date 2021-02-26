package tia

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	QueryLogs = tencentcloud.Action{
		Service: "tia",
		Version: "2018-02-26",
		Action:  "QueryLogs",
	}
	ListModels = tencentcloud.Action{
		Service: "tia",
		Version: "2018-02-26",
		Action:  "ListModels",
	}
	ListJobs = tencentcloud.Action{
		Service: "tia",
		Version: "2018-02-26",
		Action:  "ListJobs",
	}
	InstallAgent = tencentcloud.Action{
		Service: "tia",
		Version: "2018-02-26",
		Action:  "InstallAgent",
	}
	DescribeModel = tencentcloud.Action{
		Service: "tia",
		Version: "2018-02-26",
		Action:  "DescribeModel",
	}
	DescribeJob = tencentcloud.Action{
		Service: "tia",
		Version: "2018-02-26",
		Action:  "DescribeJob",
	}
	DeleteModel = tencentcloud.Action{
		Service: "tia",
		Version: "2018-02-26",
		Action:  "DeleteModel",
	}
	DeleteJob = tencentcloud.Action{
		Service: "tia",
		Version: "2018-02-26",
		Action:  "DeleteJob",
	}
	CreateModel = tencentcloud.Action{
		Service: "tia",
		Version: "2018-02-26",
		Action:  "CreateModel",
	}
	CreateJob = tencentcloud.Action{
		Service: "tia",
		Version: "2018-02-26",
		Action:  "CreateJob",
	}
)
