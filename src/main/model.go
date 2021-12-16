package main

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/guregu/null/zero"
)

/**

The following structures are used to retrieve information of the Database tables object for SQL queries.
Description are not needed because it represents the database in details.

*/

type AppAnalytics_v3 struct {
	Serial                      string `db:"serial"`
	PackageName                 string `db:"package_name"`
	InstallCount                string `db:"install_count"`
	UninstallCount              string `db:"uninstall_count"`
	RestoredFromPreload         string `db:"restored_from_preload"`
	NotRestoredFromPreload      string `db:"not_restored_from_preload"`
	InstalledFromParental       string `db:"installed_from_parental"`
	UninstalledFromParental     string `db:"uninstalled_from_parental"`
	NewAppInstalledFromParental string `db:"new_app_installed_from_parental"`
	UpdatedFromParental         string `db:"updated_from_parental"`
	Id                          int64  `db:"id"`
}

type AppAnalytics_v4 struct {
	Serial         string `db:"serial"`
	PackageName    string `db:"package_name"`
	InstallCount   string `db:"install_count"`
	UninstallCount string `db:"uninstall_count"`
	Id             int64  `db:"id"`
}

type AccessToken struct {
	IssuedTo   string `json:"issued_to"`
	Audience   string `json:"audience"`
	UserId     string `json:"user_id"`
	Scope      string `json:"scope"`
	ExpiresIn  int    `json:"expires_in"`
	AccessType string `json:"access_type"`
}

type ActivityStatusAnalytics struct {
	Serial       string `db:"serial"`
	AnalyticsUid string `db:"analytics_uid"`
	PackageName  string `db:"package_name"`
	ActivityName string `db:"activity_name"`
	Enabled      string `db:"enabled"`
	LaunchCount  string `db:"launch_count"`
	TimeSpent    string `db:"time_spent"`
	Id           string `db:"id"`
}

type ActivityStatusAnalyticsByRomId struct {
	RomId        string         `db:"rom_id"`
	PackageName  string         `db:"package_name"`
	ActivityName string         `db:"activity_name"`
	LaunchCount  int64          `db:"launch_count"`
	TimeSpent    int64          `db:"time_spent"`
	UpdatedAt    mysql.NullTime `db:"updated_at"`
	Id           int64          `db:"id"`
}

type Profile struct {
	IdProfile           int64          `db:"id_profile"`
	Serial              string         `db:"serial"`
	IsBoy               string         `db:"is_boy"`
	Birth               mysql.NullTime `db:"birth"`
	InternetAccessMode  string         `db:"internet_access_mode"`
	TimeConstrolStatus  string         `db:"time_constrol_status"`
	AllowUSBStatus      string         `db:"allow_USB_status"`
	AuthorizeAdsStatus  string         `db:"authorize_ads_status"`
	AutoAuthoriseStatus string         `db:"auto_authorise_status"`
}

type ProfileAnalytics struct {
	Idx                        int64          `db:"idx"`
	IdProfile                  int64          `db:"id_profile"`
	Serial                     string         `db:"serial"`
	AnalyticsUid               string         `db:"analytics_uid"`
	IsBoy                      string         `db:"is_boy"`
	Birth                      mysql.NullTime `db:"birth"`
	InternetAccessMode         string         `db:"internet_access_mode"`
	TimeControlStatus          string         `db:"time_control_status"`
	IsTimeSlotEnabled          string         `db:"is_time_slot_enabled"`
	DailyPlayTimeWeekDays      string         `db:"daily_play_time_week_days"`
	MaxSessionDurationWeekDays string         `db:"max_session_duration_week_days"`
	RestTimeWeekDays           string         `db:"rest_time_week_days"`
	AllowUSBStatus             string         `db:"allow_USB_status"`
	AuthorizeAdsStatus         string         `db:"authorize_ads_status"`
	AutoAuthoriseStatus        string         `db:"auto_authorise_status"`
	ProfileChanged             string         `db:"profile_changed"`
	ChildProtectedByPwd        string         `db:"child_protected_by_pwd"`
	IsWebListActivated         string         `db:"is_web_list_activated"`
	Deleted                    string         `db:"deleted"`
}

type ProfileAnalyticsV4 struct {
	Idx                        int64          `db:"idx"`
	IdProfile                  int64          `db:"id_profile"`
	Serial                     string         `db:"serial"`
	Curef                      zero.String    `db:"curef"`
	AnalyticsUid               string         `db:"analytics_uid"`
	IsBoy                      string         `db:"is_boy"`
	Birth                      mysql.NullTime `db:"birth"`
	InternetAccessMode         string         `db:"internet_access_mode"`
	TimeControlStatus          string         `db:"time_control_status"`
	IsTimeSlotEnabled          string         `db:"is_time_slot_enabled"`
	DailyPlayTimeWeekDays      string         `db:"daily_play_time_week_days"`
	MaxSessionDurationWeekDays string         `db:"max_session_duration_week_days"`
	RestTimeWeekDays           string         `db:"rest_time_week_days"`
	AllowUSBStatus             string         `db:"allow_USB_status"`
	AuthorizeAdsStatus         string         `db:"authorize_ads_status"`
	AutoAuthoriseStatus        string         `db:"auto_authorise_status"`
	ProfileChanged             string         `db:"profile_changed"`
	ChildProtectedByPwd        string         `db:"child_protected_by_pwd"`
	IsWebListActivated         string         `db:"is_web_list_activated"`
	Deleted                    string         `db:"deleted"`
}

type Serial struct {
	InsertRank                int64          `db:"insert_rank"`
	Serial                    string         `db:"serial"`
	Curef                     string         `db:"curef"`
	DateProd                  mysql.NullTime `db:"date_prod"`
	RoBuildDisplayId          string         `db:"ro_build_display_id"`
	RoBuildVersionIncremental string         `db:"ro_build_version_incremental"`
	RoProductModel            string         `db:"ro_product_model"`
	RoProductName             string         `db:"ro_product_name"`
	RoProductDevice           string         `db:"ro_product_device"`
	RoProductBoard            sql.NullString `db:"ro_product_board"`
	RoProductLocaleLanguage   sql.NullString `db:"ro_product_locale_language"`
	RoProductLocaleRegion     sql.NullString `db:"ro_product_locale_region"`
	KurioSystemVersion        sql.NullString `db:"kurio_system_version"`
	MACAddress                string         `db:"MAC_address"`
	LastIp                    string         `db:"last_ip"`
	SystemLanguage            string         `db:"system_language"`
	ServerEntree              mysql.NullTime `db:"server_entree"`
	FirstActivation           mysql.NullTime `db:"first_activation"`
	SetupDate                 mysql.NullTime `db:"setup_date"`
	LastModification          mysql.NullTime `db:"last_modification"`
	ParentEmail               string         `db:"parent_email"`
	MailActive                string         `db:"mail_active"`
	OwnerLockscreenCode       sql.NullString `db:"owner_lockscreen_code"`
	ActiveProfiles            sql.NullInt64  `db:"active_profiles"`
	ChildrenProfiles          sql.NullInt64  `db:"children_profiles"`
	PwdReset                  string         `db:"pwd_reset"`
	LastParentLog             mysql.NullTime `db:"last_parent_log"`
	AppManagementCount        string         `db:"app_management_count"`
	FaqCount                  string         `db:"faq_count"`
	ContactUsCount            string         `db:"contact_us_count"`
	UserManualCount           string         `db:"user_manual_count"`
	TimeControlUsedOnce       string         `db:"time_control_used_once"`
	UpdateAt                  mysql.NullTime `db:"update_at"`
}

type SerialAnalytics struct {
	InsertRank                int64          `db:"insert_rank"`
	Serial                    string         `db:"serial"`
	DateProd                  mysql.NullTime `db:"date_prod"`
	LastParentLog             mysql.NullTime `db:"last_parent_log"`
	RoBuildDisplayId          sql.NullString `db:"ro_build_display_id"`
	RoBuildVersionIncremental sql.NullString `db:"ro_build_version_incremental"`
	RoProductModel            sql.NullString `db:"ro_product_model"`
	RoProductName             sql.NullString `db:"ro_product_name"`
	RoProductDevice           sql.NullString `db:"ro_product_device"`
	RoProductBoard            sql.NullString `db:"ro_product_board"`
	RoProductLocaleLanguage   sql.NullString `db:"ro_product_locale_language"`
	RoProductLocaleRegion     sql.NullString `db:"ro_product_locale_region"`
	KurioSystemVersion        string         `db:"kurio_system_version"`
	MACAddress                sql.NullString `db:"MAC_address"`
	LastIp                    sql.NullString `db:"last_ip"`
	SystemLanguage            string         `db:"system_language"`
	ServerEntree              mysql.NullTime `db:"server_entree"`
	FirstActivation           mysql.NullTime `db:"first_activation"`
	SetupDate                 mysql.NullTime `db:"setup_date"`
	LastModification          mysql.NullTime `db:"last_modification"`
	ParentEmail               string         `db:"parent_email"`
	MailActive                string         `db:"mail_active"`
	OwnerLockscreenCode       sql.NullString `db:"owner_lockscreen_code"`
	ActiveProfiles            sql.NullInt64  `db:"active_profiles"`
	ChildrenProfiles          sql.NullInt64  `db:"children_profiles"`
	PwdReset                  string         `db:"pwd_reset"`
	AppManagementCount        string         `db:"app_management_count"`
	FaqCount                  string         `db:"faq_count"`
	ContactUsCount            string         `db:"contact_us_count"`
	UserManualCount           string         `db:"user_manual_count"`
	TimeControlUsedOnce       string         `db:"time_control_used_once"`
}

type SerialAnalyticsv4 struct {
	InsertRank                int64          `db:"insert_rank"`
	Serial                    string         `db:"serial"`
	Curef                     string         `db:"curef"`
	DateProd                  mysql.NullTime `db:"date_prod"`
	LastParentLog             mysql.NullTime `db:"last_parent_log"`
	RoBuildDisplayId          sql.NullString `db:"ro_build_display_id"`
	RoBuildVersionIncremental sql.NullString `db:"ro_build_version_incremental"`
	RoProductModel            sql.NullString `db:"ro_product_model"`
	RoProductName             sql.NullString `db:"ro_product_name"`
	RoProductDevice           sql.NullString `db:"ro_product_device"`
	RoProductBoard            sql.NullString `db:"ro_product_board"`
	RoProductLocaleLanguage   sql.NullString `db:"ro_product_locale_language"`
	RoProductLocaleRegion     sql.NullString `db:"ro_product_locale_region"`
	KurioSystemVersion        string         `db:"kurio_system_version"`
	MACAddress                sql.NullString `db:"MAC_address"`
	LastIp                    sql.NullString `db:"last_ip"`
	SystemLanguage            string         `db:"system_language"`
	ServerEntree              mysql.NullTime `db:"server_entree"`
	FirstActivation           mysql.NullTime `db:"first_activation"`
	SetupDate                 mysql.NullTime `db:"setup_date"`
	LastModification          mysql.NullTime `db:"last_modification"`
	ParentEmail               string         `db:"parent_email"`
	MailActive                string         `db:"mail_active"`
	OwnerLockscreenCode       sql.NullString `db:"owner_lockscreen_code"`
	ActiveProfiles            sql.NullInt64  `db:"active_profiles"`
	ChildrenProfiles          sql.NullInt64  `db:"children_profiles"`
	PwdReset                  string         `db:"pwd_reset"`
	AppManagementCount        string         `db:"app_management_count"`
	FaqCount                  string         `db:"faq_count"`
	ContactUsCount            string         `db:"contact_us_count"`
	UserManualCount           string         `db:"user_manual_count"`
	TimeControlUsedOnce       string         `db:"time_control_used_once"`
	Update_at                 []byte         `db:"update_at"`
}

type SerialOld7Tab struct {
	Serial          string `db:"serial_number"`
	MACAddress      string `db:"mac_address"`
	ParentEmail     string `db:"email"`
	BuildModel      string `db:"build_model"`
	FirstActivation string `db:"inserted_on"`
}

type Tab7SByRomId struct {
	Country string
	RomId   string `db:"rom_id"`
	Count   int64  `db:"count"`
}

type Analytics_users struct {
	Id          int64  `db:"id"`
	Login       string `db:"login"`
	Pass        string `db:"password_user"`
	Date_create string `db:"date_insert"`
}

type Tab7XByRomId struct {
	Country string
	RomId   string `db:"rom_id"`
	Count   int64  `db:"count"`
}

type TabV4ByRomId struct {
	Country string
	RomId   string `db:"rom_id"`
	Count   int64  `db:"count"`
}

type Tab7XActivation struct {
	ActivationDate string `db:"activation_date"`
	Count          int64  `db:"count"`
}

type TabV4Activation struct {
	ActivationDate string `db:"activation_date"`
	Count          int64  `db:"count"`
}
type Tab7SActivation struct {
	ActivationDate string `db:"activation_date"`
	Count          int64  `db:"count"`
}

type Tab7XAppPref struct {
	PackageName      string `db:"package_name"`
	TotalTimeSpent   string `db:"total_time_spent"`
	TotalLaunchCount string `db:"total_launch_count"`
}

type TabV4AppPref struct {
	PackageName      string `db:"package_name"`
	TotalTimeSpent   string `db:"total_time_spent"`
	TotalLaunchCount string `db:"total_launch_count"`
}

type Tab7SAppPref struct {
	PackageName      string `db:"package_name"`
	TotalTimeSpent   string `db:"total_time_spent"`
	TotalLaunchCount string `db:"total_launch_count"`
}

type AllProfileInfos struct {
	Total                   int64 `db:"total"`
	Girl                    int64 `db:"girl"`
	TimeControlEnabled      int64 `db:"time_control_enabled"`
	TimeSlotEnabled         int64 `db:"time_slot_enabled"`
	AdvancedSettingsEnabled int64 `db:"advanced_setting_enabled"`
	AllowUsbStatus          int64 `db:"allow_USB_status"`
	AuthorizeAdsStatus      int64 `db:"authorize_ads_status"`
	AutoAuthorizeStatus     int64 `db:"auto_authorise_status"`
	ProfileChanged          int64 `db:"profile_changed"`
	ChildProtectedByPWD     int64 `db:"child_protected_by_pwd"`
	WebListActivated        int64 `db:"web_list_activated"`
	Deleted                 int64 `db:"deleted"`
}

type GenericWith2Field struct {
	Name  string `db:"name"`
	Count int64  `db:"count"`
}

type AccountInformation struct {
	Login    string `db:"login`
	Password string `db:password`
}

type GenericWith3Field struct {
	Name        string `db:"name"`
	Count       int64  `db:"count"`
	CountSecond int64  `db:"countSecond"`
}
