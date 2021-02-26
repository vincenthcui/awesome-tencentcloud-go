package tiia

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	RecognizeCar = tencentcloud.Action{
		Service: "tiia",
		Version: "2019-05-29",
		Action:  "RecognizeCar",
	}
	EnhanceImage = tencentcloud.Action{
		Service: "tiia",
		Version: "2019-05-29",
		Action:  "EnhanceImage",
	}
	DetectProductBeta = tencentcloud.Action{
		Service: "tiia",
		Version: "2019-05-29",
		Action:  "DetectProductBeta",
	}
	DetectProduct = tencentcloud.Action{
		Service: "tiia",
		Version: "2019-05-29",
		Action:  "DetectProduct",
	}
	DetectMisbehavior = tencentcloud.Action{
		Service: "tiia",
		Version: "2019-05-29",
		Action:  "DetectMisbehavior",
	}
	DetectLabel = tencentcloud.Action{
		Service: "tiia",
		Version: "2019-05-29",
		Action:  "DetectLabel",
	}
	DetectDisgust = tencentcloud.Action{
		Service: "tiia",
		Version: "2019-05-29",
		Action:  "DetectDisgust",
	}
	DetectCelebrity = tencentcloud.Action{
		Service: "tiia",
		Version: "2019-05-29",
		Action:  "DetectCelebrity",
	}
	CropImage = tencentcloud.Action{
		Service: "tiia",
		Version: "2019-05-29",
		Action:  "CropImage",
	}
	AssessQuality = tencentcloud.Action{
		Service: "tiia",
		Version: "2019-05-29",
		Action:  "AssessQuality",
	}
)
