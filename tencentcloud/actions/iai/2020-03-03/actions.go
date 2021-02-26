package iai

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	VerifyPerson = tencentcloud.Action{
		Service: "iai",
		Version: "2020-03-03",
		Action:  "VerifyPerson",
	}
	VerifyFace = tencentcloud.Action{
		Service: "iai",
		Version: "2020-03-03",
		Action:  "VerifyFace",
	}
	UpgradeGroupFaceModelVersion = tencentcloud.Action{
		Service: "iai",
		Version: "2020-03-03",
		Action:  "UpgradeGroupFaceModelVersion",
	}
	SearchPersonsReturnsByGroup = tencentcloud.Action{
		Service: "iai",
		Version: "2020-03-03",
		Action:  "SearchPersonsReturnsByGroup",
	}
	SearchPersons = tencentcloud.Action{
		Service: "iai",
		Version: "2020-03-03",
		Action:  "SearchPersons",
	}
	SearchFacesReturnsByGroup = tencentcloud.Action{
		Service: "iai",
		Version: "2020-03-03",
		Action:  "SearchFacesReturnsByGroup",
	}
	SearchFaces = tencentcloud.Action{
		Service: "iai",
		Version: "2020-03-03",
		Action:  "SearchFaces",
	}
	RevertGroupFaceModelVersion = tencentcloud.Action{
		Service: "iai",
		Version: "2020-03-03",
		Action:  "RevertGroupFaceModelVersion",
	}
	ModifyPersonGroupInfo = tencentcloud.Action{
		Service: "iai",
		Version: "2020-03-03",
		Action:  "ModifyPersonGroupInfo",
	}
	ModifyPersonBaseInfo = tencentcloud.Action{
		Service: "iai",
		Version: "2020-03-03",
		Action:  "ModifyPersonBaseInfo",
	}
	ModifyGroup = tencentcloud.Action{
		Service: "iai",
		Version: "2020-03-03",
		Action:  "ModifyGroup",
	}
	GetUpgradeGroupFaceModelVersionResult = tencentcloud.Action{
		Service: "iai",
		Version: "2020-03-03",
		Action:  "GetUpgradeGroupFaceModelVersionResult",
	}
	GetUpgradeGroupFaceModelVersionJobList = tencentcloud.Action{
		Service: "iai",
		Version: "2020-03-03",
		Action:  "GetUpgradeGroupFaceModelVersionJobList",
	}
	GetSimilarPersonResult = tencentcloud.Action{
		Service: "iai",
		Version: "2020-03-03",
		Action:  "GetSimilarPersonResult",
	}
	GetPersonListNum = tencentcloud.Action{
		Service: "iai",
		Version: "2020-03-03",
		Action:  "GetPersonListNum",
	}
	GetPersonList = tencentcloud.Action{
		Service: "iai",
		Version: "2020-03-03",
		Action:  "GetPersonList",
	}
	GetPersonGroupInfo = tencentcloud.Action{
		Service: "iai",
		Version: "2020-03-03",
		Action:  "GetPersonGroupInfo",
	}
	GetPersonBaseInfo = tencentcloud.Action{
		Service: "iai",
		Version: "2020-03-03",
		Action:  "GetPersonBaseInfo",
	}
	GetGroupList = tencentcloud.Action{
		Service: "iai",
		Version: "2020-03-03",
		Action:  "GetGroupList",
	}
	GetGroupInfo = tencentcloud.Action{
		Service: "iai",
		Version: "2020-03-03",
		Action:  "GetGroupInfo",
	}
	GetCheckSimilarPersonJobIdList = tencentcloud.Action{
		Service: "iai",
		Version: "2020-03-03",
		Action:  "GetCheckSimilarPersonJobIdList",
	}
	EstimateCheckSimilarPersonCostTime = tencentcloud.Action{
		Service: "iai",
		Version: "2020-03-03",
		Action:  "EstimateCheckSimilarPersonCostTime",
	}
	DetectLiveFaceAccurate = tencentcloud.Action{
		Service: "iai",
		Version: "2020-03-03",
		Action:  "DetectLiveFaceAccurate",
	}
	DetectLiveFace = tencentcloud.Action{
		Service: "iai",
		Version: "2020-03-03",
		Action:  "DetectLiveFace",
	}
	DetectFaceAttributes = tencentcloud.Action{
		Service: "iai",
		Version: "2020-03-03",
		Action:  "DetectFaceAttributes",
	}
	DetectFace = tencentcloud.Action{
		Service: "iai",
		Version: "2020-03-03",
		Action:  "DetectFace",
	}
	DeletePersonFromGroup = tencentcloud.Action{
		Service: "iai",
		Version: "2020-03-03",
		Action:  "DeletePersonFromGroup",
	}
	DeletePerson = tencentcloud.Action{
		Service: "iai",
		Version: "2020-03-03",
		Action:  "DeletePerson",
	}
	DeleteGroup = tencentcloud.Action{
		Service: "iai",
		Version: "2020-03-03",
		Action:  "DeleteGroup",
	}
	DeleteFace = tencentcloud.Action{
		Service: "iai",
		Version: "2020-03-03",
		Action:  "DeleteFace",
	}
	CreatePerson = tencentcloud.Action{
		Service: "iai",
		Version: "2020-03-03",
		Action:  "CreatePerson",
	}
	CreateGroup = tencentcloud.Action{
		Service: "iai",
		Version: "2020-03-03",
		Action:  "CreateGroup",
	}
	CreateFace = tencentcloud.Action{
		Service: "iai",
		Version: "2020-03-03",
		Action:  "CreateFace",
	}
	CopyPerson = tencentcloud.Action{
		Service: "iai",
		Version: "2020-03-03",
		Action:  "CopyPerson",
	}
	CompareMaskFace = tencentcloud.Action{
		Service: "iai",
		Version: "2020-03-03",
		Action:  "CompareMaskFace",
	}
	CompareFace = tencentcloud.Action{
		Service: "iai",
		Version: "2020-03-03",
		Action:  "CompareFace",
	}
	CheckSimilarPerson = tencentcloud.Action{
		Service: "iai",
		Version: "2020-03-03",
		Action:  "CheckSimilarPerson",
	}
	AnalyzeFace = tencentcloud.Action{
		Service: "iai",
		Version: "2020-03-03",
		Action:  "AnalyzeFace",
	}
	AnalyzeDenseLandmarks = tencentcloud.Action{
		Service: "iai",
		Version: "2020-03-03",
		Action:  "AnalyzeDenseLandmarks",
	}
)
