package main

func ReturnStellusEnv() []GoApiEnvironment {
	env := []GoApiEnvironment{

		{
			ApiCall:       "/reservation/machine_instance/",
			GoApiCallName: "apiCallReservationMachineInstance",
			GetMethod:     "GetReservationMachineInstance",
			GetOneMethod:  "GetOneReservationMachineInstance",
			//		PostMethod:    "PostReservationMachineInstance",
			//		DelMethod:     "DelReservationMachineInstance",
			//		PatchMethod:   "PatchReservationMachineInstance",
			WideType:    "ReservationMachineInstanceAPI",
			ExampleType: "MachineInstance",

			ExampleName:      "MachineInstancies", // name of array included in API structure
			IdField:          "InstanceID",
			PrimaryFieldName: "IP",
		},
		{
			ApiCall:       "/reservation/my/",
			GoApiCallName: "apiCallReservationMy",
			GetMethod:     "GetReservationMy",
			WideType:      "ReservationMyAPI",
			//		ExampleName:   "Current",
		},
		{
			ApiCall:       "/reservation/status/choices/",
			GoApiCallName: "apiCallReservationStatusChoices",
			GetMethod:     "GetReservationStatusChoices",
			WideType:      "map[string]string",
		},

		{
			ApiCall:       "/reservation/type/choices/",
			GoApiCallName: "apiCallReservationTypeChoices",
			GetMethod:     "GetReservationTypeChoices",
			WideType:      "map[string]string",
		},

		{
			ApiCall:       "/reservation/view/",
			GoApiCallName: "apiCallReservationView",
			GetMethod:     "GetReservationReservationView",
			GetOneMethod:  "GetOneReservationReservationView",
			PostMethod:    "PostReservationReservationView",
			DelMethod:     "DelReservationReservationView",
			PatchMethod:   "PatchReservationReservationView",
			WideType:      "ReservationViewAPI",
			ExampleType:   "ReservationView",

			ExampleName:      "ReservationsView", // name of array included in API structure
			IdField:          "ReservationID",
			PrimaryFieldName: "Title",
		},
	}

	/*
	   var env = []GoApiEnvironment{
	   	{
	   		ApiCall:       "/jenkins_builds/",
	   		GoApiCallName: "apiCallJenkinsBuilds",
	   		GetMethod:     "GetJenkinsBuilds",
	   		GetOneMethod:  "GetOneJenkinsBuilds",
	   //		PostMethod:    "PostJenkinsBuilds",
	   //		DelMethod:     "DelJenkinsBuilds",
	   //		PatchMethod:   "PatchJenkinsBuilds",
	   		WideType:      "JenkinsBuildAPI",
	   		ExampleType:   "JenkinsBuild",

	   		ExampleName:      "JenkinsBuilds", // name of array included in API structure
	   		IdField:          "BuildID",
	   		PrimaryFieldName: "BuildURL",
	   	},

	   //	{
	   //		ApiCall:       "/jenkins_builds/configuration/",
	   //		GoApiCallName: "apiCallJenkinsBuildsConfigurations",
	   //		GetMethod:     "GetJenkinsBuildsConfigurations",
	   //		WideType:      "JenkinsBuildsConfigurationsAPI",
	   //		ExampleName:   "BuildStream",
	   //	},

	   	{
	   		ApiCall:       "/jenkins_projects/",
	   		GoApiCallName: "apiCallJenkinsProjects",
	   		GetMethod:     "GetJenkinsProjects",
	   		GetOneMethod:  "GetOneJenkinsProjects",
	   //		PostMethod:    "PostJenkinsProjects",
	   //		DelMethod:     "DelJenkinsProjects",
	   //		PatchMethod:   "PatchJenkinsProjects",
	   		WideType:      "JenkinsProjectsAPI",
	   		ExampleType:   "JenkinsProject",

	   		ExampleName:      "JenkinsProjects", // name of array included in API structure
	   		IdField:          "JenkinsProjectID",
	   		PrimaryFieldName: "Name",
	   	},
	   	{
	   		ApiCall:       "/jenkins_sites/",
	   		GoApiCallName: "apiCallJenkinsSites",
	   		GetMethod:     "GetJenkinsSites",
	   		GetOneMethod:  "GetOneJenkinsSites",
	   		PostMethod:    "PostJenkinsSites",
	   		DelMethod:     "DelJenkinsSites",
	   		PatchMethod:   "PatchJenkinsSites",
	   		WideType:      "JenkinsSitesAPI",
	   		ExampleType:   "JenkinsSite",

	   		ExampleName:      "JenkinsSites", // name of array included in API structure
	   		IdField:          "SiteID",
	   		PrimaryFieldName: "URL",
	   	},
	   }
	   /*
	   var env = []GoApiEnvironment{
	   		{
	   		ApiCall:       "/dashboard/notification/",
	   		GoApiCallName: "apiCallDashboardNotifications",
	   		GetMethod:     "GetDashboardNotifications",
	   		GetOneMethod:  "GetOneDashboardNotifications",
	   //		PostMethod:    "PostDashboardNotifications",
	   //		DelMethod:     "DelDashboardNotifications",
	   //		PatchMethod:   "PatchDashboardNotifications",
	   		WideType:      "DashboardNotificationsAPI",
	   		ExampleType:   "DashboardNotification", // is included in WideType

	   		ExampleName:      "DashboardNotifications", // name in structure
	   		IdField:          "NotificationsID",
	   		PrimaryFieldName: "TypeName",
	   	},
	   	{
	   		ApiCall:       "/history/by_id/",
	   		GoApiCallName: "apiCallHistoryById",
	   		GetMethod:     "GetHistoryById",
	   //		GetOneMethod:  "GetOneHistoryById",
	   //		PostMethod:    "PostHistoryById",
	   //		DelMethod:     "DelHistoryById",
	   		PatchMethod:   "PatchHistoryById",
	   		WideType:      "HistoryByIdAPI",
	   		ExampleType:   "HistoryById",
	   		//		OneGetType:    "HistoryOne",

	   		ExampleName:      "HistoriesById",
	   		IdField:          "HistoryID",
	   		PrimaryFieldName: "TestGroupName",

	   		PostRejected:  true,
	   		DelRejected:   true,
	   		PatchRejected: true,
	   	},
	   		{
	   		ApiCall:       "/history/",
	   		GoApiCallName: "apiCallHistory",
	   		GetMethod:     "GetHistory",
	   		GetOneMethod:  "GetOneHistory",
	   //		PostMethod:    "PostHistory",
	   //		DelMethod:     "DelHistory",
	   //		PatchMethod:   "PatchHistory",
	   		WideType:      "HistoryAPI",
	   		ExampleType:   "History",
	   		//		OneGetType:    "HistoryOne",

	   		ExampleName:      "Histories",
	   		IdField:          "RunnableID",
	   		PrimaryFieldName: "RunnableName",
	   	},
	   	{
	   		ApiCall:       "/history_detail/",
	   		GoApiCallName: "apiCallHistoryDetail",
	   		GetMethod:     "GetHistoryDetail",
	   		GetOneMethod:  "GetOneHistoryDetail",
	   //		PostMethod:    "PostHistoryDetail",
	   //		DelMethod:     "DelHistoryDetail",
	   //		PatchMethod:   "PatchHistoryDetail",
	   		WideType:      "HistoryDetailAPI",
	   		ExampleType:   "HistoryDetail",
	   		//		OneGetType:    "HistoryOne",

	   		ExampleName:      "HistoriesDetail", // name of array included in API structure
	   		IdField:          "HistoryDetailID",
	   		PrimaryFieldName: "BuildName",

	   	},
	   	{
	   		ApiCall:       "/usage_stats/",
	   		GoApiCallName: "apiCallUsageStats",
	   		GetMethod:     "GetTestUsageStats",
	   		GetOneMethod:  "GetOneUsageStats",
	   //		PostMethod:    "PostUsageStats",
	   //		DelMethod:     "DelUsageStats",
	   //		PatchMethod:   "PatchUsageStats",
	   		WideType:      "UsageStatsAPI",
	   		ExampleType:   "UsageStat",

	   		ExampleName:      "UsageStats",
	   		IdField:          "Count",
	   		PrimaryFieldName: "APIURL",

	   	},


	   }
	   /*
	   {

	   		ApiCall:       "/runnable/step_info/",
	   		GoApiCallName: "apiCallRunnableStepInfo",
	   		GetMethod:     "GetRunnableStepInfo",
	   		GetOneMethod:  "GetOneRunnableStepInfo",
	   //		PostMethod:    "PostRunnableStepInfo",
	   //		DelMethod:     "DelRunnableStepInfo",
	   //		PatchMethod:   "PatchRunnableStepInfo",
	   		WideType:      "RunnableStepInfoAPI",
	   		ExampleType:   "StepInfo",

	   		ExampleName:      "StepsInfo", // name of array included in API structure
	   		IdField:          "StepKey",
	   		PrimaryFieldName: "StepInfo",

	   	},
	   	{
	   		ApiCall:       "/runnable/system_info/",
	   		GoApiCallName: "apiCallRunnableSystemInfo",
	   		GetMethod:     "GetRunnableSystemInfo",
	   		GetOneMethod:  "GetOneRunnableSystemInfo",
	   //		PostMethod:    "PostRunnableSystemInfo",
	   //		DelMethod:     "DelRunnableSystemInfo",
	   //		PatchMethod:   "PatchRunnableSystemInfo",
	   		WideType:      "RunnableSystemInfoAPI",
	   		ExampleType:   "RunnableSystemInfo",

	   		ExampleName:      "RunnablesSystemInfo", // name of array included in API structure
	   		IdField:          "ID",
	   		PrimaryFieldName: "Name",
	   	},



	   }
	   /*
	   	{
	   		ApiCall:       "/jama_api/search_folder/",
	   		GoApiCallName: "apiCallJamaApiSerarchFolder",
	   		GetMethod:     "GetJamaApiSerarchFolder",
	   		GetOneMethod: "GetOneJamaApiSerarchFolder",
	   		WideType:      "JamaApiSerarchFolderAPI",
	   		ExampleType:   "JamaApiSerarchFolder",

	   		ExampleName:      "Folders", // name of array included in API structure
	   		IdField:          "ID",
	   		PrimaryFieldName: "Name",
	   	},

	   	{
	   		ApiCall:       "/jama_api/test_cases/",
	   		GoApiCallName: "apiCallJamaApiTestCases",
	   		GetMethod:     "GetJamaApiTestCases",
	   		GetOneMethod:  "GetOneJamaApiTestCases",
	   		PostMethod:    "PostJamaApiTestCases",
	   		DelMethod:     "DelJamaApiTestCases",
	   		PatchMethod:   "PatchJamaApiTestCases",
	   		WideType:      "JamaApiTestCasesAPI",
	   		ExampleType:   "JamaTestCaseType",

	   		ExampleName:      "JamaApiTestCases", // name of array included in API structure
	   		IdField:          "ID",
	   		PrimaryFieldName: "Name",
	   	},

	   	{
	   		ApiCall:       "/jama_api/test_cycles/",
	   		GoApiCallName: "apiCallJamaApiTestCycles",
	   		GetMethod:     "GetJamaApiTestCycles",
	   		GetOneMethod:  "GetOneJamaApiTestCycles",
	   		PostMethod:    "PostJamaApiTestCycles",
	   		DelMethod:     "DelJamaApiTestCycles",
	   		PatchMethod:   "PatchJamaApiTestCycles",
	   		WideType:      "JamaApiTestCyclesAPI",
	   		ExampleType:   "JamaApiTestCycle",

	   		ExampleName:      "JamaApiTestCycles", // name of array included in API structure
	   		IdField:          "ID",
	   		PrimaryFieldName: "Name",
	   	},

	   	{
	   		ApiCall:       "/jama_api/test_cycles_cumulative/",
	   		GoApiCallName: "apiCallJamaApiTestCyclesCumulative",
	   		GetMethod:     "GetJamaApiTestCyclesCumulative",
	   		WideType:      "JamaApiTestCyclesClumulativeAPI",
	   		ExampleName:   "Clumulative",
	   	},

	   	{
	   		ApiCall:       "/jama_api/test_folders/",
	   		GoApiCallName: "apiCallJamaApiTestFolders",
	   		GetMethod:     "GetJamaApiTestFolders",
	   		GetOneMethod:  "GetOneJamaApiTestFolders",
	   		PostMethod:    "PostJamaApiTestFolders",
	   		DelMethod:     "DelJamaApiTestFolders",
	   		PatchMethod:   "PatchJamaApiTestFolders",
	   		WideType:      "JamaApiTestFoldersAPI",
	   		ExampleType:   "JamaApiTestFolder",

	   		ExampleName:      "JamaApiTestFolders", // name of array included in API structure
	   		IdField:          "ID",
	   		PrimaryFieldName: "Name",
	   	},
	   	{
	   		ApiCall:       "/jama_api/test_plans/",
	   		GoApiCallName: "apiCallJamaApiTestPlans",
	   		GetMethod:     "GetJamaApiTestPlans",
	   		GetOneMethod:     "GetOneaApiTestPlans",
	   		WideType:      "JamaApiTestPlansAPI",
	   		ExampleType:      "JamaTestPlanType ",


	   		ExampleName:   "TestPlans",
	   		IdField:          "ID",
	   		PrimaryFieldName: "Name",


	   	},
	   }



	   /*
	   	{
	   		ApiCall:       "/autotests/",
	   		GoApiCallName: "apiCallAutoTests",
	   		GetMethod:     "GetAutoTests",
	   		GetOneMethod:  "GetOneAutoTests",
	   		PostMethod:    "PostAutoTest",
	   		DelMethod:     "DelAutoTest",
	   		PatchMethod:   "PatchAutoTest",
	   		WideType:      "AutoTestsAPI",
	   		ExampleType:   "AutoTest",

	   		ExampleName:      "AutoTests", // name in structure
	   		IdField:          "AutotestID",
	   		PrimaryFieldName: "LastRunTimeStr",
	   	},
	   	{
	   		ApiCall:       "/test_cases/",
	   		GoApiCallName: "apiCallTestCases",
	   		GetMethod:     "GetTestCases",
	   		GetOneMethod:  "GetOneTestCases",
	   		PostMethod:    "PostTestCases",
	   		DelMethod:     "DelTestCases",
	   		PatchMethod:   "PatchTestCases",
	   		WideType:      "TestCasesAPI",
	   		ExampleType:   "TestCase",

	   		ExampleName:      "TestCases",
	   		IdField:          "TestCaseID",
	   		PrimaryFieldName: "Name",
	   	},
	   	{
	   		ApiCall:       "/test_config/",
	   		GoApiCallName: "apiCallTestConfig",
	   		GetMethod:     "GetTestConfig",
	   		GetOneMethod:  "GetOneTestConfig",
	   		PostMethod:    "PostTestConfig",
	   		DelMethod:     "DelTestConfig",
	   		PatchMethod:   "PatchTestConfigs",
	   		WideType:      "TestConfigAPI",
	   		ExampleType:   "TestConfigType",

	   		ExampleName:      "TestConfigs", // name of array included in API structure
	   		IdField:          "TestConfigID",
	   		PrimaryFieldName: "ConfigContent",
	   	},
	   	{
	   		ApiCall:       "/test_groups/",
	   		GoApiCallName: "apiCallTestGroups",
	   		GetMethod:     "GetTestGroups",
	   		GetOneMethod:  "GetOneTestGroups",
	   		PostMethod:    "PostTestGroups",
	   		DelMethod:     "DelTestGroups",
	   		PatchMethod:   "PatchTestGroups",
	   		WideType:      "TestGroupsAPI",
	   		ExampleType:   "TestGroup",

	   		ExampleName:      "TestGroups", // name of array included in API structure
	   		IdField:          "TestGroupID",
	   		PrimaryFieldName: "Name",
	   	},

	   }
	   /*
	   	{
	   		ApiCall:       "/runnable/developer_template/",
	   		GoApiCallName: "apiCallRunnableDeveloperTemplate",
	   		GetMethod:     "GetRunnableDeveloperTemplate",
	   		GetOneMethod:  "GetOneRunnableDeveloperTemplate",
	   		PostMethod:    "PostRunnableDeveloperTemplate",
	   		PatchMethod:   "PatchRunnableDeveloperTemplate",
	   		DelMethod:     "DelRunnableDeveloperTemplate",

	   		WideType:    "RunnableDeveloperTemplateAPI",
	   		ExampleType: "RunnableDeveloperTemplateAPI",

	   		ExampleName:      "XXX",
	   		IdField:          "RunnableID",
	   		PrimaryFieldName: "Owner",
	   	},

	   	{
	   		ApiCall:       "/runnable/developer_test/latest/",
	   		GoApiCallName: "apiCallRunnableDeveloperTestLatest",
	   		GetMethod:     "GetRunnableDeveloperTestLatest",
	   		WideType:      "RunnableDeveloperTestLatestAPI",
	   		ExampleName:   "Message",
	   	},


	   	{

	   		ApiCall:       "/runnable/step_info/",
	   		GoApiCallName: "apiCallRunnableStepInfo",
	   		GetMethod:     "GetRunnableStepInfo",
	   		GetOneMethod:  "GetOneRunnableStepInfo",
	   		PostMethod:    "PostRunnableStepInfo",
	   		DelMethod:     "DelRunnableStepInfo",
	   		PatchMethod:   "PatchRunnableStepInfo",
	   		WideType:      "RunnableStepInfoAPI",
	   		ExampleType:   "StepInfo",

	   		ExampleName:      "StepsInfo", // name of array included in API structure
	   		IdField:          "Order",
	   		PrimaryFieldName: "StepInfoStr",

	   		PostRejected:  true,
	   		DelRejected:   true,
	   		PatchRejected: true,
	   	},
	   	{
	   		ApiCall:       "/runnable/loads/",
	   		GoApiCallName: "apiCallRunnableLoads",
	   		GetMethod:     "GetRunnableLoads",
	   		GetOneMethod:  "GetOneRunnableLoads",
	   		PostMethod:    "PostRunnableLoads",
	   		DelMethod:     "DelRunnableLoads",
	   		PatchMethod:   "PatchRunnableLoads",
	   		WideType:      "RunnableLoadsAPI",
	   		ExampleType:   "RunnableLoadType",

	   		ExampleName:      "RunnableLoads", // name of array included in API structure
	   		IdField:          "RunnableLoadID",
	   		PrimaryFieldName: "Name",
	   	},
	   }

	   /*

	   	{
	   		ApiCall:          "/management/machine_image/",
	   		GoApiCallName:    "apiCallManagementMamachineImage",
	   		GetMethod:        "GetMamachineImage",
	   		GetOneMethod:     "GetOneMamachineImage",
	   		PostMethod:       "PostMamachineImage",
	   		DelMethod:        "DelMamachineImage",
	   		PatchMethod:      "PatchMamachineImage",
	   		WideType:         "MamachineImageType",
	   		ExampleType:      "MamachineImageType",

	   		ExampleName:      "MachineAct",
	   		IdField:          "ActID",
	   		PrimaryFieldName: "ActName",

	   	},
	   	{
	   		ApiCall:          "/management/machine_setup/",
	   		GoApiCallName:    "apiCallManagementMachineSetup",
	   		GetMethod:        "GetManagementMachineSetup",
	   		GetOneMethod:     "GetOneManagementMachineSetup",
	   		PostMethod:       "PostManagementMachineSetup",
	   		DelMethod:        "DelManagementMachineSetup",
	   		PatchMethod:      "PatchManagementMachineSetup",
	   		WideType:         "MachineSetupType",
	   		ExampleType:      "MachineSetupType",
	   		ExampleName:      "MachineAct",
	   		IdField:          "ActID",
	   		PrimaryFieldName: "ActName",

	   	},

	   	{
	   		ApiCall:          "/management/machine_user/",
	   		GoApiCallName:    "apiCallManagementMachineUser",
	   		GetMethod:        "GetManagementMachineUser",
	   		GetOneMethod:     "GetOneManagementMachineUser",
	   		PostMethod:       "PostManagementMachineUser",
	   		DelMethod:        "DelManagementMachineUser",
	   		PatchMethod:      "PatchManagementMachineUser",
	   		WideType:         "MachineUserType",
	   		ExampleType:      "MachineUserType",
	   		ExampleName:      "MachineAct",
	   		IdField:          "ActID",
	   		PrimaryFieldName: "ActName",

	   	},

	   	{
	   		ApiCall:       "/management/power_control/",
	   		GoApiCallName: "apiCallManagementPowerCocontrol",
	   		GetMethod:     "GetManagementPowerCocontrol",
	   		GetOneMethod:  "GetOneManagementPowerCocontrol",
	   		PostMethod:    "PostManagementPowerCocontrol",
	   		DelMethod:     "DelManagementPowerCocontrol",
	   		PatchMethod:   "PatchManagementPowerCocontrol",
	   		WideType:      "PowerCocontrolType",
	   		ExampleType:   "PowerCocontrolType",

	   		ExampleName:      "MachineAct",
	   		IdField:          "ActID",
	   		PrimaryFieldName: "ActName",
	   	},

	   	{
	   		ApiCall:       "/management/power_control/choices/",
	   		GoApiCallName: "apiCallManagementPowerControlChoices",
	   		GetMethod:     "GetManagementPowerControlChoices",
	   		WideType:      "map[string]string",
	   	},
	   	{
	   		ApiCall:       "/management/system_info/",
	   		GoApiCallName: "apiCallManagementSystemInfo",
	   		GetMethod:     "ManagementSystemInfoAPI",
	   		WideType:      "[]byte",
	   	},
	   }

	   /*
	   var env = []GoApiEnvironment{
	   	{
	   		ApiCall:       "/performance/",
	   		GoApiCallName: "apiCallPerformance",
	   		GetMethod:     "GetPerformance",
	   		WideType:      "[]byte",
	   	},
	   	// query?
	   	{
	   		ApiCall:       "/latency_log/",
	   		GoApiCallName: "apiCallLatencyLog",
	   		GetMethod:     "GetLatencyLog",
	   		WideType:      "[]byte",
	   	},
	   		// pk?
	   	{
	   		ApiCall:       "/logs_count/",
	   		GoApiCallName: "apiCallLogsCount",
	   		GetMethod:     "GetLogsCount",
	   		WideType:      "[]byte",
	   	},
	   	// fail!
	   	{
	   		ApiCall:       "/roles/",
	   		GoApiCallName: "apiCallRoles",
	   		GetMethod:     "GetRoles",
	   		WideType:      "[]byte",
	   	},
	   	// no needed patametr!
	   	{
	   		ApiCall:       "/search_repo/",
	   		GoApiCallName: "apiCallSearchRepo",
	   		GetMethod:     "GetSearchRepo",
	   		WideType:      "[]byte",
	   	},
	   	// detail not found!
	   	{
	   		ApiCall:       "/sys_log/",
	   		GoApiCallName: "apiCallSysLog",
	   		GetMethod:     "GetSysLog",
	   	vng]string ar env = []GoApiEnvironment{
	   	{
	   		ApiCall:       "/management/machine_image/",
	   		GoApiCallName: "apiCallManagementMamachineImage",
	   		GetMethod:     "GetMamachineImage",
	   		GetOneMethod:  "GetOneMamachineImage",
	   		PostMethod:    "PostMamachineImage",
	   		DelMethod:     "DelMamachineImage",
	   		PatchMethod:   "PatchMamachineImage",
	   		WideType:      "MamachineImageType",

	   		//		ExampleName:      "MachineImage", // name of array included in API structure
	   		//		IdField:          "InstanceID",
	   		//		PrimaryFieldName: "IP",
	   		//		WideType:         "[]byte",
	   	},
	   	{
	   		ApiCall:       "/management/machine_setup/",
	   		GoApiCallName: "apiCallManagementMachineSetup",
	   		GetMethod:     "GetManagementMachineSetup",
	   		GetOneMethod:  "GetOneManagementMachineSetup",
	   		PostMethod:    "PostManagementMachineSetup",
	   		DelMethod:     "DelManagementMachineSetup",
	   		PatchMethod:   "PatchManagementMachineSetup",
	   		WideType:      "MachineSetupReceive",
	   		ExampleType:   "MachineSetupPost",

	   		//		ExampleName:      "MachineInstacies", // name of array included in API structure
	   		//		IdField:          "InstanceID",
	   		//		PrimaryFieldName: "IP",
	   		//		WideType:         "[]byte",
	   	},

	   	{
	   		ApiCall:       "/management/machine_user/",
	   		GoApiCallName: "apiCallManagementMachineUser",
	   		GetMethod:     "GetManagementMachineUser",
	   		GetOneMethod:  "GetOneManagementMachineUser",
	   		PostMethod:    "PostManagementMachineUser",
	   		DelMethod:     "DelManagementMachineUser",
	   		PatchMethod:   "PatchManagementMachineUser",
	   		WideType:      "ManagementMachineReceive",
	   		ExampleType:   "ManagementMachinePost",

	   		//		ExampleName:      "MachineInstacies", // name of array included in API structure
	   		//		IdField:          "InstanceID",
	   		//		PrimaryFieldName: "IP",
	   		//		WideType:         "[]byte",
	   	},

	   	{
	   		ApiCall:       "/management/power_control/",
	   		GoApiCallName: "apiCallManagementPowerCocontrol",
	   		GetMethod:     "GetManagementPowerCocontrol",
	   		GetOneMethod:  "GetOneManagementPowerCocontrol",
	   		PostMethod:    "PostManagementPowerCocontrol",
	   		DelMethod:     "DelManagementPowerCocontrol",
	   		PatchMethod:   "PatchManagementPowerCocontrol",
	   		WideType:      "PowerCocontrolReceive",
	   		ExampleType:   "MachineInstacePost",

	   		//		ExampleName:      "MachineInstacies", // name of array included in API structure
	   		//		IdField:          "InstanceID",
	   		//		PrimaryFieldName: "IP",
	   	},

	   	{
	   		ApiCall:       "/management/power_control/choices/",
	   		GoApiCallName: "apiCallManagementPowerControlChoices",
	   		GetMethod:     "GetManagementPowerControlChoices",
	   		WideType:      "map[string]string",
	   	},
	   	{
	   		ApiCall:       "/management/system_info/",
	   		GoApiCallName: "apiCallManagementSystemInfo",
	   		GetMethod:     "ManagementSystemInfoAPI",
	   		WideType:      "[]byte",
	   	},
	   }

	   	WideType:      "[]byte",
	   	},
	   	//
	   	{
	   		ApiCall:       "/usage_stats/",
	   		GoApiCallName: "apiCallUsageStats",
	   		GetMethod:     "GetUsageStats",
	   		WideType:      "[]byte",
	   	},
	   }

	   /*
	   /*


	   /*

	*/
	/*
	 */
	/*
		{
			ApiCall:       "/instance/configurations/",
			GoApiCallName: "apiCallInstanceConfigurations",
			GetMethod:     "GetInstanceConfigurations",
			WideType:      "InstanceConfigurationsAPI",
			ExampleName:   "VssdFirmware",
		},

		{
			ApiCall:       "/instance/targets/",
			GoApiCallName: "apiCallInstanceTargets",
			GetMethod:     "GetInstanceTarget",
			GetOneMethod:  "GetOneInstanceTarget",
			PostMethod:    "PostInstanceTarget",
			DelMethod:     "DelInstanceTarget",
			PatchMethod:   "PatchInstanceTarget",
			WideType:      "InstanceTargetAPI",
			ExampleType:   "Target",

			ExampleName:      "Targets", // name of array included in API structure
			IdField:          "InstanceID",
			PrimaryFieldName: "Name",

		},


		{
			ApiCall:       "/instance/targets/",
			GoApiCallName: "apiCallInstanceTargets",
			GetMethod:     "GetInstanceTarget",
			GetOneMethod:  "GetOneInstanceTarget",
			PostMethod:    "PostInstanceTarget",
			DelMethod:     "DelInstanceTarget",
			PatchMethod:   "PatchInstanceTarget",
			WideType:      "InstanceTargetAPI",
			ExampleType:   "Target",

			ExampleName:      "Targets", // name of array included in API structure
			IdField:          "InstanceID",
			PrimaryFieldName: "Name",

			//		PostRejected:  true,
			DelRejected:   true,
			PatchRejected: true,
		},


		{
			ApiCall:       "/runnables/system_info/latest/",
			GoApiCallName: "apiCallRunnableSystemInfoLatest",
			GetMethod:     "GetRunnableSystemInfoLatest",
			WideType:      "RunnableSystemInfoLatestAPI",
			ExampleType:   "",
		},

		{ // map
			ApiCall:       "/runnables/model_type/",
			GoApiCallName: "apiCallRunnablesModelType",
			GetMethod:     "GetRunnablesModelType",
			WideType:      "map[string]string",

			ExampleName: "XXX",
		},

		{
			ApiCall:       "/users/",
			GoApiCallName: "apiCallUsers",

			GetMethod:    "GetTestUsers",
			GetOneMethod: "GetOneUsers",
			PostMethod:   "PostUsers",
			DelMethod:    "DelUsers",
			PatchMethod:  "PatchUsers",

			WideType:    "UsersAPI",
			ExampleType: "User",

			ExampleName:      "Users",
			IdField:          "ID",
			PrimaryFieldName: "Username",
		},
		{
			ApiCall:       "/build_sanity_test/",
			GoApiCallName: "apiCallBuildSanityTest",
			GetMethod:     "GetBuildSanityTest",
			WideType:      "BuildSanityTestAPI",

			ExampleName: "Status",
		},

		{
			ApiCall:       "/dashboard/",
			GoApiCallName: "apiCallDashboard",
			GetMethod:     "GetDashboard",
			WideType:      "DashboardAPI",
			ExampleName:   "TotalUsers",
		},

		{
			ApiCall:       "/failures/by_type/",
			GoApiCallName: "apiFailuresByType",
			GetMethod:     "GetFailuresByType",
			WideType:      "FailuresByTypeAPI",
			ExampleName:   "XXX",
		},

		/* made manually
		{   // All and One types are not similar
			ApiCall:       "/failures/",
			GoApiCallName: "apiCallFailures",
			GetMethod:     "GetFailures",
			GetOneMethod:  "GetOneFailures",
			WideType:      "FailuresAPI",
			OneGetType:    "FailuresID",

			ExampleName:      "History",
			IdField:          "TestCaseID",
			PrimaryFieldName: "Name",
		},
		{  	// look Arrays
			ApiCall:       "/failures/cumulative/",
			GoApiCallName: "apiCallFailuresCumulative",
			GetMethod:     "GetFailuresCumulativeAr",
			WideType:      "FailuresCumulative",

			ExampleName: "Total",
		},
	*/
	/* made manually
	   	{ // array
	   		ApiCall:       "/instance/infra/",
	   		GoApiCallName: "apiCallInstanceInfra",
	   		GetMethod:     "GetInstanceInfra",
	   		WideType:      "InstanceInfraAPI",

	   		ExampleName: "IP",
	   	},
	   	{ // map
	   		ApiCall:       "/instance/targets/choices/",
	   		GoApiCallName: "apiCallInstanceTargetsChoices",
	   		GetMethod:     "GetInstanceTargetsChoices",
	   		//		WideType:      "InstanceTargetChoicesAPI",
	   		WideType: "map[string]string",

	   		ExampleName: "IP",
	   	},
	   	{
	   		ApiCall:       "/instance/targets/info/",
	   		GoApiCallName: "apiCallInstanceTargetsInfo",
	   		GetMethod:     "GetInstanceTargetsInfo2",
	   		WideType:      "InstanceTargetsInfoAPI",
	   		ExampleName:   "SystemInfo.Running",
	   	},
	   	{ // look Arrays
	   		ApiCall:       "/test_types/",
	   		GoApiCallName: "apiCallTestTypes",
	   		GetMethod:     "GetTestTypes",
	   		GetOneMethod:  "GetOneTestTypes",
	   		WideType:      "TestTypesAPI",
	   		OneGetType:    "TestTypeConfig",

	   		ExampleName: "Configs.TestTypeName",
	   	},
	   	{
	   		ApiCall:       "/user_settings/",
	   		GoApiCallName: "apiCallUserSettings",
	   		GetMethod:     "GetUserSettings",
	   		//		GetOneMethod:  "GetOneUserSettings",
	   		PostMethod:       "PostUserSettings",
	   		PatchMethod:      "PatchUserSettings",
	   		WideType:         "UserSettingsAPI",
	   		ExampleType:      "UserSettingsAPI",
	   		ExampleName:      "KeyValue",
	   		IdField:          "UserSettingID",
	   		PrimaryFieldName: "ColorSchemaStr",
	   	},

	   	{
	   		ApiCall:       "/users/me/",
	   		GoApiCallName: "apiCallUsersMe",
	   		GetMethod:     "GetUsersMe",
	   		GetOneMethod:  "GetOneUsersMe",
	   		PostMethod:    "PostUsersMe",
	   		PatchMethod:   "PatchUsersMe",
	   		WideType:      "UsersMeAPI",
	   		ExampleType:   "UsersMeAPI",
	   		ExampleName:   "Username",

	   		IdField:          "ID",
	   		PrimaryFieldName: "Role",
	   	},
	   }
	*/

	return env
}
