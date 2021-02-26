package tic

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	UpdateStackVersion = tencentcloud.Action{
		Service: "tic",
		Version: "2020-11-17",
		Action:  "UpdateStackVersion",
	}
	UpdateStack = tencentcloud.Action{
		Service: "tic",
		Version: "2020-11-17",
		Action:  "UpdateStack",
	}
	PlanStack = tencentcloud.Action{
		Service: "tic",
		Version: "2020-11-17",
		Action:  "PlanStack",
	}
	DestroyStack = tencentcloud.Action{
		Service: "tic",
		Version: "2020-11-17",
		Action:  "DestroyStack",
	}
	DescribeStacks = tencentcloud.Action{
		Service: "tic",
		Version: "2020-11-17",
		Action:  "DescribeStacks",
	}
	DescribeStackVersions = tencentcloud.Action{
		Service: "tic",
		Version: "2020-11-17",
		Action:  "DescribeStackVersions",
	}
	DescribeStackEvents = tencentcloud.Action{
		Service: "tic",
		Version: "2020-11-17",
		Action:  "DescribeStackEvents",
	}
	DescribeStackEvent = tencentcloud.Action{
		Service: "tic",
		Version: "2020-11-17",
		Action:  "DescribeStackEvent",
	}
	DeleteStackVersion = tencentcloud.Action{
		Service: "tic",
		Version: "2020-11-17",
		Action:  "DeleteStackVersion",
	}
	DeleteStack = tencentcloud.Action{
		Service: "tic",
		Version: "2020-11-17",
		Action:  "DeleteStack",
	}
	CreateStackVersion = tencentcloud.Action{
		Service: "tic",
		Version: "2020-11-17",
		Action:  "CreateStackVersion",
	}
	CreateStack = tencentcloud.Action{
		Service: "tic",
		Version: "2020-11-17",
		Action:  "CreateStack",
	}
	ApplyStack = tencentcloud.Action{
		Service: "tic",
		Version: "2020-11-17",
		Action:  "ApplyStack",
	}
)
