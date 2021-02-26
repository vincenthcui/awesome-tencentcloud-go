package iotcloud

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	UploadFirmware = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "UploadFirmware",
	}
	UpdateTopicPolicy = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "UpdateTopicPolicy",
	}
	UpdateDeviceShadow = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "UpdateDeviceShadow",
	}
	UpdateDeviceAvailableState = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "UpdateDeviceAvailableState",
	}
	UnbindDevices = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "UnbindDevices",
	}
	RetryDeviceFirmwareTask = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "RetryDeviceFirmwareTask",
	}
	ResetDeviceState = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "ResetDeviceState",
	}
	ReplaceTopicRule = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "ReplaceTopicRule",
	}
	PublishToDevice = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "PublishToDevice",
	}
	PublishRRPCMessage = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "PublishRRPCMessage",
	}
	PublishMessage = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "PublishMessage",
	}
	PublishBroadcastMessage = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "PublishBroadcastMessage",
	}
	PublishAsDevice = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "PublishAsDevice",
	}
	EnableTopicRule = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "EnableTopicRule",
	}
	EditFirmware = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "EditFirmware",
	}
	DisableTopicRule = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "DisableTopicRule",
	}
	DescribeTasks = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "DescribeTasks",
	}
	DescribeTask = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "DescribeTask",
	}
	DescribeProducts = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "DescribeProducts",
	}
	DescribeProductTasks = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "DescribeProductTasks",
	}
	DescribeProductTask = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "DescribeProductTask",
	}
	DescribeMultiDevices = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "DescribeMultiDevices",
	}
	DescribeMultiDevTask = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "DescribeMultiDevTask",
	}
	DescribeLoraDevice = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "DescribeLoraDevice",
	}
	DescribeFirmwareTasks = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "DescribeFirmwareTasks",
	}
	DescribeFirmwareTaskStatistics = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "DescribeFirmwareTaskStatistics",
	}
	DescribeFirmwareTaskDistribution = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "DescribeFirmwareTaskDistribution",
	}
	DescribeFirmwareTaskDevices = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "DescribeFirmwareTaskDevices",
	}
	DescribeFirmwareTask = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "DescribeFirmwareTask",
	}
	DescribeFirmware = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "DescribeFirmware",
	}
	DescribeDevices = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "DescribeDevices",
	}
	DescribeDeviceShadow = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "DescribeDeviceShadow",
	}
	DescribeDeviceClientKey = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "DescribeDeviceClientKey",
	}
	DescribeDevice = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "DescribeDevice",
	}
	DescribeAllDevices = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "DescribeAllDevices",
	}
	DeleteTopicRule = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "DeleteTopicRule",
	}
	DeleteProduct = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "DeleteProduct",
	}
	DeleteLoraDevice = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "DeleteLoraDevice",
	}
	DeleteDevice = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "DeleteDevice",
	}
	CreateTopicRule = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "CreateTopicRule",
	}
	CreateTopicPolicy = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "CreateTopicPolicy",
	}
	CreateTaskFileUrl = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "CreateTaskFileUrl",
	}
	CreateTask = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "CreateTask",
	}
	CreateProduct = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "CreateProduct",
	}
	CreateMultiDevicesTask = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "CreateMultiDevicesTask",
	}
	CreateMultiDevice = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "CreateMultiDevice",
	}
	CreateLoraDevice = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "CreateLoraDevice",
	}
	CreateDevice = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "CreateDevice",
	}
	CancelTask = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "CancelTask",
	}
	CancelDeviceFirmwareTask = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "CancelDeviceFirmwareTask",
	}
	BindDevices = tencentcloud.Action{
		Service: "iotcloud",
		Version: "2018-06-14",
		Action:  "BindDevices",
	}
)
