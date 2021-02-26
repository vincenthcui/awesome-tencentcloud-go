package oceanus

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	StopJobs = tencentcloud.Action{
		Service: "oceanus",
		Version: "2019-04-22",
		Action:  "StopJobs",
	}
	RunJobs = tencentcloud.Action{
		Service: "oceanus",
		Version: "2019-04-22",
		Action:  "RunJobs",
	}
	DescribeSystemResources = tencentcloud.Action{
		Service: "oceanus",
		Version: "2019-04-22",
		Action:  "DescribeSystemResources",
	}
	DescribeJobs = tencentcloud.Action{
		Service: "oceanus",
		Version: "2019-04-22",
		Action:  "DescribeJobs",
	}
	DescribeJobConfigs = tencentcloud.Action{
		Service: "oceanus",
		Version: "2019-04-22",
		Action:  "DescribeJobConfigs",
	}
	DeleteTableConfig = tencentcloud.Action{
		Service: "oceanus",
		Version: "2019-04-22",
		Action:  "DeleteTableConfig",
	}
	CreateResourceConfig = tencentcloud.Action{
		Service: "oceanus",
		Version: "2019-04-22",
		Action:  "CreateResourceConfig",
	}
	CreateResource = tencentcloud.Action{
		Service: "oceanus",
		Version: "2019-04-22",
		Action:  "CreateResource",
	}
	CreateJobConfig = tencentcloud.Action{
		Service: "oceanus",
		Version: "2019-04-22",
		Action:  "CreateJobConfig",
	}
	CreateJob = tencentcloud.Action{
		Service: "oceanus",
		Version: "2019-04-22",
		Action:  "CreateJob",
	}
)
