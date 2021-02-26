package bmlb

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	UploadCert = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "UploadCert",
	}
	UnbindTrafficMirrorReceivers = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "UnbindTrafficMirrorReceivers",
	}
	UnbindTrafficMirrorListeners = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "UnbindTrafficMirrorListeners",
	}
	UnbindL7Backends = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "UnbindL7Backends",
	}
	UnbindL4Backends = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "UnbindL4Backends",
	}
	SetTrafficMirrorHealthSwitch = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "SetTrafficMirrorHealthSwitch",
	}
	SetTrafficMirrorAlias = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "SetTrafficMirrorAlias",
	}
	ReplaceCert = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "ReplaceCert",
	}
	ModifyLoadBalancerChargeMode = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "ModifyLoadBalancerChargeMode",
	}
	ModifyLoadBalancer = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "ModifyLoadBalancer",
	}
	ModifyL7Locations = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "ModifyL7Locations",
	}
	ModifyL7Listener = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "ModifyL7Listener",
	}
	ModifyL7BackendWeight = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "ModifyL7BackendWeight",
	}
	ModifyL7BackendPort = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "ModifyL7BackendPort",
	}
	ModifyL4Listener = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "ModifyL4Listener",
	}
	ModifyL4BackendWeight = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "ModifyL4BackendWeight",
	}
	ModifyL4BackendProbePort = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "ModifyL4BackendProbePort",
	}
	ModifyL4BackendPort = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "ModifyL4BackendPort",
	}
	DescribeTrafficMirrors = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "DescribeTrafficMirrors",
	}
	DescribeTrafficMirrorReceivers = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "DescribeTrafficMirrorReceivers",
	}
	DescribeTrafficMirrorReceiverHealthStatus = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "DescribeTrafficMirrorReceiverHealthStatus",
	}
	DescribeTrafficMirrorListeners = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "DescribeTrafficMirrorListeners",
	}
	DescribeLoadBalancers = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "DescribeLoadBalancers",
	}
	DescribeLoadBalancerTaskResult = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "DescribeLoadBalancerTaskResult",
	}
	DescribeLoadBalancerPortInfo = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "DescribeLoadBalancerPortInfo",
	}
	DescribeL7Rules = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "DescribeL7Rules",
	}
	DescribeL7ListenersEx = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "DescribeL7ListenersEx",
	}
	DescribeL7Listeners = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "DescribeL7Listeners",
	}
	DescribeL7ListenerInfo = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "DescribeL7ListenerInfo",
	}
	DescribeL7Backends = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "DescribeL7Backends",
	}
	DescribeL4Listeners = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "DescribeL4Listeners",
	}
	DescribeL4ListenerInfo = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "DescribeL4ListenerInfo",
	}
	DescribeL4Backends = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "DescribeL4Backends",
	}
	DescribeDevicesBindInfo = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "DescribeDevicesBindInfo",
	}
	DescribeCertDetail = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "DescribeCertDetail",
	}
	DeleteTrafficMirror = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "DeleteTrafficMirror",
	}
	DeleteLoadBalancer = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "DeleteLoadBalancer",
	}
	DeleteListeners = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "DeleteListeners",
	}
	DeleteL7Rules = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "DeleteL7Rules",
	}
	DeleteL7Domains = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "DeleteL7Domains",
	}
	CreateTrafficMirror = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "CreateTrafficMirror",
	}
	CreateLoadBalancers = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "CreateLoadBalancers",
	}
	CreateL7Rules = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "CreateL7Rules",
	}
	CreateL7Listeners = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "CreateL7Listeners",
	}
	CreateL4Listeners = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "CreateL4Listeners",
	}
	BindTrafficMirrorReceivers = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "BindTrafficMirrorReceivers",
	}
	BindTrafficMirrorListeners = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "BindTrafficMirrorListeners",
	}
	BindL7Backends = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "BindL7Backends",
	}
	BindL4Backends = tencentcloud.Action{
		Service: "bmlb",
		Version: "2018-06-25",
		Action:  "BindL4Backends",
	}
)
