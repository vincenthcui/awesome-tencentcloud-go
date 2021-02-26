package memcached

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	DescribeInstances = tencentcloud.Action{
		Service: "memcached",
		Version: "2019-03-18",
		Action:  "DescribeInstances",
	}
)
