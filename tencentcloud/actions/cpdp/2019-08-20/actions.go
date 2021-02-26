package cpdp

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	WithdrawCashMembership = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "WithdrawCashMembership",
	}
	UnifiedOrder = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "UnifiedOrder",
	}
	UnbindRelateAcct = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "UnbindRelateAcct",
	}
	UnBindAcct = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "UnBindAcct",
	}
	TransferSinglePay = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "TransferSinglePay",
	}
	RevokeRechargeByThirdPay = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "RevokeRechargeByThirdPay",
	}
	RevokeMemberRechargeThirdPay = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "RevokeMemberRechargeThirdPay",
	}
	ReviseMbrProperty = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "ReviseMbrProperty",
	}
	RevResigterBillSupportWithdraw = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "RevResigterBillSupportWithdraw",
	}
	RegisterBillSupportWithdraw = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "RegisterBillSupportWithdraw",
	}
	RegisterBill = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "RegisterBill",
	}
	RefundMemberTransaction = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "RefundMemberTransaction",
	}
	Refund = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "Refund",
	}
	RechargeMemberThirdPay = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "RechargeMemberThirdPay",
	}
	RechargeByThirdPay = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "RechargeByThirdPay",
	}
	QueryTransferResult = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "QueryTransferResult",
	}
	QueryTransferDetail = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "QueryTransferDetail",
	}
	QueryTransferBatch = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "QueryTransferBatch",
	}
	QueryTrade = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "QueryTrade",
	}
	QuerySmallAmountTransfer = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "QuerySmallAmountTransfer",
	}
	QuerySingleTransactionStatus = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "QuerySingleTransactionStatus",
	}
	QuerySinglePay = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "QuerySinglePay",
	}
	QueryRefund = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "QueryRefund",
	}
	QueryReconciliationDocument = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "QueryReconciliationDocument",
	}
	QueryPayerInfo = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "QueryPayerInfo",
	}
	QueryOutwardOrder = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "QueryOutwardOrder",
	}
	QueryOrder = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "QueryOrder",
	}
	QueryMerchantInfoForManagement = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "QueryMerchantInfoForManagement",
	}
	QueryMerchantBalance = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "QueryMerchantBalance",
	}
	QueryMemberTransaction = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "QueryMemberTransaction",
	}
	QueryMemberBind = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "QueryMemberBind",
	}
	QueryInvoice = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "QueryInvoice",
	}
	QueryExchangeRate = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "QueryExchangeRate",
	}
	QueryCustAcctIdBalance = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "QueryCustAcctIdBalance",
	}
	QueryCommonTransferRecharge = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "QueryCommonTransferRecharge",
	}
	QueryBillDownloadURL = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "QueryBillDownloadURL",
	}
	QueryBankWithdrawCashDetails = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "QueryBankWithdrawCashDetails",
	}
	QueryBankTransactionDetails = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "QueryBankTransactionDetails",
	}
	QueryBankClear = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "QueryBankClear",
	}
	QueryBalance = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "QueryBalance",
	}
	QueryApplicationMaterial = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "QueryApplicationMaterial",
	}
	QueryAnchorContractInfo = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "QueryAnchorContractInfo",
	}
	QueryAgentTaxPaymentBatch = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "QueryAgentTaxPaymentBatch",
	}
	QueryAgentStatements = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "QueryAgentStatements",
	}
	QueryAcctInfoList = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "QueryAcctInfoList",
	}
	QueryAcctInfo = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "QueryAcctInfo",
	}
	QueryAcctBinding = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "QueryAcctBinding",
	}
	ModifyMntMbrBindRelateAcctBankCode = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "ModifyMntMbrBindRelateAcctBankCode",
	}
	ModifyAgentTaxPaymentInfo = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "ModifyAgentTaxPaymentInfo",
	}
	MigrateOrderRefundQuery = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "MigrateOrderRefundQuery",
	}
	MigrateOrderRefund = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "MigrateOrderRefund",
	}
	ExecuteMemberTransaction = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "ExecuteMemberTransaction",
	}
	DownloadBill = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "DownloadBill",
	}
	DescribeOrderStatus = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "DescribeOrderStatus",
	}
	DescribeChargeDetail = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "DescribeChargeDetail",
	}
	DeleteAgentTaxPaymentInfos = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "DeleteAgentTaxPaymentInfos",
	}
	DeleteAgentTaxPaymentInfo = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "DeleteAgentTaxPaymentInfo",
	}
	CreateTransferBatch = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "CreateTransferBatch",
	}
	CreateSinglePay = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "CreateSinglePay",
	}
	CreateRedInvoice = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "CreateRedInvoice",
	}
	CreateMerchant = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "CreateMerchant",
	}
	CreateInvoice = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "CreateInvoice",
	}
	CreateCustAcctId = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "CreateCustAcctId",
	}
	CreateAgentTaxPaymentInfos = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "CreateAgentTaxPaymentInfos",
	}
	CreateAcct = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "CreateAcct",
	}
	CloseOrder = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "CloseOrder",
	}
	CheckAmount = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "CheckAmount",
	}
	CheckAcct = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "CheckAcct",
	}
	BindRelateAcctUnionPay = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "BindRelateAcctUnionPay",
	}
	BindRelateAcctSmallAmount = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "BindRelateAcctSmallAmount",
	}
	BindRelateAccReUnionPay = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "BindRelateAccReUnionPay",
	}
	BindAcct = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "BindAcct",
	}
	ApplyWithdrawal = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "ApplyWithdrawal",
	}
	ApplyTrade = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "ApplyTrade",
	}
	ApplyReWithdrawal = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "ApplyReWithdrawal",
	}
	ApplyPayerInfo = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "ApplyPayerInfo",
	}
	ApplyOutwardOrder = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "ApplyOutwardOrder",
	}
	ApplyApplicationMaterial = tencentcloud.Action{
		Service: "cpdp",
		Version: "2019-08-20",
		Action:  "ApplyApplicationMaterial",
	}
)
