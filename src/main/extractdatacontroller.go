package main

import (
	"archive/zip"
	"fmt"
	"github.com/lioonel/gorp"
	"github.com/martini-contrib/render"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var file_zip = ""

func ExtractDataDefault(res http.ResponseWriter, req *http.Request, r render.Render) {

	/**
	Name	:	ExtractDataDefault
	Version :	1.0
	Description :
				This function is dedicated to manage several ROMID search extraction.

	State	: Used to do
	*/
	if (idsession == "Jeremy") || (idsession == "Caroline") {
		newmap := map[string]interface{}{"metatitle": "Extract Data"}
		r.HTML(200, "extract_data_default", newmap)
	}/* else {
		http.Redirect(res, req, "/analytics_bad_user", 302)
	}*/
}



func ExtractDataWeekly(res http.ResponseWriter, req *http.Request, r render.Render, dmap *gorp.DbMap) {

	/**
	Name	:	ExtractDataWeekly
	Version :	1.0
	Description :
				This function is dedicated to lauch the method which extract the weekly report for aftersales.

	State	: Used
	*/

	if idsession == "extract"  {

		newmap := map[string]interface{}{"metatitle": "Extract Data"}
		r.HTML(200, "extract_data_weekly", newmap)
	} else {
		http.Redirect(res, req, "/analytics_bad_user", 302)
	}

}

func ExtractDataWeeklySelected(res http.ResponseWriter, req *http.Request, r render.Render, dmap *gorp.DbMap) {

	/**
	Name	:	ExtractDataWeeklySelected
	Version :	1.0
	Description :
				This function is dedicated to extract the weekly report for aftersales

	State	: Used
	*/

	if idsession == "extract"  {

		// creation du fichier pour les US
		year, month, day := time.Now().Date()
		finalDate := ""
		mois := ""
		if month == time.Month(1) {
			mois = "1"
		} else if month == time.Month(2) {
			mois = "2"
		} else if month == time.Month(3) {
			mois = "3"
		} else if month == time.Month(4) {
			mois = "4"
		} else if month == time.Month(5) {
			mois = "5"
		} else if month == time.Month(6) {
			mois = "6"
		} else if month == time.Month(7) {
			mois = "7"
		} else if month == time.Month(8) {
			mois = "8"
		} else if month == time.Month(9) {
			mois = "9"
		} else if month == time.Month(10) {
			mois = "10"
		} else if month == time.Month(11) {
			mois = "11"
		} else if month == time.Month(12) {
			mois = "12"
		}
		finalDate = strconv.Itoa(year) + mois + strconv.Itoa(day)

		pathzipfile := root_directory + "Weekly_serial_" + req.FormValue("MODEL") + "_" + req.FormValue("COUNTRY") + "_" + ".zip"
		pwd, errpath := os.Getwd()
		checkErr(errpath, "errueur path os")
		//create zip file
		foZippedData, err := os.Create(pathzipfile)
		//	foZippedData, err := os.Create(pathzipfile)
		defer foZippedData.Close()
		checkErr(err, "erreur creation file")
		w := zip.NewWriter(foZippedData)

		checkErr(err, "erreur creation file")
		selectQueryAnalytics := ""
		withMail := "mail"

		if req.FormValue("MODEL") == "TAB2 2015" {
			if req.FormValue("COUNTRY") != "BE" {
				selectQueryAnalytics = "select * from Serial_analytics_v4"
				CreateQueryAnalyticsv4(&selectQueryAnalytics, req)
				CreateSerialData(w, dmap, selectQueryAnalytics, "Kurio_TAB2 serie_extraction_"+req.FormValue("COUNTRY")+"_"+finalDate+".xls", "")
			} else {
				log.Println("N/A")
			}

		} else if req.FormValue("MODEL") == "TAB2 2016" {

			selectQueryAnalytics = "select * from Serial_analytics_v4"
			CreateQueryAnalyticsv6(&selectQueryAnalytics, req)
			CreateSerialData(w, dmap, selectQueryAnalytics, "Kurio_TAB2 2016_serie_extraction_"+req.FormValue("COUNTRY")+"_"+finalDate+".xls", "")

		} else if req.FormValue("MODEL") == "TAB3 2017" {

			selectQueryAnalytics = "select * from Serial_analytics_v4"
			CreateQueryAnalyticsv8(&selectQueryAnalytics, req)
			CreateSerialData(w, dmap, selectQueryAnalytics, "Kurio_TAB3 2017_serie_extraction_"+req.FormValue("COUNTRY")+"_"+finalDate+".xls", "")

		} else if req.FormValue("MODEL") == "Serie X" {

			selectQueryAnalytics = "select * from Serial_analytics"
			CreateQueryAnalytics(&selectQueryAnalytics, req)
			CreateSerialData7X(w, dmap, selectQueryAnalytics, "Kurio_Xtreme serie_extraction_"+req.FormValue("COUNTRY")+"_"+finalDate+".xls", "")

		} else if req.FormValue("MODEL") == "Serie S" {
			if req.FormValue("COUNTRY") == "NL" {

				selectQueryAnalytics = "select * from Serial"
				CreateQueryNLSerial7S(&selectQueryAnalytics)
				CreateSerialDataS(w, dmap, selectQueryAnalytics, "Kurio_7S serie_extraction_NL "+finalDate+".xls", "")
				selectQueryAnalytics = "select * from Serial"
				CreateQueryNLSerial4S(&selectQueryAnalytics)
				CreateSerialDataS(w, dmap, selectQueryAnalytics, "Kurio_4S serie_extraction_NL "+finalDate+".xls", "")
				selectQueryAnalytics = "select * from Serial"
				CreateQueryNLSerial10S(&selectQueryAnalytics)
				CreateSerialDataS(w, dmap, selectQueryAnalytics, "Kurio_10S serie_extraction_NL "+finalDate+".xls", "")

			} else if req.FormValue("COUNTRY") == "DK" {

				selectQueryAnalytics = "select * from Serial"
				CreateQueryDKSerial7S(&selectQueryAnalytics)
				CreateSerialDataS(w, dmap, selectQueryAnalytics, "Kurio_7S serie_extraction_DK "+finalDate+".xls", "")
				selectQueryAnalytics = "select * from Serial"
				CreateQueryDKSerial4S(&selectQueryAnalytics)
				CreateSerialDataS(w, dmap, selectQueryAnalytics, "Kurio_4S serie_extraction_DK "+finalDate+".xls", "")
				selectQueryAnalytics = "select * from Serial"
				CreateQueryDKSerial10S(&selectQueryAnalytics)
				CreateSerialDataS(w, dmap, selectQueryAnalytics, "Kurio_10S serie_extraction_DK "+finalDate+".xls", "")
			} else if req.FormValue("COUNTRY") == "US" {

				selectQueryAnalytics = "select * from Serial"
				CreateQueryUSSerial4S(&selectQueryAnalytics)
				CreateSerialDataS(w, dmap, selectQueryAnalytics, "Kurio_4S serie_extraction_US "+finalDate+".xls", withMail)
				selectQueryAnalytics = "select * from Serial"
				CreateQueryUSSerial7S(&selectQueryAnalytics)
				CreateSerialDataS(w, dmap, selectQueryAnalytics, "Kurio_7S serie_extraction_US "+finalDate+".xls", withMail)
				selectQueryAnalytics = "select * from Serial"
				CreateQueryUSSerial10S(&selectQueryAnalytics)
				CreateSerialDataS(w, dmap, selectQueryAnalytics, "Kurio_10S serie_extraction_US "+finalDate+".xls", withMail)
			} else if req.FormValue("COUNTRY") == "DE" {

				selectQueryAnalytics = "select * from Serial"
				CreateQueryDESerial7S(&selectQueryAnalytics)
				CreateSerialDataS(w, dmap, selectQueryAnalytics, "Kurio_7S serie_extraction_DE "+finalDate+".xls", "")
				selectQueryAnalytics = "select * from Serial"
				CreateQueryDESerial10S(&selectQueryAnalytics)
				CreateSerialDataS(w, dmap, selectQueryAnalytics, "Kurio_10S serie_extraction_DE "+finalDate+".xls", "")
			} else if req.FormValue("COUNTRY") == "FR" {

				selectQueryAnalytics = "select * from Serial"
				CreateQueryDESerial7S(&selectQueryAnalytics)
				CreateSerialDataS(w, dmap, selectQueryAnalytics, "Kurio_7S serie_extraction_DE "+finalDate+".xls", "")
				selectQueryAnalytics = "select * from Serial"
				CreateQueryDESerial10S(&selectQueryAnalytics)
				CreateSerialDataS(w, dmap, selectQueryAnalytics, "Kurio_10S serie_extraction_DE "+finalDate+".xls", "")

			} else if req.FormValue("COUNTRY") == "ES" {

				selectQueryAnalytics = "select * from Serial"
				CreateQueryDESerial7S(&selectQueryAnalytics)
				CreateSerialDataS(w, dmap, selectQueryAnalytics, "Kurio_7S serie_extraction_DE "+finalDate+".xls", "")
				selectQueryAnalytics = "select * from Serial"
				CreateQueryDESerial10S(&selectQueryAnalytics)
				CreateSerialDataS(w, dmap, selectQueryAnalytics, "Kurio_10S serie_extraction_DE "+finalDate+".xls", "")

			} else if req.FormValue("COUNTRY") == "UK" {

				selectQueryAnalytics = "select * from Serial"
				CreateQueryDESerial7S(&selectQueryAnalytics)
				CreateSerialDataS(w, dmap, selectQueryAnalytics, "Kurio_7S serie_extraction_DE "+finalDate+".xls", "")
				selectQueryAnalytics = "select * from Serial"
				CreateQueryDESerial10S(&selectQueryAnalytics)
				CreateSerialDataS(w, dmap, selectQueryAnalytics, "Kurio_10S serie_extraction_DE "+finalDate+".xls", "")

			} else if req.FormValue("COUNTRY") == "LATAM" {

				selectQueryAnalytics = "select * from Serial"
				CreateQueryDESerial7S(&selectQueryAnalytics)
				CreateSerialDataS(w, dmap, selectQueryAnalytics, "Kurio_7S serie_extraction_DE "+finalDate+".xls", "")
				selectQueryAnalytics = "select * from Serial"
				CreateQueryDESerial10S(&selectQueryAnalytics)
				CreateSerialDataS(w, dmap, selectQueryAnalytics, "Kurio_10S serie_extraction_DE "+finalDate+".xls", "")

			} else if req.FormValue("COUNTRY") == "EE" {

				selectQueryAnalytics = "select * from Serial"
				CreateQueryDESerial7S(&selectQueryAnalytics)
				CreateSerialDataS(w, dmap, selectQueryAnalytics, "Kurio_7S serie_extraction_DE "+finalDate+".xls", "")
				selectQueryAnalytics = "select * from Serial"
				CreateQueryDESerial10S(&selectQueryAnalytics)
				CreateSerialDataS(w, dmap, selectQueryAnalytics, "Kurio_10S serie_extraction_DE "+finalDate+".xls", "")

			} else if req.FormValue("COUNTRY") == "ZA" {

				selectQueryAnalytics = "select * from Serial"
				CreateQueryDESerial7S(&selectQueryAnalytics)
				CreateSerialDataS(w, dmap, selectQueryAnalytics, "Kurio_7S serie_extraction_DE "+finalDate+".xls", "")
				selectQueryAnalytics = "select * from Serial"
				CreateQueryDESerial10S(&selectQueryAnalytics)
				CreateSerialDataS(w, dmap, selectQueryAnalytics, "Kurio_10S serie_extraction_DE "+finalDate+".xls", "")

			}
		} else if req.FormValue("MODEL") == "Old Tab" {
			/* to finalize */
		}

		w.Close()
		res.Header().Set("Content-Description", "File Transfer")
		res.Header().Set("Content-Type", "application/zip")
		res.Header().Set("Content-Transfer-Encoding", "binary")
		res.Header().Set("Content-Disposition", "attachment; filename=Weekly_serial_"+req.FormValue("MODEL")+"_"+req.FormValue("COUNTRY")+"_"+".zip")
		log.Println("Data Created")
		http.ServeFile(res, req, root_directory+"Weekly_serial_"+req.FormValue("MODEL")+"_"+req.FormValue("COUNTRY")+"_"+".zip")
		log.Println("Path is  : " + pwd)
		_, err = io.Copy(foZippedData, req.Body)
		checkErr(err, "erreur copy")
	} else {
		http.Redirect(res, req, "/analytics_bad_user", 302)
	}
}

func ExtractDataWeeklyChoice(res http.ResponseWriter, req *http.Request, r render.Render, dmap *gorp.DbMap) {

	/**
	Name	:	ExtractDataWeeklyChoice
	Version :	1.0
	Description :
				This function is dedicated to manage the choice of submit on "extract_data_weekly.tmpl" file

	State	: Used
	*/

	Choice := req.FormValue("Choice")
	if idsession == "extract"  {
		if Choice == "Download Country Selected" {
			log.Println("Download Country Selected")
			ExtractDataWeeklySelected(res, req, r, dmap)
		} else {
			ExtractDataWeeklyAll(res, req, r, dmap)
		}

	} else {
		http.Redirect(res, req, "/analytics_bad_user", 302)
	}
}

func ExtractDataPremiumChoice(res http.ResponseWriter, req *http.Request, r render.Render, dmap *gorp.DbMap) {

	/**
	Name	:	ExtractDataWeeklyChoice
	Version :	1.0
	Description :
				This function is dedicated to manage the choice of submit on "extract_data_weekly.tmpl" file

	State	: Used
	*/

	Choice := req.FormValue("Choice")
	if idsession == "extract" {
		if Choice == "Download Country Selected" {
			log.Println("Download Country Selected")
			ExtractDataWeeklySelected(res, req, r, dmap)
		} else {
			ExtractDataWeeklyAll(res, req, r, dmap)
		}

	} else {
		http.Redirect(res, req, "/analytics_bad_user", 302)
	}
}


func ExtractDataSchemaChoice(res http.ResponseWriter, req *http.Request, r render.Render, dmap *gorp.DbMap) {

	/**
	Name	:	ExtractDataSchemaChoice
	Version :	1.0
	Description :
				This function is dedicated to manage the choice of dowload on Analytics_viewer.tmpl file

	State	: Used
	*/

	Choice := req.FormValue("Choice")
	log.Println(Choice)
	if (idsession == "cidesei3") || (idsession == "Jeremy") || (idsession == "Caroline") {
		if Choice == "Download Device Activation to CSV" {
			ExtractAnalyticsDashboardPieCharts(res, req, r, dmap)
		} else if Choice == "Download Activity to CSV" {
			ExtractAnalyticsDashboardTableCharts(res, req, r, dmap)
		} else if Choice == "Download install count by app to CSV" {
			ExtractAnalyticsDashboardInstall(res, req, r, dmap)
		} else if Choice == "Content delivery U40 Stat" {
			ExtractAnalyticsContentDelivery(res, req, r, dmap)
		}

	} else {
		http.Redirect(res, req, "/analytics_bad_user", 302)
	}
}

func ExtractDataOlivierChoice(res http.ResponseWriter, req *http.Request, r render.Render, dmap *gorp.DbMap) {

	/**
	Name	:	ExtractDataOlivierChoice
	Version :	1.0
	Description :
				This function is dedicated to manage the choice of dowload on Analytics_viewer.tmpl file

	State	: Used
	*/

	Choice := req.FormValue("Choice")
	if idsession == "extract"{
		if Choice == "Download" {
			ExtractDataOlivierSpec(res, req, r, dmap)
		}
	} else {
		http.Redirect(res, req, "/analytics_bad_user", 302)
	}
}

func ExtractDataOlivierSpec(res http.ResponseWriter, req *http.Request, r render.Render, dmap *gorp.DbMap) {

	/**
	Name	:	ExtractDataOlivierSpec
	Version :	1.0
	Description :
				This function is dedicated to perform the request on the database to extract the device's corresponding data.

	State	: Used
	*/

	curefSearch := req.FormValue("curefToSearch")
	serialSearch := req.FormValue("serialToSearch")
	emailSearch := req.FormValue("emailToSearch")
	macAdressSearch := req.FormValue("macAdressToSearch")
	tableToSearch := req.FormValue("tabToSearch")
	if idsession == "extract" {
		// gestion de la date

		year, month, day := time.Now().Date()
		finalDate := ""
		mois := ""
		if month == time.Month(1) {
			mois = "1"
		} else if month == time.Month(2) {
			mois = "2"
		} else if month == time.Month(3) {
			mois = "3"
		} else if month == time.Month(4) {
			mois = "4"
		} else if month == time.Month(5) {
			mois = "5"
		} else if month == time.Month(6) {
			mois = "6"
		} else if month == time.Month(7) {
			mois = "7"
		} else if month == time.Month(8) {
			mois = "8"
		} else if month == time.Month(9) {
			mois = "9"
		} else if month == time.Month(10) {
			mois = "10"
		} else if month == time.Month(11) {
			mois = "11"
		} else if month == time.Month(12) {
			mois = "12"
		}
		log.Println("mois fait")
		finalDate = strconv.Itoa(year) + mois + strconv.Itoa(day)

		pathzipfile := root_directory + "Weekly_serial_" + tableToSearch + ".zip"
		pwd, errpath := os.Getwd()
		checkErr(errpath, "errueur path os")
		//create zip file
		foZippedData, err := os.Create(pathzipfile)
		defer foZippedData.Close()
		checkErr(err, "erreur creation file")
		w := zip.NewWriter(foZippedData)

		checkErr(err, "erreur creation file")
		selectQueryAnalytics := ""

		if curefSearch != "" || serialSearch != "" || emailSearch != "" || macAdressSearch != "" {

			if tableToSearch == "register" {
				dmap2 := InitOldDb()
				selectQueryAnalytics = "select * from register"
				CreateQueryextractregister(&selectQueryAnalytics, serialSearch, emailSearch, macAdressSearch)
				CreateSerialData(w, dmap2, selectQueryAnalytics, tableToSearch+finalDate+".xls", "")

			} else if tableToSearch == "Serial" {

				selectQueryAnalytics = "select * from Serial"
				CreateQueryextractSerial(&selectQueryAnalytics, serialSearch, emailSearch, macAdressSearch)
				CreateSerialDataS(w, dmap, selectQueryAnalytics, tableToSearch+finalDate+".xls", "")

			} else if tableToSearch == "Serial_analytics" {

				selectQueryAnalytics = "select * from Serial_analytics"
				CreateQueryextractSerial_analytics(&selectQueryAnalytics, curefSearch, serialSearch, emailSearch, macAdressSearch)
				CreateSerialData7X(w, dmap, selectQueryAnalytics, tableToSearch+finalDate+".xls", "")

			} else if tableToSearch == "Serial_analytics_v4" {

				selectQueryAnalytics = "select * from Serial_analytics_v4"
				CreateQueryextractSerial_analytics_v4(&selectQueryAnalytics, curefSearch, serialSearch, emailSearch, macAdressSearch)
				CreateSerialData(w, dmap, selectQueryAnalytics, tableToSearch+finalDate+".xls", "")

			}
			log.Println("requete faite")

		} else {
			log.Println("il s est rien passe")
		}

		w.Close()
		res.Header().Set("Content-Description", "File Transfer")
		res.Header().Set("Content-Type", "application/zip")
		res.Header().Set("Content-Transfer-Encoding", "binary")
		res.Header().Set("Content-Disposition", "attachment; filename=Weekly_serial_"+req.FormValue("MODEL")+"_"+req.FormValue("COUNTRY")+"_"+".zip")
		log.Println("Data Created")
		http.ServeFile(res, req, root_directory+"Weekly_serial_"+req.FormValue("MODEL")+"_"+req.FormValue("COUNTRY")+"_"+".zip")
		log.Println("Path is  : " + pwd)

		_, err = io.Copy(foZippedData, req.Body)
		checkErr(err, "erreur copy")
	} else {
		http.Redirect(res, req, "/analytics_bad_user", 302)
	}
}

func ExtractDataWeeklyAll(res http.ResponseWriter, req *http.Request, r render.Render, dmap *gorp.DbMap) {

	/**
	Name	:	ExtractDataWeeklyAll
	Version :	1.0
	Description :
				This function is dedicated to perform the request on the database to extract the weekly data for all aftersales.

	State	: Used
	*/

	if idsession == "extract" {

		// creation du fichier pour les US
		year, month, day := time.Now().Date()
		finalDate := ""
		mois := ""
		if month == time.Month(1) {
			mois = "1"
		} else if month == time.Month(2) {
			mois = "2"
		} else if month == time.Month(3) {
			mois = "3"
		} else if month == time.Month(4) {
			mois = "4"
		} else if month == time.Month(5) {
			mois = "5"
		} else if month == time.Month(6) {
			mois = "6"
		} else if month == time.Month(7) {
			mois = "7"
		} else if month == time.Month(8) {
			mois = "8"
		} else if month == time.Month(9) {
			mois = "9"
		} else if month == time.Month(10) {
			mois = "10"
		} else if month == time.Month(11) {
			mois = "11"
		} else if month == time.Month(12) {
			mois = "12"
		}
		finalDate = strconv.Itoa(year) + mois + strconv.Itoa(day)

		pathzipfile := root_directory + "Weekly_serial_TAB2.zip"
		pwd, errpath := os.Getwd()
		checkErr(errpath, "errueur path os")
		//create zip file
		foZippedData, err := os.Create(pathzipfile)
		defer foZippedData.Close()
		checkErr(err, "erreur creation file")
		w := zip.NewWriter(foZippedData)

		checkErr(err, "erreur creation file")

		withMail := "mail"

		selectQueryAnalytics := "select * from Serial_analytics_v4 where serial not like '%40'"
		CreateSerialData(w, dmap, selectQueryAnalytics, "Kurio_Xtreme 2015 serie_extraction_ALL "+finalDate+".xls", "")
		selectQueryAnalytics = "select * from Serial_analytics_v4 where serial  like '%40'"
		CreateSerialData(w, dmap, selectQueryAnalytics, "Kurio_Xtreme 2016 serie_extraction_ALL "+finalDate+".xls", "")

		selectQueryAnalytics = "select * from Serial_analytics where serial like '%126'"
		CreateSerialData7X(w, dmap, selectQueryAnalytics, "Kurio_X serie_extraction_NL "+finalDate+".xls", "")
		selectQueryAnalytics = "select * from Serial_analytics where serial like '%135'"
		CreateSerialData7X(w, dmap, selectQueryAnalytics, "Kurio_X serie_extraction_DE "+finalDate+".xls", "")
		selectQueryAnalytics = "select * from Serial_analytics where serial like '%105'"
		CreateSerialData7X(w, dmap, selectQueryAnalytics, "Kurio_X serie_extraction_DK "+finalDate+".xls", "")
		selectQueryAnalytics = "select * from Serial_analytics where serial like '%102'"
		CreateSerialData7X(w, dmap, selectQueryAnalytics, "Kurio_X serie_extraction_US "+finalDate+".xls", withMail)

		selectQueryAnalytics = "select * from Serial"
		CreateQueryNLSerial7S(&selectQueryAnalytics)
		CreateSerialDataS(w, dmap, selectQueryAnalytics, "Kurio_7S serie_extraction_NL "+finalDate+".xls", "")
		selectQueryAnalytics = "select * from Serial"
		CreateQueryDESerial7S(&selectQueryAnalytics)
		CreateSerialDataS(w, dmap, selectQueryAnalytics, "Kurio_7S serie_extraction_DE "+finalDate+".xls", "")
		selectQueryAnalytics = "select * from Serial"
		CreateQueryDKSerial7S(&selectQueryAnalytics)
		CreateSerialDataS(w, dmap, selectQueryAnalytics, "Kurio_7S serie_extraction_DK "+finalDate+".xls", "")
		selectQueryAnalytics = "select * from Serial"
		CreateQueryUSSerial7S(&selectQueryAnalytics)
		CreateSerialDataS(w, dmap, selectQueryAnalytics, "Kurio_7S serie_extraction_US "+finalDate+".xls", withMail)

		selectQueryAnalytics = "select * from Serial"
		CreateQueryNLSerial4S(&selectQueryAnalytics)
		CreateSerialDataS(w, dmap, selectQueryAnalytics, "Kurio_4S serie_extraction_NL "+finalDate+".xls", "")
		selectQueryAnalytics = "select * from Serial"
		CreateQueryDKSerial4S(&selectQueryAnalytics)
		CreateSerialDataS(w, dmap, selectQueryAnalytics, "Kurio_4S serie_extraction_DK "+finalDate+".xls", "")
		selectQueryAnalytics = "select * from Serial"
		CreateQueryUSSerial4S(&selectQueryAnalytics)
		CreateSerialDataS(w, dmap, selectQueryAnalytics, "Kurio_4S serie_extraction_US "+finalDate+".xls", withMail)

		selectQueryAnalytics = "select * from Serial"
		CreateQueryNLSerial10S(&selectQueryAnalytics)
		CreateSerialDataS(w, dmap, selectQueryAnalytics, "Kurio_10S serie_extraction_NL "+finalDate+".xls", "")
		selectQueryAnalytics = "select * from Serial"
		CreateQueryDESerial10S(&selectQueryAnalytics)
		CreateSerialDataS(w, dmap, selectQueryAnalytics, "Kurio_10S serie_extraction_DE "+finalDate+".xls", "")
		selectQueryAnalytics = "select * from Serial"
		CreateQueryDKSerial10S(&selectQueryAnalytics)
		CreateSerialDataS(w, dmap, selectQueryAnalytics, "Kurio_10S serie_extraction_DK "+finalDate+".xls", "")
		selectQueryAnalytics = "select * from Serial"
		CreateQueryUSSerial10S(&selectQueryAnalytics)
		CreateSerialDataS(w, dmap, selectQueryAnalytics, "Kurio_10S serie_extraction_US "+finalDate+".xls", withMail)

		w.Close()
		res.Header().Set("Content-Description", "File Transfer")
		res.Header().Set("Content-Type", "application/zip")
		res.Header().Set("Content-Transfer-Encoding", "binary")
		res.Header().Set("Content-Disposition", "attachment; filename=Weekly_serial_TAB2.zip")
		log.Println("Data Created")
		http.ServeFile(res, req, root_directory+"/Weekly_serial_TAB2.zip")
		log.Println("Path is  : " + pwd)
		_, err = io.Copy(foZippedData, req.Body)
		checkErr(err, "erreur copy")
	} else {
		http.Redirect(res, req, "/analytics_bad_user", 302)
	}

}
func ExtractDataJulien(res http.ResponseWriter, req *http.Request, r render.Render) {

	/**
	Name	:	ExtractDataOlivier
	Version :	1.0
	Description :
				This function is dedicated to display the "ExtractDataOlivier.tmpl" file which manage the data search on each Kurio Devices.

	State	: Used
	*/

	if idsession == "extract" {
		newmap := map[string]interface{}{"metatitle": "Extract Data"}
		r.HTML(200, "extract_data_julien", newmap)
	} else {
		http.Redirect(res, req, "/analytics_bad_user", 302)
	}

}
func ExtractSerialDownload(res http.ResponseWriter, req *http.Request, r render.Render, dmap *gorp.DbMap) {

	/**
	  Name	:	ExtractSerialDownload
	  Version :	1.0
	  Description :
	  			This function is dedicated to download the serial data for a specific TAB.

	  State	: On Use
	*/

	pathzipfile := root_directory + "Data_serial_type.zip"
	tabChoice := req.FormValue("tabToSearch")
	withMail := ""

	count := 1
	var serialArray [1]string
	serialArray[0] = req.FormValue("serialToSearch")
	var serialToQuery = make([]string, count)

	value := serialArray[0]
	serialToQuery[0] = value
	log.Println("cest le numéro de série " + serialArray[0])
	pwd, errpath := os.Getwd()
	checkErr(errpath, "errueur path os")

	//create zip file
	foZippedData, err := os.Create(pathzipfile)
	defer foZippedData.Close()
	checkErr(err, "erreur creation file")
	w := zip.NewWriter(foZippedData)

	checkErr(err, "erreur creation file")

	if tabChoice == "Serial" {
		selectQuery := "select * from Serial"
		CreateQueryWithSerial(&selectQuery, serialToQuery)
		CreateSerialData(w, dmap, selectQuery, "serial7S.xls", withMail)
	} else if tabChoice == "Serial_analytics" {
		selectQueryAnalytics := "select * from Serial_analytics"
		CreateQueryWithSerialAnalytics(&selectQueryAnalytics, serialToQuery)
		CreateSerialData(w, dmap, selectQueryAnalytics, "serial7X.xls", withMail)
	} else if tabChoice == "Serial_analytics_v4" {
		selectQueryAnalytics := "select * from Serial_analytics_v4"
		CreateQueryWithSerialAnalytics(&selectQueryAnalytics, serialToQuery)
		CreateSerialData(w, dmap, selectQueryAnalytics, "serialTab2.xls", withMail)
	} else {
		log.Println("Fonction non dévellopée")
	}
	w.Close()
	res.Header().Set("Content-Description", "File Transfer")
	res.Header().Set("Content-Type", "application/zip")
	res.Header().Set("Content-Transfer-Encoding", "binary")
	res.Header().Set("Content-Disposition", "attachment; filename=Data_serial_type.zip")
	log.Println("Data Created")
	http.ServeFile(res, req, root_directory+"Data_serial_type.zip")
	log.Println("Path is  : " + pwd)
	_, err = io.Copy(foZippedData, req.Body)
	checkErr(err, "erreur copy")
}

func ExtractDataDownload(res http.ResponseWriter, req *http.Request, r render.Render, dmap *gorp.DbMap) {

	/**
	  Name	:	ExtractDataDownload
	  Version :	1.0
	  Description :
	  			This function is dedicated to download the serial data for a all tab for a specific

	  State	: On Use
	*/

	pathzipfile := root_directory + "Data_serial_type.zip"

	var serialArray [10]string
	serialArray[0] = "serial1"
	serialArray[1] = "serial2"
	serialArray[2] = "serial3"
	serialArray[3] = "serial4"
	serialArray[4] = "serial5"
	serialArray[5] = "serial6"
	serialArray[6] = "serial7"
	serialArray[7] = "serial8"
	serialArray[8] = "serial9"
	serialArray[9] = "serial10"

	tabChoice := req.FormValue("tabChoice")
	withMail := req.FormValue("mail")

	count := 0

	for i := 0; i < len(serialArray); i++ {
		value := req.FormValue(serialArray[i])
		if value != "" {
			count++
		} else {
			break
		}
	}

	var serialToQuery = make([]string, count)

	for i := 0; i < len(serialArray); i++ {
		value := req.FormValue(serialArray[i])
		if value != "" {
			for i := len(value); i < 3; i++ {
				value = strconv.Itoa(0) + value
			}
			serialToQuery[i] = value
			fmt.Println("value : ", value)
		} else {
			break
		}
	}

	pwd, errpath := os.Getwd()
	checkErr(errpath, "errueur path os")

	//create zip file
	foZippedData, err := os.Create(pathzipfile)
	defer foZippedData.Close()
	checkErr(err, "erreur creation file")
	w := zip.NewWriter(foZippedData)

	checkErr(err, "erreur creation file")

	if tabChoice == "7S" {
		selectQuery := "select * from Serial"
		CreateQueryWithRomId(&selectQuery, serialToQuery)
		CreateSerialData(w, dmap, selectQuery, "serial7S.xls", withMail)
	} else if tabChoice == "7X" {
		selectQueryAnalytics := "select * from Serial_analytics"
		CreateQueryWithRomIdAnalytics(&selectQueryAnalytics, serialToQuery)
		CreateSerialData(w, dmap, selectQueryAnalytics, "serial7X.xls", withMail)
	} else {
		selectQueryAnalytics := "select * from Serial_analytics_v4"
		CreateQueryWithRomIdAnalytics(&selectQueryAnalytics, serialToQuery)
		CreateSerialData(w, dmap, selectQueryAnalytics, "serialTab2.xls", withMail)

	}
	w.Close()
	res.Header().Set("Content-Description", "File Transfer")
	res.Header().Set("Content-Type", "application/zip")
	res.Header().Set("Content-Transfer-Encoding", "binary")
	res.Header().Set("Content-Disposition", "attachment; filename=Data_serial_type.zip")
	log.Println("Data Created")
	http.ServeFile(res, req, root_directory+"Data_serial_type.zip")
	log.Println("Path is  : " + pwd)

	_, err = io.Copy(foZippedData, req.Body)
	checkErr(err, "erreur copy")
}

func CreateSerialData7X(w *zip.Writer, dmapExtract *gorp.DbMap, selectQuery string, fileName string, withMail string) {

	/**
	  Name	:	CreateSerialData7X
	  Version :	1.0
	  Description :
	  			This function is dedicated to retrieve the data for X series  for a specific query.

	  State	: On Use
	*/

	var serials []Serial

	_, err := dmapExtract.Select(&serials, selectQuery)
	checkErr(err, "select erreur ")
	fmt.Println("serials size ", len(serials))

	f, _ := w.Create(fileName)

	header := "serial	"
	header = header + "RoBuildDisplayId	"
	header = header + "RoBuildVersionIncremental	"
	header = header + "RoProductModel	"
	header = header + "RoProductName	"
	header = header + "RoProductDevice	"
	header = header + "MACAddress	"
	header = header + "LastIp	"
	header = header + "SystemLanguage	"
	header = header + "ServerEntree	"
	header = header + "FirstActivation	"
	header = header + "SetupDate	"
	if withMail != "" {
		header = header + "LastModification	"
		header = header + "ParentEmail	\n"
	} else {
		header = header + "LastModification	\n"
	}

	f.Write([]byte(header))

	for _, serial := range serials {
		dataOut := serial.Serial + "	"
		dataOut = dataOut + serial.RoBuildDisplayId + "	"
		dataOut = dataOut + serial.RoBuildVersionIncremental + "	"
		dataOut = dataOut + serial.RoProductModel + "	"
		dataOut = dataOut + serial.RoProductName + "	"
		dataOut = dataOut + serial.RoProductDevice + "	"
		dataOut = dataOut + serial.MACAddress + "	"
		dataOut = dataOut + serial.LastIp + "	"
		dataOut = dataOut + serial.SystemLanguage + "	"

		serverEntree := serial.ServerEntree.Time.Format("02/01/2006")
		dataOut = dataOut + serverEntree + "	"
		dataOut = dataOut + serial.FirstActivation.Time.Format("02/01/2006") + "	"
		dataOut = dataOut + serial.SetupDate.Time.Format("02/01/2006") + "	"
		if withMail != "" {
			dataOut = dataOut + serial.LastModification.Time.Format("02/01/2006") + "	"
			dataOut = dataOut + serial.ParentEmail + "	" + "\n"
		} else {
			dataOut = dataOut + serial.LastModification.Time.Format("02/01/2006") + "	\n"
		}
		f.Write([]byte(dataOut))
	}
}

func CreateSerialData(w *zip.Writer, dmapExtract *gorp.DbMap, selectQuery string, fileName string, withMail string) {

	/**
	  Name	:	CreateSerialData
	  Version :	1.0
	  Description :
	  			This function is dedicated to retrieve the data for Tab 2  for a specific query.

	  State	: On Use
	*/

	var serials []Serial

	_, err := dmapExtract.Select(&serials, selectQuery)
	checkErr(err, "select erreur ")
	fmt.Println("serials size ", len(serials))

	f, _ := w.Create(fileName)

	header := "serial	"
	header = header + "curef	"
	header = header + "KurioSystemVersion	"
	header = header + "RoBuildDisplayId	"
	header = header + "RoBuildVersionIncremental	"
	header = header + "RoProductModel	"
	header = header + "RoProductName	"
	header = header + "RoProductDevice	"
	header = header + "MACAddress	"
	header = header + "LastIp	"
	header = header + "SystemLanguage	"
	header = header + "ServerEntree	"
	header = header + "FirstActivation	"
	header = header + "SetupDate	"
	header = header + "LastModification	"
	header = header + "ParentEmail	\n"
	
	f.Write([]byte(header))

	for _, serial := range serials {
		dataOut := serial.Serial + "	"
		dataOut = dataOut + serial.Curef + "	"
		dataOut = dataOut + serial.KurioSystemVersion.String + "	"
		dataOut = dataOut + serial.RoBuildDisplayId + "	"
		dataOut = dataOut + serial.RoBuildVersionIncremental + "	"
		dataOut = dataOut + serial.RoProductModel + "	"
		dataOut = dataOut + serial.RoProductName + "	"
		dataOut = dataOut + serial.RoProductDevice + "	"
		dataOut = dataOut + serial.MACAddress + "	"
		dataOut = dataOut + serial.LastIp + "	"
		dataOut = dataOut + serial.SystemLanguage + "	"
		serverEntree := serial.ServerEntree.Time.Format("02/01/2006")
		dataOut = dataOut + serverEntree + "	"
		dataOut = dataOut + serial.FirstActivation.Time.Format("02/01/2006") + "	"
		dataOut = dataOut + serial.SetupDate.Time.Format("02/01/2006") + "	"
		dataOut = dataOut + serial.LastModification.Time.Format("02/01/2006") + "	"
		dataOut = dataOut + serial.ParentEmail + "	" + "\n"
		f.Write([]byte(dataOut))
	}
}

func CreateSerialDataS(w *zip.Writer, dmapExtract *gorp.DbMap, selectQuery string, fileName string, withMail string) {

	/**
	  Name	:	CreateSerialDataS
	  Version :	1.0
	  Description :
	  			This function is dedicated to retrieve the data for S series  for a specific query.

	  State	: On Use
	*/

	var serials []Serial

	_, err := dmapExtract.Select(&serials, selectQuery)
	checkErr(err, "select erreur ")
	fmt.Println("serials size ", len(serials))

	f, _ := w.Create(fileName)

	header := "serial	"
	header = header + "RoBuildDisplayId	"
	header = header + "RoBuildVersionIncremental	"
	header = header + "RoProductModel	"
	header = header + "RoProductName	"
	header = header + "RoProductDevice	"
	header = header + "MACAddress	"
	header = header + "LastIp	"
	header = header + "SystemLanguage	"
	header = header + "ServerEntree	"
	header = header + "FirstActivation	"
	header = header + "SetupDate	"
	header = header + "LastModification	"
	header = header + "ParentEmail	\n"
	
	f.Write([]byte(header))

	for _, serial := range serials {
		dataOut := serial.Serial + "	"
		dataOut = dataOut + serial.RoBuildDisplayId + "	"
		dataOut = dataOut + serial.RoBuildVersionIncremental + "	"
		dataOut = dataOut + serial.RoProductModel + "	"
		dataOut = dataOut + serial.RoProductName + "	"
		dataOut = dataOut + serial.RoProductDevice + "	"
		dataOut = dataOut + serial.MACAddress + "	"
		dataOut = dataOut + serial.LastIp + "	"
		dataOut = dataOut + serial.SystemLanguage + "	"
		serverEntree := serial.ServerEntree.Time.Format("02/01/2006")
		dataOut = dataOut + serverEntree + "	"
		dataOut = dataOut + serial.FirstActivation.Time.Format("02/01/2006") + "	"
		dataOut = dataOut + serial.SetupDate.Time.Format("02/01/2006") + "	"
		dataOut = dataOut + serial.LastModification.Time.Format("02/01/2006") + "	"
		dataOut = dataOut + serial.ParentEmail + "	" + "\n"
		f.Write([]byte(dataOut))
	}
}

//called first to select rom id to extract
func ExtractAllData(res http.ResponseWriter, req *http.Request, r render.Render) {

	/**
	  Name	:	ExtractAllData
	  Version :	1.0
	  Description :
	  			This function is dedicated to display the extract_all_data.tmpl page .

	  State	: On Use
	*/

	if idsession == "Jeremy" || (idsession == "Caroline") {
		password := req.FormValue("password")
		newmap := map[string]interface{}{"metatitle": "Extract All Data", "password": password}
		r.HTML(200, "extract_all_data", newmap)
	} else {
		http.Redirect(res, req, "/analytics_bad_user", 302)
	}

}

//create a xls file and download data
func ExtractAllDataDownload(res http.ResponseWriter, req *http.Request, r render.Render, dmap *gorp.DbMap) {

	/**
	  Name	:	ExtractAllDataDownload
	  Version :	1.0
	  Description :
	  			This function is dedicated to retrieve the serial and profile data for X series  .

	  State	: On Use
	*/

	pathZipFile := root_directory + "allData.zip"

	var serialArray [10]string
	serialArray[0] = "serial1"
	serialArray[1] = "serial2"
	serialArray[2] = "serial3"
	serialArray[3] = "serial4"
	serialArray[4] = "serial5"
	serialArray[5] = "serial6"
	serialArray[6] = "serial7"
	serialArray[7] = "serial8"
	serialArray[8] = "serial9"
	serialArray[9] = "serial10"

	count := 0
	//count number of field fill
	for i := 0; i < len(serialArray); i++ {
		value := req.FormValue(serialArray[i])
		if value != "" {
			count++
		} else {
			break
		}
	}

	var serialToQuery = make([]string, count)

	for i := 0; i < len(serialArray); i++ {
		value := req.FormValue(serialArray[i])
		if value != "" {
			serialToQuery[i] = value
		} else {
			break
		}
	}

	//create zip file
	foZippedData, err := os.Create(pathZipFile)
	defer foZippedData.Close()
	checkErr(err, "erreur creation file")

	w := zip.NewWriter(foZippedData)

	/***** SErial analytics ******/
	CreateAllSerialData(w, dmap, serialToQuery)
	CreateAllProfileData(w, dmap, serialToQuery)

	/*** Creation downloading file ***/
	w.Close()

	res.Header().Set("Content-Description", "File Transfer")
	res.Header().Set("Content-Type", "application/zip")
	res.Header().Set("Content-Transfer-Encoding", "binary")
	res.Header().Set("Content-Disposition", "attachment; filename=allData.zip")
	http.ServeFile(res, req, "/opt/goLang/src/kurioanalytics"+pathZipFile)
	_, err = io.Copy(foZippedData, req.Body)
	checkErr(err, "erreur copy")

}

func CreateAllSerialData(w *zip.Writer, dmapExtract *gorp.DbMap, serialToQuery []string) {

	/**
	  Name	:	CreateAllSerialData
	  Version :	1.0
	  Description :
	  			This function is dedicated to make the query for the serial data for X series  .

	  State	: On Use
	*/

	var serials []Serial
	selectQuery := "select * from Serial"
	CreateQueryWithRomId(&selectQuery, serialToQuery)

	_, err := dmapExtract.Select(&serials, selectQuery)
	checkErr(err, "select erreur analytics ")
	f, _ := w.Create("dataSerial.xls")

	header := "serial	"
	header = header + "DateProd	"
	header = header + "RoBuildDisplayId	"
	header = header + "RoBuildVersionIncremental	"
	header = header + "RoProductModel	"
	header = header + "RoProductName	"
	header = header + "RoProductDevice	"
	header = header + "RoProductLocaleLanguage	"
	header = header + "RoProductLocaleRegion	"
	header = header + "KurioSystemVersion	"
	header = header + "MACAddress	"
	header = header + "LastIp	"
	header = header + "SystemLanguage	"
	header = header + "ServerEntree	"
	header = header + "FirstActivation	"
	header = header + "SetupDate	"
	header = header + "LastModification	"
	header = header + "ParentEmail	"
	header = header + "MailActive	"
	header = header + "OwnerLockscreenCode	"
	header = header + "ActiveProfiles	"
	header = header + "ChildrenProfiles	"
	header = header + "PwdReset	\n"
	f.Write([]byte(header))

	for _, serial := range serials {
		dataOut := serial.Serial + "	"
		dataOut = dataOut + serial.DateProd.Time.Format("2006/01/02") + "	"
		dataOut = dataOut + serial.RoBuildDisplayId + "	"
		dataOut = dataOut + serial.RoBuildVersionIncremental + "	"
		dataOut = dataOut + serial.RoProductModel + "	"
		dataOut = dataOut + serial.RoProductName + "	"
		dataOut = dataOut + serial.RoProductDevice + "	"
		dataOut = dataOut + serial.RoProductLocaleLanguage.String + "	"
		dataOut = dataOut + serial.RoProductLocaleRegion.String + "	"
		dataOut = dataOut + serial.KurioSystemVersion.String + "	"
		dataOut = dataOut + serial.MACAddress + "	"
		dataOut = dataOut + serial.LastIp + "	"
		dataOut = dataOut + serial.SystemLanguage + "	"
		dataOut = dataOut + serial.ServerEntree.Time.Format("2006/01/02") + "	"
		dataOut = dataOut + serial.FirstActivation.Time.Format("2006/01/02") + "	"
		dataOut = dataOut + serial.SetupDate.Time.Format("2006/01/02") + "	"
		dataOut = dataOut + serial.LastModification.Time.Format("2006/01/02") + "	"
		dataOut = dataOut + serial.ParentEmail + "	"
		dataOut = dataOut + serial.MailActive + "	"
		dataOut = dataOut + serial.OwnerLockscreenCode.String + "	"
		dataOut = dataOut + strconv.FormatInt(serial.ActiveProfiles.Int64, 10) + "	"
		dataOut = dataOut + strconv.FormatInt(serial.ChildrenProfiles.Int64, 10) + "	"
		dataOut = dataOut + serial.PwdReset + "	\n"
		f.Write([]byte(dataOut))
	}
}

func CreateAllProfileData(w *zip.Writer, dmapExtract *gorp.DbMap, serialToQuery []string) {

	/**
	  Name	:	CreateAllProfileData
	  Version :	1.0
	  Description :
	  			This function is dedicated to make the query for the profile data for X series  .

	  State	: On Use
	*/

	var profiles []Profile
	selectQuery := "select * from Profile"
	CreateQueryWithRomId(&selectQuery, serialToQuery)

	_, err := dmapExtract.Select(&profiles, selectQuery)
	checkErr(err, "select erreur analytics ")
	f, _ := w.Create("dataProfile.xls")

	header := "serial	"
	header = header + "IsBoy	"
	header = header + "Birth	"
	header = header + "InternetAccessMode	"
	header = header + "TimeConstrolStatus	"
	header = header + "AllowUSBStatus	"
	header = header + "AuthorizeAdsStatus	"
	header = header + "AutoAuthoriseStatus	\n"
	f.Write([]byte(header))

	for _, profile := range profiles {
		dataOut := profile.Serial + "	"
		dataOut = dataOut + profile.IsBoy + "	"
		dataOut = dataOut + profile.Birth.Time.Format("2006/01/02") + "	"
		dataOut = dataOut + profile.InternetAccessMode + "	"
		dataOut = dataOut + profile.TimeConstrolStatus + "	"
		dataOut = dataOut + profile.AllowUSBStatus + "	"
		dataOut = dataOut + profile.AuthorizeAdsStatus + "	"
		dataOut = dataOut + profile.AutoAuthoriseStatus + "	" + "	\n"
		f.Write([]byte(dataOut))
	}
}

func CreateQueryWithRomId(selectQuery *string, serialToQuery []string) {

	/**
	  Name	:	CreateQueryWithRomId
	  Version :	1.0
	  Description :
	  			This function is dedicated to make the query for a specific Serial number .

	  State	: On Use
	*/

	for countSerial, serialValue := range serialToQuery {
		if serialValue != "" {
			if countSerial == 0 {
				*selectQuery = *selectQuery + " where serial  like '%" + serialValue + "_'"
			} else {
				*selectQuery = *selectQuery + " serial  like '%" + serialValue + "_'"
			}

			if countSerial != len(serialToQuery)-1 {
				*selectQuery = *selectQuery + " OR "
			}
		}
	}
}
func CreateQueryWithSerial(selectQuery *string, serialToQuery []string) {

	/**
	  Name	:	CreateQueryWithSerial
	  Version :	1.0
	  Description :
	  			This function is dedicated to make the query for a specific Serial number .

	  State	: On Use
	*/

	for countSerial, serialValue := range serialToQuery {
		if serialValue != "" {
			if countSerial == 0 {
				*selectQuery = *selectQuery + " where serial  like '%" + serialValue + "%'"
			} else {
				*selectQuery = *selectQuery + " serial  like '%" + serialValue + "%'"
			}

			if countSerial != len(serialToQuery)-1 {
				*selectQuery = *selectQuery + " OR "
			}
		}
	}
}
func CreateQueryWithSeveralRomId(selectQuery *string, serialToQuery []string) {

	/**
	  Name	:	CreateQueryWithSeveralRomId
	  Version :	1.0
	  Description :
	  			This function is dedicated to make the query for a specific range of Serial number .

	  State	: On Use
	*/

	for countSerial, serialValue := range serialToQuery {
		if serialValue != "" {
			if countSerial == 0 {
				*selectQuery = *selectQuery + " where serial  like '%" + serialValue + "_'"
			} else {
				*selectQuery = *selectQuery + " serial  like '%" + serialValue + "_'"
			}

			if countSerial != len(serialToQuery)-1 {
				*selectQuery = *selectQuery + " OR "
			}
		}
	}
}

func CreateQueryWithCuref(selectQuery *string, CurefToQuery []string) {

	/**
	  Name	:	CreateQueryWithCuref
	  Version :	1.0
	  Description :
	  			This function is dedicated to make the query for a specific range of Curef .

	  State	: On Use
	*/

	for countSerial, curefValue := range CurefToQuery {
		if curefValue != "" {
			if countSerial == 0 {
				*selectQuery = *selectQuery + " where Curef  like '%" + curefValue + "%" + "_'"
			} else {
				*selectQuery = *selectQuery + " Curef  like '%" + curefValue + "%" + "_'"
			}

			if countSerial != len(CurefToQuery)-1 {
				*selectQuery = *selectQuery + " OR "
			}
		}
	}
}

func CreateQueryWithCountryRegister(selectQuery *string, CountryToQuery []string) {

	/**
	  Name	:	CreateQueryWithCountryRegister
	  Version :	1.0
	  Description :
	  			This function is dedicated to make the query for a specific range of build_.

	  State	: On Use
	*/

	for countCountry, CountryValue := range CountryToQuery {
		if CountryValue != "" {
			if countCountry == 0 {
				*selectQuery = *selectQuery + " where build_ like '%" + CountryValue + "_'"
			} else {
				*selectQuery = *selectQuery + " build_like '%" + CountryValue + "_'"
			}

			if countCountry != len(CountryToQuery)-1 {
				*selectQuery = *selectQuery + " OR "
			}
		}
	}
}
