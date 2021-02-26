package live

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	UpdateLiveWatermark = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "UpdateLiveWatermark",
	}
	UnBindLiveDomainCert = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "UnBindLiveDomainCert",
	}
	StopRecordTask = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "StopRecordTask",
	}
	StopLiveRecord = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "StopLiveRecord",
	}
	ResumeLiveStream = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "ResumeLiveStream",
	}
	ResumeDelayLiveStream = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "ResumeDelayLiveStream",
	}
	ModifyPullStreamStatus = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "ModifyPullStreamStatus",
	}
	ModifyPullStreamConfig = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "ModifyPullStreamConfig",
	}
	ModifyLiveTranscodeTemplate = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "ModifyLiveTranscodeTemplate",
	}
	ModifyLiveSnapshotTemplate = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "ModifyLiveSnapshotTemplate",
	}
	ModifyLiveRecordTemplate = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "ModifyLiveRecordTemplate",
	}
	ModifyLivePushAuthKey = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "ModifyLivePushAuthKey",
	}
	ModifyLivePlayDomain = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "ModifyLivePlayDomain",
	}
	ModifyLivePlayAuthKey = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "ModifyLivePlayAuthKey",
	}
	ModifyLiveDomainCert = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "ModifyLiveDomainCert",
	}
	ModifyLiveCert = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "ModifyLiveCert",
	}
	ModifyLiveCallbackTemplate = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "ModifyLiveCallbackTemplate",
	}
	ForbidLiveStream = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "ForbidLiveStream",
	}
	ForbidLiveDomain = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "ForbidLiveDomain",
	}
	EnableLiveDomain = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "EnableLiveDomain",
	}
	DropLiveStream = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DropLiveStream",
	}
	DescribeVisitTopSumInfoList = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeVisitTopSumInfoList",
	}
	DescribeUploadStreamNums = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeUploadStreamNums",
	}
	DescribeTopClientIpSumInfoList = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeTopClientIpSumInfoList",
	}
	DescribeStreamPushInfoList = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeStreamPushInfoList",
	}
	DescribeStreamPlayInfoList = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeStreamPlayInfoList",
	}
	DescribeStreamDayPlayInfoList = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeStreamDayPlayInfoList",
	}
	DescribeScreenShotSheetNumList = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeScreenShotSheetNumList",
	}
	DescribePullStreamConfigs = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribePullStreamConfigs",
	}
	DescribeProvinceIspPlayInfoList = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeProvinceIspPlayInfoList",
	}
	DescribeProIspPlaySumInfoList = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeProIspPlaySumInfoList",
	}
	DescribePlayErrorCodeSumInfoList = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribePlayErrorCodeSumInfoList",
	}
	DescribePlayErrorCodeDetailInfoList = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribePlayErrorCodeDetailInfoList",
	}
	DescribeLogDownloadList = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeLogDownloadList",
	}
	DescribeLiveWatermarks = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeLiveWatermarks",
	}
	DescribeLiveWatermarkRules = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeLiveWatermarkRules",
	}
	DescribeLiveWatermark = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeLiveWatermark",
	}
	DescribeLiveTranscodeTemplates = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeLiveTranscodeTemplates",
	}
	DescribeLiveTranscodeTemplate = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeLiveTranscodeTemplate",
	}
	DescribeLiveTranscodeRules = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeLiveTranscodeRules",
	}
	DescribeLiveTranscodeDetailInfo = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeLiveTranscodeDetailInfo",
	}
	DescribeLiveStreamState = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeLiveStreamState",
	}
	DescribeLiveStreamPushInfoList = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeLiveStreamPushInfoList",
	}
	DescribeLiveStreamPublishedList = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeLiveStreamPublishedList",
	}
	DescribeLiveStreamOnlineList = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeLiveStreamOnlineList",
	}
	DescribeLiveStreamEventList = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeLiveStreamEventList",
	}
	DescribeLiveSnapshotTemplates = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeLiveSnapshotTemplates",
	}
	DescribeLiveSnapshotTemplate = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeLiveSnapshotTemplate",
	}
	DescribeLiveSnapshotRules = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeLiveSnapshotRules",
	}
	DescribeLiveRecordTemplates = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeLiveRecordTemplates",
	}
	DescribeLiveRecordTemplate = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeLiveRecordTemplate",
	}
	DescribeLiveRecordRules = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeLiveRecordRules",
	}
	DescribeLivePushAuthKey = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeLivePushAuthKey",
	}
	DescribeLivePlayAuthKey = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeLivePlayAuthKey",
	}
	DescribeLivePackageInfo = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeLivePackageInfo",
	}
	DescribeLiveForbidStreamList = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeLiveForbidStreamList",
	}
	DescribeLiveDomains = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeLiveDomains",
	}
	DescribeLiveDomainPlayInfoList = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeLiveDomainPlayInfoList",
	}
	DescribeLiveDomainCert = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeLiveDomainCert",
	}
	DescribeLiveDomain = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeLiveDomain",
	}
	DescribeLiveDelayInfoList = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeLiveDelayInfoList",
	}
	DescribeLiveCerts = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeLiveCerts",
	}
	DescribeLiveCert = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeLiveCert",
	}
	DescribeLiveCallbackTemplates = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeLiveCallbackTemplates",
	}
	DescribeLiveCallbackTemplate = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeLiveCallbackTemplate",
	}
	DescribeLiveCallbackRules = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeLiveCallbackRules",
	}
	DescribeHttpStatusInfoList = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeHttpStatusInfoList",
	}
	DescribeGroupProIspPlayInfoList = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeGroupProIspPlayInfoList",
	}
	DescribeDeliverBandwidthList = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeDeliverBandwidthList",
	}
	DescribeConcurrentRecordStreamNum = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeConcurrentRecordStreamNum",
	}
	DescribeCallbackRecordsList = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeCallbackRecordsList",
	}
	DescribeBillBandwidthAndFluxList = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeBillBandwidthAndFluxList",
	}
	DescribeAreaBillBandwidthAndFluxList = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeAreaBillBandwidthAndFluxList",
	}
	DescribeAllStreamPlayInfoList = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DescribeAllStreamPlayInfoList",
	}
	DeleteRecordTask = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DeleteRecordTask",
	}
	DeletePullStreamConfig = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DeletePullStreamConfig",
	}
	DeleteLiveWatermarkRule = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DeleteLiveWatermarkRule",
	}
	DeleteLiveWatermark = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DeleteLiveWatermark",
	}
	DeleteLiveTranscodeTemplate = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DeleteLiveTranscodeTemplate",
	}
	DeleteLiveTranscodeRule = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DeleteLiveTranscodeRule",
	}
	DeleteLiveSnapshotTemplate = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DeleteLiveSnapshotTemplate",
	}
	DeleteLiveSnapshotRule = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DeleteLiveSnapshotRule",
	}
	DeleteLiveRecordTemplate = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DeleteLiveRecordTemplate",
	}
	DeleteLiveRecordRule = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DeleteLiveRecordRule",
	}
	DeleteLiveRecord = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DeleteLiveRecord",
	}
	DeleteLiveDomain = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DeleteLiveDomain",
	}
	DeleteLiveCert = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DeleteLiveCert",
	}
	DeleteLiveCallbackTemplate = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DeleteLiveCallbackTemplate",
	}
	DeleteLiveCallbackRule = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "DeleteLiveCallbackRule",
	}
	CreateRecordTask = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "CreateRecordTask",
	}
	CreatePullStreamConfig = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "CreatePullStreamConfig",
	}
	CreateLiveWatermarkRule = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "CreateLiveWatermarkRule",
	}
	CreateLiveTranscodeTemplate = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "CreateLiveTranscodeTemplate",
	}
	CreateLiveTranscodeRule = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "CreateLiveTranscodeRule",
	}
	CreateLiveSnapshotTemplate = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "CreateLiveSnapshotTemplate",
	}
	CreateLiveSnapshotRule = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "CreateLiveSnapshotRule",
	}
	CreateLiveRecordTemplate = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "CreateLiveRecordTemplate",
	}
	CreateLiveRecordRule = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "CreateLiveRecordRule",
	}
	CreateLiveRecord = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "CreateLiveRecord",
	}
	CreateLiveCert = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "CreateLiveCert",
	}
	CreateLiveCallbackTemplate = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "CreateLiveCallbackTemplate",
	}
	CreateLiveCallbackRule = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "CreateLiveCallbackRule",
	}
	CreateCommonMixStream = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "CreateCommonMixStream",
	}
	CancelCommonMixStream = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "CancelCommonMixStream",
	}
	BindLiveDomainCert = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "BindLiveDomainCert",
	}
	AddLiveWatermark = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "AddLiveWatermark",
	}
	AddLiveDomain = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "AddLiveDomain",
	}
	AddDelayLiveStream = tencentcloud.Action{
		Service: "live",
		Version: "2018-08-01",
		Action:  "AddDelayLiveStream",
	}
)
