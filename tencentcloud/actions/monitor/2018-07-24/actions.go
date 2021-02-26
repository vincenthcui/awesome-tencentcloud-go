package monitor

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	UpdateServiceDiscovery = tencentcloud.Action{
		Service: "monitor",
		Version: "2018-07-24",
		Action:  "UpdateServiceDiscovery",
	}
	UnBindingPolicyObject = tencentcloud.Action{
		Service: "monitor",
		Version: "2018-07-24",
		Action:  "UnBindingPolicyObject",
	}
	UnBindingAllPolicyObject = tencentcloud.Action{
		Service: "monitor",
		Version: "2018-07-24",
		Action:  "UnBindingAllPolicyObject",
	}
	SetDefaultAlarmPolicy = tencentcloud.Action{
		Service: "monitor",
		Version: "2018-07-24",
		Action:  "SetDefaultAlarmPolicy",
	}
	SendCustomAlarmMsg = tencentcloud.Action{
		Service: "monitor",
		Version: "2018-07-24",
		Action:  "SendCustomAlarmMsg",
	}
	PutMonitorData = tencentcloud.Action{
		Service: "monitor",
		Version: "2018-07-24",
		Action:  "PutMonitorData",
	}
	ModifyPolicyGroup = tencentcloud.Action{
		Service: "monitor",
		Version: "2018-07-24",
		Action:  "ModifyPolicyGroup",
	}
	ModifyAlarmReceivers = tencentcloud.Action{
		Service: "monitor",
		Version: "2018-07-24",
		Action:  "ModifyAlarmReceivers",
	}
	ModifyAlarmPolicyTasks = tencentcloud.Action{
		Service: "monitor",
		Version: "2018-07-24",
		Action:  "ModifyAlarmPolicyTasks",
	}
	ModifyAlarmPolicyStatus = tencentcloud.Action{
		Service: "monitor",
		Version: "2018-07-24",
		Action:  "ModifyAlarmPolicyStatus",
	}
	ModifyAlarmPolicyNotice = tencentcloud.Action{
		Service: "monitor",
		Version: "2018-07-24",
		Action:  "ModifyAlarmPolicyNotice",
	}
	ModifyAlarmPolicyInfo = tencentcloud.Action{
		Service: "monitor",
		Version: "2018-07-24",
		Action:  "ModifyAlarmPolicyInfo",
	}
	ModifyAlarmPolicyCondition = tencentcloud.Action{
		Service: "monitor",
		Version: "2018-07-24",
		Action:  "ModifyAlarmPolicyCondition",
	}
	ModifyAlarmNotice = tencentcloud.Action{
		Service: "monitor",
		Version: "2018-07-24",
		Action:  "ModifyAlarmNotice",
	}
	GetMonitorData = tencentcloud.Action{
		Service: "monitor",
		Version: "2018-07-24",
		Action:  "GetMonitorData",
	}
	DescribeStatisticData = tencentcloud.Action{
		Service: "monitor",
		Version: "2018-07-24",
		Action:  "DescribeStatisticData",
	}
	DescribeServiceDiscovery = tencentcloud.Action{
		Service: "monitor",
		Version: "2018-07-24",
		Action:  "DescribeServiceDiscovery",
	}
	DescribeProductList = tencentcloud.Action{
		Service: "monitor",
		Version: "2018-07-24",
		Action:  "DescribeProductList",
	}
	DescribeProductEventList = tencentcloud.Action{
		Service: "monitor",
		Version: "2018-07-24",
		Action:  "DescribeProductEventList",
	}
	DescribePolicyGroupList = tencentcloud.Action{
		Service: "monitor",
		Version: "2018-07-24",
		Action:  "DescribePolicyGroupList",
	}
	DescribePolicyGroupInfo = tencentcloud.Action{
		Service: "monitor",
		Version: "2018-07-24",
		Action:  "DescribePolicyGroupInfo",
	}
	DescribePolicyConditionList = tencentcloud.Action{
		Service: "monitor",
		Version: "2018-07-24",
		Action:  "DescribePolicyConditionList",
	}
	DescribeMonitorTypes = tencentcloud.Action{
		Service: "monitor",
		Version: "2018-07-24",
		Action:  "DescribeMonitorTypes",
	}
	DescribeBindingPolicyObjectList = tencentcloud.Action{
		Service: "monitor",
		Version: "2018-07-24",
		Action:  "DescribeBindingPolicyObjectList",
	}
	DescribeBasicAlarmList = tencentcloud.Action{
		Service: "monitor",
		Version: "2018-07-24",
		Action:  "DescribeBasicAlarmList",
	}
	DescribeBaseMetrics = tencentcloud.Action{
		Service: "monitor",
		Version: "2018-07-24",
		Action:  "DescribeBaseMetrics",
	}
	DescribeAllNamespaces = tencentcloud.Action{
		Service: "monitor",
		Version: "2018-07-24",
		Action:  "DescribeAllNamespaces",
	}
	DescribeAlarmPolicy = tencentcloud.Action{
		Service: "monitor",
		Version: "2018-07-24",
		Action:  "DescribeAlarmPolicy",
	}
	DescribeAlarmPolicies = tencentcloud.Action{
		Service: "monitor",
		Version: "2018-07-24",
		Action:  "DescribeAlarmPolicies",
	}
	DescribeAlarmNotices = tencentcloud.Action{
		Service: "monitor",
		Version: "2018-07-24",
		Action:  "DescribeAlarmNotices",
	}
	DescribeAlarmNoticeCallbacks = tencentcloud.Action{
		Service: "monitor",
		Version: "2018-07-24",
		Action:  "DescribeAlarmNoticeCallbacks",
	}
	DescribeAlarmNotice = tencentcloud.Action{
		Service: "monitor",
		Version: "2018-07-24",
		Action:  "DescribeAlarmNotice",
	}
	DescribeAlarmMetrics = tencentcloud.Action{
		Service: "monitor",
		Version: "2018-07-24",
		Action:  "DescribeAlarmMetrics",
	}
	DescribeAlarmHistories = tencentcloud.Action{
		Service: "monitor",
		Version: "2018-07-24",
		Action:  "DescribeAlarmHistories",
	}
	DescribeAlarmEvents = tencentcloud.Action{
		Service: "monitor",
		Version: "2018-07-24",
		Action:  "DescribeAlarmEvents",
	}
	DescribeAccidentEventList = tencentcloud.Action{
		Service: "monitor",
		Version: "2018-07-24",
		Action:  "DescribeAccidentEventList",
	}
	DeleteServiceDiscovery = tencentcloud.Action{
		Service: "monitor",
		Version: "2018-07-24",
		Action:  "DeleteServiceDiscovery",
	}
	DeletePolicyGroup = tencentcloud.Action{
		Service: "monitor",
		Version: "2018-07-24",
		Action:  "DeletePolicyGroup",
	}
	DeleteAlarmPolicy = tencentcloud.Action{
		Service: "monitor",
		Version: "2018-07-24",
		Action:  "DeleteAlarmPolicy",
	}
	DeleteAlarmNotices = tencentcloud.Action{
		Service: "monitor",
		Version: "2018-07-24",
		Action:  "DeleteAlarmNotices",
	}
	CreateServiceDiscovery = tencentcloud.Action{
		Service: "monitor",
		Version: "2018-07-24",
		Action:  "CreateServiceDiscovery",
	}
	CreatePolicyGroup = tencentcloud.Action{
		Service: "monitor",
		Version: "2018-07-24",
		Action:  "CreatePolicyGroup",
	}
	CreateAlarmPolicy = tencentcloud.Action{
		Service: "monitor",
		Version: "2018-07-24",
		Action:  "CreateAlarmPolicy",
	}
	CreateAlarmNotice = tencentcloud.Action{
		Service: "monitor",
		Version: "2018-07-24",
		Action:  "CreateAlarmNotice",
	}
	BindingPolicyObject = tencentcloud.Action{
		Service: "monitor",
		Version: "2018-07-24",
		Action:  "BindingPolicyObject",
	}
)
