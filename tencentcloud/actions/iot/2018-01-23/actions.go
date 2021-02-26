package iot

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	UpdateRule = tencentcloud.Action{
		Service: "iot",
		Version: "2018-01-23",
		Action:  "UpdateRule",
	}
	UpdateProduct = tencentcloud.Action{
		Service: "iot",
		Version: "2018-01-23",
		Action:  "UpdateProduct",
	}
	UnassociateSubDeviceFromGatewayProduct = tencentcloud.Action{
		Service: "iot",
		Version: "2018-01-23",
		Action:  "UnassociateSubDeviceFromGatewayProduct",
	}
	ResetDevice = tencentcloud.Action{
		Service: "iot",
		Version: "2018-01-23",
		Action:  "ResetDevice",
	}
	PublishMsg = tencentcloud.Action{
		Service: "iot",
		Version: "2018-01-23",
		Action:  "PublishMsg",
	}
	IssueDeviceControl = tencentcloud.Action{
		Service: "iot",
		Version: "2018-01-23",
		Action:  "IssueDeviceControl",
	}
	GetTopics = tencentcloud.Action{
		Service: "iot",
		Version: "2018-01-23",
		Action:  "GetTopics",
	}
	GetTopic = tencentcloud.Action{
		Service: "iot",
		Version: "2018-01-23",
		Action:  "GetTopic",
	}
	GetRules = tencentcloud.Action{
		Service: "iot",
		Version: "2018-01-23",
		Action:  "GetRules",
	}
	GetRule = tencentcloud.Action{
		Service: "iot",
		Version: "2018-01-23",
		Action:  "GetRule",
	}
	GetProducts = tencentcloud.Action{
		Service: "iot",
		Version: "2018-01-23",
		Action:  "GetProducts",
	}
	GetProduct = tencentcloud.Action{
		Service: "iot",
		Version: "2018-01-23",
		Action:  "GetProduct",
	}
	GetDevices = tencentcloud.Action{
		Service: "iot",
		Version: "2018-01-23",
		Action:  "GetDevices",
	}
	GetDeviceStatuses = tencentcloud.Action{
		Service: "iot",
		Version: "2018-01-23",
		Action:  "GetDeviceStatuses",
	}
	GetDeviceStatistics = tencentcloud.Action{
		Service: "iot",
		Version: "2018-01-23",
		Action:  "GetDeviceStatistics",
	}
	GetDeviceSignatures = tencentcloud.Action{
		Service: "iot",
		Version: "2018-01-23",
		Action:  "GetDeviceSignatures",
	}
	GetDeviceLog = tencentcloud.Action{
		Service: "iot",
		Version: "2018-01-23",
		Action:  "GetDeviceLog",
	}
	GetDeviceData = tencentcloud.Action{
		Service: "iot",
		Version: "2018-01-23",
		Action:  "GetDeviceData",
	}
	GetDevice = tencentcloud.Action{
		Service: "iot",
		Version: "2018-01-23",
		Action:  "GetDevice",
	}
	GetDebugLog = tencentcloud.Action{
		Service: "iot",
		Version: "2018-01-23",
		Action:  "GetDebugLog",
	}
	GetDataHistory = tencentcloud.Action{
		Service: "iot",
		Version: "2018-01-23",
		Action:  "GetDataHistory",
	}
	DeleteTopic = tencentcloud.Action{
		Service: "iot",
		Version: "2018-01-23",
		Action:  "DeleteTopic",
	}
	DeleteRule = tencentcloud.Action{
		Service: "iot",
		Version: "2018-01-23",
		Action:  "DeleteRule",
	}
	DeleteProduct = tencentcloud.Action{
		Service: "iot",
		Version: "2018-01-23",
		Action:  "DeleteProduct",
	}
	DeleteDevice = tencentcloud.Action{
		Service: "iot",
		Version: "2018-01-23",
		Action:  "DeleteDevice",
	}
	DeactivateRule = tencentcloud.Action{
		Service: "iot",
		Version: "2018-01-23",
		Action:  "DeactivateRule",
	}
	AssociateSubDeviceToGatewayProduct = tencentcloud.Action{
		Service: "iot",
		Version: "2018-01-23",
		Action:  "AssociateSubDeviceToGatewayProduct",
	}
	AppUpdateUser = tencentcloud.Action{
		Service: "iot",
		Version: "2018-01-23",
		Action:  "AppUpdateUser",
	}
	AppUpdateDevice = tencentcloud.Action{
		Service: "iot",
		Version: "2018-01-23",
		Action:  "AppUpdateDevice",
	}
	AppSecureAddDevice = tencentcloud.Action{
		Service: "iot",
		Version: "2018-01-23",
		Action:  "AppSecureAddDevice",
	}
	AppResetPassword = tencentcloud.Action{
		Service: "iot",
		Version: "2018-01-23",
		Action:  "AppResetPassword",
	}
	AppIssueDeviceControl = tencentcloud.Action{
		Service: "iot",
		Version: "2018-01-23",
		Action:  "AppIssueDeviceControl",
	}
	AppGetUser = tencentcloud.Action{
		Service: "iot",
		Version: "2018-01-23",
		Action:  "AppGetUser",
	}
	AppGetToken = tencentcloud.Action{
		Service: "iot",
		Version: "2018-01-23",
		Action:  "AppGetToken",
	}
	AppGetDevices = tencentcloud.Action{
		Service: "iot",
		Version: "2018-01-23",
		Action:  "AppGetDevices",
	}
	AppGetDeviceStatuses = tencentcloud.Action{
		Service: "iot",
		Version: "2018-01-23",
		Action:  "AppGetDeviceStatuses",
	}
	AppGetDeviceData = tencentcloud.Action{
		Service: "iot",
		Version: "2018-01-23",
		Action:  "AppGetDeviceData",
	}
	AppGetDevice = tencentcloud.Action{
		Service: "iot",
		Version: "2018-01-23",
		Action:  "AppGetDevice",
	}
	AppDeleteDevice = tencentcloud.Action{
		Service: "iot",
		Version: "2018-01-23",
		Action:  "AppDeleteDevice",
	}
	AppAddUser = tencentcloud.Action{
		Service: "iot",
		Version: "2018-01-23",
		Action:  "AppAddUser",
	}
	AddTopic = tencentcloud.Action{
		Service: "iot",
		Version: "2018-01-23",
		Action:  "AddTopic",
	}
	AddRule = tencentcloud.Action{
		Service: "iot",
		Version: "2018-01-23",
		Action:  "AddRule",
	}
	AddProduct = tencentcloud.Action{
		Service: "iot",
		Version: "2018-01-23",
		Action:  "AddProduct",
	}
	AddDevice = tencentcloud.Action{
		Service: "iot",
		Version: "2018-01-23",
		Action:  "AddDevice",
	}
	ActivateRule = tencentcloud.Action{
		Service: "iot",
		Version: "2018-01-23",
		Action:  "ActivateRule",
	}
)
