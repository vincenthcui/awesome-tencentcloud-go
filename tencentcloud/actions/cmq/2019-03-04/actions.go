package cmq

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	UnbindDeadLetter = tencentcloud.Action{
		Service: "cmq",
		Version: "2019-03-04",
		Action:  "UnbindDeadLetter",
	}
	RewindQueue = tencentcloud.Action{
		Service: "cmq",
		Version: "2019-03-04",
		Action:  "RewindQueue",
	}
	ModifyTopicAttribute = tencentcloud.Action{
		Service: "cmq",
		Version: "2019-03-04",
		Action:  "ModifyTopicAttribute",
	}
	ModifySubscriptionAttribute = tencentcloud.Action{
		Service: "cmq",
		Version: "2019-03-04",
		Action:  "ModifySubscriptionAttribute",
	}
	ModifyQueueAttribute = tencentcloud.Action{
		Service: "cmq",
		Version: "2019-03-04",
		Action:  "ModifyQueueAttribute",
	}
	DescribeTopicDetail = tencentcloud.Action{
		Service: "cmq",
		Version: "2019-03-04",
		Action:  "DescribeTopicDetail",
	}
	DescribeSubscriptionDetail = tencentcloud.Action{
		Service: "cmq",
		Version: "2019-03-04",
		Action:  "DescribeSubscriptionDetail",
	}
	DescribeQueueDetail = tencentcloud.Action{
		Service: "cmq",
		Version: "2019-03-04",
		Action:  "DescribeQueueDetail",
	}
	DescribeDeadLetterSourceQueues = tencentcloud.Action{
		Service: "cmq",
		Version: "2019-03-04",
		Action:  "DescribeDeadLetterSourceQueues",
	}
	DeleteTopic = tencentcloud.Action{
		Service: "cmq",
		Version: "2019-03-04",
		Action:  "DeleteTopic",
	}
	DeleteSubscribe = tencentcloud.Action{
		Service: "cmq",
		Version: "2019-03-04",
		Action:  "DeleteSubscribe",
	}
	DeleteQueue = tencentcloud.Action{
		Service: "cmq",
		Version: "2019-03-04",
		Action:  "DeleteQueue",
	}
	CreateTopic = tencentcloud.Action{
		Service: "cmq",
		Version: "2019-03-04",
		Action:  "CreateTopic",
	}
	CreateSubscribe = tencentcloud.Action{
		Service: "cmq",
		Version: "2019-03-04",
		Action:  "CreateSubscribe",
	}
	CreateQueue = tencentcloud.Action{
		Service: "cmq",
		Version: "2019-03-04",
		Action:  "CreateQueue",
	}
	ClearSubscriptionFilterTags = tencentcloud.Action{
		Service: "cmq",
		Version: "2019-03-04",
		Action:  "ClearSubscriptionFilterTags",
	}
	ClearQueue = tencentcloud.Action{
		Service: "cmq",
		Version: "2019-03-04",
		Action:  "ClearQueue",
	}
)
