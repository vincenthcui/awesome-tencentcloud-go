package cam

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	UpdateUser = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "UpdateUser",
	}
	UpdateSAMLProvider = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "UpdateSAMLProvider",
	}
	UpdateRoleDescription = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "UpdateRoleDescription",
	}
	UpdateRoleConsoleLogin = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "UpdateRoleConsoleLogin",
	}
	UpdatePolicy = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "UpdatePolicy",
	}
	UpdateGroup = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "UpdateGroup",
	}
	UpdateAssumeRolePolicy = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "UpdateAssumeRolePolicy",
	}
	SetMfaFlag = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "SetMfaFlag",
	}
	SetDefaultPolicyVersion = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "SetDefaultPolicyVersion",
	}
	RemoveUserFromGroup = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "RemoveUserFromGroup",
	}
	PutUserPermissionsBoundary = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "PutUserPermissionsBoundary",
	}
	PutRolePermissionsBoundary = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "PutRolePermissionsBoundary",
	}
	ListWeChatWorkSubAccounts = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "ListWeChatWorkSubAccounts",
	}
	ListUsersForGroup = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "ListUsersForGroup",
	}
	ListUsers = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "ListUsers",
	}
	ListSAMLProviders = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "ListSAMLProviders",
	}
	ListPolicyVersions = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "ListPolicyVersions",
	}
	ListPolicies = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "ListPolicies",
	}
	ListGroupsForUser = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "ListGroupsForUser",
	}
	ListGroups = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "ListGroups",
	}
	ListEntitiesForPolicy = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "ListEntitiesForPolicy",
	}
	ListCollaborators = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "ListCollaborators",
	}
	ListAttachedUserPolicies = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "ListAttachedUserPolicies",
	}
	ListAttachedRolePolicies = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "ListAttachedRolePolicies",
	}
	ListAttachedGroupPolicies = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "ListAttachedGroupPolicies",
	}
	ListAccessKeys = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "ListAccessKeys",
	}
	GetUser = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "GetUser",
	}
	GetServiceLinkedRoleDeletionStatus = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "GetServiceLinkedRoleDeletionStatus",
	}
	GetSAMLProvider = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "GetSAMLProvider",
	}
	GetRole = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "GetRole",
	}
	GetPolicyVersion = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "GetPolicyVersion",
	}
	GetPolicy = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "GetPolicy",
	}
	GetGroup = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "GetGroup",
	}
	GetCustomMFATokenInfo = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "GetCustomMFATokenInfo",
	}
	DetachUserPolicy = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "DetachUserPolicy",
	}
	DetachRolePolicy = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "DetachRolePolicy",
	}
	DetachGroupPolicy = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "DetachGroupPolicy",
	}
	DescribeSafeAuthFlagColl = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "DescribeSafeAuthFlagColl",
	}
	DescribeSafeAuthFlag = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "DescribeSafeAuthFlag",
	}
	DescribeRoleList = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "DescribeRoleList",
	}
	DeleteUserPermissionsBoundary = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "DeleteUserPermissionsBoundary",
	}
	DeleteUser = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "DeleteUser",
	}
	DeleteServiceLinkedRole = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "DeleteServiceLinkedRole",
	}
	DeleteSAMLProvider = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "DeleteSAMLProvider",
	}
	DeleteRolePermissionsBoundary = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "DeleteRolePermissionsBoundary",
	}
	DeleteRole = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "DeleteRole",
	}
	DeletePolicyVersion = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "DeletePolicyVersion",
	}
	DeletePolicy = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "DeletePolicy",
	}
	DeleteGroup = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "DeleteGroup",
	}
	CreateServiceLinkedRole = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "CreateServiceLinkedRole",
	}
	CreateSAMLProvider = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "CreateSAMLProvider",
	}
	CreateRole = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "CreateRole",
	}
	CreatePolicyVersion = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "CreatePolicyVersion",
	}
	CreatePolicy = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "CreatePolicy",
	}
	CreateGroup = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "CreateGroup",
	}
	ConsumeCustomMFAToken = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "ConsumeCustomMFAToken",
	}
	AttachUserPolicy = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "AttachUserPolicy",
	}
	AttachRolePolicy = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "AttachRolePolicy",
	}
	AttachGroupPolicy = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "AttachGroupPolicy",
	}
	AddUserToGroup = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "AddUserToGroup",
	}
	AddUser = tencentcloud.Action{
		Service: "cam",
		Version: "2019-01-16",
		Action:  "AddUser",
	}
)
