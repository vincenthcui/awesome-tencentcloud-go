package sms

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	SmsPackagesStatistics = tencentcloud.Action{
		Service: "sms",
		Version: "2019-07-11",
		Action:  "SmsPackagesStatistics",
	}
	SendStatusStatistics = tencentcloud.Action{
		Service: "sms",
		Version: "2019-07-11",
		Action:  "SendStatusStatistics",
	}
	SendSms = tencentcloud.Action{
		Service: "sms",
		Version: "2019-07-11",
		Action:  "SendSms",
	}
	PullSmsSendStatusByPhoneNumber = tencentcloud.Action{
		Service: "sms",
		Version: "2019-07-11",
		Action:  "PullSmsSendStatusByPhoneNumber",
	}
	PullSmsSendStatus = tencentcloud.Action{
		Service: "sms",
		Version: "2019-07-11",
		Action:  "PullSmsSendStatus",
	}
	PullSmsReplyStatusByPhoneNumber = tencentcloud.Action{
		Service: "sms",
		Version: "2019-07-11",
		Action:  "PullSmsReplyStatusByPhoneNumber",
	}
	PullSmsReplyStatus = tencentcloud.Action{
		Service: "sms",
		Version: "2019-07-11",
		Action:  "PullSmsReplyStatus",
	}
	ModifySmsTemplate = tencentcloud.Action{
		Service: "sms",
		Version: "2019-07-11",
		Action:  "ModifySmsTemplate",
	}
	ModifySmsSign = tencentcloud.Action{
		Service: "sms",
		Version: "2019-07-11",
		Action:  "ModifySmsSign",
	}
	DescribeSmsTemplateList = tencentcloud.Action{
		Service: "sms",
		Version: "2019-07-11",
		Action:  "DescribeSmsTemplateList",
	}
	DescribeSmsSignList = tencentcloud.Action{
		Service: "sms",
		Version: "2019-07-11",
		Action:  "DescribeSmsSignList",
	}
	DeleteSmsTemplate = tencentcloud.Action{
		Service: "sms",
		Version: "2019-07-11",
		Action:  "DeleteSmsTemplate",
	}
	DeleteSmsSign = tencentcloud.Action{
		Service: "sms",
		Version: "2019-07-11",
		Action:  "DeleteSmsSign",
	}
	CallbackStatusStatistics = tencentcloud.Action{
		Service: "sms",
		Version: "2019-07-11",
		Action:  "CallbackStatusStatistics",
	}
	AddSmsTemplate = tencentcloud.Action{
		Service: "sms",
		Version: "2019-07-11",
		Action:  "AddSmsTemplate",
	}
	AddSmsSign = tencentcloud.Action{
		Service: "sms",
		Version: "2019-07-11",
		Action:  "AddSmsSign",
	}
)
