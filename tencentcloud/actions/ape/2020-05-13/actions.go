package ape

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	DescribeImages = tencentcloud.Action{
		Service: "ape",
		Version: "2020-05-13",
		Action:  "DescribeImages",
	}
	DescribeImage = tencentcloud.Action{
		Service: "ape",
		Version: "2020-05-13",
		Action:  "DescribeImage",
	}
	DescribeDownloadInfos = tencentcloud.Action{
		Service: "ape",
		Version: "2020-05-13",
		Action:  "DescribeDownloadInfos",
	}
	DescribeAuthUsers = tencentcloud.Action{
		Service: "ape",
		Version: "2020-05-13",
		Action:  "DescribeAuthUsers",
	}
	CreateOrderAndPay = tencentcloud.Action{
		Service: "ape",
		Version: "2020-05-13",
		Action:  "CreateOrderAndPay",
	}
	CreateOrderAndDownloads = tencentcloud.Action{
		Service: "ape",
		Version: "2020-05-13",
		Action:  "CreateOrderAndDownloads",
	}
	BatchDescribeOrderImage = tencentcloud.Action{
		Service: "ape",
		Version: "2020-05-13",
		Action:  "BatchDescribeOrderImage",
	}
	BatchDescribeOrderCertificate = tencentcloud.Action{
		Service: "ape",
		Version: "2020-05-13",
		Action:  "BatchDescribeOrderCertificate",
	}
)
