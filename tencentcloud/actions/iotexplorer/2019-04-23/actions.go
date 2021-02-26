package iotexplorer

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	SearchTopicRule = tencentcloud.Action{
		Service: "iotexplorer",
		Version: "2019-04-23",
		Action:  "SearchTopicRule",
	}
	SearchStudioProduct = tencentcloud.Action{
		Service: "iotexplorer",
		Version: "2019-04-23",
		Action:  "SearchStudioProduct",
	}
	ReleaseStudioProduct = tencentcloud.Action{
		Service: "iotexplorer",
		Version: "2019-04-23",
		Action:  "ReleaseStudioProduct",
	}
	PublishMessage = tencentcloud.Action{
		Service: "iotexplorer",
		Version: "2019-04-23",
		Action:  "PublishMessage",
	}
	ModifyTopicRule = tencentcloud.Action{
		Service: "iotexplorer",
		Version: "2019-04-23",
		Action:  "ModifyTopicRule",
	}
	ModifyStudioProduct = tencentcloud.Action{
		Service: "iotexplorer",
		Version: "2019-04-23",
		Action:  "ModifyStudioProduct",
	}
	ModifyProject = tencentcloud.Action{
		Service: "iotexplorer",
		Version: "2019-04-23",
		Action:  "ModifyProject",
	}
	ModifyModelDefinition = tencentcloud.Action{
		Service: "iotexplorer",
		Version: "2019-04-23",
		Action:  "ModifyModelDefinition",
	}
	ModifyLoRaGateway = tencentcloud.Action{
		Service: "iotexplorer",
		Version: "2019-04-23",
		Action:  "ModifyLoRaGateway",
	}
	ModifyLoRaFrequency = tencentcloud.Action{
		Service: "iotexplorer",
		Version: "2019-04-23",
		Action:  "ModifyLoRaFrequency",
	}
	ListEventHistory = tencentcloud.Action{
		Service: "iotexplorer",
		Version: "2019-04-23",
		Action:  "ListEventHistory",
	}
	GetTopicRuleList = tencentcloud.Action{
		Service: "iotexplorer",
		Version: "2019-04-23",
		Action:  "GetTopicRuleList",
	}
	GetStudioProductList = tencentcloud.Action{
		Service: "iotexplorer",
		Version: "2019-04-23",
		Action:  "GetStudioProductList",
	}
	GetProjectList = tencentcloud.Action{
		Service: "iotexplorer",
		Version: "2019-04-23",
		Action:  "GetProjectList",
	}
	GetLoRaGatewayList = tencentcloud.Action{
		Service: "iotexplorer",
		Version: "2019-04-23",
		Action:  "GetLoRaGatewayList",
	}
	GetDeviceList = tencentcloud.Action{
		Service: "iotexplorer",
		Version: "2019-04-23",
		Action:  "GetDeviceList",
	}
	EnableTopicRule = tencentcloud.Action{
		Service: "iotexplorer",
		Version: "2019-04-23",
		Action:  "EnableTopicRule",
	}
	DisableTopicRule = tencentcloud.Action{
		Service: "iotexplorer",
		Version: "2019-04-23",
		Action:  "DisableTopicRule",
	}
	DescribeTopicRule = tencentcloud.Action{
		Service: "iotexplorer",
		Version: "2019-04-23",
		Action:  "DescribeTopicRule",
	}
	DescribeStudioProduct = tencentcloud.Action{
		Service: "iotexplorer",
		Version: "2019-04-23",
		Action:  "DescribeStudioProduct",
	}
	DescribeProject = tencentcloud.Action{
		Service: "iotexplorer",
		Version: "2019-04-23",
		Action:  "DescribeProject",
	}
	DescribeModelDefinition = tencentcloud.Action{
		Service: "iotexplorer",
		Version: "2019-04-23",
		Action:  "DescribeModelDefinition",
	}
	DescribeLoRaFrequency = tencentcloud.Action{
		Service: "iotexplorer",
		Version: "2019-04-23",
		Action:  "DescribeLoRaFrequency",
	}
	DescribeDeviceDataHistory = tencentcloud.Action{
		Service: "iotexplorer",
		Version: "2019-04-23",
		Action:  "DescribeDeviceDataHistory",
	}
	DescribeDeviceData = tencentcloud.Action{
		Service: "iotexplorer",
		Version: "2019-04-23",
		Action:  "DescribeDeviceData",
	}
	DescribeDevice = tencentcloud.Action{
		Service: "iotexplorer",
		Version: "2019-04-23",
		Action:  "DescribeDevice",
	}
	DeleteTopicRule = tencentcloud.Action{
		Service: "iotexplorer",
		Version: "2019-04-23",
		Action:  "DeleteTopicRule",
	}
	DeleteStudioProduct = tencentcloud.Action{
		Service: "iotexplorer",
		Version: "2019-04-23",
		Action:  "DeleteStudioProduct",
	}
	DeleteProject = tencentcloud.Action{
		Service: "iotexplorer",
		Version: "2019-04-23",
		Action:  "DeleteProject",
	}
	DeleteLoRaGateway = tencentcloud.Action{
		Service: "iotexplorer",
		Version: "2019-04-23",
		Action:  "DeleteLoRaGateway",
	}
	DeleteLoRaFrequency = tencentcloud.Action{
		Service: "iotexplorer",
		Version: "2019-04-23",
		Action:  "DeleteLoRaFrequency",
	}
	DeleteDevice = tencentcloud.Action{
		Service: "iotexplorer",
		Version: "2019-04-23",
		Action:  "DeleteDevice",
	}
	CreateTopicRule = tencentcloud.Action{
		Service: "iotexplorer",
		Version: "2019-04-23",
		Action:  "CreateTopicRule",
	}
	CreateStudioProduct = tencentcloud.Action{
		Service: "iotexplorer",
		Version: "2019-04-23",
		Action:  "CreateStudioProduct",
	}
	CreateProject = tencentcloud.Action{
		Service: "iotexplorer",
		Version: "2019-04-23",
		Action:  "CreateProject",
	}
	CreateLoRaGateway = tencentcloud.Action{
		Service: "iotexplorer",
		Version: "2019-04-23",
		Action:  "CreateLoRaGateway",
	}
	CreateLoRaFrequency = tencentcloud.Action{
		Service: "iotexplorer",
		Version: "2019-04-23",
		Action:  "CreateLoRaFrequency",
	}
	CreateDevice = tencentcloud.Action{
		Service: "iotexplorer",
		Version: "2019-04-23",
		Action:  "CreateDevice",
	}
	ControlDeviceData = tencentcloud.Action{
		Service: "iotexplorer",
		Version: "2019-04-23",
		Action:  "ControlDeviceData",
	}
	CallDeviceActionSync = tencentcloud.Action{
		Service: "iotexplorer",
		Version: "2019-04-23",
		Action:  "CallDeviceActionSync",
	}
	CallDeviceActionAsync = tencentcloud.Action{
		Service: "iotexplorer",
		Version: "2019-04-23",
		Action:  "CallDeviceActionAsync",
	}
)
