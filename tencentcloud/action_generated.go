package tencentcloud

// suppose to auto-generate action list here

var (
	CVMDescribeInstances = action{
		service: "cvm",
		version: "2017-03-12",
		action:  "DescribeInstances",
	}
	CVMDescribeImages = action{
		service: "cvm",
		version: "2017-03-12",
		action:  "DescribeImages",
	}
	CVMRunInstances = action{
		service: "cvm",
		version: "2017-03-12",
		action:  "RunInstances",
	}
	CVMTerminateInstances = action{
		service: "cvm",
		version: "2017-03-12",
		action:  "TerminateInstances",
	}
	CVMResetInstance = action{
		service: "cvm",
		version: "2017-03-12",
		action:  "ResetInstance",
	}
	CVMDescribeInstanceTypeConfigs = action{
		service: "cvm",
		version: "2017-03-12",
		action:  "DescribeInstanceTypeConfigs",
	}
	TAGModifyResourcesTagValue = action{
		service: "tag",
		version: "2018-08-13",
		action:  "ModifyResourcesTagValue",
	}
	CAMAttachUserPolicy = action{
		service: "cam",
		version: "2019-01-16",
		action:  "AttachUserPolicy",
	}
	CAMCreatePolicy = action{
		service: "cam",
		version: "2019-01-16",
		action:  "CreatePolicy",
	}
	CAMDeletePolicy = action{
		service: "cam",
		version: "2019-01-16",
		action:  "DeletePolicy",
	}
	CDBRunInstanceHour = action{
		service: "cdb",
		version: "2017-03-20",
		action:  "CreateDBInstanceHour",
	}
	CDBDescribeInstances = action{
		service: "cdb",
		version: "2017-03-20",
		action:  "DescribeDBInstances",
	}
	OpenWanService = action{
		service: "cdb",
		version: "2017-03-20",
		action:  "OpenWanService",
	}
	CDBIsolateInstance = action{
		service: "cdb",
		version: "2017-03-20",
		action:  "IsolateDBInstance",
	}
	CDBOfflineIsolatedInstances = action{
		service: "cdb",
		version: "2017-03-20",
		action:  "OfflineIsolatedInstances",
	}
)
