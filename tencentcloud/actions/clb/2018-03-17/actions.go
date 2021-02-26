package clb

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	SetSecurityGroupForLoadbalancers = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "SetSecurityGroupForLoadbalancers",
	}
	SetLoadBalancerSecurityGroups = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "SetLoadBalancerSecurityGroups",
	}
	SetLoadBalancerClsLog = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "SetLoadBalancerClsLog",
	}
	ReplaceCertForLoadBalancers = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "ReplaceCertForLoadBalancers",
	}
	RegisterTargetsWithClassicalLB = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "RegisterTargetsWithClassicalLB",
	}
	RegisterTargets = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "RegisterTargets",
	}
	RegisterTargetGroupInstances = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "RegisterTargetGroupInstances",
	}
	ModifyTargetWeight = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "ModifyTargetWeight",
	}
	ModifyTargetPort = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "ModifyTargetPort",
	}
	ModifyTargetGroupInstancesWeight = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "ModifyTargetGroupInstancesWeight",
	}
	ModifyTargetGroupInstancesPort = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "ModifyTargetGroupInstancesPort",
	}
	ModifyTargetGroupAttribute = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "ModifyTargetGroupAttribute",
	}
	ModifyRule = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "ModifyRule",
	}
	ModifyLoadBalancerAttributes = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "ModifyLoadBalancerAttributes",
	}
	ModifyListener = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "ModifyListener",
	}
	ModifyDomainAttributes = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "ModifyDomainAttributes",
	}
	ModifyDomain = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "ModifyDomain",
	}
	ModifyBlockIPList = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "ModifyBlockIPList",
	}
	ManualRewrite = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "ManualRewrite",
	}
	DisassociateTargetGroups = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "DisassociateTargetGroups",
	}
	DescribeTaskStatus = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "DescribeTaskStatus",
	}
	DescribeTargets = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "DescribeTargets",
	}
	DescribeTargetHealth = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "DescribeTargetHealth",
	}
	DescribeTargetGroups = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "DescribeTargetGroups",
	}
	DescribeTargetGroupList = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "DescribeTargetGroupList",
	}
	DescribeTargetGroupInstances = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "DescribeTargetGroupInstances",
	}
	DescribeRewrite = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "DescribeRewrite",
	}
	DescribeQuota = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "DescribeQuota",
	}
	DescribeLoadBalancersDetail = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "DescribeLoadBalancersDetail",
	}
	DescribeLoadBalancers = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "DescribeLoadBalancers",
	}
	DescribeLoadBalancerTraffic = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "DescribeLoadBalancerTraffic",
	}
	DescribeLoadBalancerListByCertId = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "DescribeLoadBalancerListByCertId",
	}
	DescribeListeners = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "DescribeListeners",
	}
	DescribeExclusiveClusters = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "DescribeExclusiveClusters",
	}
	DescribeClusterResources = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "DescribeClusterResources",
	}
	DescribeClsLogSet = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "DescribeClsLogSet",
	}
	DescribeClassicalLBTargets = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "DescribeClassicalLBTargets",
	}
	DescribeClassicalLBListeners = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "DescribeClassicalLBListeners",
	}
	DescribeClassicalLBHealthStatus = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "DescribeClassicalLBHealthStatus",
	}
	DescribeClassicalLBByInstanceId = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "DescribeClassicalLBByInstanceId",
	}
	DescribeBlockIPTask = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "DescribeBlockIPTask",
	}
	DescribeBlockIPList = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "DescribeBlockIPList",
	}
	DeregisterTargetsFromClassicalLB = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "DeregisterTargetsFromClassicalLB",
	}
	DeregisterTargets = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "DeregisterTargets",
	}
	DeregisterTargetGroupInstances = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "DeregisterTargetGroupInstances",
	}
	DeleteTargetGroups = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "DeleteTargetGroups",
	}
	DeleteRule = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "DeleteRule",
	}
	DeleteRewrite = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "DeleteRewrite",
	}
	DeleteLoadBalancerSnatIps = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "DeleteLoadBalancerSnatIps",
	}
	DeleteLoadBalancerListeners = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "DeleteLoadBalancerListeners",
	}
	DeleteLoadBalancer = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "DeleteLoadBalancer",
	}
	DeleteListener = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "DeleteListener",
	}
	CreateTopic = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "CreateTopic",
	}
	CreateTargetGroup = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "CreateTargetGroup",
	}
	CreateRule = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "CreateRule",
	}
	CreateLoadBalancerSnatIps = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "CreateLoadBalancerSnatIps",
	}
	CreateLoadBalancer = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "CreateLoadBalancer",
	}
	CreateListener = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "CreateListener",
	}
	CreateClsLogSet = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "CreateClsLogSet",
	}
	BatchRegisterTargets = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "BatchRegisterTargets",
	}
	BatchModifyTargetWeight = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "BatchModifyTargetWeight",
	}
	BatchDeregisterTargets = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "BatchDeregisterTargets",
	}
	AutoRewrite = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "AutoRewrite",
	}
	AssociateTargetGroups = tencentcloud.Action{
		Service: "clb",
		Version: "2018-03-17",
		Action:  "AssociateTargetGroups",
	}
)
