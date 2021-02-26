package tcb

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	ReinstateEnv = tencentcloud.Action{
		Service: "tcb",
		Version: "2018-06-08",
		Action:  "ReinstateEnv",
	}
	ModifyEnv = tencentcloud.Action{
		Service: "tcb",
		Version: "2018-06-08",
		Action:  "ModifyEnv",
	}
	ModifyEndUser = tencentcloud.Action{
		Service: "tcb",
		Version: "2018-06-08",
		Action:  "ModifyEndUser",
	}
	ModifyDatabaseACL = tencentcloud.Action{
		Service: "tcb",
		Version: "2018-06-08",
		Action:  "ModifyDatabaseACL",
	}
	EstablishCloudBaseRunServer = tencentcloud.Action{
		Service: "tcb",
		Version: "2018-06-08",
		Action:  "EstablishCloudBaseRunServer",
	}
	DestroyStaticStore = tencentcloud.Action{
		Service: "tcb",
		Version: "2018-06-08",
		Action:  "DestroyStaticStore",
	}
	DestroyEnv = tencentcloud.Action{
		Service: "tcb",
		Version: "2018-06-08",
		Action:  "DestroyEnv",
	}
	DescribeSmsQuotas = tencentcloud.Action{
		Service: "tcb",
		Version: "2018-06-08",
		Action:  "DescribeSmsQuotas",
	}
	DescribeQuotaData = tencentcloud.Action{
		Service: "tcb",
		Version: "2018-06-08",
		Action:  "DescribeQuotaData",
	}
	DescribePostpayPackageFreeQuotas = tencentcloud.Action{
		Service: "tcb",
		Version: "2018-06-08",
		Action:  "DescribePostpayPackageFreeQuotas",
	}
	DescribePostpayFreeQuotas = tencentcloud.Action{
		Service: "tcb",
		Version: "2018-06-08",
		Action:  "DescribePostpayFreeQuotas",
	}
	DescribeExtraPkgBillingInfo = tencentcloud.Action{
		Service: "tcb",
		Version: "2018-06-08",
		Action:  "DescribeExtraPkgBillingInfo",
	}
	DescribeEnvs = tencentcloud.Action{
		Service: "tcb",
		Version: "2018-06-08",
		Action:  "DescribeEnvs",
	}
	DescribeEnvLimit = tencentcloud.Action{
		Service: "tcb",
		Version: "2018-06-08",
		Action:  "DescribeEnvLimit",
	}
	DescribeEnvFreeQuota = tencentcloud.Action{
		Service: "tcb",
		Version: "2018-06-08",
		Action:  "DescribeEnvFreeQuota",
	}
	DescribeEndUsers = tencentcloud.Action{
		Service: "tcb",
		Version: "2018-06-08",
		Action:  "DescribeEndUsers",
	}
	DescribeEndUserStatistic = tencentcloud.Action{
		Service: "tcb",
		Version: "2018-06-08",
		Action:  "DescribeEndUserStatistic",
	}
	DescribeEndUserLoginStatistic = tencentcloud.Action{
		Service: "tcb",
		Version: "2018-06-08",
		Action:  "DescribeEndUserLoginStatistic",
	}
	DescribeDownloadFile = tencentcloud.Action{
		Service: "tcb",
		Version: "2018-06-08",
		Action:  "DescribeDownloadFile",
	}
	DescribeDatabaseACL = tencentcloud.Action{
		Service: "tcb",
		Version: "2018-06-08",
		Action:  "DescribeDatabaseACL",
	}
	DescribeCloudBaseRunVersionSnapshot = tencentcloud.Action{
		Service: "tcb",
		Version: "2018-06-08",
		Action:  "DescribeCloudBaseRunVersionSnapshot",
	}
	DescribeCloudBaseRunServerVersion = tencentcloud.Action{
		Service: "tcb",
		Version: "2018-06-08",
		Action:  "DescribeCloudBaseRunServerVersion",
	}
	DescribeCloudBaseRunResourceForExtend = tencentcloud.Action{
		Service: "tcb",
		Version: "2018-06-08",
		Action:  "DescribeCloudBaseRunResourceForExtend",
	}
	DescribeCloudBaseRunResource = tencentcloud.Action{
		Service: "tcb",
		Version: "2018-06-08",
		Action:  "DescribeCloudBaseRunResource",
	}
	DescribeCloudBaseProjectLatestVersionList = tencentcloud.Action{
		Service: "tcb",
		Version: "2018-06-08",
		Action:  "DescribeCloudBaseProjectLatestVersionList",
	}
	DescribeCloudBaseBuildService = tencentcloud.Action{
		Service: "tcb",
		Version: "2018-06-08",
		Action:  "DescribeCloudBaseBuildService",
	}
	DescribeAuthDomains = tencentcloud.Action{
		Service: "tcb",
		Version: "2018-06-08",
		Action:  "DescribeAuthDomains",
	}
	DeleteEndUser = tencentcloud.Action{
		Service: "tcb",
		Version: "2018-06-08",
		Action:  "DeleteEndUser",
	}
	DeleteCloudBaseProjectLatestVersion = tencentcloud.Action{
		Service: "tcb",
		Version: "2018-06-08",
		Action:  "DeleteCloudBaseProjectLatestVersion",
	}
	CreateStaticStore = tencentcloud.Action{
		Service: "tcb",
		Version: "2018-06-08",
		Action:  "CreateStaticStore",
	}
	CreatePostpayPackage = tencentcloud.Action{
		Service: "tcb",
		Version: "2018-06-08",
		Action:  "CreatePostpayPackage",
	}
	CreateHostingDomain = tencentcloud.Action{
		Service: "tcb",
		Version: "2018-06-08",
		Action:  "CreateHostingDomain",
	}
	CreateCloudBaseRunServerVersion = tencentcloud.Action{
		Service: "tcb",
		Version: "2018-06-08",
		Action:  "CreateCloudBaseRunServerVersion",
	}
	CreateCloudBaseRunResource = tencentcloud.Action{
		Service: "tcb",
		Version: "2018-06-08",
		Action:  "CreateCloudBaseRunResource",
	}
	CreateAuthDomain = tencentcloud.Action{
		Service: "tcb",
		Version: "2018-06-08",
		Action:  "CreateAuthDomain",
	}
	CreateAndDeployCloudBaseProject = tencentcloud.Action{
		Service: "tcb",
		Version: "2018-06-08",
		Action:  "CreateAndDeployCloudBaseProject",
	}
	CommonServiceAPI = tencentcloud.Action{
		Service: "tcb",
		Version: "2018-06-08",
		Action:  "CommonServiceAPI",
	}
	CheckTcbService = tencentcloud.Action{
		Service: "tcb",
		Version: "2018-06-08",
		Action:  "CheckTcbService",
	}
)
