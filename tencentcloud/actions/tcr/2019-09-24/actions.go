package tcr

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	ValidateRepositoryExistPersonal = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "ValidateRepositoryExistPersonal",
	}
	ValidateNamespaceExistPersonal = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "ValidateNamespaceExistPersonal",
	}
	RenewInstance = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "RenewInstance",
	}
	ModifyWebhookTrigger = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "ModifyWebhookTrigger",
	}
	ModifyUserPasswordPersonal = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "ModifyUserPasswordPersonal",
	}
	ModifyRepositoryInfoPersonal = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "ModifyRepositoryInfoPersonal",
	}
	ModifyRepositoryAccessPersonal = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "ModifyRepositoryAccessPersonal",
	}
	ModifyRepository = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "ModifyRepository",
	}
	ModifyNamespace = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "ModifyNamespace",
	}
	ModifyInstanceToken = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "ModifyInstanceToken",
	}
	ModifyApplicationTriggerPersonal = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "ModifyApplicationTriggerPersonal",
	}
	ManageInternalEndpoint = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "ManageInternalEndpoint",
	}
	ManageImageLifecycleGlobalPersonal = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "ManageImageLifecycleGlobalPersonal",
	}
	DuplicateImagePersonal = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "DuplicateImagePersonal",
	}
	DescribeWebhookTriggerLog = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "DescribeWebhookTriggerLog",
	}
	DescribeWebhookTrigger = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "DescribeWebhookTrigger",
	}
	DescribeUserQuotaPersonal = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "DescribeUserQuotaPersonal",
	}
	DescribeRepositoryPersonal = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "DescribeRepositoryPersonal",
	}
	DescribeRepositoryOwnerPersonal = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "DescribeRepositoryOwnerPersonal",
	}
	DescribeRepositoryFilterPersonal = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "DescribeRepositoryFilterPersonal",
	}
	DescribeRepositories = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "DescribeRepositories",
	}
	DescribeReplicationInstances = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "DescribeReplicationInstances",
	}
	DescribeReplicationInstanceCreateTasks = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "DescribeReplicationInstanceCreateTasks",
	}
	DescribeNamespaces = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "DescribeNamespaces",
	}
	DescribeNamespacePersonal = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "DescribeNamespacePersonal",
	}
	DescribeInternalEndpoints = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "DescribeInternalEndpoints",
	}
	DescribeInternalEndpointDnsStatus = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "DescribeInternalEndpointDnsStatus",
	}
	DescribeInstances = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "DescribeInstances",
	}
	DescribeInstanceToken = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "DescribeInstanceToken",
	}
	DescribeInstanceStatus = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "DescribeInstanceStatus",
	}
	DescribeImages = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "DescribeImages",
	}
	DescribeImagePersonal = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "DescribeImagePersonal",
	}
	DescribeImageManifests = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "DescribeImageManifests",
	}
	DescribeImageLifecyclePersonal = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "DescribeImageLifecyclePersonal",
	}
	DescribeImageLifecycleGlobalPersonal = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "DescribeImageLifecycleGlobalPersonal",
	}
	DescribeImageFilterPersonal = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "DescribeImageFilterPersonal",
	}
	DescribeFavorRepositoryPersonal = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "DescribeFavorRepositoryPersonal",
	}
	DescribeApplicationTriggerPersonal = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "DescribeApplicationTriggerPersonal",
	}
	DescribeApplicationTriggerLogPersonal = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "DescribeApplicationTriggerLogPersonal",
	}
	DeleteWebhookTrigger = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "DeleteWebhookTrigger",
	}
	DeleteRepositoryPersonal = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "DeleteRepositoryPersonal",
	}
	DeleteRepository = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "DeleteRepository",
	}
	DeleteNamespacePersonal = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "DeleteNamespacePersonal",
	}
	DeleteNamespace = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "DeleteNamespace",
	}
	DeleteInternalEndpointDns = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "DeleteInternalEndpointDns",
	}
	DeleteInstanceToken = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "DeleteInstanceToken",
	}
	DeleteInstance = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "DeleteInstance",
	}
	DeleteImagePersonal = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "DeleteImagePersonal",
	}
	DeleteImageLifecyclePersonal = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "DeleteImageLifecyclePersonal",
	}
	DeleteImageLifecycleGlobalPersonal = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "DeleteImageLifecycleGlobalPersonal",
	}
	DeleteImage = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "DeleteImage",
	}
	DeleteApplicationTriggerPersonal = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "DeleteApplicationTriggerPersonal",
	}
	CreateWebhookTrigger = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "CreateWebhookTrigger",
	}
	CreateUserPersonal = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "CreateUserPersonal",
	}
	CreateRepositoryPersonal = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "CreateRepositoryPersonal",
	}
	CreateRepository = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "CreateRepository",
	}
	CreateNamespacePersonal = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "CreateNamespacePersonal",
	}
	CreateNamespace = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "CreateNamespace",
	}
	CreateInternalEndpointDns = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "CreateInternalEndpointDns",
	}
	CreateInstanceToken = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "CreateInstanceToken",
	}
	CreateInstance = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "CreateInstance",
	}
	CreateImageLifecyclePersonal = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "CreateImageLifecyclePersonal",
	}
	CreateApplicationTriggerPersonal = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "CreateApplicationTriggerPersonal",
	}
	CheckInstanceName = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "CheckInstanceName",
	}
	BatchDeleteRepositoryPersonal = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "BatchDeleteRepositoryPersonal",
	}
	BatchDeleteImagePersonal = tencentcloud.Action{
		Service: "tcr",
		Version: "2019-09-24",
		Action:  "BatchDeleteImagePersonal",
	}
)
