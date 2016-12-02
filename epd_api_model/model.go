package epd_api_model
const (
	AuditHistoryViewCall = "audit_history/view"
	ChoicesBlockSizeCall = "choices/block_size"
	ChoicesCapacityCall = "choices/capacity"
	ChoicesLifeCycleCall = "choices/life_cycle"
	ChoicesPortCall = "choices/port"
	ChoicesTestDurationTypeCall = "choices/test_duration_type"
	ChoicesTestPurposeCall = "choices/test_purpose"
	DashboardCall = "dashboard"
	DashboardNotificationsCall = "dashboard/notifications"
	ExportHelpFileUploadCall = "export_help_file/upload"
	ImportFileUploadCall = "import_file/upload"
	KeyForCallsViewCall = "key_for_calls/view"
	RolesCall = "roles"
	SettingsEnvCall = "settings/env"
	SubscriptionGroupViewCall = "subscription/group/view"
	SubscriptionUserViewCall = "subscription/user/view"
	TctDetailViewCall = "tct_detail/view"
	TestCaseCategoriesCall = "test_case/categories"
	TestCaseViewCall = "test_case/view"
	TestCycleReportCall = "test_cycle/report"
	TestCycleViewCall = "test_cycle/view"
	TestCycleTemplateViewCall = "test_cycle_template/view"
	TestGroupViewCall = "test_group/view"
	TestLauncherLibraryCall = "test_launcher/library"
	TestLauncherLibraryDependencyCall = "test_launcher/library_dependency"
	TestLauncherTestPeripheralDependencyCall = "test_launcher/test_peripheral_dependency"
	TestLauncherTestPeripheralTypeCall = "test_launcher/test_peripheral_type"
	TestLauncherTestToolCall = "test_launcher/test_tool"
	TestLauncherToolDependencyCall = "test_launcher/tool_dependency"
	TestLauncherViewCall = "test_launcher/view"
	TestPlanViewCall = "test_plan/view"
	TestProjectViewCall = "test_project/view"
	TestResultChoicesStatusCall = "test_result/choices/status"
	TestResultCommentTagViewCall = "test_result/comment_tag/view"
	TestResultFioTestResultViewCall = "test_result/fio_test_result/view"
	TestResultIometerTestReportViewCall = "test_result/iometer_test_report/view"
	TestResultIometerTestResultViewCall = "test_result/iometer_test_result/view"
	TestResultLogCall = "test_result/log"
	TestResultReportTemplateViewCall = "test_result/report_template/view"
	TestResultTestRunResultCommentViewCall = "test_result/test_run_result_comment/view"
	TestResultViewCall = "test_result/view"
	TestRunDuplicateCall = "test_run/duplicate"
	TestRunViewCall = "test_run/view"
	TestSystemChoicesDataScopeCall = "test_system/choices/data_scope"
	TestSystemChoicesDeviceUnderTestingBackendCall = "test_system/choices/device_under_testing_backend"
	TestSystemChoicesDeviceUnderTestingCapacityCall = "test_system/choices/device_under_testing_capacity"
	TestSystemChoicesDeviceUnderTestingCustomerCall = "test_system/choices/device_under_testing_customer"
	TestSystemChoicesDeviceUnderTestingFormFactorCall = "test_system/choices/device_under_testing_form_factor"
	TestSystemChoicesDeviceUnderTestingFrontendCall = "test_system/choices/device_under_testing_frontend"
	TestSystemChoicesDeviceUnderTestingHealthStatusCall = "test_system/choices/device_under_testing_health_status"
	TestSystemChoicesDeviceUnderTestingHostInterfaceCall = "test_system/choices/device_under_testing_host_interface"
	TestSystemChoicesDeviceUnderTestingProductCall = "test_system/choices/device_under_testing_product"
	TestSystemChoicesDeviceUnderTestingStatusCall = "test_system/choices/device_under_testing_status"
	TestSystemChoicesDownloadTypeCall = "test_system/choices/download_type"
	TestSystemChoicesFirmwareTypeCall = "test_system/choices/firmware_type"
	TestSystemChoicesNandTypeCall = "test_system/choices/nand_type"
	TestSystemChoicesNandVersionCall = "test_system/choices/nand_version"
	TestSystemChoicesOperatingSystemCall = "test_system/choices/operating_system"
	TestSystemChoicesOverProvisioningCall = "test_system/choices/over_provisioning"
	TestSystemChoicesPortCall = "test_system/choices/port"
	TestSystemChoicesRaidCall = "test_system/choices/raid"
	TestSystemChoicesResultCall = "test_system/choices/result"
	TestSystemChoicesSecurityCall = "test_system/choices/security"
	TestSystemChoicesSpeedCall = "test_system/choices/speed"
	TestSystemChoicesStatusCall = "test_system/choices/status"
	TestSystemChoicesStorageInterfaceFormFactorCall = "test_system/choices/storage_interface_form_factor"
	TestSystemChoicesTestProjectStatusCall = "test_system/choices/test_project_status"
	TestSystemChoicesTestRunPriorityStatusCall = "test_system/choices/test_run_priority_status"
	TestSystemChoicesTestRunResultCall = "test_system/choices/test_run_result"
	TestSystemChoicesTestToolTypeCall = "test_system/choices/test_tool_type"
	TestSystemChoicesTestTypeCall = "test_system/choices/test_type"
	TestSystemControllerCall = "test_system/controller"
	TestSystemDeviceUnderTestingCall = "test_system/device_under_testing"
	TestSystemFirmwareCall = "test_system/firmware"
	TestSystemNodeCall = "test_system/node"
	TestSystemOperatingSystemCall = "test_system/operating_system"
	TestSystemProductCall = "test_system/product"
	TestSystemProductFamilyCall = "test_system/product_family"
	TestSystemProductLinkCall = "test_system/product_link"
	TestSystemQualPlanCall = "test_system/qual_plan"
	TestSystemStorageDriverCall = "test_system/storage_driver"
	TestSystemStorageInterfaceCall = "test_system/storage_interface"
	TestSystemStorageInterfaceTypeCall = "test_system/storage_interface_type"
	TestSystemTestPeripheralCall = "test_system/test_peripheral"
	TestSystemTestRequestCall = "test_system/test_request"
	TestSystemTestRequestDetailCall = "test_system/test_request_detail"
	TestSystemViewCall = "test_system/view"
	TriggerEventChoicesCall = "trigger_event_choices"
	UsageStatsCall = "usage_stats"
	UserGroupCall = "user_group"
	UserSettingsCall = "user_settings"
	UserUploadViewCall = "user_upload/view"
	UserUploadViewTypeCall = "user_upload/view_type"
	UsersCall = "users"
	UsersMeCall = "users/me"
	UsersResetPasswordCall = "users/reset_password"
	V2FirmwareReportCall = "v2/firmware_report"

)


type AuditHistoryViewType struct {
	Count    int64       `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		AuditHistoryID int64 `json:"audit_history_id"`
		Changes        struct {
			Action               string  `json:"action"`
			DataScopes           []int64 `json:"data_scopes"`
			Description          string  `json:"description"`
			DeviceUnderTestingID int64   `json:"device_under_testing_id"`
			FirmwareID           int64   `json:"firmware_id"`
			IsActive             bool    `json:"is_active"`
			Name                 string  `json:"name"`
			RerunNum             int64   `json:"rerun_num"`
			SlotNumber           int64   `json:"slot_number"`
			TestCycleID          int64   `json:"test_cycle_id"`
			TestLauncherID       int64   `json:"test_launcher_id"`
			TestPlanID           int64   `json:"test_plan_id"`
			TestProjectID        int64   `json:"test_project_id"`
			TestRunID            int64   `json:"test_run_id"`
			TestToolID           int64   `json:"test_tool_id"`
		} `json:"changes"`
		DateModify     string `json:"date_modify"`
		KeyCall        int64  `json:"key_call"`
		KeyCallStr     string `json:"key_call_str"`
		TypeChanges    int64  `json:"type_changes"`
		TypeChangesStr string `json:"type_changes_str"`
		UserModify     struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"user_modify"`
	} `json:"results"`
	Total int64 `json:"total"`
}


type ChoicesBlockSizeType struct {
	Count    int64       `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		BlockSizeChoiceID int64   `json:"block_size_choice_id"`
		DataScopes        []int64 `json:"data_scopes"`
		Name              string  `json:"name"`
		Order             int64   `json:"order"`
		Value             int64   `json:"value"`
	} `json:"results"`
	Total int64 `json:"total"`
}


type ChoicesCapacityType struct {
	Count    int64       `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		CapacityID int64   `json:"capacity_id"`
		DataScopes []int64 `json:"data_scopes"`
		Name       string  `json:"name"`
		Order      int64   `json:"order"`
		Value      int64   `json:"value"`
	} `json:"results"`
	Total int64 `json:"total"`
}


type ChoicesLifeCycleType struct {
	Count    int64       `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		DataScopes  []int64 `json:"data_scopes"`
		LifeCycleID int64   `json:"life_cycle_id"`
		Name        string  `json:"name"`
		Order       int64   `json:"order"`
		Value       int64   `json:"value"`
	} `json:"results"`
	Total int64 `json:"total"`
}


type ChoicesPortType struct {
	Count    int64       `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		DataScopes   []int64 `json:"data_scopes"`
		Name         string  `json:"name"`
		Order        int64   `json:"order"`
		PortChoiceID int64   `json:"port_choice_id"`
		Value        int64   `json:"value"`
	} `json:"results"`
	Total int64 `json:"total"`
}


type ChoicesTestDurationTypeType struct {
	Count    int64       `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		DataScopes         []int64 `json:"data_scopes"`
		Name               string  `json:"name"`
		Order              int64   `json:"order"`
		TestDurationTypeID int64   `json:"test_duration_type_id"`
		Value              int64   `json:"value"`
	} `json:"results"`
	Total int64 `json:"total"`
}


type ChoicesTestPurposeType struct {
	Count    int64       `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		DataScopes    []int64 `json:"data_scopes"`
		Name          string  `json:"name"`
		Order         int64   `json:"order"`
		TestPurposeID int64   `json:"test_purpose_id"`
		Value         int64   `json:"value"`
	} `json:"results"`
	Total int64 `json:"total"`
}


type DashboardType struct {
	AvailableHosts int64 `json:"available_hosts"`
	TotalHosts     int64 `json:"total_hosts"`
	TotalTestCases int64 `json:"total_test_cases"`
	TotalTests     int64 `json:"total_tests"`
	TotalUsers     int64 `json:"total_users"`
}


type DashboardNotificationsType struct {
	Count    int64       `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		CreatedTime     string `json:"created_time"`
		ItemID          int64  `json:"item_id"`
		Message         string `json:"message"`
		NotificationsID int64  `json:"notifications_id"`
		StartTime       string `json:"start_time"`
		Status          string `json:"status"`
		TypeName        string `json:"type_name"`
	} `json:"results"`
	Total int64 `json:"total"`
}


type KeyForTypesViewCall struct {
	One   string `json:"1"`
	One0  string `json:"10"`
	One1  string `json:"11"`
	One2  string `json:"12"`
	One3  string `json:"13"`
	One4  string `json:"14"`
	One5  string `json:"15"`
	One6  string `json:"16"`
	One7  string `json:"17"`
	One8  string `json:"18"`
	One9  string `json:"19"`
	Two   string `json:"2"`
	Two0  string `json:"20"`
	Two1  string `json:"21"`
	Three string `json:"3"`
	Four  string `json:"4"`
	Five  string `json:"5"`
	Six   string `json:"6"`
	Seven string `json:"7"`
	Eight string `json:"8"`
	Nine  string `json:"9"`
}


type RolesType struct {
	Admin        string `json:"admin"`
	Manager      string `json:"manager"`
	Technician   string `json:"technician"`
	TestEngineer string `json:"test_engineer"`
	TestLead     string `json:"test_lead"`
}


type SettingsEnvType struct {
	UploadDir string `json:"upload_dir"`
}


type SubscriptionUserViewType struct {
	Count    int64       `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		TestProject struct {
			TestProjectID   int64  `json:"test_project_id"`
			TestProjectName string `json:"test_project_name"`
		} `json:"test_project"`
		TestRequest struct {
			Description   string `json:"description"`
			TestRequestID int64  `json:"test_request_id"`
		} `json:"test_request"`
		User struct {
			UserID   int64  `json:"user_id"`
			Username string `json:"username"`
		} `json:"user"`
		UserSubscriptionID int64 `json:"user_subscription_id"`
	} `json:"results"`
	Total int64 `json:"total"`
}


type TctDetailViewType struct {
	Count    int64       `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		EnableNotification bool  `json:"enable_notification"`
		Order              int64 `json:"order"`
		TctDetailID        int64 `json:"tct_detail_id"`
		TctRuns            []struct {
			AssignedUser         interface{} `json:"assigned_user"`
			DeviceUnderTesting   interface{} `json:"device_under_testing"`
			DownloadType         string      `json:"download_type"`
			EmailNotification    bool        `json:"email_notification"`
			IsManual             bool        `json:"is_manual"`
			Priority             interface{} `json:"priority"`
			SlotNumber           int64       `json:"slot_number"`
			TargetCompletionDate interface{} `json:"target_completion_date"`
			TctDetail            struct {
				TctDetailID int64 `json:"tct_detail_id"`
			} `json:"tct_detail"`
			TctRunID   int64       `json:"tct_run_id"`
			TestSystem interface{} `json:"test_system"`
			WontRun    bool        `json:"wont_run"`
		} `json:"tct_runs"`
		TestCase struct {
			DefaultAssignee struct {
				ID       int64  `json:"id"`
				Username string `json:"username"`
			} `json:"default_assignee"`
			IsManual   bool   `json:"is_manual"`
			Name       string `json:"name"`
			TestCaseID int64  `json:"test_case_id"`
		} `json:"test_case"`
		TestCycleTemplate struct {
			Name                string      `json:"name"`
			TestCycle           interface{} `json:"test_cycle"`
			TestCycleTemplateID int64       `json:"test_cycle_template_id"`
			TestPlan            struct {
				Description string `json:"description"`
				Name        string `json:"name"`
				TestPlanID  int64  `json:"test_plan_id"`
				TestType    int64  `json:"test_type"`
				TestTypeStr string `json:"test_type_str"`
				Version     string `json:"version"`
			} `json:"test_plan"`
		} `json:"test_cycle_template"`
		TestGroup struct {
			IsManual bool   `json:"is_manual"`
			Name     string `json:"name"`
			TestCase []struct {
				IsManual   bool   `json:"is_manual"`
				Name       string `json:"name"`
				Order      int64  `json:"order"`
				TestCaseID int64  `json:"test_case_id"`
			} `json:"test_case"`
			TestGroupID int64 `json:"test_group_id"`
		} `json:"test_group"`
		TestRuns    []interface{} `json:"test_runs"`
		User        interface{}   `json:"user"`
		UserGroupID interface{}   `json:"user_group_id"`
		WontRun     bool          `json:"wont_run"`
	} `json:"results"`
	Total int64 `json:"total"`
}


type TestCaseCategoriesType struct {
	Zero  string `json:"0"`
	One   string `json:"1"`
	Two   string `json:"2"`
	Three string `json:"3"`
	Four  string `json:"4"`
	Five  string `json:"5"`
	Six   string `json:"6"`
	Seven string `json:"7"`
	Eight string `json:"8"`
	Nine  string `json:"9"`
}


type TestCaseViewType struct {
	Count    int64       `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		BlockSizeChoices []struct {
			BlockSizeChoiceID int64  `json:"block_size_choice_id"`
			Name              string `json:"name"`
			Value             int64  `json:"value"`
		} `json:"block_size_choices"`
		DataScopes      []int64  `json:"data_scopes"`
		DataScopesStr   []string `json:"data_scopes_str"`
		DateCreate      string   `json:"date_create"`
		DateLastModify  string   `json:"date_last_modify"`
		DefaultAssignee struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"default_assignee"`
		Description string `json:"description"`
		Interface   struct {
			InterfaceID int64  `json:"interface_id"`
			Name        string `json:"name"`
			Value       int64  `json:"value"`
		} `json:"interface"`
		IsActive bool   `json:"is_active"`
		IsManual bool   `json:"is_manual"`
		Name     string `json:"name"`
		Os       []struct {
			Name  string `json:"name"`
			OsID  int64  `json:"os_id"`
			Value int64  `json:"value"`
		} `json:"os"`
		Owner int64 `json:"owner"`
		Ports []struct {
			Name         string `json:"name"`
			PortChoiceID int64  `json:"port_choice_id"`
			Value        int64  `json:"value"`
		} `json:"ports"`
		Raids []struct {
			Name   string `json:"name"`
			RaidID int64  `json:"raid_id"`
			Value  int64  `json:"value"`
		} `json:"raids"`
		Speeds []struct {
			Name    string `json:"name"`
			SpeedID int64  `json:"speed_id"`
			Value   int64  `json:"value"`
		} `json:"speeds"`
		TestCaseID             int64 `json:"test_case_id"`
		TestCoverageCategories []struct {
			Name                   string `json:"name"`
			TestCoverageCategoryID int64  `json:"test_coverage_category_id"`
			Value                  int64  `json:"value"`
		} `json:"test_coverage_categories"`
		TestDuration         string `json:"test_duration"`
		TestInstruction      string `json:"test_instruction"`
		TestInstructionFiles []struct {
			ArchivePath  string `json:"archive_path"`
			UserUploadID int64  `json:"user_upload_id"`
		} `json:"test_instruction_files"`
		TestLauncher struct {
			Name           string `json:"name"`
			SvnTag         string `json:"svn_tag"`
			TestLauncherID int64  `json:"test_launcher_id"`
		} `json:"test_launcher"`
		TestLauncherConfig struct {
			TestLauncherConfigID int64 `json:"test_launcher_config_id"`
			UserUploads          []struct {
				ArchivePath  string `json:"archive_path"`
				UserUploadID int64  `json:"user_upload_id"`
			} `json:"user_uploads"`
		} `json:"test_launcher_config"`
		TestPurpose struct {
			Name          string `json:"name"`
			TestPurposeID int64  `json:"test_purpose_id"`
			Value         int64  `json:"value"`
		} `json:"test_purpose"`
		TestTool struct {
			Name       string `json:"name"`
			Revision   string `json:"revision"`
			TestToolID int64  `json:"test_tool_id"`
		} `json:"test_tool"`
		UserCreate struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"user_create"`
		UserLastModify struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"user_last_modify"`
	} `json:"results"`
	Total int64 `json:"total"`
}


type TestCycleReportType struct {
	Count int64 `json:"count"`
}


type TestCycleTemplateViewType struct {
	Count    int64       `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		DateCreate     string `json:"date_create"`
		DateLastModify string `json:"date_last_modify"`
		Description    string `json:"description"`
		Name           string `json:"name"`
		TestCycle      struct {
			Name        string `json:"name"`
			TestCycleID int64  `json:"test_cycle_id"`
			TestPlan    struct {
				Description string `json:"description"`
				Name        string `json:"name"`
				TestPlanID  int64  `json:"test_plan_id"`
				TestType    int64  `json:"test_type"`
				TestTypeStr string `json:"test_type_str"`
				Version     string `json:"version"`
			} `json:"test_plan"`
			Title string `json:"title"`
		} `json:"test_cycle"`
		TestCycleTemplateID int64 `json:"test_cycle_template_id"`
		TestPlan            struct {
			Name       string `json:"name"`
			TestPlanID int64  `json:"test_plan_id"`
		} `json:"test_plan"`
		UserCreate struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"user_create"`
		UserLastModify struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"user_last_modify"`
	} `json:"results"`
	Total int64 `json:"total"`
}


type TestGroupViewType struct {
	Count    int64       `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		CategoryIds    []int64  `json:"category_ids"`
		CategoryIdsStr []string `json:"category_ids_str"`
		DataScopes     []int64  `json:"data_scopes"`
		DataScopesStr  []string `json:"data_scopes_str"`
		DateCreate     string   `json:"date_create"`
		DateLastModify string   `json:"date_last_modify"`
		IsActive       bool     `json:"is_active"`
		Name           string   `json:"name"`
		TestCase       []struct {
			Name       string `json:"name"`
			Order      int64  `json:"order"`
			TestCaseID int64  `json:"test_case_id"`
		} `json:"test_case"`
		TestGroupID int64 `json:"test_group_id"`
		UserCreate  struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"user_create"`
		UserLastModify struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"user_last_modify"`
	} `json:"results"`
	Total int64 `json:"total"`
}


type TestLauncherLibraryType struct {
	Count    int64       `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		DataScopes     []int64     `json:"data_scopes"`
		DateCreate     string      `json:"date_create"`
		DateLastModify string      `json:"date_last_modify"`
		Developer      string      `json:"developer"`
		IsActive       bool        `json:"is_active"`
		LibraryID      int64       `json:"library_id"`
		Name           string      `json:"name"`
		Tag            interface{} `json:"tag"`
		UserCreate     struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"user_create"`
		UserLastModify struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"user_last_modify"`
		UserUpload struct {
			ArchivePath  string `json:"archive_path"`
			UserUploadID int64  `json:"user_upload_id"`
		} `json:"user_upload"`
		Version string `json:"version"`
	} `json:"results"`
	Total int64 `json:"total"`
}


type TestLauncherLibraryDependencyType struct {
	Count    int64       `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		DependencyID int64       `json:"dependency_id"`
		Library      int64       `json:"library"`
		Tag          interface{} `json:"tag"`
		TestLauncher int64       `json:"test_launcher"`
	} `json:"results"`
	Total int64 `json:"total"`
}


type TestLauncherTestPeripheralDependencyType struct {
	Count    int64       `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		DependencyID       int64       `json:"dependency_id"`
		Tag                interface{} `json:"tag"`
		TestLauncher       int64       `json:"test_launcher"`
		TestPeripheralType int64       `json:"test_peripheral_type"`
	} `json:"results"`
	Total int64 `json:"total"`
}


type TestLauncherTestPeripheralTypeType struct {
	Count    int64       `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		Function             string `json:"function"`
		Model                string `json:"model"`
		Name                 string `json:"name"`
		Tag                  string `json:"tag"`
		TestPeripheralTypeID int64  `json:"test_peripheral_type_id"`
		Vendor               string `json:"vendor"`
	} `json:"results"`
	Total int64 `json:"total"`
}


type TestLauncherTestToolType struct {
	Count    int64       `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		BlockSizeChoices []struct {
			BlockSizeChoiceID int64  `json:"block_size_choice_id"`
			Name              string `json:"name"`
			Value             int64  `json:"value"`
		} `json:"block_size_choices"`
		DataScopes     []int64  `json:"data_scopes"`
		DataScopesStr  []string `json:"data_scopes_str"`
		DateCreate     string   `json:"date_create"`
		DateLastModify string   `json:"date_last_modify"`
		Description    string   `json:"description"`
		IsActive       bool     `json:"is_active"`
		Name           string   `json:"name"`
		OsRequirement  struct {
			Name              string `json:"name"`
			OperatingSystemID int64  `json:"operating_system_id"`
		} `json:"os_requirement"`
		Revision                  string      `json:"revision"`
		Tag                       interface{} `json:"tag"`
		TestCoverageCategories    []int64     `json:"test_coverage_categories"`
		TestCoverageCategoriesStr []string    `json:"test_coverage_categories_str"`
		TestToolID                int64       `json:"test_tool_id"`
		UserCreate                struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"user_create"`
		UserLastModify struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"user_last_modify"`
		UserUpload struct {
			ArchivePath  string `json:"archive_path"`
			UserUploadID int64  `json:"user_upload_id"`
		} `json:"user_upload"`
	} `json:"results"`
	Total int64 `json:"total"`
}


type TestLauncherViewType struct {
	Count    int64       `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		BlockSizeChoices []struct {
			BlockSizeChoiceID int64  `json:"block_size_choice_id"`
			Name              string `json:"name"`
			Value             int64  `json:"value"`
		} `json:"block_size_choices"`
		DataScopes     []int64  `json:"data_scopes"`
		DataScopesStr  []string `json:"data_scopes_str"`
		DateCreate     string   `json:"date_create"`
		DateLastModify string   `json:"date_last_modify"`
		IsActive       bool     `json:"is_active"`
		Libraries      []struct {
			LibraryID      int64  `json:"library_id"`
			LibraryStr     string `json:"library_str"`
			LibraryVersion string `json:"library_version"`
		} `json:"libraries"`
		Name string `json:"name"`
		Os   []struct {
			Name  string `json:"name"`
			OsID  int64  `json:"os_id"`
			Value int64  `json:"value"`
		} `json:"os"`
		Ports []struct {
			Name         string `json:"name"`
			PortChoiceID int64  `json:"port_choice_id"`
			Value        int64  `json:"value"`
		} `json:"ports"`
		ReleaseBy   string      `json:"release_by"`
		ReleaseDate interface{} `json:"release_date"`
		Speeds      []struct {
			Name    string `json:"name"`
			SpeedID int64  `json:"speed_id"`
			Value   int64  `json:"value"`
		} `json:"speeds"`
		SvnTag                 string      `json:"svn_tag"`
		Tag                    interface{} `json:"tag"`
		TestCoverageCategories []struct {
			Name                   string `json:"name"`
			TestCoverageCategoryID int64  `json:"test_coverage_category_id"`
			Value                  int64  `json:"value"`
		} `json:"test_coverage_categories"`
		TestDurationType struct {
			Name               string `json:"name"`
			TestDurationTypeID int64  `json:"test_duration_type_id"`
			Value              int64  `json:"value"`
		} `json:"test_duration_type"`
		TestLauncherID      int64 `json:"test_launcher_id"`
		TestPeripheralTypes []struct {
			TestPeripheralTypeID  int64  `json:"test_peripheral_type_id"`
			TestPeripheralTypeStr string `json:"test_peripheral_type_str"`
		} `json:"test_peripheral_types"`
		TestTools []struct {
			Name       string `json:"name"`
			Revision   string `json:"revision"`
			TestToolID int64  `json:"test_tool_id"`
		} `json:"test_tools"`
		UserCreate struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"user_create"`
		UserLastModify struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"user_last_modify"`
		UserUpload struct {
			ArchivePath  string `json:"archive_path"`
			UserUploadID int64  `json:"user_upload_id"`
		} `json:"user_upload"`
	} `json:"results"`
	Total int64 `json:"total"`
}


type TestPlanViewType struct {
	Count    int64       `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		BlockSizeChoice struct {
			BlockSizeChoiceID int64  `json:"block_size_choice_id"`
			Name              string `json:"name"`
			Value             int64  `json:"value"`
		} `json:"block_size_choice"`
		DataScopes     []int64       `json:"data_scopes"`
		DataScopesStr  []string      `json:"data_scopes_str"`
		DateCreate     string        `json:"date_create"`
		DateLastModify string        `json:"date_last_modify"`
		Description    string        `json:"description"`
		IsActive       bool          `json:"is_active"`
		Name           string        `json:"name"`
		QualPlan       []interface{} `json:"qual_plan"`
		TestPlanGroup  []struct {
			Order    int64 `json:"order"`
			TestCase struct {
				Name       string `json:"name"`
				TestCaseID int64  `json:"test_case_id"`
			} `json:"test_case"`
			TestGroup struct {
				IsManual bool   `json:"is_manual"`
				Name     string `json:"name"`
				TestCase []struct {
					IsManual   bool   `json:"is_manual"`
					Name       string `json:"name"`
					Order      int64  `json:"order"`
					TestCaseID int64  `json:"test_case_id"`
				} `json:"test_case"`
				TestGroupID int64 `json:"test_group_id"`
			} `json:"test_group"`
		} `json:"test_plan_group"`
		TestPlanID  int64  `json:"test_plan_id"`
		TestType    int64  `json:"test_type"`
		TestTypeStr string `json:"test_type_str"`
		UserCreate  struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"user_create"`
		UserLastModify struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"user_last_modify"`
		Version string `json:"version"`
	} `json:"results"`
	Total int64 `json:"total"`
}


type TestProjectViewType struct {
	Count    int64       `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		BlockSizeChoice struct {
			BlockSizeChoiceID int64  `json:"block_size_choice_id"`
			Name              string `json:"name"`
			Value             int64  `json:"value"`
		} `json:"block_size_choice"`
		Capacity            []int64  `json:"capacity"`
		CapacityStr         []string `json:"capacity_str"`
		CreateTime          string   `json:"create_time"`
		CreatedBy           int64    `json:"created_by"`
		Customer            int64    `json:"customer"`
		CustomerStr         string   `json:"customer_str"`
		DataScopes          []int64  `json:"data_scopes"`
		DataScopesStr       []string `json:"data_scopes_str"`
		DateCreate          string   `json:"date_create"`
		DateLastModify      string   `json:"date_last_modify"`
		FilterProductID     int64    `json:"filter_product_id"`
		FormFactor          []int64  `json:"form_factor"`
		FormFactorStr       []string `json:"form_factor_str"`
		HostInterface       int64    `json:"host_interface"`
		HostInterfaceStr    string   `json:"host_interface_str"`
		IsActive            bool     `json:"is_active"`
		OverProvisioning    int64    `json:"over_provisioning"`
		OverProvisioningStr string   `json:"over_provisioning_str"`
		Owner               struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"owner"`
		Product struct {
			Comments          string   `json:"comments"`
			DataScopes        []int64  `json:"data_scopes"`
			DataScopesStr     []string `json:"data_scopes_str"`
			HostController    int64    `json:"host_controller"`
			HostControllerStr string   `json:"host_controller_str"`
			HostInterface     int64    `json:"host_interface"`
			HostInterfaceStr  string   `json:"host_interface_str"`
			IsActive          bool     `json:"is_active"`
			Name              string   `json:"name"`
			NandType          int64    `json:"nand_type"`
			NandTypeStr       string   `json:"nand_type_str"`
			NandVersion       int64    `json:"nand_version"`
			NandVersionStr    string   `json:"nand_version_str"`
			ProductID         int64    `json:"product_id"`
		} `json:"product"`
		Security           int64  `json:"security"`
		SecurityStr        string `json:"security_str"`
		StartDate          string `json:"start_date"`
		Status             int64  `json:"status"`
		StatusStr          string `json:"status_str"`
		SubscriptionGroups []struct {
			GroupID int64  `json:"group_id"`
			Name    string `json:"name"`
			Users   []struct {
				UserID   int64  `json:"user_id"`
				Username string `json:"username"`
			} `json:"users"`
		} `json:"subscription_groups"`
		SubscriptionUsers []struct {
			UserID   int64  `json:"user_id"`
			Username string `json:"username"`
		} `json:"subscription_users"`
		TestProjectID   int64  `json:"test_project_id"`
		TestProjectName string `json:"test_project_name"`
		UserCreate      struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"user_create"`
		UserLastModify struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"user_last_modify"`
	} `json:"results"`
	Total int64 `json:"total"`
}


type TestResultChoicesStatusType struct {
	Zero  string `json:"0"`
	One   string `json:"1"`
	Two   string `json:"2"`
	Three string `json:"3"`
	Four  string `json:"4"`
	Five  string `json:"5"`
	Six   string `json:"6"`
}


type TestResultCommentTagViewType struct {
	Count    int64       `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		CommentTagID int64  `json:"comment_tag_id"`
		Tag          string `json:"tag"`
	} `json:"results"`
	Total int64 `json:"total"`
}


type TestResultFioTestResultViewType struct {
	Count    int64         `json:"count"`
	Next     interface{}   `json:"next"`
	Previous interface{}   `json:"previous"`
	Results  []interface{} `json:"results"`
	Total    int64         `json:"total"`
}


type TestResultIometerTestReportViewType []interface{}


type TestResultIometerTestResultViewType struct {
	Count    int64         `json:"count"`
	Next     interface{}   `json:"next"`
	Previous interface{}   `json:"previous"`
	Results  []interface{} `json:"results"`
	Total    int64         `json:"total"`
}


type TestResultLogType struct {
	Count    int64       `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		ArchivePath   string `json:"archive_path"`
		Comment       string `json:"comment"`
		ContentType   string `json:"content_type"`
		LogType       int64  `json:"log_type"`
		LogURL        string `json:"log_url"`
		Tag           string `json:"tag"`
		TestLogID     int64  `json:"test_log_id"`
		TestRunResult struct {
			TestRun struct {
				TestRunID int64 `json:"test_run_id"`
			} `json:"test_run"`
			TestRunResultID int64 `json:"test_run_result_id"`
		} `json:"test_run_result"`
	} `json:"results"`
	Total int64 `json:"total"`
}


type TestResultReportTemplateViewType struct {
	Count    int64       `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		Description      string `json:"description"`
		Image            string `json:"image"`
		Name             string `json:"name"`
		PivotTableConfig string `json:"pivot_table_config"`
		ReportTemplateID int64  `json:"report_template_id"`
		ReportType       string `json:"report_type"`
		TestTool         string `json:"test_tool"`
	} `json:"results"`
	Total int64 `json:"total"`
}


type TestResultTestRunResultCommentViewType struct {
	Count    int64       `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		Comment       string        `json:"comment"`
		CommentID     int64         `json:"comment_id"`
		LastModified  interface{}   `json:"last_modified"`
		Tags          []interface{} `json:"tags"`
		TestRunResult struct {
			TestRun struct {
				TestRunID int64 `json:"test_run_id"`
			} `json:"test_run"`
			TestRunResultID int64 `json:"test_run_result_id"`
		} `json:"test_run_result"`
		Timestamp string `json:"timestamp"`
		User      struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"user"`
	} `json:"results"`
	Total int64 `json:"total"`
}


type TestResultViewType struct {
	Count    int64       `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		ControlStatus    int64    `json:"control_status"`
		ControlStatusStr string   `json:"control_status_str"`
		DataScopes       []int64  `json:"data_scopes"`
		DataScopesStr    []string `json:"data_scopes_str"`
		EndTime          string   `json:"end_time"`
		ErrorCode        string   `json:"error_code"`
		ErrorMessage     string   `json:"error_message"`
		ExecutedBy       struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"executed_by"`
		GroupOrder int64  `json:"group_order"`
		IsActive   bool   `json:"is_active"`
		IsValid    bool   `json:"is_valid"`
		JiraID     string `json:"jira_id"`
		JiraIDStr  string `json:"jira_id_str"`
		Progress   int64  `json:"progress"`
		Reasons    []struct {
			Comment   string `json:"comment"`
			IsValid   bool   `json:"is_valid"`
			ReasonID  int64  `json:"reason_id"`
			Timestamp string `json:"timestamp"`
			User      struct {
				ID       int64  `json:"id"`
				Username string `json:"username"`
			} `json:"user"`
		} `json:"reasons"`
		Result     int64  `json:"result"`
		ResultStr  string `json:"result_str"`
		StartTime  string `json:"start_time"`
		Status     int64  `json:"status"`
		StatusStr  string `json:"status_str"`
		TestResult string `json:"test_result"`
		TestRun    struct {
			IsManual  bool  `json:"is_manual"`
			TestRunID int64 `json:"test_run_id"`
		} `json:"test_run"`
		TestRunResultID int64       `json:"test_run_result_id"`
		TestStatus      int64       `json:"test_status"`
		TestStatusStr   string      `json:"test_status_str"`
		Type            int64       `json:"type"`
		TypeStr         string      `json:"type_str"`
		UserUpload      interface{} `json:"user_upload"`
	} `json:"results"`
	Total int64 `json:"total"`
}

type TestRunViewType struct {
	Count    int64       `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		AssignedUser struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"assigned_user"`
		DataScopes         []int64     `json:"data_scopes"`
		DateCreate         interface{} `json:"date_create"`
		DateLastModify     string      `json:"date_last_modify"`
		DeviceUnderTesting struct {
			Capacity             int64  `json:"capacity"`
			CapacityStr          string `json:"capacity_str"`
			DeviceUnderTestingID int64  `json:"device_under_testing_id"`
			HealthStatus         int64  `json:"health_status"`
			HealthStatusStr      string `json:"health_status_str"`
			Model                string `json:"model"`
			ProductFamily        struct {
				Name            string `json:"name"`
				ProductFamilyID int64  `json:"product_family_id"`
			} `json:"product_family"`
			SerialNumber string `json:"serial_number"`
		} `json:"device_under_testing"`
		DownloadType         string      `json:"download_type"`
		DownloadTypeStr      string      `json:"download_type_str"`
		EmailNotification    bool        `json:"email_notification"`
		Fields               interface{} `json:"fields"`
		IsActive             bool        `json:"is_active"`
		IsManual             bool        `json:"is_manual"`
		Order                int64       `json:"order"`
		Priority             interface{} `json:"priority"`
		PriorityStr          interface{} `json:"priority_str"`
		RerunNum             int64       `json:"rerun_num"`
		Result               int64       `json:"result"`
		ResultStr            string      `json:"result_str"`
		SlotNumber           int64       `json:"slot_number"`
		Status               int64       `json:"status"`
		StatusStr            string      `json:"status_str"`
		TargetCompletionDate interface{} `json:"target_completion_date"`
		TestCase             struct {
			DefaultAssignee struct {
				ID       int64  `json:"id"`
				Username string `json:"username"`
			} `json:"default_assignee"`
			IsManual   bool   `json:"is_manual"`
			Name       string `json:"name"`
			TestCaseID int64  `json:"test_case_id"`
		} `json:"test_case"`
		TestCycle struct {
			Name        string `json:"name"`
			TestCycleID int64  `json:"test_cycle_id"`
			TestPlan    struct {
				Description string `json:"description"`
				Name        string `json:"name"`
				TestPlanID  int64  `json:"test_plan_id"`
				TestType    int64  `json:"test_type"`
				TestTypeStr string `json:"test_type_str"`
				Version     string `json:"version"`
			} `json:"test_plan"`
		} `json:"test_cycle"`
		TestGroup struct {
			IsManual bool   `json:"is_manual"`
			Name     string `json:"name"`
			TestCase []struct {
				IsManual   bool   `json:"is_manual"`
				Name       string `json:"name"`
				Order      int64  `json:"order"`
				TestCaseID int64  `json:"test_case_id"`
			} `json:"test_case"`
			TestGroupID int64 `json:"test_group_id"`
		} `json:"test_group"`
		TestRunID  int64 `json:"test_run_id"`
		TestSystem struct {
			HostName     string `json:"host_name"`
			TestSystemID int64  `json:"test_system_id"`
		} `json:"test_system"`
		UserCreate     interface{} `json:"user_create"`
		UserLastModify struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"user_last_modify"`
		WasSent bool `json:"was_sent"`
		WontRun bool `json:"wont_run"`
	} `json:"results"`
	Total int64 `json:"total"`
}


type TestSystemChoicesDataScopeType struct {
	One   string `json:"1"`
	Two   string `json:"2"`
	Three string `json:"3"`
}


type TestSystemChoicesDeviceUnderTestingBackendType struct {
	One   string `json:"1"`
	Two   string `json:"2"`
	Three string `json:"3"`
	Four  string `json:"4"`
	Five  string `json:"5"`
}


type TestSystemChoicesDeviceUnderTestingCapacityType struct {
	Zero  string `json:"0"`
	One   string `json:"1"`
	One0  string `json:"10"`
	One1  string `json:"11"`
	One2  string `json:"12"`
	One3  string `json:"13"`
	Two   string `json:"2"`
	Three string `json:"3"`
	Four  string `json:"4"`
	Five  string `json:"5"`
	Six   string `json:"6"`
	Seven string `json:"7"`
	Eight string `json:"8"`
	Nine  string `json:"9"`
}


type TestSystemChoicesDeviceUnderTestingCustomerType struct {
	Zero  string `json:"0"`
	One   string `json:"1"`
	One0  string `json:"10"`
	One1  string `json:"11"`
	One2  string `json:"12"`
	Two   string `json:"2"`
	Three string `json:"3"`
	Four  string `json:"4"`
	Five  string `json:"5"`
	Six   string `json:"6"`
	Seven string `json:"7"`
	Eight string `json:"8"`
	Nine  string `json:"9"`
}


type TestSystemChoicesDeviceUnderTestingFormFactorType struct {
	One string `json:"1"`
}


type TestSystemChoicesDeviceUnderTestingFrontendType struct {
	Zero string `json:"0"`
	One  string `json:"1"`
	Two  string `json:"2"`
}


type TestSystemChoicesDeviceUnderTestingHealthStatusType struct {
	Zero string `json:"0"`
	One  string `json:"1"`
}


type TestSystemChoicesDeviceUnderTestingHostInterfaceType struct {
	One string `json:"1"`
	Two string `json:"2"`
}


type TestSystemChoicesDeviceUnderTestingProductType struct {
	One   string `json:"1"`
	One0  string `json:"10"`
	One1  string `json:"11"`
	Two   string `json:"2"`
	Three string `json:"3"`
	Four  string `json:"4"`
	Five  string `json:"5"`
	Six   string `json:"6"`
	Seven string `json:"7"`
	Eight string `json:"8"`
	Nine  string `json:"9"`
}


type TestSystemChoicesDeviceUnderTestingStatusType struct {
	Zero  string `json:"0"`
	One   string `json:"1"`
	Two   string `json:"2"`
	Three string `json:"3"`
}


type TestSystemChoicesDownloadTypeType struct {
	CoreImage    string `json:"core_image"`
	Default      string `json:"default"`
	None         string `json:"none"`
	ProductImage string `json:"product_image"`
}


type TestSystemChoicesFirmwareTypeType struct {
	Zero  string `json:"0"`
	One   string `json:"1"`
	Two   string `json:"2"`
	Three string `json:"3"`
	Four  string `json:"4"`
}


type TestSystemChoicesNandTypeType struct {
	Three string `json:"3"`
	Four  string `json:"4"`
}


type TestSystemChoicesNandVersionType struct {
	One  string `json:"1"`
	Two  string `json:"2"`
	Five string `json:"5"`
}


type TestSystemChoicesOperatingSystemType struct {
	Count    int64       `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		Name              string `json:"name"`
		OperatingSystemID int64  `json:"operating_system_id"`
	} `json:"results"`
	Total int64 `json:"total"`
}


type TestSystemChoicesOverProvisioningType struct {
	Zero  string `json:"0"`
	One   string `json:"1"`
	Two   string `json:"2"`
	Three string `json:"3"`
	Four  string `json:"4"`
}


type TestSystemChoicesPortType struct {
	Zero string `json:"0"`
	One  string `json:"1"`
}


type TestSystemChoicesRaidType struct {
	Zero  string `json:"0"`
	One   string `json:"1"`
	Two   string `json:"2"`
	Three string `json:"3"`
	Four  string `json:"4"`
	Five  string `json:"5"`
	Six   string `json:"6"`
}


type TestSystemChoicesResultType struct {
	One string `json:"1"`
	Two string `json:"2"`
}


type TestSystemChoicesSecurityType struct {
	Zero string `json:"0"`
	One  string `json:"1"`
	Two  string `json:"2"`
}


type TestSystemChoicesSpeedType struct {
	Zero string `json:"0"`
	One  string `json:"1"`
	Two  string `json:"2"`
}


type TestSystemChoicesStatusType struct {
	One   string `json:"1"`
	Two   string `json:"2"`
	Three string `json:"3"`
	Four  string `json:"4"`
	Five  string `json:"5"`
	Six   string `json:"6"`
	Seven string `json:"7"`
}


type TestSystemChoicesStorageInterfaceFormFactorType struct {
	One string `json:"1"`
}


type TestSystemChoicesTestProjectStatusType struct {
	Zero  string `json:"0"`
	One   string `json:"1"`
	Two   string `json:"2"`
	Three string `json:"3"`
	Four  string `json:"4"`
}


type TestSystemChoicesTestRunPriorityStatusType struct {
	Zero  string `json:"0"`
	One   string `json:"1"`
	Two   string `json:"2"`
	Three string `json:"3"`
}


type TestSystemChoicesTestRunResultType struct {
	Zero string `json:"0"`
	One  string `json:"1"`
}


type TestSystemChoicesTestToolTypeType struct {
	Zero string `json:"0"`
	One  string `json:"1"`
}


type TestSystemChoicesTestTypeType struct {
	Zero  string `json:"0"`
	One   string `json:"1"`
	One0  string `json:"10"`
	Two   string `json:"2"`
	Three string `json:"3"`
	Four  string `json:"4"`
	Five  string `json:"5"`
	Six   string `json:"6"`
	Seven string `json:"7"`
	Eight string `json:"8"`
	Nine  string `json:"9"`
}


type TestSystemControllerType struct {
	Count    int64       `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		ControllerID   int64         `json:"controller_id"`
		DataScopes     []int64       `json:"data_scopes"`
		DataScopesStr  []string      `json:"data_scopes_str"`
		DateCreate     string        `json:"date_create"`
		DateLastModify string        `json:"date_last_modify"`
		HostName       string        `json:"host_name"`
		IP             string        `json:"ip"`
		IsActive       bool          `json:"is_active"`
		Tag            string        `json:"tag"`
		TestSystems    []interface{} `json:"test_systems"`
		UserCreate     struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"user_create"`
		UserLastModify struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"user_last_modify"`
	} `json:"results"`
	Total int64 `json:"total"`
}


type TestSystemDeviceUnderTestingType struct {
	Count    int64       `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		Capacity             int64   `json:"capacity"`
		CapacityStr          string  `json:"capacity_str"`
		DataScopes           []int64 `json:"data_scopes"`
		DateCreate           string  `json:"date_create"`
		DateLastModify       string  `json:"date_last_modify"`
		DeviceUnderTestingID int64   `json:"device_under_testing_id"`
		FormFactor           int64   `json:"form_factor"`
		FormFactorStr        string  `json:"form_factor_str"`
		HealthStatus         int64   `json:"health_status"`
		HealthStatusStr      string  `json:"health_status_str"`
		HostController       int64   `json:"host_controller"`
		HostControllerStr    string  `json:"host_controller_str"`
		HostInterfaceStr     string  `json:"host_interface_str"`
		IsActive             bool    `json:"is_active"`
		Model                string  `json:"model"`
		NandType             int64   `json:"nand_type"`
		NandTypeStr          string  `json:"nand_type_str"`
		NandVersion          int64   `json:"nand_version"`
		NandVersionStr       string  `json:"nand_version_str"`
		OemCustomer          int64   `json:"oem_customer"`
		OemCustomerStr       string  `json:"oem_customer_str"`
		Owner                int64   `json:"owner"`
		ProductFamily        struct {
			Name            string `json:"name"`
			ProductFamilyID int64  `json:"product_family_id"`
		} `json:"product_family"`
		SerialNumber string `json:"serial_number"`
		Status       int64  `json:"status"`
		StatusStr    string `json:"status_str"`
		TestSystems  []struct {
			HostName     string `json:"host_name"`
			IP           string `json:"ip"`
			TestSystemID int64  `json:"test_system_id"`
		} `json:"test_systems"`
		UserCreate struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"user_create"`
		UserLastModify struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"user_last_modify"`
	} `json:"results"`
	Total int64 `json:"total"`
}


type TestSystemFirmwareType struct {
	Count    int64       `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		BlockSizeChoice struct {
			BlockSizeChoiceID int64  `json:"block_size_choice_id"`
			Name              string `json:"name"`
			Value             int64  `json:"value"`
		} `json:"block_size_choice"`
		BootramImage struct {
			ArchivePath  string `json:"archive_path"`
			UserUploadID int64  `json:"user_upload_id"`
		} `json:"bootram_image"`
		BridgeImage struct {
			ArchivePath  string `json:"archive_path"`
			UserUploadID int64  `json:"user_upload_id"`
		} `json:"bridge_image"`
		Configuration string `json:"configuration"`
		CoreImage     struct {
			ArchivePath  string `json:"archive_path"`
			UserUploadID int64  `json:"user_upload_id"`
		} `json:"core_image"`
		DataScopes            []int64  `json:"data_scopes"`
		DataScopesStr         []string `json:"data_scopes_str"`
		DateCreate            string   `json:"date_create"`
		DateLastModify        string   `json:"date_last_modify"`
		Dependency            string   `json:"dependency"`
		Description           string   `json:"description"`
		FirmwareID            int64    `json:"firmware_id"`
		FwImageLastUploadDate string   `json:"fw_image_last_upload_date"`
		FwType                int64    `json:"fw_type"`
		FwTypeStr             string   `json:"fw_type_str"`
		IsActive              bool     `json:"is_active"`
		Name                  string   `json:"name"`
		OemCustomer           string   `json:"oem_customer"`
		OverProvisioning      int64    `json:"over_provisioning"`
		OverProvisioningStr   string   `json:"over_provisioning_str"`
		Product               struct {
			Comments          string   `json:"comments"`
			DataScopes        []int64  `json:"data_scopes"`
			DataScopesStr     []string `json:"data_scopes_str"`
			HostController    int64    `json:"host_controller"`
			HostControllerStr string   `json:"host_controller_str"`
			HostInterface     int64    `json:"host_interface"`
			HostInterfaceStr  string   `json:"host_interface_str"`
			IsActive          bool     `json:"is_active"`
			Name              string   `json:"name"`
			NandType          int64    `json:"nand_type"`
			NandTypeStr       string   `json:"nand_type_str"`
			NandVersion       int64    `json:"nand_version"`
			NandVersionStr    string   `json:"nand_version_str"`
			ProductID         int64    `json:"product_id"`
		} `json:"product"`
		ProductImage struct {
			ArchivePath  string `json:"archive_path"`
			UserUploadID int64  `json:"user_upload_id"`
		} `json:"product_image"`
		ReleaseJira string `json:"release_jira"`
		Security    string `json:"security"`
		TestProject struct {
			TestProjectID   int64  `json:"test_project_id"`
			TestProjectName string `json:"test_project_name"`
		} `json:"test_project"`
		UserCreate struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"user_create"`
		UserLastModify struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"user_last_modify"`
	} `json:"results"`
	Total int64 `json:"total"`
}


type TestSystemNodeType struct {
	Count    int64         `json:"count"`
	Next     interface{}   `json:"next"`
	Previous interface{}   `json:"previous"`
	Results  []interface{} `json:"results"`
	Total    int64         `json:"total"`
}


type TestSystemOperatingSystemType struct {
	Count    int64       `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		Bitness           int64    `json:"bitness"`
		BitnessStr        string   `json:"bitness_str"`
		DataScopes        []int64  `json:"data_scopes"`
		DataScopesStr     []string `json:"data_scopes_str"`
		DateCreate        string   `json:"date_create"`
		DateLastModify    string   `json:"date_last_modify"`
		Description       string   `json:"description"`
		IsActive          bool     `json:"is_active"`
		Name              string   `json:"name"`
		OperatingSystemID int64    `json:"operating_system_id"`
		Tag               string   `json:"tag"`
		Type              int64    `json:"type"`
		TypeStr           string   `json:"type_str"`
		UserCreate        struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"user_create"`
		UserLastModify struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"user_last_modify"`
	} `json:"results"`
	Total int64 `json:"total"`
}


type TestSystemProductType struct {
	Count    int64       `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		Comments          string   `json:"comments"`
		DataScopes        []int64  `json:"data_scopes"`
		DataScopesStr     []string `json:"data_scopes_str"`
		HostController    int64    `json:"host_controller"`
		HostControllerStr string   `json:"host_controller_str"`
		HostInterface     int64    `json:"host_interface"`
		HostInterfaceStr  string   `json:"host_interface_str"`
		IsActive          bool     `json:"is_active"`
		Name              string   `json:"name"`
		NandType          int64    `json:"nand_type"`
		NandTypeStr       string   `json:"nand_type_str"`
		NandVersion       int64    `json:"nand_version"`
		NandVersionStr    string   `json:"nand_version_str"`
		ProductFamily     struct {
			Name            string `json:"name"`
			ProductFamilyID int64  `json:"product_family_id"`
		} `json:"product_family"`
		ProductID   int64 `json:"product_id"`
		ProductLink []struct {
			Capacity            []int64  `json:"capacity"`
			CapacityStr         []string `json:"capacity_str"`
			DataScopes          []int64  `json:"data_scopes"`
			FormFactor          []int64  `json:"form_factor"`
			FormFactorStr       []string `json:"form_factor_str"`
			IsActive            bool     `json:"is_active"`
			OverProvisioning    []int64  `json:"over_provisioning"`
			OverProvisioningStr []string `json:"over_provisioning_str"`
			ProductLinkID       int64    `json:"product_link_id"`
		} `json:"product_link"`
	} `json:"results"`
	Total int64 `json:"total"`
}


type TestSystemProductFamilyType struct {
	Count    int64       `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		Capacity            []int64  `json:"capacity"`
		CapacityStr         []string `json:"capacity_str"`
		DataScopes          []int64  `json:"data_scopes"`
		DataScopesStr       []string `json:"data_scopes_str"`
		FormFactor          []int64  `json:"form_factor"`
		FormFactorStr       []string `json:"form_factor_str"`
		HostController      string   `json:"host_controller"`
		IsActive            bool     `json:"is_active"`
		Name                string   `json:"name"`
		NandType            string   `json:"nand_type"`
		NandVersion         string   `json:"nand_version"`
		OverProvisioning    []int64  `json:"over_provisioning"`
		OverProvisioningStr []string `json:"over_provisioning_str"`
		ProductFamilyID     int64    `json:"product_family_id"`
		Products            []struct {
			Comments          string   `json:"comments"`
			DataScopes        []int64  `json:"data_scopes"`
			DataScopesStr     []string `json:"data_scopes_str"`
			HostController    int64    `json:"host_controller"`
			HostControllerStr string   `json:"host_controller_str"`
			HostInterface     int64    `json:"host_interface"`
			HostInterfaceStr  string   `json:"host_interface_str"`
			IsActive          bool     `json:"is_active"`
			Name              string   `json:"name"`
			NandType          int64    `json:"nand_type"`
			NandTypeStr       string   `json:"nand_type_str"`
			NandVersion       int64    `json:"nand_version"`
			NandVersionStr    string   `json:"nand_version_str"`
			ProductID         int64    `json:"product_id"`
		} `json:"products"`
	} `json:"results"`
	Total int64 `json:"total"`
}


type TestSystemProductLinkType struct {
	Count    int64       `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		Capacity            []int64  `json:"capacity"`
		CapacityStr         []string `json:"capacity_str"`
		DataScopes          []int64  `json:"data_scopes"`
		FormFactor          []int64  `json:"form_factor"`
		FormFactorStr       []string `json:"form_factor_str"`
		IsActive            bool     `json:"is_active"`
		OverProvisioning    []int64  `json:"over_provisioning"`
		OverProvisioningStr []string `json:"over_provisioning_str"`
		Product             int64    `json:"product"`
		ProductLinkID       int64    `json:"product_link_id"`
	} `json:"results"`
	Total int64 `json:"total"`
}


type TestSystemQualPlanType struct {
	Count    int64       `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		Capacity         []int64  `json:"capacity"`
		CapacityStr      []string `json:"capacity_str"`
		Customer         int64    `json:"customer"`
		CustomerStr      string   `json:"customer_str"`
		DataScopes       []int64  `json:"data_scopes"`
		DataScopesStr    []string `json:"data_scopes_str"`
		DateCreate       string   `json:"date_create"`
		DateLastModify   string   `json:"date_last_modify"`
		HostInterface    int64    `json:"host_interface"`
		HostInterfaceStr string   `json:"host_interface_str"`
		IsActive         bool     `json:"is_active"`
		Name             string   `json:"name"`
		Product          struct {
			Comments          string   `json:"comments"`
			DataScopes        []int64  `json:"data_scopes"`
			DataScopesStr     []string `json:"data_scopes_str"`
			HostController    int64    `json:"host_controller"`
			HostControllerStr string   `json:"host_controller_str"`
			HostInterface     int64    `json:"host_interface"`
			HostInterfaceStr  string   `json:"host_interface_str"`
			IsActive          bool     `json:"is_active"`
			Name              string   `json:"name"`
			NandType          int64    `json:"nand_type"`
			NandTypeStr       string   `json:"nand_type_str"`
			NandVersion       int64    `json:"nand_version"`
			NandVersionStr    string   `json:"nand_version_str"`
			ProductID         int64    `json:"product_id"`
		} `json:"product"`
		QualPlanID int64  `json:"qual_plan_id"`
		Security   string `json:"security"`
		TestPlans  []struct {
			Name       string `json:"name"`
			TestPlanID int64  `json:"test_plan_id"`
		} `json:"test_plans"`
		TestProject struct {
			TestProjectID   int64  `json:"test_project_id"`
			TestProjectName string `json:"test_project_name"`
		} `json:"test_project"`
		UserCreate struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"user_create"`
		UserLastModify struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"user_last_modify"`
	} `json:"results"`
	Total int64 `json:"total"`
}


type TestSystemStorageDriverType struct {
	Count    int64       `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		DataScopes      []int64  `json:"data_scopes"`
		DataScopesStr   []string `json:"data_scopes_str"`
		DateCreate      string   `json:"date_create"`
		DateLastModify  string   `json:"date_last_modify"`
		Description     string   `json:"description"`
		DriverName      string   `json:"driver_name"`
		IsActive        bool     `json:"is_active"`
		StorageDriverID int64    `json:"storage_driver_id"`
		Tag             string   `json:"tag"`
		UserCreate      struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"user_create"`
		UserLastModify struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"user_last_modify"`
		UserUpload struct {
			ArchivePath  string `json:"archive_path"`
			UserUploadID int64  `json:"user_upload_id"`
		} `json:"user_upload"`
		Version string `json:"version"`
	} `json:"results"`
	Total int64 `json:"total"`
}


type TestSystemStorageInterfaceType struct {
	Count    int64       `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		DataScopes         []int64  `json:"data_scopes"`
		DataScopesStr      []string `json:"data_scopes_str"`
		DateCreate         string   `json:"date_create"`
		DateLastModify     string   `json:"date_last_modify"`
		FormFactor         int64    `json:"form_factor"`
		FormFactorStr      string   `json:"form_factor_str"`
		InterfaceSpeed     int64    `json:"interface_speed"`
		InterfaceSpeedStr  string   `json:"interface_speed_str"`
		InterfaceType      int64    `json:"interface_type"`
		InterfaceTypeStr   string   `json:"interface_type_str"`
		IsActive           bool     `json:"is_active"`
		Raid               int64    `json:"raid"`
		RaidStr            string   `json:"raid_str"`
		StorageInterfaceID int64    `json:"storage_interface_id"`
		UserCreate         struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"user_create"`
		UserLastModify struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"user_last_modify"`
	} `json:"results"`
	Total int64 `json:"total"`
}


type TestSystemStorageInterfaceTypeType struct {
	Zero string `json:"0"`
	One  string `json:"1"`
	Two  string `json:"2"`
}


type TestSystemTestPeripheralType struct {
	Count    int64       `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		DataScopes       []int64 `json:"data_scopes"`
		DateCreate       string  `json:"date_create"`
		DateLastModify   string  `json:"date_last_modify"`
		IsActive         bool    `json:"is_active"`
		Name             string  `json:"name"`
		TestPeripheralID int64   `json:"test_peripheral_id"`
		UserCreate       struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"user_create"`
		UserLastModify struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"user_last_modify"`
		Value struct {
			Connection struct {
				BaudRate     string `json:"BAUD_RATE"`
				Interface    string `json:"INTERFACE"`
				IPAddress    string `json:"IP_ADDRESS"`
				PortName     string `json:"PORT_NAME"`
				PortNumber   string `json:"PORT_NUMBER"`
				SerialNumber string `json:"SERIAL_NUMBER"`
			} `json:"CONNECTION"`
			DeviceName string `json:"DEVICE_NAME"`
		} `json:"value"`
	} `json:"results"`
	Total int64 `json:"total"`
}


type TestSystemTestRequestType struct {
	Count    int64       `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		CycleEnd       []string `json:"cycle_end"`
		DataScopes     []int64  `json:"data_scopes"`
		DataScopesStr  []string `json:"data_scopes_str"`
		DateCreate     string   `json:"date_create"`
		DateLastModify string   `json:"date_last_modify"`
		Description    string   `json:"description"`
		Firmware       struct {
			Configuration string `json:"configuration"`
			FirmwareID    int64  `json:"firmware_id"`
			Name          string `json:"name"`
		} `json:"firmware"`
		IsActive  bool `json:"is_active"`
		LifeCycle struct {
			LifeCycleID int64  `json:"life_cycle_id"`
			Name        string `json:"name"`
			Value       int64  `json:"value"`
		} `json:"life_cycle"`
		Result             interface{} `json:"result"`
		ResultStr          string      `json:"result_str"`
		Status             string      `json:"status"`
		StatusStr          string      `json:"status_str"`
		StatusUpdate       []string    `json:"status_update"`
		StatusUpdateFreq   []int64     `json:"status_update_freq"`
		SubscriptionGroups []struct {
			GroupID int64  `json:"group_id"`
			Name    string `json:"name"`
			Users   []struct {
				UserID   int64  `json:"user_id"`
				Username string `json:"username"`
			} `json:"users"`
		} `json:"subscription_groups"`
		SubscriptionUsers []struct {
			UserID   int64  `json:"user_id"`
			Username string `json:"username"`
		} `json:"subscription_users"`
		TestProject struct {
			TestProjectID   int64  `json:"test_project_id"`
			TestProjectName string `json:"test_project_name"`
		} `json:"test_project"`
		TestRequestDetails []struct {
			CycleEnd         string      `json:"cycle_end"`
			Order            int64       `json:"order"`
			Priority         int64       `json:"priority"`
			PriorityStr      string      `json:"priority_str"`
			QualPlan         interface{} `json:"qual_plan"`
			StatusUpdate     string      `json:"status_update"`
			StatusUpdateFreq int64       `json:"status_update_freq"`
			TestCycle        []struct {
				Name        string      `json:"name"`
				Result      interface{} `json:"result"`
				ResultStr   interface{} `json:"result_str"`
				Status      int64       `json:"status"`
				StatusStr   string      `json:"status_str"`
				TestCycleID int64       `json:"test_cycle_id"`
			} `json:"test_cycle"`
			TestPlan struct {
				Name       string `json:"name"`
				TestPlanID int64  `json:"test_plan_id"`
			} `json:"test_plan"`
			TestRequestDetailID int64 `json:"test_request_detail_id"`
			User                struct {
				ID       int64  `json:"id"`
				Username string `json:"username"`
			} `json:"user"`
		} `json:"test_request_details"`
		TestRequestID int64    `json:"test_request_id"`
		TestTypes     []int64  `json:"test_types"`
		TestTypesStr  []string `json:"test_types_str"`
		UserCreate    struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"user_create"`
		UserLastModify struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"user_last_modify"`
	} `json:"results"`
	Total int64 `json:"total"`
}


type TestSystemTestRequestDetailType struct {
	Count    int64       `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		CycleEnd         string      `json:"cycle_end"`
		Order            int64       `json:"order"`
		Priority         int64       `json:"priority"`
		PriorityStr      string      `json:"priority_str"`
		QualPlan         interface{} `json:"qual_plan"`
		StatusUpdate     string      `json:"status_update"`
		StatusUpdateFreq int64       `json:"status_update_freq"`
		TestPlan         struct {
			Name       string `json:"name"`
			TestPlanID int64  `json:"test_plan_id"`
		} `json:"test_plan"`
		TestRequest struct {
			TestProject struct {
				TestProjectID   int64  `json:"test_project_id"`
				TestProjectName string `json:"test_project_name"`
			} `json:"test_project"`
			TestRequestID int64 `json:"test_request_id"`
		} `json:"test_request"`
		TestRequestDetailID int64 `json:"test_request_detail_id"`
		User                struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"user"`
	} `json:"results"`
	Total int64 `json:"total"`
}


type TestSystemViewType struct {
	Count    int64       `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		AgentConnected bool   `json:"agent_connected"`
		Autodetect     bool   `json:"autodetect"`
		BenchID        string `json:"bench_id"`
		Busy           bool   `json:"busy"`
		Capability     []struct {
			StorageDriver interface{} `json:"storage_driver"`
			TestTool      struct {
				Name       string `json:"name"`
				Revision   string `json:"revision"`
				TestToolID int64  `json:"test_tool_id"`
			} `json:"test_tool"`
		} `json:"capability"`
		Connected    bool `json:"connected"`
		Connectivity []struct {
			DeviceUnderTesting struct {
				Capacity             int64  `json:"capacity"`
				CapacityStr          string `json:"capacity_str"`
				DeviceUnderTestingID int64  `json:"device_under_testing_id"`
				HealthStatus         int64  `json:"health_status"`
				HealthStatusStr      string `json:"health_status_str"`
				Model                string `json:"model"`
				ProductFamily        struct {
					Name            string `json:"name"`
					ProductFamilyID int64  `json:"product_family_id"`
				} `json:"product_family"`
				SerialNumber string `json:"serial_number"`
			} `json:"device_under_testing"`
		} `json:"connectivity"`
		Controller       interface{} `json:"controller"`
		DataScopes       []int64     `json:"data_scopes"`
		DataScopesStr    []string    `json:"data_scopes_str"`
		DateCreate       string      `json:"date_create"`
		DateLastModify   string      `json:"date_last_modify"`
		DirectAccess     bool        `json:"direct_access"`
		HostName         string      `json:"host_name"`
		IP               string      `json:"ip"`
		IsActive         bool        `json:"is_active"`
		MainPurpose      string      `json:"main_purpose"`
		Model            string      `json:"model"`
		RackID           string      `json:"rack_id"`
		StorageInterface []struct {
			Interface          string `json:"interface"`
			InterfaceSpeed     int64  `json:"interface_speed"`
			InterfaceSpeedStr  string `json:"interface_speed_str"`
			InterfaceType      int64  `json:"interface_type"`
			InterfaceTypeStr   string `json:"interface_type_str"`
			Raid               int64  `json:"raid"`
			RaidStr            string `json:"raid_str"`
			StorageInterfaceID int64  `json:"storage_interface_id"`
		} `json:"storage_interface"`
		TestPeripherals []struct {
			Name             string `json:"name"`
			TestPeripheralID int64  `json:"test_peripheral_id"`
		} `json:"test_peripherals"`
		TestRunResults struct {
			CountInProgress int64 `json:"count_in_progress"`
			CountQueued     int64 `json:"count_queued"`
			Tests           []struct {
				Status        int64  `json:"status"`
				StatusStr     string `json:"status_str"`
				TestCycleID   int64  `json:"test_cycle_id"`
				TestCycleName string `json:"test_cycle_name"`
				TestName      string `json:"test_name"`
			} `json:"tests"`
		} `json:"test_run_results"`
		TestSystemID int64 `json:"test_system_id"`
		Usage        int64 `json:"usage"`
		UserCreate   struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"user_create"`
		UserLastModify struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"user_last_modify"`
		Vendor string `json:"vendor"`
	} `json:"results"`
	Total int64 `json:"total"`
}


type TriggerEventChoicesType struct {
	One   string `json:"1"`
	Two   string `json:"2"`
	Three string `json:"3"`
	Four  string `json:"4"`
	Five  string `json:"5"`
	Six   string `json:"6"`
	Seven string `json:"7"`
	Eight string `json:"8"`
}


type UsageStatsType struct {
	Count    int64       `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		APIURL        string  `json:"api_url"`
		Count         int64   `json:"count"`
		DBQueries     int64   `json:"db_queries"`
		DBTime        float64 `json:"db_time"`
		PythonTime    float64 `json:"python_time"`
		RequestMethod string  `json:"request_method"`
		StartTime     string  `json:"start_time"`
		TotalTime     float64 `json:"total_time"`
	} `json:"results"`
	Total int64 `json:"total"`
}


type UserGroupType struct {
	Count    int64       `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		DataScope         []int64  `json:"data_scope"`
		DataScopeStr      []string `json:"data_scope_str"`
		DateCreate        string   `json:"date_create"`
		DateLastModify    string   `json:"date_last_modify"`
		EmailNotification int64    `json:"email_notification"`
		ID                int64    `json:"id"`
		IsActive          bool     `json:"is_active"`
		Name              string   `json:"name"`
		TriggerEvents     []int64  `json:"trigger_events"`
		TriggerEventsStr  []string `json:"trigger_events_str"`
		UserCreate        struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"user_create"`
		UserLastModify struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"user_last_modify"`
		Users []struct {
			Role     string `json:"role"`
			UserID   int64  `json:"user_id"`
			Username string `json:"username"`
		} `json:"users"`
	} `json:"results"`
	Total int64 `json:"total"`
}


type UserSettingsType struct {
	ColorSchema    int64       `json:"color_schema"`
	ColorSchemaStr string      `json:"color_schema_str"`
	Fields         interface{} `json:"fields"`
	UserSettingID  int64       `json:"user_setting_id"`
}


type UserUploadViewType struct {
	Count    int64       `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		ArchivePath string `json:"archive_path"`
		ContentType string `json:"content_type"`
		CreatedDate string `json:"created_date"`
		Owner       struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"owner"`
		Type         int64 `json:"type"`
		UserUploadID int64 `json:"user_upload_id"`
	} `json:"results"`
	Total int64 `json:"total"`
}


type UserUploadViewTypeType struct {
	One   string `json:"1"`
	One0  string `json:"10"`
	One1  string `json:"11"`
	One2  string `json:"12"`
	One3  string `json:"13"`
	One4  string `json:"14"`
	Two   string `json:"2"`
	Three string `json:"3"`
	Four  string `json:"4"`
	Five  string `json:"5"`
	Six   string `json:"6"`
	Seven string `json:"7"`
	Eight string `json:"8"`
	Nine  string `json:"9"`
}


type UsersType struct {
	Count    int64       `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		ActiveDataScope  int64         `json:"active_data_scope"`
		DataScopes       []int64       `json:"data_scopes"`
		DataScopesStr    []string      `json:"data_scopes_str"`
		DateCreate       string        `json:"date_create"`
		DateLastModify   string        `json:"date_last_modify"`
		Email            string        `json:"email"`
		ID               int64         `json:"id"`
		Role             string        `json:"role"`
		Sessionid        string        `json:"sessionid"`
		Timezone         string        `json:"timezone"`
		TriggerEvents    []interface{} `json:"trigger_events"`
		TriggerEventsStr []interface{} `json:"trigger_events_str"`
		UserCreate       struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"user_create"`
		UserGroups []struct {
			Name        string `json:"name"`
			UserGroupID int64  `json:"user_group_id"`
		} `json:"user_groups"`
		UserLastModify struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"user_last_modify"`
		Username string `json:"username"`
	} `json:"results"`
	Total int64 `json:"total"`
}


type UsersMeType struct {
	ActiveDataScope  int64         `json:"active_data_scope"`
	DataScopes       []int64       `json:"data_scopes"`
	DataScopesStr    []string      `json:"data_scopes_str"`
	DateCreate       interface{}   `json:"date_create"`
	DateLastModify   interface{}   `json:"date_last_modify"`
	Email            string        `json:"email"`
	ID               int64         `json:"id"`
	Role             string        `json:"role"`
	Sessionid        interface{}   `json:"sessionid"`
	Timezone         string        `json:"timezone"`
	TriggerEvents    []interface{} `json:"trigger_events"`
	TriggerEventsStr []interface{} `json:"trigger_events_str"`
	UserCreate       interface{}   `json:"user_create"`
	UserGroups       []struct {
		Name        string `json:"name"`
		UserGroupID int64  `json:"user_group_id"`
	} `json:"user_groups"`
	UserLastModify interface{} `json:"user_last_modify"`
	Username       string      `json:"username"`
}

type V2FirmwareReportType struct {
	Configurations []string `json:"configurations"`
	Failed         []int64  `json:"failed"`
	Success        []int64  `json:"success"`
}


