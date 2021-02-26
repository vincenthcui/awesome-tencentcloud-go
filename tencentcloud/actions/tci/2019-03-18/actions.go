package tci

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	TransmitAudioStream = tencentcloud.Action{
		Service: "tci",
		Version: "2019-03-18",
		Action:  "TransmitAudioStream",
	}
	SubmitTraditionalClassTask = tencentcloud.Action{
		Service: "tci",
		Version: "2019-03-18",
		Action:  "SubmitTraditionalClassTask",
	}
	SubmitPartialBodyClassTask = tencentcloud.Action{
		Service: "tci",
		Version: "2019-03-18",
		Action:  "SubmitPartialBodyClassTask",
	}
	SubmitOpenClassTask = tencentcloud.Action{
		Service: "tci",
		Version: "2019-03-18",
		Action:  "SubmitOpenClassTask",
	}
	SubmitOneByOneClassTask = tencentcloud.Action{
		Service: "tci",
		Version: "2019-03-18",
		Action:  "SubmitOneByOneClassTask",
	}
	SubmitImageTaskPlus = tencentcloud.Action{
		Service: "tci",
		Version: "2019-03-18",
		Action:  "SubmitImageTaskPlus",
	}
	SubmitImageTask = tencentcloud.Action{
		Service: "tci",
		Version: "2019-03-18",
		Action:  "SubmitImageTask",
	}
	SubmitHighlights = tencentcloud.Action{
		Service: "tci",
		Version: "2019-03-18",
		Action:  "SubmitHighlights",
	}
	SubmitFullBodyClassTask = tencentcloud.Action{
		Service: "tci",
		Version: "2019-03-18",
		Action:  "SubmitFullBodyClassTask",
	}
	SubmitDoubleVideoHighlights = tencentcloud.Action{
		Service: "tci",
		Version: "2019-03-18",
		Action:  "SubmitDoubleVideoHighlights",
	}
	SubmitConversationTask = tencentcloud.Action{
		Service: "tci",
		Version: "2019-03-18",
		Action:  "SubmitConversationTask",
	}
	SubmitCheckAttendanceTaskPlus = tencentcloud.Action{
		Service: "tci",
		Version: "2019-03-18",
		Action:  "SubmitCheckAttendanceTaskPlus",
	}
	SubmitCheckAttendanceTask = tencentcloud.Action{
		Service: "tci",
		Version: "2019-03-18",
		Action:  "SubmitCheckAttendanceTask",
	}
	SubmitAudioTask = tencentcloud.Action{
		Service: "tci",
		Version: "2019-03-18",
		Action:  "SubmitAudioTask",
	}
	ModifyPerson = tencentcloud.Action{
		Service: "tci",
		Version: "2019-03-18",
		Action:  "ModifyPerson",
	}
	ModifyLibrary = tencentcloud.Action{
		Service: "tci",
		Version: "2019-03-18",
		Action:  "ModifyLibrary",
	}
	DescribeVocabLib = tencentcloud.Action{
		Service: "tci",
		Version: "2019-03-18",
		Action:  "DescribeVocabLib",
	}
	DescribeVocab = tencentcloud.Action{
		Service: "tci",
		Version: "2019-03-18",
		Action:  "DescribeVocab",
	}
	DescribePersons = tencentcloud.Action{
		Service: "tci",
		Version: "2019-03-18",
		Action:  "DescribePersons",
	}
	DescribePerson = tencentcloud.Action{
		Service: "tci",
		Version: "2019-03-18",
		Action:  "DescribePerson",
	}
	DescribeLibraries = tencentcloud.Action{
		Service: "tci",
		Version: "2019-03-18",
		Action:  "DescribeLibraries",
	}
	DescribeImageTaskStatistic = tencentcloud.Action{
		Service: "tci",
		Version: "2019-03-18",
		Action:  "DescribeImageTaskStatistic",
	}
	DescribeImageTask = tencentcloud.Action{
		Service: "tci",
		Version: "2019-03-18",
		Action:  "DescribeImageTask",
	}
	DescribeHighlightResult = tencentcloud.Action{
		Service: "tci",
		Version: "2019-03-18",
		Action:  "DescribeHighlightResult",
	}
	DescribeConversationTask = tencentcloud.Action{
		Service: "tci",
		Version: "2019-03-18",
		Action:  "DescribeConversationTask",
	}
	DescribeAudioTask = tencentcloud.Action{
		Service: "tci",
		Version: "2019-03-18",
		Action:  "DescribeAudioTask",
	}
	DescribeAttendanceResult = tencentcloud.Action{
		Service: "tci",
		Version: "2019-03-18",
		Action:  "DescribeAttendanceResult",
	}
	DescribeAITaskResult = tencentcloud.Action{
		Service: "tci",
		Version: "2019-03-18",
		Action:  "DescribeAITaskResult",
	}
	DeleteVocabLib = tencentcloud.Action{
		Service: "tci",
		Version: "2019-03-18",
		Action:  "DeleteVocabLib",
	}
	DeleteVocab = tencentcloud.Action{
		Service: "tci",
		Version: "2019-03-18",
		Action:  "DeleteVocab",
	}
	DeletePerson = tencentcloud.Action{
		Service: "tci",
		Version: "2019-03-18",
		Action:  "DeletePerson",
	}
	DeleteLibrary = tencentcloud.Action{
		Service: "tci",
		Version: "2019-03-18",
		Action:  "DeleteLibrary",
	}
	DeleteFace = tencentcloud.Action{
		Service: "tci",
		Version: "2019-03-18",
		Action:  "DeleteFace",
	}
	CreateVocabLib = tencentcloud.Action{
		Service: "tci",
		Version: "2019-03-18",
		Action:  "CreateVocabLib",
	}
	CreateVocab = tencentcloud.Action{
		Service: "tci",
		Version: "2019-03-18",
		Action:  "CreateVocab",
	}
	CreatePerson = tencentcloud.Action{
		Service: "tci",
		Version: "2019-03-18",
		Action:  "CreatePerson",
	}
	CreateLibrary = tencentcloud.Action{
		Service: "tci",
		Version: "2019-03-18",
		Action:  "CreateLibrary",
	}
	CreateFace = tencentcloud.Action{
		Service: "tci",
		Version: "2019-03-18",
		Action:  "CreateFace",
	}
	CheckFacePhoto = tencentcloud.Action{
		Service: "tci",
		Version: "2019-03-18",
		Action:  "CheckFacePhoto",
	}
	CancelTask = tencentcloud.Action{
		Service: "tci",
		Version: "2019-03-18",
		Action:  "CancelTask",
	}
	AIAssistant = tencentcloud.Action{
		Service: "tci",
		Version: "2019-03-18",
		Action:  "AIAssistant",
	}
)
