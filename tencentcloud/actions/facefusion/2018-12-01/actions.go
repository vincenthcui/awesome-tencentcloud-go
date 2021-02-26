package facefusion

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	FuseFace = tencentcloud.Action{
		Service: "facefusion",
		Version: "2018-12-01",
		Action:  "FuseFace",
	}
	FaceFusionLite = tencentcloud.Action{
		Service: "facefusion",
		Version: "2018-12-01",
		Action:  "FaceFusionLite",
	}
	FaceFusion = tencentcloud.Action{
		Service: "facefusion",
		Version: "2018-12-01",
		Action:  "FaceFusion",
	}
	DescribeMaterialList = tencentcloud.Action{
		Service: "facefusion",
		Version: "2018-12-01",
		Action:  "DescribeMaterialList",
	}
)
