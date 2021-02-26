package taf

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	SendTrafficSecuritySmsMessage = tencentcloud.Action{
		Service: "taf",
		Version: "2020-02-10",
		Action:  "SendTrafficSecuritySmsMessage",
	}
	RecognizeTargetAudience = tencentcloud.Action{
		Service: "taf",
		Version: "2020-02-10",
		Action:  "RecognizeTargetAudience",
	}
	RecognizePreciseTargetAudience = tencentcloud.Action{
		Service: "taf",
		Version: "2020-02-10",
		Action:  "RecognizePreciseTargetAudience",
	}
	RecognizeEffectiveFlow = tencentcloud.Action{
		Service: "taf",
		Version: "2020-02-10",
		Action:  "RecognizeEffectiveFlow",
	}
	RecognizeCustomizedAudience = tencentcloud.Action{
		Service: "taf",
		Version: "2020-02-10",
		Action:  "RecognizeCustomizedAudience",
	}
	EnhanceTaDegree = tencentcloud.Action{
		Service: "taf",
		Version: "2020-02-10",
		Action:  "EnhanceTaDegree",
	}
	DetectFraudKOL = tencentcloud.Action{
		Service: "taf",
		Version: "2020-02-10",
		Action:  "DetectFraudKOL",
	}
)
