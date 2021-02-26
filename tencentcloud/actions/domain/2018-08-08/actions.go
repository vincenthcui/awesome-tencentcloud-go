package domain

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	UploadImage = tencentcloud.Action{
		Service: "domain",
		Version: "2018-08-08",
		Action:  "UploadImage",
	}
	UpdateProhibitionBatch = tencentcloud.Action{
		Service: "domain",
		Version: "2018-08-08",
		Action:  "UpdateProhibitionBatch",
	}
	TransferProhibitionBatch = tencentcloud.Action{
		Service: "domain",
		Version: "2018-08-08",
		Action:  "TransferProhibitionBatch",
	}
	TransferInDomainBatch = tencentcloud.Action{
		Service: "domain",
		Version: "2018-08-08",
		Action:  "TransferInDomainBatch",
	}
	SetDomainAutoRenew = tencentcloud.Action{
		Service: "domain",
		Version: "2018-08-08",
		Action:  "SetDomainAutoRenew",
	}
	RenewDomainBatch = tencentcloud.Action{
		Service: "domain",
		Version: "2018-08-08",
		Action:  "RenewDomainBatch",
	}
	ModifyDomainOwnerBatch = tencentcloud.Action{
		Service: "domain",
		Version: "2018-08-08",
		Action:  "ModifyDomainOwnerBatch",
	}
	ModifyDomainDNSBatch = tencentcloud.Action{
		Service: "domain",
		Version: "2018-08-08",
		Action:  "ModifyDomainDNSBatch",
	}
	DescribeTemplateList = tencentcloud.Action{
		Service: "domain",
		Version: "2018-08-08",
		Action:  "DescribeTemplateList",
	}
	DescribeTemplate = tencentcloud.Action{
		Service: "domain",
		Version: "2018-08-08",
		Action:  "DescribeTemplate",
	}
	DescribeDomainPriceList = tencentcloud.Action{
		Service: "domain",
		Version: "2018-08-08",
		Action:  "DescribeDomainPriceList",
	}
	DescribeDomainNameList = tencentcloud.Action{
		Service: "domain",
		Version: "2018-08-08",
		Action:  "DescribeDomainNameList",
	}
	DescribeDomainBaseInfo = tencentcloud.Action{
		Service: "domain",
		Version: "2018-08-08",
		Action:  "DescribeDomainBaseInfo",
	}
	DescribeBatchOperationLogs = tencentcloud.Action{
		Service: "domain",
		Version: "2018-08-08",
		Action:  "DescribeBatchOperationLogs",
	}
	DescribeBatchOperationLogDetails = tencentcloud.Action{
		Service: "domain",
		Version: "2018-08-08",
		Action:  "DescribeBatchOperationLogDetails",
	}
	DeleteTemplate = tencentcloud.Action{
		Service: "domain",
		Version: "2018-08-08",
		Action:  "DeleteTemplate",
	}
	CreateTemplate = tencentcloud.Action{
		Service: "domain",
		Version: "2018-08-08",
		Action:  "CreateTemplate",
	}
	CreateDomainBatch = tencentcloud.Action{
		Service: "domain",
		Version: "2018-08-08",
		Action:  "CreateDomainBatch",
	}
	CheckDomain = tencentcloud.Action{
		Service: "domain",
		Version: "2018-08-08",
		Action:  "CheckDomain",
	}
	CheckBatchStatus = tencentcloud.Action{
		Service: "domain",
		Version: "2018-08-08",
		Action:  "CheckBatchStatus",
	}
	BatchModifyDomainInfo = tencentcloud.Action{
		Service: "domain",
		Version: "2018-08-08",
		Action:  "BatchModifyDomainInfo",
	}
)
