package mps

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	ResetWorkflow = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "ResetWorkflow",
	}
	RecognizeMediaForZhiXue = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "RecognizeMediaForZhiXue",
	}
	ProcessMedia = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "ProcessMedia",
	}
	ProcessLiveStream = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "ProcessLiveStream",
	}
	ParseNotification = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "ParseNotification",
	}
	ParseLiveStreamProcessNotification = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "ParseLiveStreamProcessNotification",
	}
	ModifyWordSample = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "ModifyWordSample",
	}
	ModifyWatermarkTemplate = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "ModifyWatermarkTemplate",
	}
	ModifyTranscodeTemplate = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "ModifyTranscodeTemplate",
	}
	ModifySnapshotByTimeOffsetTemplate = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "ModifySnapshotByTimeOffsetTemplate",
	}
	ModifySampleSnapshotTemplate = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "ModifySampleSnapshotTemplate",
	}
	ModifyPersonSample = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "ModifyPersonSample",
	}
	ModifyImageSpriteTemplate = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "ModifyImageSpriteTemplate",
	}
	ModifyContentReviewTemplate = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "ModifyContentReviewTemplate",
	}
	ModifyAnimatedGraphicsTemplate = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "ModifyAnimatedGraphicsTemplate",
	}
	ModifyAdaptiveDynamicStreamingTemplate = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "ModifyAdaptiveDynamicStreamingTemplate",
	}
	ModifyAIRecognitionTemplate = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "ModifyAIRecognitionTemplate",
	}
	ModifyAIAnalysisTemplate = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "ModifyAIAnalysisTemplate",
	}
	ManageTask = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "ManageTask",
	}
	ExecuteFunction = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "ExecuteFunction",
	}
	EnableWorkflow = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "EnableWorkflow",
	}
	EditMedia = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "EditMedia",
	}
	DisableWorkflow = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "DisableWorkflow",
	}
	DescribeWorkflows = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "DescribeWorkflows",
	}
	DescribeWordSamples = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "DescribeWordSamples",
	}
	DescribeWatermarkTemplates = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "DescribeWatermarkTemplates",
	}
	DescribeTranscodeTemplates = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "DescribeTranscodeTemplates",
	}
	DescribeTasks = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "DescribeTasks",
	}
	DescribeTaskDetail = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "DescribeTaskDetail",
	}
	DescribeSnapshotByTimeOffsetTemplates = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "DescribeSnapshotByTimeOffsetTemplates",
	}
	DescribeSampleSnapshotTemplates = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "DescribeSampleSnapshotTemplates",
	}
	DescribePersonSamples = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "DescribePersonSamples",
	}
	DescribeMediaMetaData = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "DescribeMediaMetaData",
	}
	DescribeImageSpriteTemplates = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "DescribeImageSpriteTemplates",
	}
	DescribeContentReviewTemplates = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "DescribeContentReviewTemplates",
	}
	DescribeAnimatedGraphicsTemplates = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "DescribeAnimatedGraphicsTemplates",
	}
	DescribeAdaptiveDynamicStreamingTemplates = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "DescribeAdaptiveDynamicStreamingTemplates",
	}
	DescribeAIRecognitionTemplates = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "DescribeAIRecognitionTemplates",
	}
	DescribeAIAnalysisTemplates = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "DescribeAIAnalysisTemplates",
	}
	DeleteWorkflow = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "DeleteWorkflow",
	}
	DeleteWordSamples = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "DeleteWordSamples",
	}
	DeleteWatermarkTemplate = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "DeleteWatermarkTemplate",
	}
	DeleteTranscodeTemplate = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "DeleteTranscodeTemplate",
	}
	DeleteSnapshotByTimeOffsetTemplate = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "DeleteSnapshotByTimeOffsetTemplate",
	}
	DeleteSampleSnapshotTemplate = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "DeleteSampleSnapshotTemplate",
	}
	DeletePersonSample = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "DeletePersonSample",
	}
	DeleteImageSpriteTemplate = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "DeleteImageSpriteTemplate",
	}
	DeleteContentReviewTemplate = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "DeleteContentReviewTemplate",
	}
	DeleteAnimatedGraphicsTemplate = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "DeleteAnimatedGraphicsTemplate",
	}
	DeleteAdaptiveDynamicStreamingTemplate = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "DeleteAdaptiveDynamicStreamingTemplate",
	}
	DeleteAIRecognitionTemplate = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "DeleteAIRecognitionTemplate",
	}
	DeleteAIAnalysisTemplate = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "DeleteAIAnalysisTemplate",
	}
	CreateWorkflow = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "CreateWorkflow",
	}
	CreateWordSamples = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "CreateWordSamples",
	}
	CreateWatermarkTemplate = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "CreateWatermarkTemplate",
	}
	CreateTranscodeTemplate = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "CreateTranscodeTemplate",
	}
	CreateSnapshotByTimeOffsetTemplate = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "CreateSnapshotByTimeOffsetTemplate",
	}
	CreateSampleSnapshotTemplate = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "CreateSampleSnapshotTemplate",
	}
	CreatePersonSample = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "CreatePersonSample",
	}
	CreateImageSpriteTemplate = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "CreateImageSpriteTemplate",
	}
	CreateContentReviewTemplate = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "CreateContentReviewTemplate",
	}
	CreateAnimatedGraphicsTemplate = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "CreateAnimatedGraphicsTemplate",
	}
	CreateAdaptiveDynamicStreamingTemplate = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "CreateAdaptiveDynamicStreamingTemplate",
	}
	CreateAIRecognitionTemplate = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "CreateAIRecognitionTemplate",
	}
	CreateAIAnalysisTemplate = tencentcloud.Action{
		Service: "mps",
		Version: "2019-06-12",
		Action:  "CreateAIAnalysisTemplate",
	}
)
