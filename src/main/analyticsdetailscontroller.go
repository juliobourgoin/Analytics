package main

import (
	"github.com/lioonel/gorp"
	"github.com/martini-contrib/render"
	"log"
	"net/http"
	"reflect"
	//	"strconv"
)

//called first to select rom id to extract
func AnalyticsViewerDetails(dmap *gorp.DbMap, res http.ResponseWriter, req *http.Request, r render.Render) {

	/**
	Name	:	AnalyticsViewerDetails
	Version :	1.0
	Description :
				This function is dedicated to initialize information on details link on menu of analytics_viewer

	State	: Used
	*/

	if (idsession == "Jeremy") || (idsession == "Caroline") {

		var queryWithSerialWhere string
		var queryWithSerialAnd string
		activateDBTrace(dmap)
		serialSelected := req.FormValue("serial_selected")
		versionSelected := req.FormValue("version_selected")

		if serialSelected == "tmpCuref" {
			queryWithSerialWhere = " where LENGTH(serial) < 20"
			queryWithSerialAnd = " and LENGTH(serial) < 20"
		} else if serialSelected != "" && len(serialSelected) != 4 && serialSelected != "all" {
			queryWithSerialWhere = " where serial like '%" + serialSelected + "'"
			queryWithSerialAnd = " and serial like '%" + serialSelected + "'"
		} else if serialSelected != "" && serialSelected != "all" {
			queryWithSerialWhere = " where curef like '%" + serialSelected + "'"
			queryWithSerialAnd = " and curef like '%" + serialSelected + "'"
		}

		if versionSelected == "" {
			versionSelected = "4" // Default version
		}

		var countValue map[string]string
		var otaRepartition []GenericWith2Field
		var systemVersion []GenericWith2Field
		var profilesByAge []GenericWith2Field
		var nbChildProfiles []GenericWith2Field

		if versionSelected == "3" {
			countValue = GetCountValues_v3(dmap, queryWithSerialWhere, queryWithSerialAnd)
			otaRepartition = GetOTARepartition_v3(dmap, queryWithSerialWhere)
			systemVersion = GetSystemVersion_v3(dmap, queryWithSerialWhere)
			profilesByAge = GetProfilesByAge_v3(dmap, queryWithSerialAnd)
			nbChildProfiles = GetNumberOfChilds_v3(dmap, queryWithSerialWhere)
		} else if versionSelected == "4" {
			countValue = GetCountValues_v4(dmap, queryWithSerialWhere, queryWithSerialAnd)
			otaRepartition = GetOTARepartition_v4(dmap, queryWithSerialWhere)
			systemVersion = GetSystemVersion_v4(dmap, queryWithSerialWhere)
			profilesByAge = GetProfilesByAge_v4(dmap, queryWithSerialAnd)
			nbChildProfiles = GetNumberOfChilds_v4(dmap, queryWithSerialWhere)
		} else if versionSelected == "5" {
			countValue = GetCountValues_v6(dmap, queryWithSerialWhere, queryWithSerialAnd)
			otaRepartition = GetOTARepartition_v6(dmap, queryWithSerialWhere)
			systemVersion = GetSystemVersion_v6(dmap, queryWithSerialWhere)
			profilesByAge = GetProfilesByAge_v6(dmap, queryWithSerialAnd)
			nbChildProfiles = GetNumberOfChilds_v6(dmap, queryWithSerialWhere)
		} else if versionSelected == "6" {
			countValue = GetCountValues_v8(dmap, queryWithSerialWhere, queryWithSerialAnd)
			otaRepartition = GetOTARepartition_v8(dmap, queryWithSerialWhere)
			systemVersion = GetSystemVersion_v8(dmap, queryWithSerialWhere)
			profilesByAge = GetProfilesByAge_v8(dmap, queryWithSerialAnd)
			nbChildProfiles = GetNumberOfChilds_v8(dmap, queryWithSerialWhere)
		}

		newmap := map[string]interface{}{"metatitle": "Extract Analytics Data", "serial_selected": serialSelected, "version_selected": versionSelected,
			"countValue": countValue, "otaRepartition": otaRepartition, "systemVersion": systemVersion, "profilesByAge": profilesByAge, "nbChildProfiles": nbChildProfiles}

		r.HTML(200, "analytics_viewer_details", newmap)

	} else {
		http.Redirect(res, req, "/analytics_bad_user", 302)
	}
}

func Application_details(dmap *gorp.DbMap, res http.ResponseWriter, req *http.Request, r render.Render) {

	/**
	Name	:	AnalyticsViewerDetails
	Version :	1.0
	Description :
				This function is dedicated to initialize information on details link on menu of analytics_viewer

	State	: Used
	*/

	if (idsession == "Jeremy") || (idsession == "Caroline") {

		newmap := map[string]interface{}{"metatitle": "Analytics details"}

		r.HTML(200, "Application_details", newmap)

	} else {
		http.Redirect(res, req, "/analytics_bad_user", 302)
	}
}


func AnalyticsViewerRomId(dmap *gorp.DbMap, res http.ResponseWriter, req *http.Request, r render.Render) {

	/**
	Name	:	AnalyticsViewerRomId
	Version :	1.0
	Description :
				This function is dedicated to initialize information on details link on menu of analytics_viewer when user change ROMID

	State	: Used
	*/

	if idsession == "Olivier" {

		var queryWithSerialWhere string
		var queryWithSerialAnd string
		activateDBTrace(dmap)
		serialSelected := req.FormValue("serial_selected_Olivier")
		versionSelected := req.FormValue("version_selected_Olivier")

		if serialSelected == "tmpCuref" {
			queryWithSerialWhere = " where LENGTH(serial) < 20"
			queryWithSerialAnd = " and LENGTH(serial) < 20"
		} else if serialSelected != "" && len(serialSelected) != 4 && serialSelected != "all" {
			queryWithSerialWhere = " where serial like '%" + serialSelected + "'"
			queryWithSerialAnd = " and serial like '%" + serialSelected + "'"
		} else if serialSelected != "" && serialSelected != "all" {
			queryWithSerialWhere = " where curef like '%" + serialSelected + "'"
			queryWithSerialAnd = " and curef like '%" + serialSelected + "'"
		}

		if versionSelected == "" {
			versionSelected = "4" // Default version
		}

		var countValue map[string]string
		var otaRepartition []GenericWith2Field
		var systemVersion []GenericWith2Field
		var profilesByAge []GenericWith2Field
		var nbChildProfiles []GenericWith2Field

		if versionSelected == "3" {
			countValue = GetCountValues_v3(dmap, queryWithSerialWhere, queryWithSerialAnd)
			otaRepartition = GetOTARepartition_v3(dmap, queryWithSerialWhere)
			systemVersion = GetSystemVersion_v3(dmap, queryWithSerialWhere)
			profilesByAge = GetProfilesByAge_v3(dmap, queryWithSerialAnd)
			nbChildProfiles = GetNumberOfChilds_v3(dmap, queryWithSerialWhere)
		} else if versionSelected == "4" {
			countValue = GetCountValues_v4(dmap, queryWithSerialWhere, queryWithSerialAnd)
			otaRepartition = GetOTARepartition_v4(dmap, queryWithSerialWhere)
			systemVersion = GetSystemVersion_v4(dmap, queryWithSerialWhere)
			profilesByAge = GetProfilesByAge_v4(dmap, queryWithSerialAnd)
			nbChildProfiles = GetNumberOfChilds_v4(dmap, queryWithSerialWhere)
		}

		newmap := map[string]interface{}{"metatitle": "Extract Analytics Data", "serial_selected_Olivier": serialSelected, "version_selected_Olivier": versionSelected,
			"countValue": countValue, "otaRepartition": otaRepartition, "systemVersion": systemVersion, "profilesByAge": profilesByAge, "nbChildProfiles": nbChildProfiles}

		r.HTML(200, "extract_data_weekly", newmap)

	} else {
		http.Redirect(res, req, "/analytics_bad_user", 302)
	}
}

func GetCountValues_v3(dmap *gorp.DbMap, queryWithSerialWhere, queryWithSerialAnd string) map[string]string {

	/**
	Name	:	GetCountValues_v3
	Version :	1.0
	Description :
				This function is dedicated to initialize information in tab of the analytics details page for X series.

	State	: Used
	*/

	var (
		selectQuery            string
		totalCount             string
		boys                   string
		girls                  string
		pwdProtected           string
		internetFilteringCount string
		timeManagementCount    string
		timeSlotCount          string
		advancedSettingsCount  string
		dailyPlaytimeCount     string
		maxSessionCount        string
		appManagementCount     string
		authorizeUSBCount      string
		authorizeAdsCount      string
		changeProfileCount     string
		weblistCount           string
		faqCount               string
		contactUsCount         string
		userManualCount        string
		adultProfiles          string
		childrenProfiles       string
	)

	selectQuery = "select count(*) from Profile_analytics "
	totalCount = GetCount(dmap, queryWithSerialWhere, selectQuery)

	selectQuery = "select count(*) from Profile_analytics where is_boy = 'true'"
	boys = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(*) from Profile_analytics where is_boy = 'false'"
	girls = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(*) from Profile_analytics where child_protected_by_pwd = 'true'"
	pwdProtected = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(internet_access_mode) from Profile_analytics where internet_access_mode = 'CUSTOM'"
	internetFilteringCount = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(time_control_name_id) from Profile_analytics where time_control_name_id like 'user_preset%'"
	timeManagementCount = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(is_time_slot_enabled) from Profile_analytics where is_time_slot_enabled = 'true'"
	timeSlotCount = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(is_advanced_settings_enabled) from Profile_analytics where is_advanced_settings_enabled = 'true'"
	advancedSettingsCount = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(daily_play_time_week_days) from Profile_analytics where (daily_play_time_week_days != '0' OR daily_play_time_week_end != '0')"
	dailyPlaytimeCount = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(max_session_duration_week_days) from Profile_analytics where (max_session_duration_week_days != '0' OR max_session_duration_week_end != '0')"
	maxSessionCount = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(app_management_count) from Serial_analytics where (app_management_count != '0' and app_management_count != '')"
	appManagementCount = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(allow_usb_status) from Profile_analytics where allow_usb_status = 'true'"
	authorizeUSBCount = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(authorize_ads_status) from Profile_analytics where authorize_ads_status = 'true'"
	authorizeAdsCount = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(profile_changed) from Profile_analytics where profile_changed = 'true'"
	changeProfileCount = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(is_web_list_activated) from Profile_analytics where is_web_list_activated = 'true'"
	weblistCount = GetCount(dmap, queryWithSerialAnd, selectQuery)

	/**** requetes sql  a revoir ******************/

	// selectQuery = "SELECT count(package_name) FROM KurioS_.7x_app_static where installed_from_parental > 0"
	// appInstalledPreloadCount := GetCount(dmap, queryWithSerialAnd, selectQuery)

	appInstalledPreloadCount := "-1"
	// selectQuery = "SELECT count(package_name) FROM KurioS_.7x_app_static where (install_count - restored_from_preload - installed_from_parental - new_app_installed_from_parental) > 0"
	// appInstalledGPlayCount := GetCount(dmap, queryWithSerialAnd, selectQuery)
	appInstalledGPlayCount := "-1"
	/*************************************/

	selectQuery = "select count(faq_count) from Serial_analytics where faq_count > '0' and faq_count != ''"
	faqCount = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(contact_us_count) from Serial_analytics where contact_us_count > '0' and contact_us_count != ''"
	contactUsCount = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(user_manual_count) from Serial_analytics where user_manual_count > '0' and user_manual_count != ''"
	userManualCount = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "SELECT count(active_profiles) as count FROM KurioS_.Serial_analytics where (active_profiles - children_profiles) > 1"
	adultProfiles = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "SELECT count(children_profiles) as count FROM KurioS_.Serial_analytics where children_profiles > 1"
	childrenProfiles = GetCount(dmap, queryWithSerialAnd, selectQuery)

	countValue := make(map[string]string)

	countValue["totalCount"] = totalCount
	countValue["boys"] = boys
	countValue["girls"] = girls
	countValue["pwdProtected"] = pwdProtected
	countValue["internetFilteringCount"] = internetFilteringCount
	countValue["timeManagementCount"] = timeManagementCount
	countValue["timeSlotCount"] = timeSlotCount
	countValue["advancedSettingsCount"] = advancedSettingsCount
	countValue["dailyPlaytimeCount"] = dailyPlaytimeCount
	countValue["maxSessionCount"] = maxSessionCount
	countValue["appManagementCount"] = appManagementCount
	countValue["authorizeUSBCount"] = authorizeUSBCount
	countValue["authorizeAdsCount"] = authorizeAdsCount
	countValue["changeProfileCount"] = changeProfileCount
	countValue["weblistCount"] = weblistCount
	countValue["appInstalledPreloadCount"] = appInstalledPreloadCount
	countValue["appInstalledGPlayCount"] = appInstalledGPlayCount
	countValue["faqCount"] = faqCount
	countValue["contactUsCount"] = contactUsCount
	countValue["userManualCount"] = userManualCount
	countValue["adultProfiles"] = adultProfiles
	countValue["childrenProfiles"] = childrenProfiles

	return countValue
}

func GetCountValues_v4(dmap *gorp.DbMap, queryWithSerialWhere, queryWithSerialAnd string) map[string]string {

	/**
	Name	:	GetCountValues_v4
	Version :	1.0
	Description :
				This function is dedicated to initialize information in tab of the analytics details page for TAB2 2015 series.

	State	: Used
	*/

	var (
		selectQuery        string
		totalCount         string
		boys               string
		girls              string
		pwdProtected       string
		appManagementCount string
		autoAuthorize      string
		authorizeUSBCount  string
		changeProfileCount string
		timeControl        string
		weblistCount       string
		faqCount           string
		contactUsCount     string
		userManualCount    string
		adultProfiles      string
		childrenProfiles   string
	)
	rom_id_2015 := " and serial not like '%40' and serial not like '%60'"
	selectQuery = "select count(*) from Profile_analytics_v4 where deleted = 'false'" + rom_id_2015
	totalCount = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(*) from Profile_analytics_v4 where deleted = 'false' and is_boy = 'true'" + rom_id_2015
	boys = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(*) from Profile_analytics_v4 where deleted = 'false' and is_boy = 'false'" + rom_id_2015
	girls = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(*) from Profile_analytics_v4 where deleted = 'false' and child_protected_by_pwd = 'true'" + rom_id_2015
	pwdProtected = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(app_management_count) from Serial_analytics_v4 where (app_management_count != '0' and app_management_count != '')" + rom_id_2015
	appManagementCount = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(auto_authorise_status) from Profile_analytics_v4 where auto_authorise_status = 'true' and deleted = 'false'" + rom_id_2015
	autoAuthorize = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(allow_usb_status) from Profile_analytics_v4 where allow_usb_status = 'true' and deleted = 'false'" + rom_id_2015
	authorizeUSBCount = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(allow_usb_status) from Profile_analytics_v4 where allow_google_status = 'true' and deleted = 'false'" + rom_id_2015
	authorizeUSBCount = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(profile_changed) from Profile_analytics_v4 where profile_changed = 'true' and deleted = 'false'" + rom_id_2015
	changeProfileCount = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(time_control_status) from Profile_analytics_v4 where time_control_status = 'true' and deleted = 'false'" + rom_id_2015
	timeControl = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(is_web_list_activated) from Profile_analytics_v4 where is_web_list_activated = 'true' and deleted = 'false'" + rom_id_2015
	weblistCount = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(faq_count) from Serial_analytics_v4 where faq_count > '0' and faq_count != ''" + rom_id_2015
	faqCount = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(contact_us_count) from Serial_analytics_v4 where contact_us_count > '0' and contact_us_count != ''" + rom_id_2015
	contactUsCount = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(user_manual_count) from Serial_analytics_v4 where user_manual_count > '0' and user_manual_count != ''" + rom_id_2015
	userManualCount = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "SELECT count(active_profiles) as count FROM KurioS_.Serial_analytics_v4 where (active_profiles - children_profiles) > 1" + rom_id_2015
	adultProfiles = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "SELECT count(children_profiles) as count FROM KurioS_.Serial_analytics_v4 where children_profiles > 1" + rom_id_2015
	childrenProfiles = GetCount(dmap, queryWithSerialAnd, selectQuery)

	countValue := make(map[string]string)

	countValue["totalCount"] = totalCount
	countValue["boys"] = boys
	countValue["girls"] = girls
	countValue["pwdProtected"] = pwdProtected
	countValue["appManagementCount"] = appManagementCount
	countValue["autoAuthorize"] = autoAuthorize
	countValue["authorizeUSBCount"] = authorizeUSBCount
	countValue["changeProfileCount"] = changeProfileCount
	countValue["timeControl"] = timeControl
	countValue["weblistCount"] = weblistCount
	countValue["faqCount"] = faqCount
	countValue["contactUsCount"] = contactUsCount
	countValue["userManualCount"] = userManualCount
	countValue["adultProfiles"] = adultProfiles
	countValue["childrenProfiles"] = childrenProfiles

	return countValue
}

func GetCountValues_v8(dmap *gorp.DbMap, queryWithSerialWhere, queryWithSerialAnd string) map[string]string {

	/**
	Name	:	GetCountValues_v8
	Version :	1.0
	Description :
				This function is dedicated to initialize information in tab of the analytics details page for TAB2 2016 series.

	State	: Used
	*/

	var (
		selectQuery        string
		totalCount         string
		boys               string
		girls              string
		pwdProtected       string
		appManagementCount string
		autoAuthorize      string
		authorizeUSBCount  string
		changeProfileCount string
		timeControl        string
		weblistCount       string
		faqCount           string
		contactUsCount     string
		userManualCount    string
		adultProfiles      string
		childrenProfiles   string
	)
	rom_id_2016 := " and serial like '%60'"
	selectQuery = "select count(*) from Profile_analytics_v4 where deleted = 'false'" + rom_id_2016
	totalCount = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(*) from Profile_analytics_v4 where deleted = 'false' and is_boy = 'true'" + rom_id_2016
	boys = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(*) from Profile_analytics_v4 where deleted = 'false' and is_boy = 'false'" + rom_id_2016
	girls = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(*) from Profile_analytics_v4 where deleted = 'false' and child_protected_by_pwd = 'true'" + rom_id_2016
	pwdProtected = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(app_management_count) from Serial_analytics_v4 where (app_management_count != '0' and app_management_count != '')" + rom_id_2016
	appManagementCount = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(auto_authorise_status) from Profile_analytics_v4 where auto_authorise_status = 'true' and deleted = 'false'" + rom_id_2016
	autoAuthorize = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(allow_usb_status) from Profile_analytics_v4 where allow_usb_status = 'true' and deleted = 'false'" + rom_id_2016
	authorizeUSBCount = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(profile_changed) from Profile_analytics_v4 where profile_changed = 'true' and deleted = 'false'" + rom_id_2016
	changeProfileCount = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(time_control_status) from Profile_analytics_v4 where time_control_status = 'true' and deleted = 'false'" + rom_id_2016
	timeControl = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(is_web_list_activated) from Profile_analytics_v4 where is_web_list_activated = 'true' and deleted = 'false'" + rom_id_2016
	weblistCount = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(faq_count) from Serial_analytics_v4 where faq_count > '0' and faq_count != ''" + rom_id_2016
	faqCount = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(contact_us_count) from Serial_analytics_v4 where contact_us_count > '0' and contact_us_count != ''" + rom_id_2016
	contactUsCount = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(user_manual_count) from Serial_analytics_v4 where user_manual_count > '0' and user_manual_count != ''" + rom_id_2016
	userManualCount = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "SELECT count(active_profiles) as count FROM KurioS_.Serial_analytics_v4 where (active_profiles - children_profiles) > 1" + rom_id_2016
	adultProfiles = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "SELECT count(children_profiles) as count FROM KurioS_.Serial_analytics_v4 where children_profiles > 1" + rom_id_2016
	childrenProfiles = GetCount(dmap, queryWithSerialAnd, selectQuery)

	countValue := make(map[string]string)

	countValue["totalCount"] = totalCount
	countValue["boys"] = boys
	countValue["girls"] = girls
	countValue["pwdProtected"] = pwdProtected
	countValue["appManagementCount"] = appManagementCount
	countValue["autoAuthorize"] = autoAuthorize
	countValue["authorizeUSBCount"] = authorizeUSBCount
	countValue["changeProfileCount"] = changeProfileCount
	countValue["timeControl"] = timeControl
	countValue["weblistCount"] = weblistCount
	countValue["faqCount"] = faqCount
	countValue["contactUsCount"] = contactUsCount
	countValue["userManualCount"] = userManualCount
	countValue["adultProfiles"] = adultProfiles
	countValue["childrenProfiles"] = childrenProfiles

	return countValue
}
func GetCountValues_v6(dmap *gorp.DbMap, queryWithSerialWhere, queryWithSerialAnd string) map[string]string {

	/**
	Name	:	GetCountValues_v6
	Version :	1.0
	Description :
				This function is dedicated to initialize information in tab of the analytics details page for TAB2 2016 series.

	State	: Used
	*/

	var (
		selectQuery        string
		totalCount         string
		boys               string
		girls              string
		pwdProtected       string
		appManagementCount string
		autoAuthorize      string
		authorizeUSBCount  string
		changeProfileCount string
		timeControl        string
		weblistCount       string
		faqCount           string
		contactUsCount     string
		userManualCount    string
		adultProfiles      string
		childrenProfiles   string
	)
	rom_id_2016 := " and serial like '%40' and serial not like '%340'"
	selectQuery = "select count(*) from Profile_analytics_v4 where deleted = 'false'" + rom_id_2016
	totalCount = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(*) from Profile_analytics_v4 where deleted = 'false' and is_boy = 'true'" + rom_id_2016
	boys = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(*) from Profile_analytics_v4 where deleted = 'false' and is_boy = 'false'" + rom_id_2016
	girls = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(*) from Profile_analytics_v4 where deleted = 'false' and child_protected_by_pwd = 'true'" + rom_id_2016
	pwdProtected = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(app_management_count) from Serial_analytics_v4 where (app_management_count != '0' and app_management_count != '')" + rom_id_2016
	appManagementCount = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(auto_authorise_status) from Profile_analytics_v4 where auto_authorise_status = 'true' and deleted = 'false'" + rom_id_2016
	autoAuthorize = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(allow_usb_status) from Profile_analytics_v4 where allow_usb_status = 'true' and deleted = 'false'" + rom_id_2016
	authorizeUSBCount = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(profile_changed) from Profile_analytics_v4 where profile_changed = 'true' and deleted = 'false'" + rom_id_2016
	changeProfileCount = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(time_control_status) from Profile_analytics_v4 where time_control_status = 'true' and deleted = 'false'" + rom_id_2016
	timeControl = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(is_web_list_activated) from Profile_analytics_v4 where is_web_list_activated = 'true' and deleted = 'false'" + rom_id_2016
	weblistCount = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(faq_count) from Serial_analytics_v4 where faq_count > '0' and faq_count != ''" + rom_id_2016
	faqCount = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(contact_us_count) from Serial_analytics_v4 where contact_us_count > '0' and contact_us_count != ''" + rom_id_2016
	contactUsCount = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "select count(user_manual_count) from Serial_analytics_v4 where user_manual_count > '0' and user_manual_count != ''" + rom_id_2016
	userManualCount = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "SELECT count(active_profiles) as count FROM KurioS_.Serial_analytics_v4 where (active_profiles - children_profiles) > 1" + rom_id_2016
	adultProfiles = GetCount(dmap, queryWithSerialAnd, selectQuery)

	selectQuery = "SELECT count(children_profiles) as count FROM KurioS_.Serial_analytics_v4 where children_profiles > 1" + rom_id_2016
	childrenProfiles = GetCount(dmap, queryWithSerialAnd, selectQuery)

	countValue := make(map[string]string)

	countValue["totalCount"] = totalCount
	countValue["boys"] = boys
	countValue["girls"] = girls
	countValue["pwdProtected"] = pwdProtected
	countValue["appManagementCount"] = appManagementCount
	countValue["autoAuthorize"] = autoAuthorize
	countValue["authorizeUSBCount"] = authorizeUSBCount
	countValue["changeProfileCount"] = changeProfileCount
	countValue["timeControl"] = timeControl
	countValue["weblistCount"] = weblistCount
	countValue["faqCount"] = faqCount
	countValue["contactUsCount"] = contactUsCount
	countValue["userManualCount"] = userManualCount
	countValue["adultProfiles"] = adultProfiles
	countValue["childrenProfiles"] = childrenProfiles

	return countValue
}

func GetCount(dmap *gorp.DbMap, queryWithSerial, selectQuery string) string {

	/**
	Name	:	GetCount
	Version :	1.0
	Description :
				This function help to make the queries used in count tab of analytyics details.

	State	: Used
	*/

	var count string

	if queryWithSerial != "" {
		selectQuery = selectQuery + queryWithSerial
	}
	dmap.SelectOne(&count, selectQuery)
	if count == "" {
		count = "0"
	}
	return count
}

func CreateGraphFromRequest(res http.ResponseWriter, req *http.Request, r render.Render, dmap *gorp.DbMap) {

	/**
	Name	:	CreateGraphFromRequest
	Version :	1.0
	Description :
				This function is used to make the tab for request made on Graph menu.

	State	: Used
	*/

	numberOfFields := req.FormValue("numberOfFields")
	graphType := req.FormValue("graphType")
	countLegendOne := req.FormValue("countLegendOne")
	countLegendTwo := req.FormValue("countLegendTwo")

	if numberOfFields != "" {
		query := req.FormValue("query")
		results := make(map[int]map[int]map[string]interface{})

		switch numberOfFields {

		case "one":
			var resultsFromDB []GenericWith2Field
			_, err := dmap.Select(&resultsFromDB, query)
			checkErr(err, "error select Serial_analytics/Serial_analytics_v4")
			for count, resultDB := range resultsFromDB {
				s := reflect.ValueOf(&resultDB).Elem()
				results[count] = GetValuesAndTypeNew(s, countLegendOne, countLegendTwo)
			}
			break

		case "two":
			var resultsFromDB []GenericWith3Field
			_, err := dmap.Select(&resultsFromDB, query)
			checkErr(err, "error select Serial_analytics/Serial_analytics_v4")
			for count, resultDB := range resultsFromDB {
				s := reflect.ValueOf(&resultDB).Elem()
				results[count] = GetValuesAndTypeNew(s, countLegendOne, countLegendTwo)
			}
			break
		}

		newmap := map[string]interface{}{"metatitle": "Extract Analytics Data", "results": results, "graphType": graphType}
		r.HTML(200, "analytics_viewer_graph", newmap)

	} else {
		newmap := map[string]interface{}{"metatitle": "Extract Analytics Data"}
		r.HTML(200, "analytics_viewer_graph_query", newmap)
	}

}

func GetValuesAndTypeNew(s reflect.Value, countLegendOne string, countLegendTwo string) map[int]map[string]interface{} {

	/**
	Name	:	GetValuesAndTypeNew
	Version :	1.0
	Description :
				This function is used to make the map used by the database link when query is performed .

	State	: Used
	*/

	newMap := make(map[int]map[string]interface{})

	typeOfT := s.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)

		newMap[i] = make(map[string]interface{})

		if f.Type().Name() == "int64" {
			newMap[i]["type"] = "number"
		}

		if f.Type().Name() == "string" {
			newMap[i]["type"] = "string"
		}

		if i == 1 && countLegendOne != "" {
			newMap[i]["name"] = countLegendOne
		} else if i == 2 && countLegendTwo != "" {
			newMap[i]["name"] = countLegendTwo
		} else {
			newMap[i]["name"] = typeOfT.Field(i).Name
		}

		newMap[i]["value"] = f.Interface()

	}
	return newMap
}

func GetOTARepartition_v3(dmap *gorp.DbMap, queryWithSerial string) []GenericWith2Field {

	/**
	Name	:	GetOTARepartition_v3
	Version :	1.0
	Description :
				This function is used to make the query for OTA repartition on X series models .

	State	: Used
	*/

	query := "select ro_build_display_id as name, count(ro_build_display_id) as count from Serial_analytics" + queryWithSerial + " group by name"

	var resultsFromDB []GenericWith2Field
	_, err := dmap.Select(&resultsFromDB, query)
	checkErr(err, "error select GetOTARepartition")

	return resultsFromDB
}

func GetOTARepartition_v4(dmap *gorp.DbMap, queryWithSerial string) []GenericWith2Field {

	/**
	Name	:	GetOTARepartition_v4
	Version :	1.0
	Description :
				This function is used to make the query for OTA repartition on TAB 2 2015 Models.

	State	: Used
	*/

	var specificRomid string
	if queryWithSerial == "" {
		specificRomid = " where serial not like '%40' and serial not like '%60' "
	} else {
		specificRomid = " and serial not like '%40' and serial not like '%60' "
	}
	query := "select ro_build_display_id as name, count(ro_build_display_id) as count from Serial_analytics_v4 " + queryWithSerial + specificRomid + " group by name having count(*) >100"
	log.Println("on recupere l OTA")

	var resultsFromDB []GenericWith2Field
	_, err := dmap.Select(&resultsFromDB, query)
	checkErr(err, "error select GetOTARepartition")

	return resultsFromDB
}

func GetOTARepartition_v6(dmap *gorp.DbMap, queryWithSerial string) []GenericWith2Field {

	/**
	Name	:	GetOTARepartition_v6
	Version :	1.0
	Description :
				This function is used to make the query for OTA repartition on TAB 2 2016 Models.

	State	: Used
	*/

	var specificRomid string
	if queryWithSerial == "" {
		specificRomid = " where serial like '%40' and serial not like '%340' and serial not like '%60'"
	} else {
		specificRomid = " and serial like '%40' and serial not like '%340' and serial not like '%60'"
	}

	query := "select ro_build_display_id as name, count(ro_build_display_id) as count from Serial_analytics_v4" + queryWithSerial + specificRomid + " group by name"

	var resultsFromDB []GenericWith2Field
	_, err := dmap.Select(&resultsFromDB, query)
	checkErr(err, "error select GetOTARepartition")

	return resultsFromDB
}

func GetOTARepartition_v8(dmap *gorp.DbMap, queryWithSerial string) []GenericWith2Field {

	/**
	Name	:	GetOTARepartition_v8
	Version :	1.0
	Description :
				This function is used to make the query for OTA repartition on TAB 2 2016 Models.

	State	: Used
	*/

	var specificRomid string
	if queryWithSerial == "" {
		specificRomid = " where serial like '%60'"
	} else {
		specificRomid = " and serial like '%60'"
	}

	query := "select ro_build_display_id as name, count(ro_build_display_id) as count from Serial_analytics_v4" + queryWithSerial + specificRomid + " group by name"

	var resultsFromDB []GenericWith2Field
	_, err := dmap.Select(&resultsFromDB, query)
	checkErr(err, "error select GetOTARepartition")

	return resultsFromDB
}

func GetSystemVersion_v3(dmap *gorp.DbMap, queryWithSerial string) []GenericWith2Field {

	/**
	Name	:	GetSystemVersion_v3
	Version :	1.0
	Description :
				This function is used to make the query for System repartition on X Series Models.

	State	: Used
	*/

	query := "select kurio_system_version as name, count(kurio_system_version) as count from Serial_analytics" + queryWithSerial + " group by name order by name DESC"

	var resultsFromDB []GenericWith2Field
	_, err := dmap.Select(&resultsFromDB, query)
	checkErr(err, "error select GetSystemVersion")

	return resultsFromDB
}

func GetSystemVersion_v4(dmap *gorp.DbMap, queryWithSerial string) []GenericWith2Field {

	/**
	Name	:	GetSystemVersion_v4
	Version :	1.0
	Description :
				This function is used to make the query for System repartition on TAB 2 2015 Models.

	State	: Used
	*/

	var specificRomid string
	if queryWithSerial == "" {
		specificRomid = " where serial not like '%40' and serial not like '%60' and kurio_system_version like '%4.0%'"
	} else {
		specificRomid = " and serial not like '%40' and serial not like '%60' and kurio_system_version like '%4.0%'"
	}
	query := "select kurio_system_version as name, count(kurio_system_version) as count from Serial_analytics_v4" + queryWithSerial + specificRomid + " group by name having count(*) >100 order by name DESC"

	var resultsFromDB []GenericWith2Field
	_, err := dmap.Select(&resultsFromDB, query)
	checkErr(err, "error select GetSystemVersion")

	return resultsFromDB
}

func GetSystemVersion_v6(dmap *gorp.DbMap, queryWithSerial string) []GenericWith2Field {

	/**
	Name	:	GetSystemVersion_v6
	Version :	1.0
	Description :
				This function is used to make the query for System repartition on TAB 2 2016 Models.

	State	: Used
	*/

	var specificRomid string
	if queryWithSerial == "" {
		specificRomid = " where serial like '%40' and serial not like '%340' and serial not like '%60' and kurio_system_version like '%4.0%'"
	} else {
		specificRomid = " and serial like '%40' and serial not like '%340' and serial not like '%60' and kurio_system_version like '%4.0%'"
	}
	query := "select kurio_system_version as name, count(kurio_system_version) as count from Serial_analytics_v4" + queryWithSerial + specificRomid + " group by name order by name DESC"

	var resultsFromDB []GenericWith2Field
	_, err := dmap.Select(&resultsFromDB, query)
	checkErr(err, "error select GetSystemVersion")

	return resultsFromDB
}
func GetSystemVersion_v8(dmap *gorp.DbMap, queryWithSerial string) []GenericWith2Field {

	/**
	Name	:	GetSystemVersion_v6
	Version :	1.0
	Description :
				This function is used to make the query for System repartition on TAB 2 2016 Models.

	State	: Used
	*/

	var specificRomid string
	if queryWithSerial == "" {
		specificRomid = " where serial like '%60'  and kurio_system_version like '6.0%'"
	} else {
		specificRomid = " and serial like '%60'  and kurio_system_version like '6.0%'"
	}
	query := "select kurio_system_version as name, count(kurio_system_version) as count from Serial_analytics_v4" + queryWithSerial + specificRomid + " group by name order by name DESC"

	var resultsFromDB []GenericWith2Field
	_, err := dmap.Select(&resultsFromDB, query)
	checkErr(err, "error select GetSystemVersion")

	return resultsFromDB
}

func GetProfilesByAge_v3(dmap *gorp.DbMap, queryWithSerial string) []GenericWith2Field {

	/**
	Name	:	GetProfilesByAge_v3
	Version :	1.0
	Description :
				This function is used to make the query for profile by age on X Series Models.

	State	: Used
	*/

	query := "select (year(now()) - left(`Profile_analytics`.`birth`, 4)) AS `name`, count(left(`Profile_analytics`.`birth`, 4)) AS `count` from `Profile_analytics` where (year(now()) - left(`Profile_analytics`.`birth`, 4)) < 22" + queryWithSerial + " and deleted ='false' group by (year(now()) - left(`Profile_analytics`.`birth`, 4))"

	var resultsFromDB []GenericWith2Field
	_, err := dmap.Select(&resultsFromDB, query)
	checkErr(err, "error select GetProfilesByAge")

	return resultsFromDB
}

func GetProfilesByAge_v4(dmap *gorp.DbMap, queryWithSerial string) []GenericWith2Field {

	/**
	Name	:	GetProfilesByAge_v4
	Version :	1.0
	Description :
				This function is used to make the query for profile by age on Tab 2 2015 Models.

	State	: Used
	*/

	specificRomid := " and serial not like '%40' and serial not like '%60' "
	query := "select (year(now()) - left(`Profile_analytics_v4`.`birth`, 4)) AS `name`, count(left(`Profile_analytics_v4`.`birth`, 4)) AS `count` from `Profile_analytics_v4` where (year(now()) - left(`Profile_analytics_v4`.`birth`, 4)) < 22" + queryWithSerial + specificRomid + "and deleted ='false'  group by (year(now()) - left(`Profile_analytics_v4`.`birth`, 4))"

	var resultsFromDB []GenericWith2Field
	_, err := dmap.Select(&resultsFromDB, query)
	checkErr(err, "error select GetProfilesByAge")

	return resultsFromDB
}
func GetProfilesByAge_v6(dmap *gorp.DbMap, queryWithSerial string) []GenericWith2Field {

	/**
	Name	:	GetProfilesByAge_v6
	Version :	1.0
	Description :
				This function is used to make the query for profile by age on Tab 2 2016 Models.

	State	: Used
	*/

	specificRomid := " and serial like '%40' and serial not like '%340' and serial not like '%60'"
	query := "select (year(now()) - left(`Profile_analytics_v4`.`birth`, 4)) AS `name`, count(left(`Profile_analytics_v4`.`birth`, 4)) AS `count` from `Profile_analytics_v4` where (year(now()) - left(`Profile_analytics_v4`.`birth`, 4)) < 22" + queryWithSerial + specificRomid + " and deleted ='false' group by (year(now()) - left(`Profile_analytics_v4`.`birth`, 4))"

	var resultsFromDB []GenericWith2Field
	_, err := dmap.Select(&resultsFromDB, query)
	checkErr(err, "error select GetProfilesByAge")

	return resultsFromDB
}
func GetProfilesByAge_v8(dmap *gorp.DbMap, queryWithSerial string) []GenericWith2Field {

	/**
	Name	:	GetProfilesByAge_v8
	Version :	1.0
	Description :
				This function is used to make the query for profile by age on Tab 2 2016 Models.

	State	: Used
	*/

	specificRomid := " and serial like '%60'"
	query := "select (year(now()) - left(`Profile_analytics_v4`.`birth`, 4)) AS `name`, count(left(`Profile_analytics_v4`.`birth`, 4)) AS `count` from `Profile_analytics_v4` where (year(now()) - left(`Profile_analytics_v4`.`birth`, 4)) < 22" + queryWithSerial + specificRomid + "and deleted ='false' group by (year(now()) - left(`Profile_analytics_v4`.`birth`, 4))"

	var resultsFromDB []GenericWith2Field
	_, err := dmap.Select(&resultsFromDB, query)
	checkErr(err, "error select GetProfilesByAge")

	return resultsFromDB
}
func GetNumberOfChilds_v3(dmap *gorp.DbMap, queryWithSerial string) []GenericWith2Field {

	/**
	Name	:	GetNumberOfChilds_v3
	Version :	1.0
	Description :
				This function is used to make the query for Number of Childs on X Series Models.

	State	: Used
	*/

	query := "SELECT children_profiles as name, count(children_profiles) as count FROM KurioS_.Serial_analytics group by name"

	var resultsFromDB []GenericWith2Field
	_, err := dmap.Select(&resultsFromDB, query)
	checkErr(err, "error select GetNumberOfChilds")

	return resultsFromDB
}

func GetNumberOfChilds_v4(dmap *gorp.DbMap, queryWithSerial string) []GenericWith2Field {

	/**
	Name	:	GetNumberOfChilds_v4
	Version :	1.0
	Description :
				This function is used to make the query for Number of Childs on Tab 2 2015 Models.

	State	: Used
	*/

	var specificRomid string
	if queryWithSerial == "" {
		specificRomid = " where serial not like '%40' and serial not like '%60' "
	} else {
		specificRomid = " and serial not like '%40' and serial not like '%60' "
	}
	query := "SELECT children_profiles as name, count(children_profiles) as count FROM KurioS_.Serial_analytics_v4 " + queryWithSerial + specificRomid + "group by name"

	var resultsFromDB []GenericWith2Field
	_, err := dmap.Select(&resultsFromDB, query)
	checkErr(err, "error select GetNumberOfChilds")

	return resultsFromDB
}

func GetNumberOfChilds_v6(dmap *gorp.DbMap, queryWithSerial string) []GenericWith2Field {

	/**
	Name	:	GetNumberOfChilds_v6
	Version :	1.0
	Description :
				This function is used to make the query for Number of Childs on Tab 2 2016 Models.

	State	: Used
	*/

	var specificRomid string
	if queryWithSerial == "" {
		specificRomid = " where serial like '%40' and serial not like '%340' and serial not like '%60'"
	} else {
		specificRomid = " and serial like '%40' and serial not like '%340' and serial not like '%60'"
	}
	query := "SELECT children_profiles as name, count(children_profiles) as count FROM KurioS_.Serial_analytics_v4 " + queryWithSerial + specificRomid + "group by name"

	var resultsFromDB []GenericWith2Field
	_, err := dmap.Select(&resultsFromDB, query)
	checkErr(err, "error select GetNumberOfChilds")

	return resultsFromDB
}
func GetNumberOfChilds_v8(dmap *gorp.DbMap, queryWithSerial string) []GenericWith2Field {

	/**
	Name	:	GetNumberOfChilds_v8
	Version :	1.0
	Description :
				This function is used to make the query for Number of Childs on Tab 2 2016 Models.

	State	: Used
	*/

	var specificRomid string
	if queryWithSerial == "" {
		specificRomid = " where serial like '%60'"
	} else {
		specificRomid = " and serial like '%60'"
	}
	query := "SELECT children_profiles as name, count(children_profiles) as count FROM KurioS_.Serial_analytics_v4 " + queryWithSerial + specificRomid + "group by name"

	var resultsFromDB []GenericWith2Field
	_, err := dmap.Select(&resultsFromDB, query)
	checkErr(err, "error select GetNumberOfChilds")

	return resultsFromDB
}
