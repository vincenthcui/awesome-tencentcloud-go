package vod

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	WeChatMiniProgramPublish = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "WeChatMiniProgramPublish",
	}
	SplitMedia = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "SplitMedia",
	}
	SimpleHlsClip = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "SimpleHlsClip",
	}
	SearchMedia = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "SearchMedia",
	}
	ResetProcedureTemplate = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "ResetProcedureTemplate",
	}
	PushUrlCache = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "PushUrlCache",
	}
	PullUpload = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "PullUpload",
	}
	PullEvents = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "PullEvents",
	}
	ProcessMediaByUrl = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "ProcessMediaByUrl",
	}
	ProcessMediaByProcedure = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "ProcessMediaByProcedure",
	}
	ProcessMedia = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "ProcessMedia",
	}
	ParseStreamingManifest = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "ParseStreamingManifest",
	}
	ModifyWordSample = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "ModifyWordSample",
	}
	ModifyWatermarkTemplate = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "ModifyWatermarkTemplate",
	}
	ModifyTranscodeTemplate = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "ModifyTranscodeTemplate",
	}
	ModifySuperPlayerConfig = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "ModifySuperPlayerConfig",
	}
	ModifySubAppIdStatus = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "ModifySubAppIdStatus",
	}
	ModifySubAppIdInfo = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "ModifySubAppIdInfo",
	}
	ModifySnapshotByTimeOffsetTemplate = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "ModifySnapshotByTimeOffsetTemplate",
	}
	ModifySampleSnapshotTemplate = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "ModifySampleSnapshotTemplate",
	}
	ModifyPersonSample = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "ModifyPersonSample",
	}
	ModifyMediaInfo = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "ModifyMediaInfo",
	}
	ModifyImageSpriteTemplate = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "ModifyImageSpriteTemplate",
	}
	ModifyContentReviewTemplate = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "ModifyContentReviewTemplate",
	}
	ModifyClass = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "ModifyClass",
	}
	ModifyAnimatedGraphicsTemplate = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "ModifyAnimatedGraphicsTemplate",
	}
	ModifyAdaptiveDynamicStreamingTemplate = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "ModifyAdaptiveDynamicStreamingTemplate",
	}
	ModifyAIRecognitionTemplate = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "ModifyAIRecognitionTemplate",
	}
	ModifyAIAnalysisTemplate = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "ModifyAIAnalysisTemplate",
	}
	ManageTask = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "ManageTask",
	}
	LiveRealTimeClip = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "LiveRealTimeClip",
	}
	ForbidMediaDistribution = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "ForbidMediaDistribution",
	}
	ExecuteFunction = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "ExecuteFunction",
	}
	EditMedia = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "EditMedia",
	}
	DescribeWordSamples = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "DescribeWordSamples",
	}
	DescribeWatermarkTemplates = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "DescribeWatermarkTemplates",
	}
	DescribeTranscodeTemplates = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "DescribeTranscodeTemplates",
	}
	DescribeTasks = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "DescribeTasks",
	}
	DescribeTaskDetail = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "DescribeTaskDetail",
	}
	DescribeSuperPlayerConfigs = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "DescribeSuperPlayerConfigs",
	}
	DescribeSubAppIds = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "DescribeSubAppIds",
	}
	DescribeStorageDetails = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "DescribeStorageDetails",
	}
	DescribeStorageData = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "DescribeStorageData",
	}
	DescribeSnapshotByTimeOffsetTemplates = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "DescribeSnapshotByTimeOffsetTemplates",
	}
	DescribeSampleSnapshotTemplates = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "DescribeSampleSnapshotTemplates",
	}
	DescribeReviewDetails = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "DescribeReviewDetails",
	}
	DescribeProcedureTemplates = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "DescribeProcedureTemplates",
	}
	DescribePersonSamples = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "DescribePersonSamples",
	}
	DescribeMediaProcessUsageData = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "DescribeMediaProcessUsageData",
	}
	DescribeMediaInfos = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "DescribeMediaInfos",
	}
	DescribeImageSpriteTemplates = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "DescribeImageSpriteTemplates",
	}
	DescribeImageProcessingTemplates = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "DescribeImageProcessingTemplates",
	}
	DescribeEventsState = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "DescribeEventsState",
	}
	DescribeDailyPlayStatFileList = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "DescribeDailyPlayStatFileList",
	}
	DescribeContentReviewTemplates = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "DescribeContentReviewTemplates",
	}
	DescribeCdnLogs = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "DescribeCdnLogs",
	}
	DescribeCDNUsageData = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "DescribeCDNUsageData",
	}
	DescribeCDNStatDetails = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "DescribeCDNStatDetails",
	}
	DescribeAnimatedGraphicsTemplates = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "DescribeAnimatedGraphicsTemplates",
	}
	DescribeAllClass = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "DescribeAllClass",
	}
	DescribeAdaptiveDynamicStreamingTemplates = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "DescribeAdaptiveDynamicStreamingTemplates",
	}
	DescribeAIRecognitionTemplates = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "DescribeAIRecognitionTemplates",
	}
	DescribeAIAnalysisTemplates = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "DescribeAIAnalysisTemplates",
	}
	DeleteWordSamples = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "DeleteWordSamples",
	}
	DeleteWatermarkTemplate = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "DeleteWatermarkTemplate",
	}
	DeleteTranscodeTemplate = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "DeleteTranscodeTemplate",
	}
	DeleteSuperPlayerConfig = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "DeleteSuperPlayerConfig",
	}
	DeleteSnapshotByTimeOffsetTemplate = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "DeleteSnapshotByTimeOffsetTemplate",
	}
	DeleteSampleSnapshotTemplate = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "DeleteSampleSnapshotTemplate",
	}
	DeleteProcedureTemplate = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "DeleteProcedureTemplate",
	}
	DeletePersonSample = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "DeletePersonSample",
	}
	DeleteMedia = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "DeleteMedia",
	}
	DeleteImageSpriteTemplate = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "DeleteImageSpriteTemplate",
	}
	DeleteImageProcessingTemplate = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "DeleteImageProcessingTemplate",
	}
	DeleteContentReviewTemplate = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "DeleteContentReviewTemplate",
	}
	DeleteClass = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "DeleteClass",
	}
	DeleteAnimatedGraphicsTemplate = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "DeleteAnimatedGraphicsTemplate",
	}
	DeleteAdaptiveDynamicStreamingTemplate = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "DeleteAdaptiveDynamicStreamingTemplate",
	}
	DeleteAIRecognitionTemplate = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "DeleteAIRecognitionTemplate",
	}
	DeleteAIAnalysisTemplate = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "DeleteAIAnalysisTemplate",
	}
	CreateWordSamples = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "CreateWordSamples",
	}
	CreateWatermarkTemplate = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "CreateWatermarkTemplate",
	}
	CreateTranscodeTemplate = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "CreateTranscodeTemplate",
	}
	CreateSuperPlayerConfig = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "CreateSuperPlayerConfig",
	}
	CreateSubAppId = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "CreateSubAppId",
	}
	CreateSnapshotByTimeOffsetTemplate = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "CreateSnapshotByTimeOffsetTemplate",
	}
	CreateSampleSnapshotTemplate = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "CreateSampleSnapshotTemplate",
	}
	CreateProcedureTemplate = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "CreateProcedureTemplate",
	}
	CreatePersonSample = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "CreatePersonSample",
	}
	CreateImageSpriteTemplate = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "CreateImageSpriteTemplate",
	}
	CreateImageProcessingTemplate = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "CreateImageProcessingTemplate",
	}
	CreateContentReviewTemplate = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "CreateContentReviewTemplate",
	}
	CreateClass = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "CreateClass",
	}
	CreateAnimatedGraphicsTemplate = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "CreateAnimatedGraphicsTemplate",
	}
	CreateAdaptiveDynamicStreamingTemplate = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "CreateAdaptiveDynamicStreamingTemplate",
	}
	CreateAIRecognitionTemplate = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "CreateAIRecognitionTemplate",
	}
	CreateAIAnalysisTemplate = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "CreateAIAnalysisTemplate",
	}
	ConfirmEvents = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "ConfirmEvents",
	}
	ComposeMedia = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "ComposeMedia",
	}
	CommitUpload = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "CommitUpload",
	}
	ApplyUpload = tencentcloud.Action{
		Service: "vod",
		Version: "2018-07-17",
		Action:  "ApplyUpload",
	}
)
