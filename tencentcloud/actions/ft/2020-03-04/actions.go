package ft

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	SwapGenderPic = tencentcloud.Action{
		Service: "ft",
		Version: "2020-03-04",
		Action:  "SwapGenderPic",
	}
	QueryFaceMorphJob = tencentcloud.Action{
		Service: "ft",
		Version: "2020-03-04",
		Action:  "QueryFaceMorphJob",
	}
	MorphFace = tencentcloud.Action{
		Service: "ft",
		Version: "2020-03-04",
		Action:  "MorphFace",
	}
	FaceCartoonPic = tencentcloud.Action{
		Service: "ft",
		Version: "2020-03-04",
		Action:  "FaceCartoonPic",
	}
	ChangeAgePic = tencentcloud.Action{
		Service: "ft",
		Version: "2020-03-04",
		Action:  "ChangeAgePic",
	}
	CancelFaceMorphJob = tencentcloud.Action{
		Service: "ft",
		Version: "2020-03-04",
		Action:  "CancelFaceMorphJob",
	}
)
