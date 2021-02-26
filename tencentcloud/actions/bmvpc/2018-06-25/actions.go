package bmvpc

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	UpgradeNatGateway = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "UpgradeNatGateway",
	}
	UnbindSubnetsFromNatGateway = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "UnbindSubnetsFromNatGateway",
	}
	UnbindIpsFromNatGateway = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "UnbindIpsFromNatGateway",
	}
	UnbindEipsFromNatGateway = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "UnbindEipsFromNatGateway",
	}
	ResetVpnConnection = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "ResetVpnConnection",
	}
	RejectVpcPeerConnection = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "RejectVpcPeerConnection",
	}
	ModifyVpnGatewayAttribute = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "ModifyVpnGatewayAttribute",
	}
	ModifyVpnConnectionAttribute = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "ModifyVpnConnectionAttribute",
	}
	ModifyVpcPeerConnection = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "ModifyVpcPeerConnection",
	}
	ModifyVpcAttribute = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "ModifyVpcAttribute",
	}
	ModifySubnetDHCPRelay = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "ModifySubnetDHCPRelay",
	}
	ModifySubnetAttribute = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "ModifySubnetAttribute",
	}
	ModifyRouteTable = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "ModifyRouteTable",
	}
	ModifyRoutePolicy = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "ModifyRoutePolicy",
	}
	ModifyCustomerGatewayAttribute = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "ModifyCustomerGatewayAttribute",
	}
	DownloadCustomerGatewayConfiguration = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "DownloadCustomerGatewayConfiguration",
	}
	DescribeVpnGateways = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "DescribeVpnGateways",
	}
	DescribeVpnConnections = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "DescribeVpnConnections",
	}
	DescribeVpcs = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "DescribeVpcs",
	}
	DescribeVpcView = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "DescribeVpcView",
	}
	DescribeVpcResource = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "DescribeVpcResource",
	}
	DescribeVpcQuota = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "DescribeVpcQuota",
	}
	DescribeVpcPeerConnections = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "DescribeVpcPeerConnections",
	}
	DescribeTaskStatus = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "DescribeTaskStatus",
	}
	DescribeSubnets = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "DescribeSubnets",
	}
	DescribeSubnetByHostedDevice = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "DescribeSubnetByHostedDevice",
	}
	DescribeSubnetByDevice = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "DescribeSubnetByDevice",
	}
	DescribeSubnetAvailableIps = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "DescribeSubnetAvailableIps",
	}
	DescribeRouteTables = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "DescribeRouteTables",
	}
	DescribeRoutePolicies = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "DescribeRoutePolicies",
	}
	DescribeNatSubnets = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "DescribeNatSubnets",
	}
	DescribeNatGateways = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "DescribeNatGateways",
	}
	DescribeCustomerGateways = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "DescribeCustomerGateways",
	}
	DeregisterIps = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "DeregisterIps",
	}
	DeleteVpnGateway = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "DeleteVpnGateway",
	}
	DeleteVpnConnection = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "DeleteVpnConnection",
	}
	DeleteVpcPeerConnection = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "DeleteVpcPeerConnection",
	}
	DeleteVpc = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "DeleteVpc",
	}
	DeleteVirtualIp = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "DeleteVirtualIp",
	}
	DeleteSubnet = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "DeleteSubnet",
	}
	DeleteRoutePolicy = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "DeleteRoutePolicy",
	}
	DeleteNatGateway = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "DeleteNatGateway",
	}
	DeleteInterfaces = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "DeleteInterfaces",
	}
	DeleteHostedInterfaces = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "DeleteHostedInterfaces",
	}
	DeleteHostedInterface = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "DeleteHostedInterface",
	}
	DeleteCustomerGateway = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "DeleteCustomerGateway",
	}
	CreateVpcPeerConnection = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "CreateVpcPeerConnection",
	}
	CreateVpc = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "CreateVpc",
	}
	CreateVirtualSubnetWithVlan = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "CreateVirtualSubnetWithVlan",
	}
	CreateSubnet = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "CreateSubnet",
	}
	CreateRoutePolicies = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "CreateRoutePolicies",
	}
	CreateNatGateway = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "CreateNatGateway",
	}
	CreateInterfaces = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "CreateInterfaces",
	}
	CreateHostedInterface = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "CreateHostedInterface",
	}
	CreateDockerSubnetWithVlan = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "CreateDockerSubnetWithVlan",
	}
	CreateCustomerGateway = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "CreateCustomerGateway",
	}
	BindSubnetsToNatGateway = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "BindSubnetsToNatGateway",
	}
	BindIpsToNatGateway = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "BindIpsToNatGateway",
	}
	BindEipsToNatGateway = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "BindEipsToNatGateway",
	}
	AsyncRegisterIps = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "AsyncRegisterIps",
	}
	AcceptVpcPeerConnection = tencentcloud.Action{
		Service: "bmvpc",
		Version: "2018-06-25",
		Action:  "AcceptVpcPeerConnection",
	}
)
