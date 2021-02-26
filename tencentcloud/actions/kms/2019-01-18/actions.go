package kms

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	VerifyByAsymmetricKey = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "VerifyByAsymmetricKey",
	}
	UpdateKeyDescription = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "UpdateKeyDescription",
	}
	UpdateAlias = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "UpdateAlias",
	}
	UnbindCloudResource = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "UnbindCloudResource",
	}
	SignByAsymmetricKey = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "SignByAsymmetricKey",
	}
	ScheduleKeyDeletion = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "ScheduleKeyDeletion",
	}
	ReEncrypt = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "ReEncrypt",
	}
	OverwriteWhiteBoxDeviceFingerprints = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "OverwriteWhiteBoxDeviceFingerprints",
	}
	ListKeys = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "ListKeys",
	}
	ListKeyDetail = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "ListKeyDetail",
	}
	ListAlgorithms = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "ListAlgorithms",
	}
	ImportKeyMaterial = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "ImportKeyMaterial",
	}
	GetServiceStatus = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "GetServiceStatus",
	}
	GetRegions = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "GetRegions",
	}
	GetPublicKey = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "GetPublicKey",
	}
	GetParametersForImport = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "GetParametersForImport",
	}
	GetKeyRotationStatus = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "GetKeyRotationStatus",
	}
	GenerateRandom = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "GenerateRandom",
	}
	GenerateDataKey = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "GenerateDataKey",
	}
	EncryptByWhiteBox = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "EncryptByWhiteBox",
	}
	Encrypt = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "Encrypt",
	}
	EnableWhiteBoxKeys = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "EnableWhiteBoxKeys",
	}
	EnableWhiteBoxKey = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "EnableWhiteBoxKey",
	}
	EnableKeys = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "EnableKeys",
	}
	EnableKeyRotation = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "EnableKeyRotation",
	}
	EnableKey = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "EnableKey",
	}
	DisableWhiteBoxKeys = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "DisableWhiteBoxKeys",
	}
	DisableWhiteBoxKey = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "DisableWhiteBoxKey",
	}
	DisableKeys = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "DisableKeys",
	}
	DisableKeyRotation = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "DisableKeyRotation",
	}
	DisableKey = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "DisableKey",
	}
	DescribeWhiteBoxServiceStatus = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "DescribeWhiteBoxServiceStatus",
	}
	DescribeWhiteBoxKeyDetails = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "DescribeWhiteBoxKeyDetails",
	}
	DescribeWhiteBoxKey = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "DescribeWhiteBoxKey",
	}
	DescribeWhiteBoxDeviceFingerprints = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "DescribeWhiteBoxDeviceFingerprints",
	}
	DescribeWhiteBoxDecryptKey = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "DescribeWhiteBoxDecryptKey",
	}
	DescribeKeys = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "DescribeKeys",
	}
	DescribeKey = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "DescribeKey",
	}
	DeleteWhiteBoxKey = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "DeleteWhiteBoxKey",
	}
	DeleteImportedKeyMaterial = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "DeleteImportedKeyMaterial",
	}
	Decrypt = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "Decrypt",
	}
	CreateWhiteBoxKey = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "CreateWhiteBoxKey",
	}
	CreateKey = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "CreateKey",
	}
	CancelKeyDeletion = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "CancelKeyDeletion",
	}
	CancelKeyArchive = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "CancelKeyArchive",
	}
	BindCloudResource = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "BindCloudResource",
	}
	AsymmetricSm2Decrypt = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "AsymmetricSm2Decrypt",
	}
	AsymmetricRsaDecrypt = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "AsymmetricRsaDecrypt",
	}
	ArchiveKey = tencentcloud.Action{
		Service: "kms",
		Version: "2019-01-18",
		Action:  "ArchiveKey",
	}
)
