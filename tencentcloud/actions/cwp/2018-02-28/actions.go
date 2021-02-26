package cwp

import "github.com/vincenthcui/awesome-tencentcloud-go/tencentcloud"

var (
	UntrustMalwares = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "UntrustMalwares",
	}
	UntrustMaliciousRequest = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "UntrustMaliciousRequest",
	}
	TrustMalwares = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "TrustMalwares",
	}
	TrustMaliciousRequest = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "TrustMaliciousRequest",
	}
	SwitchBashRules = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "SwitchBashRules",
	}
	SetBashEventsStatus = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "SetBashEventsStatus",
	}
	SeparateMalwares = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "SeparateMalwares",
	}
	RescanImpactedHost = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "RescanImpactedHost",
	}
	RenewProVersion = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "RenewProVersion",
	}
	RecoverMalwares = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "RecoverMalwares",
	}
	OpenProVersionPrepaid = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "OpenProVersionPrepaid",
	}
	OpenProVersion = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "OpenProVersion",
	}
	ModifyProVersionRenewFlag = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "ModifyProVersionRenewFlag",
	}
	ModifyMalwareTimingScanSettings = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "ModifyMalwareTimingScanSettings",
	}
	ModifyLoginWhiteList = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "ModifyLoginWhiteList",
	}
	ModifyAutoOpenProVersionConfig = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "ModifyAutoOpenProVersionConfig",
	}
	ModifyAlarmAttribute = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "ModifyAlarmAttribute",
	}
	MisAlarmNonlocalLoginPlaces = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "MisAlarmNonlocalLoginPlaces",
	}
	InquiryPriceOpenProVersionPrepaid = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "InquiryPriceOpenProVersionPrepaid",
	}
	IgnoreImpactedHosts = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "IgnoreImpactedHosts",
	}
	ExportTasks = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "ExportTasks",
	}
	ExportReverseShellEvents = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "ExportReverseShellEvents",
	}
	ExportPrivilegeEvents = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "ExportPrivilegeEvents",
	}
	ExportNonlocalLoginPlaces = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "ExportNonlocalLoginPlaces",
	}
	ExportMalwares = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "ExportMalwares",
	}
	ExportMaliciousRequests = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "ExportMaliciousRequests",
	}
	ExportBruteAttacks = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "ExportBruteAttacks",
	}
	ExportBashEvents = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "ExportBashEvents",
	}
	ExportAttackLogs = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "ExportAttackLogs",
	}
	EditTags = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "EditTags",
	}
	EditReverseShellRule = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "EditReverseShellRule",
	}
	EditPrivilegeRule = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "EditPrivilegeRule",
	}
	EditBashRule = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "EditBashRule",
	}
	DescribeWeeklyReports = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeWeeklyReports",
	}
	DescribeWeeklyReportVuls = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeWeeklyReportVuls",
	}
	DescribeWeeklyReportNonlocalLoginPlaces = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeWeeklyReportNonlocalLoginPlaces",
	}
	DescribeWeeklyReportMalwares = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeWeeklyReportMalwares",
	}
	DescribeWeeklyReportInfo = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeWeeklyReportInfo",
	}
	DescribeWeeklyReportBruteAttacks = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeWeeklyReportBruteAttacks",
	}
	DescribeVuls = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeVuls",
	}
	DescribeVulScanResult = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeVulScanResult",
	}
	DescribeVulInfo = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeVulInfo",
	}
	DescribeUsualLoginPlaces = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeUsualLoginPlaces",
	}
	DescribeTags = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeTags",
	}
	DescribeTagMachines = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeTagMachines",
	}
	DescribeSecurityTrends = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeSecurityTrends",
	}
	DescribeSecurityEventsCnt = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeSecurityEventsCnt",
	}
	DescribeSecurityDynamics = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeSecurityDynamics",
	}
	DescribeScanMalwareSchedule = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeScanMalwareSchedule",
	}
	DescribeReverseShellRules = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeReverseShellRules",
	}
	DescribeReverseShellEvents = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeReverseShellEvents",
	}
	DescribeProcesses = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeProcesses",
	}
	DescribeProcessTaskStatus = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeProcessTaskStatus",
	}
	DescribeProcessStatistics = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeProcessStatistics",
	}
	DescribeProVersionInfo = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeProVersionInfo",
	}
	DescribePrivilegeRules = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribePrivilegeRules",
	}
	DescribePrivilegeEvents = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribePrivilegeEvents",
	}
	DescribeOverviewStatistics = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeOverviewStatistics",
	}
	DescribeOpenPorts = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeOpenPorts",
	}
	DescribeOpenPortTaskStatus = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeOpenPortTaskStatus",
	}
	DescribeOpenPortStatistics = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeOpenPortStatistics",
	}
	DescribeNonlocalLoginPlaces = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeNonlocalLoginPlaces",
	}
	DescribeMalwares = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeMalwares",
	}
	DescribeMalwareInfo = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeMalwareInfo",
	}
	DescribeMaliciousRequests = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeMaliciousRequests",
	}
	DescribeMachines = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeMachines",
	}
	DescribeMachineList = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeMachineList",
	}
	DescribeMachineInfo = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeMachineInfo",
	}
	DescribeLoginWhiteList = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeLoginWhiteList",
	}
	DescribeImpactedHosts = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeImpactedHosts",
	}
	DescribeHistoryAccounts = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeHistoryAccounts",
	}
	DescribeGeneralStat = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeGeneralStat",
	}
	DescribeExportMachines = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeExportMachines",
	}
	DescribeComponents = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeComponents",
	}
	DescribeComponentStatistics = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeComponentStatistics",
	}
	DescribeComponentInfo = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeComponentInfo",
	}
	DescribeBruteAttacks = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeBruteAttacks",
	}
	DescribeBashRules = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeBashRules",
	}
	DescribeBashEvents = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeBashEvents",
	}
	DescribeAttackLogs = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeAttackLogs",
	}
	DescribeAttackLogInfo = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeAttackLogInfo",
	}
	DescribeAlarmAttribute = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeAlarmAttribute",
	}
	DescribeAgentVuls = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeAgentVuls",
	}
	DescribeAccounts = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeAccounts",
	}
	DescribeAccountStatistics = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DescribeAccountStatistics",
	}
	DeleteUsualLoginPlaces = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DeleteUsualLoginPlaces",
	}
	DeleteTags = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DeleteTags",
	}
	DeleteReverseShellRules = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DeleteReverseShellRules",
	}
	DeleteReverseShellEvents = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DeleteReverseShellEvents",
	}
	DeletePrivilegeRules = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DeletePrivilegeRules",
	}
	DeletePrivilegeEvents = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DeletePrivilegeEvents",
	}
	DeleteNonlocalLoginPlaces = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DeleteNonlocalLoginPlaces",
	}
	DeleteMalwares = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DeleteMalwares",
	}
	DeleteMaliciousRequests = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DeleteMaliciousRequests",
	}
	DeleteMachineTag = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DeleteMachineTag",
	}
	DeleteMachine = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DeleteMachine",
	}
	DeleteLoginWhiteList = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DeleteLoginWhiteList",
	}
	DeleteBruteAttacks = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DeleteBruteAttacks",
	}
	DeleteBashRules = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DeleteBashRules",
	}
	DeleteBashEvents = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DeleteBashEvents",
	}
	DeleteAttackLogs = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "DeleteAttackLogs",
	}
	CreateUsualLoginPlaces = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "CreateUsualLoginPlaces",
	}
	CreateProcessTask = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "CreateProcessTask",
	}
	CreateOpenPortTask = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "CreateOpenPortTask",
	}
	CreateBaselineStrategy = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "CreateBaselineStrategy",
	}
	CloseProVersion = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "CloseProVersion",
	}
	AddMachineTag = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "AddMachineTag",
	}
	AddLoginWhiteList = tencentcloud.Action{
		Service: "cwp",
		Version: "2018-02-28",
		Action:  "AddLoginWhiteList",
	}
)
