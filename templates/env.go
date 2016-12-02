package main

const (
	apiCallAutoTests                   = "/autotests/"
	apiCallBuildSanityTest             = "/build_sanity_test/"
	apiCallDashboard                   = "/dashboard/"
	apiCallDashboardNotifications      = "/dashboard/notifications/"
	apiCallFailures                    = "/failures/"
	apiCallFailuresByType              = "/failures/by_type/"
	apiCallFailuresCumulative          = "/failures/cumulative"
	apiCallFailuresTypes               = "/failures/types/"
	apiCallHistory                     = "/history/"
	apiCallHistoryById                 = "/history/by_id/"
	apiCallHistoryDetail               = "/history_detail/"
	apiCallInstanceConfigurations      = "/instance/configurations/"
	apiCallInstanceInfra               = "/instance/infra/"
	apiCallInstanceTargets             = "/instance/targets/"
	apiCallInstanceTargetsChoices      = "/instance/targets/choices/"
	apiCallInstanceTargetsInfo         = "/instance/targets/info/"
	apiCallJamaApiTestCases            = "/jama_api/test_cases/"
	apiCallJamaApiTestCycles           = "/jama_api/test_cycles/"
	apiCallJamaApiTestCyclesCumulative = "/jama_api/test_cycles_cumulative/"
	apiCallJamaApiTestFolders          = "/jama_api/test_folders/"
	apiCallJamaApiTestPlans            = "/jama_api/test_plans/"
	apiCallJenkinsBuilds               = "/jenkins_builds/"
	apiCallJenkinsProjects             = "/jenkins_projects/"
	apiCallJenkinsSites                = "/jenkins_sites/"
	apiCallLlatencyLog                 = "/latency_log/"
	apiCallLogsCount                   = "/logs_count/"

	apiCallManagementMamachineImage      = "/management/machine_image/"
	apiCallManagementMachineSetup        = "/management/machine_setup/"
	apiCallManagementMachineUser         = "/management/machine_user/"
	apiCallManagementPowerCocontrol      = "/management/power_control/"
	apiCallManagementPowerControlChoices = "/management/power_control/choices/"
	apiCallManagementSystemInfo          = "/management/system_info/"
	apiCallPerformance                   = "/performance/"

	apiCallReservationEexportAll      = "/reservation/export/all/"
	apiCallReservationMachineInstance = "/reservation/machine_instance/"
	apiCallReservationMy              = "/reservation/my/"
	apiCallReservationStatusChoices   = "/reservation/status/choices/"
	apiCallReservationTypeChoices     = "/reservation/type/choices/"
	apiCallReservationView            = "/reservation/view/"

	apiCallRoles                       = "/roles/"
	apiCallRunnableBuildReport         = "/runnable/build_report/"
	apiCallRunnableDeveloperTemplate   = "/runnable/developer_template/"
	apiCallRunnableDeveloperTestLatest = "/runnable/developer_test/latest/"
	apiCallRunnableReport              = "/runnable/report/"
	apiCallRunnableStepInfo            = "/runnable/step_info/"
	apiCallRunnableSystemInfo          = "/runnable/system_info/"
	apiCallRunnableSystemInfoLatest    = "/runnable/system_info/latest/"
	apiCallRunnableLoads               = "/runnable_loads/"
	apiCallRunnables                   = "/runnables/"
	apiCallRunnablesModelType          = "/runnables/model_type/"
	apiCallSearchRepo                  = "/search_repo/"
	apiCallSysLog                      = "/sys_log/"
	apiCallTestCases                   = "/test_cases/"
	apiCallTestConfig                  = "/test_config/"
	apiCallTestConfigUpload            = "/test_config_upload/"
	apiCallTestGroups                  = "/test_groups/"
	apiCallTestTypes                   = "/test_types/"
	apiCallUsageStats                  = "/usage_stats/"
	apiCallUserSettings                = "/user_settings/"
	apiCallUsers                       = "/users/"
	apiCallUsersLogout                 = "/users/logout/"
	apiCallUsersMe                     = "/users/me/"
	apiCallUsersResetPassword          = "/users/reset_password/"
)

type GoApiEnvironment struct {
	ApiCall          string // url value  "/test_cases/"
	GoApiCallName    string
	GetMethod        string
	GetOneMethod     string
	GetExampleMethod string
	PostMethod       string
	PutMethod        string
	DelMethod        string
	OptMethod        string
	PatchMethod      string
	HeadMethod       string
	WideType         string // Common type, is returned by GET all.
	ExampleType      string // Is included in WideType as slice element.
	OneGetType       string // Uses If ExampleType != OneGetType.

	ExampleName      string // Field name of slice in getted structure.
	IdField          string
	PrimaryFieldName string

	GetRejected      bool
	GetOneRejected   bool
	PostRejected     bool
	DelRejected      bool
	PatchRejected    bool
	FieldsNameMissed bool
	PutRejected      bool
	HeadRejected     bool
	OptionRejected  bool
}

var env = ReturnEPDEnv()
