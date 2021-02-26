package lighthouse

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	StopInstances = tencentcloud.Action{
		Service: "lighthouse",
		Version: "2020-03-24",
		Action:  "StopInstances",
	}
	StartInstances = tencentcloud.Action{
		Service: "lighthouse",
		Version: "2020-03-24",
		Action:  "StartInstances",
	}
	ResetInstance = tencentcloud.Action{
		Service: "lighthouse",
		Version: "2020-03-24",
		Action:  "ResetInstance",
	}
	RebootInstances = tencentcloud.Action{
		Service: "lighthouse",
		Version: "2020-03-24",
		Action:  "RebootInstances",
	}
	DescribeInstancesTrafficPackages = tencentcloud.Action{
		Service: "lighthouse",
		Version: "2020-03-24",
		Action:  "DescribeInstancesTrafficPackages",
	}
	DescribeInstances = tencentcloud.Action{
		Service: "lighthouse",
		Version: "2020-03-24",
		Action:  "DescribeInstances",
	}
	DescribeFirewallRules = tencentcloud.Action{
		Service: "lighthouse",
		Version: "2020-03-24",
		Action:  "DescribeFirewallRules",
	}
	DescribeBundles = tencentcloud.Action{
		Service: "lighthouse",
		Version: "2020-03-24",
		Action:  "DescribeBundles",
	}
	DescribeBlueprints = tencentcloud.Action{
		Service: "lighthouse",
		Version: "2020-03-24",
		Action:  "DescribeBlueprints",
	}
	DeleteFirewallRules = tencentcloud.Action{
		Service: "lighthouse",
		Version: "2020-03-24",
		Action:  "DeleteFirewallRules",
	}
	CreateFirewallRules = tencentcloud.Action{
		Service: "lighthouse",
		Version: "2020-03-24",
		Action:  "CreateFirewallRules",
	}
)
