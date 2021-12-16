package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"html/template"
	"log"
	"runtime"
	/*"net/http"
	"github.com/martini-contrib/secure"*///	"runtime"
)

var m = martini.Classic()
//var myip string = thelocalip()
var myip string = "192.168.7.50"
var passwdKurioS string = getKurioSPassword()
var passwdRegister string = getRegisterpassword()



func main() {
	log.Println(passwdRegister)
	log.Println("voici mon ip : " + myip)
	/**
	  Name	:	Main
	  Version :	1.0
	  Description :
	  			This function is the main program which manage the rest of the app.

	  State	: No more used that is all
	*/

	//  Martini is a powerful package for quickly writing modular web applications/services in Golang.

	//	Operating System Verification to check where the goland app is installed
	if runtime.GOOS == "windows" {
		root_directory = "C:/tmp/"
	} else {
		root_directory = "/tmp/"
	}
	log.Println(root_directory)

	// WEB Framework Initialization

	//	m := martini.Classic()
	kurioSession := sessions.Sessions("my_session", sessions.NewCookieStore([]byte("secret123")))
	m.Use(kurioSession)
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
		Funcs: []template.FuncMap{
			{
				"lenminus": func(a map[int]map[string]interface{}) int {
					size := len(a)
					return size - 1
				},
			},
		},
	}))

	m.Use(martini.Static("assets"))

	m.Use(CheckSession)

	// This is the place where the framework is linked to the MYSQL database
	dbmap := initDb()
	defer dbmap.Db.Close()
	m.Map(dbmap)

	//used to extract ALL data from Serial table (pour Herve)
	m.Post("/extract_all_data_download", ExtractAllDataDownload)
	m.Post("/extract_all_data", ExtractAllData)

	//used to extract data from Serial table (pour Olivier)
	m.Post("/extract_data_download", ExtractDataDownload)
	m.Post("/extract_data_default", ExtractDataDefault)
	m.Post("/extract_data_julien", ExtractDataJulien)
	m.Post("/extract_data_weekly", ExtractDataWeekly)
	m.Post("/extract_data_weekly_choice", ExtractDataWeeklyChoice)
	m.Post("/extract_data_premium_choice", ExtractDataPremiumChoice)
	m.Post("/extract_data_olivier_choice", ExtractDataOlivierChoice)
	m.Post("/extract_data_weekly_all", ExtractDataWeeklyAll)
	m.Post("/extract_data_weekly_selected", ExtractDataWeeklySelected)
	m.Post("/extract_data_olivier_spec", ExtractDataOlivierSpec)

	//used for analytics extraction
	m.Post("/extract_analytics_data_download", ExtractAnalyticsDataDownload)
	m.Post("/extract_analytics_basic_data_download", ExtractAnalyticsBasicDataDownload)
	m.Post("/extract_analytics_Dashboard_charts", ExtractAnalyticsDashboardCharts)
	m.Post("/extract_data_schema_choice", ExtractDataSchemaChoice)
	m.Post("/analytics_viewer_details", AnalyticsViewerDetails)

	//use for the user authentification
	m.Post("/login", AnalyticsViewerLogin)
	m.Get("/", AnalyticsViewerConnect)
	m.Get("/analytics_bad_user", AnalyticsBadUser)
	m.Get("/analytics_bad_password", AnalyticsBadPassword)
	m.Get("/login", AnalyticsViewerConnect)

	//used to launch the page for the analytics for QA
	m.Get("/analytics_viewer_julien", AnalyticsViewerJulien)
	m.Get("/analytics_viewer_jeremie", AnalyticsViewerjeremie)

	//used to launch the page for the analytics for Other users
	m.Get("/analytics_viewer", AnalyticsViewer)
	m.Get("/Application_details", Application_details)
	m.Get("/analytics_viewer_choice", AnalyticsViewerChoice)
	m.Get("/analytics_viewer_details", AnalyticsViewerDetails)
	
	/* fonction pour l'extraction hebdomadaire des données pour chaque pays */

	m.Get("/analytics_api_get_device/version/(?P<version>[0-9]+)", AnalyticsGetRomIdStatus)
	m.Get("/analytics_api_get_profiles/version/(?P<version>[0-9]+)", AnalyticsGetProfiles)
	m.Get("/analytics_api_get_activation/version/(?P<version>[0-9])/serial/(?P<serial>[0-9A-Za-z]+)", AnalyticsGetActivation)
	m.Get("/analytics_api_get_activity/version/(?P<version>[0-9])/serial/(?P<serial>[0-9A-Za-z]+)", AnalyticsGetActivityStatus)
	m.Get("/analytics_api_get_appstatic/version/(?P<version>[0-9])/serial/(?P<serial>[0-9A-Za-z]+)/type/(?P<type>[A-Za-z_]+)", AnalyticsGetAppStatic)
	m.Post("/extract_serial_download", ExtractSerialDownload)

	//	m.Get("/analytics_api_get_app_by_age_sex", AnalyticsGetAppByAgeSex) // this one doesn't work an
	m.Get("/analytics_api_search_from_serial_mail_mac", AnalyticsSearchFromSerialMailMac)
	m.Get("/analytics_api_search_from_package_name", AnalyticsSearchFromPackageName)
	m.Get("/analytics_api_search_from_package_name_date", AnalyticsSearchFromPackageNameDate)
	
	m.Get("/analytics_viewer_graph", CreateGraphFromRequest)
	m.Post("/analytics_viewer_graph", CreateGraphFromRequest)
	activateDBTrace(dbmap)
	m.Run()

}
