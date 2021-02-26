package tics

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	DescribeThreatInfo = tencentcloud.Action{
		Service: "tics",
		Version: "2018-11-15",
		Action:  "DescribeThreatInfo",
	}
	DescribeIpInfo = tencentcloud.Action{
		Service: "tics",
		Version: "2018-11-15",
		Action:  "DescribeIpInfo",
	}
	DescribeFileInfo = tencentcloud.Action{
		Service: "tics",
		Version: "2018-11-15",
		Action:  "DescribeFileInfo",
	}
	DescribeDomainInfo = tencentcloud.Action{
		Service: "tics",
		Version: "2018-11-15",
		Action:  "DescribeDomainInfo",
	}
)
