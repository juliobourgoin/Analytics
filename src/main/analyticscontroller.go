package main

import (
	"archive/zip"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"github.com/go-martini/martini"
	"github.com/guregu/null/zero"
	"github.com/lioonel/gorp"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func ExtractAnalyticsData(res http.ResponseWriter, req *http.Request, r render.Render) {

	/**
	Name	:	ExtractAnalyticsData
	Version :	1.0
	Description :
				This function is used to launch the html page which help to extract all information for each Romod. .

	State	: Used
	*/

	newmap := map[string]interface{}{"metatitle": "Analytics Application"}
	r.HTML(200, "extract_analytics_data", newmap)
}

//create a xls file and download data
func ExtractAnalyticsDataDownload(res http.ResponseWriter, req *http.Request, r render.Render, dmap *gorp.DbMap) {

	/**
	Name	:	ExtractAnalyticsDataDownload
	Version :	1.0
	Description :
				This function is used to to extract all analytics information for tab 2 with each Romid on the list.

	State	: Used
	*/
	pathZipFile := root_directory + "/analyticsData.zip"

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
	CreateSerialAnalyticsFile(w, dmap, serialToQuery)

	/***** Profile analytics ******/
	CreateProfileAnalyticsFile(w, dmap, serialToQuery)

	CreateActivityStatusAnalyticsFile(w, dmap, serialToQuery)

	/*** Creation downloading file ***/
	w.Close()
	res.Header().Set("Content-Description", "File Transfer")
	res.Header().Set("Content-Type", "application/zip")
	res.Header().Set("Content-Transfer-Encoding", "binary")
	res.Header().Set("Content-Disposition", "attachment; filename=dataAnalytics.zip")
	http.ServeFile(res, req, pathZipFile)
	_, err = io.Copy(foZippedData, req.Body)
	checkErr(err, "erreur copy")

}

//create a xls file and download data
func ExtractAnalyticsBasicDataDownload(res http.ResponseWriter, req *http.Request, r render.Render, dmap *gorp.DbMap) {
	/**
	Name	:	ExtractAnalyticsBasicDataDownload
	Version :	1.0
	Description :
				This function is used to to extract all analytics information of X Series for each Romid on the list.

	State	: Used
	*/
	pathZipFile := root_directory + "/analyticsBasicData.zip"

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
	CreateSerialAnalyticsFile(w, dmap, serialToQuery)

	/***** Profile analytics ******/
	CreateProfileAnalyticsFile(w, dmap, serialToQuery)

	CreateActivityStatusAnalyticsBasicFile(w, dmap, serialToQuery)

	/*** Creation downloading file ***/
	w.Close()
	res.Header().Set("Content-Description", "File Transfer")
	res.Header().Set("Content-Type", "application/zip")
	res.Header().Set("Content-Transfer-Encoding", "binary")
	res.Header().Set("Content-Disposition", "attachment; filename=analyticsBasicData.zip")
	http.ServeFile(res, req, pathZipFile)
	_, err = io.Copy(foZippedData, req.Body)
	checkErr(err, "erreur copy")

}

func ExtractAnalyticsDashboardPieCharts(res http.ResponseWriter, req *http.Request, r render.Render, dmap *gorp.DbMap) {

	/**
	Name	:	ExtractAnalyticsDashboardPieCharts
	Version :	1.0
	Description :
				This function is used to to extract all tablet listed on the pie chart for a version of device.

	State	: Used
	*/
	log.Println("ExtractAnalyticsDashboardPieCharts")

	// creation de la date
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

	annee := ""

	if req.FormValue("version1") == "4" {
		annee = "2015"
	} else if req.FormValue("version1") == "5" {
		annee = "2016"
	} else {
		annee = "2014"
	}

	pathZipFile := root_directory + "/Device_Activation_" + annee + "_" + finalDate + "_" + ".zip"
	//create zip file
	foZippedData, err := os.Create(pathZipFile)
	defer foZippedData.Close()
	checkErr(err, "erreur creation file")

	w := zip.NewWriter(foZippedData)
	/***** SErial analytics ******/
	if req.FormValue("version1") == "4" {
		CreateSerialPieCharFilev4(w, dmap)
		log.Println("cest la version " + annee)
	} else if req.FormValue("version1") == "5" {
		CreateSerialPieCharFilev5(w, dmap)
		log.Println("cest la version " + annee)
	} else {
		CreateSerialPieCharFilev3(w, dmap)
		log.Println("cest la version " + annee)
	}

	/*** Creation downloading file ***/
	w.Close()

	res.Header().Set("Content-Description", "File Transfer")
	res.Header().Set("Content-Type", "application/zip")
	res.Header().Set("Content-Transfer-Encoding", "binary")
	res.Header().Set("Content-Disposition", "attachment; filename=Device_Activation_"+"_"+finalDate+"_"+annee+".zip")
	http.ServeFile(res, req, pathZipFile)
	_, err = io.Copy(foZippedData, req.Body)
	checkErr(err, "erreur copy")

}

func ExtractAnalyticsContentDelivery(res http.ResponseWriter, req *http.Request, r render.Render, dmap *gorp.DbMap) {

	/**
	Name	:	ExtractAnalyticsContentDelivery
	Version :	1.0
	Description :
				This function is used to to extract all information of the content delivery.

	State	: Used
	*/
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
	// creation de la date
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

	pathZipFile := root_directory + "/U40stat2016" + "_" + finalDate + ".zip"
	log.Println(serialArray[0])
	//count number of field fill
	log.Println(len(serialArray))
	for i := 0; i < len(serialArray); i++ {
		value := req.FormValue(serialArray[i])
		log.Println(value)
		if value != "" {
			count++
		} else {
			break
		}
	}

	var serialToQuery = make([]string, count)

	for i := 0; i < len(serialArray); i++ {
		value := req.FormValue(serialArray[i])
		log.Println(value)
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

	CreateContendeliveryTable(w, dmap)

	/*** Creation downloading file ***/
	w.Close()

	res.Header().Set("Content-Description", "File Transfer")
	res.Header().Set("Content-Type", "application/zip")
	res.Header().Set("Content-Transfer-Encoding", "binary")
	res.Header().Set("Content-Disposition", "attachment; filename=U40stat2016"+"_"+finalDate+".zip")
	http.ServeFile(res, req, pathZipFile)
	_, err = io.Copy(foZippedData, req.Body)
	checkErr(err, "erreur copy")

}

func ExtractAnalyticsDashboardInstall(res http.ResponseWriter, req *http.Request, r render.Render, dmap *gorp.DbMap) {

	/**
	Name	:	ExtractAnalyticsDashboardInstall
	Version :	1.0
	Description :
				This function is used to to extract all information of install for each application for a specific version of tablet and
				for a specific country.

	State	: Used
	*/

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
	// creation de la date
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

	pathZipFile := root_directory + "/ActivityStatusData" + "_" + finalDate + ".zip"
	log.Println(serialArray[0])
	//count number of field fill
	log.Println(len(serialArray))
	for i := 0; i < len(serialArray); i++ {
		value := req.FormValue(serialArray[i])
		log.Println(value)
		if value != "" {
			count++
		} else {
			break
		}
	}

	var serialToQuery = make([]string, count)

	for i := 0; i < len(serialArray); i++ {
		value := req.FormValue(serialArray[i])
		log.Println(value)
		if value != "" {
			serialToQuery[i] = value
		} else {
			break
		}
	}

	annee := ""

	if req.FormValue("version1") == "4" {
		annee = "2015"
	} else if req.FormValue("version1") == "5" {
		annee = "2016"
	} else {
		annee = "2014"
	}

	//create zip file
	foZippedData, err := os.Create(pathZipFile)
	defer foZippedData.Close()
	checkErr(err, "erreur creation file")

	w := zip.NewWriter(foZippedData)

	/***** SErial analytics ******/
	if req.FormValue("version1") == "4" {
		CreateInstallTableFilev4(w, req.FormValue("serial1"), dmap)
		log.Println("cest la version " + annee)
	} else if req.FormValue("version1") == "5" {
		CreateInstallTableFilev5(w, req.FormValue("serial1"), dmap)
		log.Println("cest la version " + annee)
	} else {
		CreateInstallTableFilev3(w, req.FormValue("serial1"), dmap)
		log.Println("cest la version " + annee)
	}

	/*** Creation downloading file ***/
	w.Close()

	res.Header().Set("Content-Description", "File Transfer")
	res.Header().Set("Content-Type", "application/zip")
	res.Header().Set("Content-Transfer-Encoding", "binary")
	res.Header().Set("Content-Disposition", "attachment; filename=ActivityStatusData"+"_"+finalDate+".zip")
	http.ServeFile(res, req, pathZipFile)
	_, err = io.Copy(foZippedData, req.Body)
	checkErr(err, "erreur copy")

}

func ExtractAnalyticsDashboardTableCharts(res http.ResponseWriter, req *http.Request, r render.Render, dmap *gorp.DbMap) {

	/**
	Name	:	ExtractAnalyticsDashboardTableCharts
	Version :	1.0
	Description :
				This function is used to to extract all information of launch time  for each application for a specific version of tablet and
				for a specific country.

	State	: Used
	*/

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
	// creation de la date
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

	pathZipFile := root_directory + "/ActivityStatusData" + "_" + finalDate + ".zip"
	log.Println(serialArray[0])
	//count number of field fill
	log.Println(len(serialArray))
	for i := 0; i < len(serialArray); i++ {
		value := req.FormValue(serialArray[i])
		log.Println(value)
		if value != "" {
			count++
		} else {
			break
		}
	}

	var serialToQuery = make([]string, count)

	for i := 0; i < len(serialArray); i++ {
		value := req.FormValue(serialArray[i])
		log.Println(value)
		if value != "" {
			serialToQuery[i] = value
		} else {
			break
		}
	}

	annee := ""

	if req.FormValue("version1") == "4" {
		annee = "2015"
	} else if req.FormValue("version1") == "5" {
		annee = "2016"
	} else {
		annee = "2014"
	}

	//create zip file
	foZippedData, err := os.Create(pathZipFile)
	defer foZippedData.Close()
	checkErr(err, "erreur creation file")

	w := zip.NewWriter(foZippedData)

	/***** SErial analytics ******/
	if req.FormValue("version1") == "4" {
		CreateActivityStatusTableFilev4(w, req.FormValue("serial1"), dmap)
		log.Println("cest la version " + annee)
	} else if req.FormValue("version1") == "5" {
		CreateActivityStatusTableFilev5(w, req.FormValue("serial1"), dmap)
		log.Println("cest la version " + annee)
	} else {
		CreateActivityStatusTableFilev3(w, req.FormValue("serial1"), dmap)
		log.Println("cest la version " + annee)
	}

	/*** Creation downloading file ***/
	w.Close()

	res.Header().Set("Content-Description", "File Transfer")
	res.Header().Set("Content-Type", "application/zip")
	res.Header().Set("Content-Transfer-Encoding", "binary")
	res.Header().Set("Content-Disposition", "attachment; filename=ActivityStatusData"+"_"+finalDate+".zip")
	http.ServeFile(res, req, pathZipFile)
	_, err = io.Copy(foZippedData, req.Body)
	checkErr(err, "erreur copy")

}

func ExtractAnalyticsDashboardCharts(res http.ResponseWriter, req *http.Request, r render.Render, dmap *gorp.DbMap) {

	/**
	Name	:	ExtractAnalyticsDashboardCharts
	Version :	1.0
	Description :
				This function is used to to extract all information of the dashboard.

	State	: Used
	*/

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
	// creation de la date
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

	pathZipFile := root_directory + "/analyticsBasicData" + "_" + finalDate + ".zip"
	log.Println(serialArray[0])
	//count number of field fill
	log.Println(len(serialArray))
	for i := 0; i < len(serialArray); i++ {
		value := req.FormValue(serialArray[i])
		log.Println(value)
		if value != "" {
			count++
		} else {
			break
		}
	}

	var serialToQuery = make([]string, count)

	for i := 0; i < len(serialArray); i++ {
		value := req.FormValue(serialArray[i])
		log.Println(value)
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
	CreateSerialAnalyticsFile(w, dmap, serialToQuery)

	/***** Profile analytics ******/
	CreateProfileAnalyticsFile(w, dmap, serialToQuery)

	CreateActivityStatusAnalyticsBasicFile(w, dmap, serialToQuery)

	//CreateAppAnalyticsBasicFile(w, dmap, serialToQuery)
	/*** Creation downloading file ***/
	w.Close()
	res.Header().Set("Content-Description", "File Transfer")
	res.Header().Set("Content-Type", "application/zip")
	res.Header().Set("Content-Transfer-Encoding", "binary")
	res.Header().Set("Content-Disposition", "attachment; filename=analyticsBasicData"+"_"+finalDate+".zip")
	http.ServeFile(res, req, pathZipFile)
	_, err = io.Copy(foZippedData, req.Body)
	checkErr(err, "erreur copy")

}

func CreateQueryWithRomIdAnalytics(selectQuery *string, serialToQuery []string) {

	/**
	Name	:	CreateQueryWithRomIdAnalytics
	Version :	1.0
	Description :
				This function is used to to create the query for the serial extraction for a specific range of RomId.

	State	: Used
	*/

	for countSerial, serialValue := range serialToQuery {
		if serialValue != "" {
			if countSerial == 0 {
				*selectQuery = *selectQuery + " where serial  like '%" + serialValue + "'"
			} else {
				*selectQuery = *selectQuery + " serial  like '%" + serialValue + "'"
			}

			if countSerial != len(serialToQuery)-1 {
				*selectQuery = *selectQuery + " OR "
			}
		}
	}
}

func CreateQueryWithSerialAnalytics(selectQuery *string, serialToQuery []string) {

	/**
	Name	:	CreateQueryWithSerialAnalytics
	Version :	1.0
	Description :
				This function is used to to create the query for the serial extraction for a specific range of RomId.

	State	: Used
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

/************ Extraction selon champ *******/

/***************  TAB 2  ***************/

func CreateQueryextractSerial_analytics_v4(selectQuery *string, curefsearch string, serialSearch string, emailSearch string, macAdressSearch string) {

	/**
	Name	:	CreateQueryextractSerial_analytics_v4
	Version :	1.0
	Description :
				This function is used to to create the query for the serial extraction for a specicic tablet regarding to the field choosen for
				tab2 Models.

	State	: Used
	*/

	*selectQuery = *selectQuery + " where curef like '" + curefsearch + "'" + " and serial like '" + serialSearch + "'" + " and email like '" + emailSearch + "'" + " and mac_adress like '" + macAdressSearch + "'"
}

/***************  7X  ***************/

func CreateQueryextractSerial_analytics(selectQuery *string, curefsearch string, serialSearch string, emailSearch string, macAdressSearch string) {

	/**
	Name	:	CreateQueryextractSerial_analytics
	Version :	1.0
	Description :
				This function is used to to create the query for the serial extraction for a specicic tablet regarding to the field choosen for
				Xseries Models.

	State	: Used
	*/

	*selectQuery = *selectQuery + " where curef like '" + curefsearch + "'" + " and serial like '" + serialSearch + "'" + " and email like '" + emailSearch + "'" + " and mac_adress like '" + macAdressSearch + "'"
}

/***************  7S  ***************/

func CreateQueryextractSerial(selectQuery *string, serialSearch string, emailSearch string, macAdressSearch string) {

	/**
	Name	:	CreateQueryextractSerial
	Version :	1.0
	Description :
				This function is used to to create the query for the serial extraction for a specicic tablet regarding to the field choosen for
				S series Models.

	State	: Used
	*/

	*selectQuery = *selectQuery + " where  serial like '" + serialSearch + "'" + " and email like '" + emailSearch + "'" + " and mac_adress like '" + macAdressSearch + "'"
}

/***************  OLD TAB  ***************/

func CreateQueryextractregister(selectQuery *string, serialSearch string, emailSearch string, macAdressSearch string) {

	/**
	Name	:	CreateQueryextractSerial
	Version :	1.0
	Description :
				This function is used to to create the query for the serial extraction for a specicic tablet regarding to the field choosen for
				Old Tab Models.

	State	: Used
	*/

	*selectQuery = *selectQuery + " where serial like '" + serialSearch + "'" + " and email like '" + emailSearch + "'" + " and mac_adress like '" + macAdressSearch + "'"
}

/***************  TAB 2  ***************/

func CreateQueryAnalyticsv4(selectQuery *string, req *http.Request) {

	/**
	Name	:	CreateQueryAnalyticsv4
	Version :	1.0
	Description :
				This function is used to to create the query for the serial display of a tab for a specicic tablet regarding to the country
				choosen for Tab2 Models.

	State	: Used
	*/

	if req.FormValue("COUNTRY") == "NL" {
		*selectQuery = *selectQuery + " where serial like ('%326')"
	} else if req.FormValue("COUNTRY") == "DE" {
		*selectQuery = *selectQuery + " where serial like ('%335')"
	} else if req.FormValue("COUNTRY") == "DK" {
		*selectQuery = *selectQuery + " where serial like ('%305')"
	} else if req.FormValue("COUNTRY") == "US" {
		*selectQuery = *selectQuery + " where serial like ('%302') or serial like('403') or serial like ('422')"
	} else if req.FormValue("COUNTRY") == "FR" {
		*selectQuery = *selectQuery + " where serial like ('%310')"
	} else if req.FormValue("COUNTRY") == "CH" {
		*selectQuery = *selectQuery + " where serial like ('%321')"
	} else if req.FormValue("COUNTRY") == "BE" {
		*selectQuery = *selectQuery + " where serial like ('%XXX')"
	} else if req.FormValue("COUNTRY") == "ES" {
		*selectQuery = *selectQuery + " where serial like ('%330')"
	} else if req.FormValue("COUNTRY") == "UK" {
		*selectQuery = *selectQuery + " where serial like ('%301')"
	} else if req.FormValue("COUNTRY") == "LATAM" {
		*selectQuery = *selectQuery + " where curef like '8053%'"
	} else if req.FormValue("COUNTRY") == "EE" {
		*selectQuery = *selectQuery + " where curef like '%E8GB1%'"
	} else if req.FormValue("COUNTRY") == "ZA" {
		*selectQuery = *selectQuery + " where serial like ('%347')"
	}

}
func CreateQueryAnalyticsv6(selectQuery *string, req *http.Request) {

	/**
	Name	:	CreateQueryAnalyticsv6
	Version :	1.0
	Description :
				This function is used to to create the query for the serial display of a tab for a specicic tablet regarding to the country
				choosen for Tab2 Models 2016.

	State	: Used
	*/

	if req.FormValue("COUNTRY") == "NL" {
		*selectQuery = *selectQuery + " where serial like ('%N40')"
	} else if req.FormValue("COUNTRY") == "DE" {
		*selectQuery = *selectQuery + " where serial like ('%D40')"
	} else if req.FormValue("COUNTRY") == "DK" {
		*selectQuery = *selectQuery + " where serial like ('%S40')"
	} else if req.FormValue("COUNTRY") == "US" {
		*selectQuery = *selectQuery + " where serial like ('%U40')"
	} else if req.FormValue("COUNTRY") == "FR" {
		*selectQuery = *selectQuery + " where serial like ('%F40')"
	} else if req.FormValue("COUNTRY") == "CH" {
		*selectQuery = *selectQuery + " where serial like ('%W40')"
	} else if req.FormValue("COUNTRY") == "BE" {
		*selectQuery = *selectQuery + " where serial like ('%B40')"
	} else if req.FormValue("COUNTRY") == "ES" {
		*selectQuery = *selectQuery + " where serial like ('%E40')"
	} else if req.FormValue("COUNTRY") == "UK" {
		*selectQuery = *selectQuery + " where serial like ('%G40')"
	} else if req.FormValue("COUNTRY") == "LATAM" {
		*selectQuery = *selectQuery + " where curef like '8053%'"
	} else if req.FormValue("COUNTRY") == "EE" {
		*selectQuery = *selectQuery + " where curef like '%E8GB2%'"
	}

}

func CreateQueryAnalyticsv8(selectQuery *string, req *http.Request) {

	/**
	Name	:	CreateQueryAnalyticsv6
	Version :	1.0
	Description :
				This function is used to to create the query for the serial display of a tab for a specicic tablet regarding to the country
				choosen for Tab2 Models 2016.

	State	: Used
	*/

	if req.FormValue("COUNTRY") == "NL" {
		*selectQuery = *selectQuery + " where serial like ('%N60')"
	} else if req.FormValue("COUNTRY") == "DE" {
		*selectQuery = *selectQuery + " where serial like ('%D60')"
	} else if req.FormValue("COUNTRY") == "DK" {
		*selectQuery = *selectQuery + " where serial like ('%S60')"
	} else if req.FormValue("COUNTRY") == "US" {
		*selectQuery = *selectQuery + " where serial like ('%U60')"
	} else if req.FormValue("COUNTRY") == "FR" {
		*selectQuery = *selectQuery + " where serial like ('%F60')"
	} else if req.FormValue("COUNTRY") == "CH" {
		*selectQuery = *selectQuery + " where serial like ('%W60')"
	} else if req.FormValue("COUNTRY") == "BE" {
		*selectQuery = *selectQuery + " where serial like ('%B60')"
	} else if req.FormValue("COUNTRY") == "ES" {
		*selectQuery = *selectQuery + " where serial like ('%E60')"
	} else if req.FormValue("COUNTRY") == "UK" {
		*selectQuery = *selectQuery + " where serial like ('%G60')"
	} else if req.FormValue("COUNTRY") == "LATAM" {
		*selectQuery = *selectQuery + " where curef like '8053%'"
	} else if req.FormValue("COUNTRY") == "EE" {
		*selectQuery = *selectQuery + " where curef like '%E8GB2%'"
	}

}

/***************  Serie X  ***************/

func CreateQueryAnalytics(selectQuery *string, req *http.Request) {

	/**
	Name	:	CreateQueryAnalytics
	Version :	1.0
	Description :
				This function is used to to create the query for the serial display of a tab for a specicic tablet regarding to the country
				choosen for Serie X Models.

	State	: Used
	*/

	if req.FormValue("COUNTRY") == "NL" {
		*selectQuery = *selectQuery + " where serial like ('%126')"
	} else if req.FormValue("COUNTRY") == "DE" {
		*selectQuery = *selectQuery + " where serial like ('%135')"
	} else if req.FormValue("COUNTRY") == "DK" {
		*selectQuery = *selectQuery + " where serial like ('%105')"
	} else if req.FormValue("COUNTRY") == "US" {
		*selectQuery = *selectQuery + " where serial like ('%102')"
	} else if req.FormValue("COUNTRY") == "FR" {
		*selectQuery = *selectQuery + " where serial like ('%110')"
	} else if req.FormValue("COUNTRY") == "CH" {
		*selectQuery = *selectQuery + " where serial like ('%121')"
	} else if req.FormValue("COUNTRY") == "BE" {
		*selectQuery = *selectQuery + " where serial like ('%NOSERIAL')"
	} else if req.FormValue("COUNTRY") == "ES" {
		*selectQuery = *selectQuery + " where serial like ('%130')"
	} else if req.FormValue("COUNTRY") == "UK" {
		*selectQuery = *selectQuery + " where serial like ('%101')"
	} else if req.FormValue("COUNTRY") == "LATAM" {
		*selectQuery = *selectQuery + " where curef like '8053%'"
	} else if req.FormValue("COUNTRY") == "EE" {
		*selectQuery = *selectQuery + " where curef like '%E8GB2%'"
	} else if req.FormValue("COUNTRY") == "ZA" {
		*selectQuery = *selectQuery + " where serial like ('%147')"
	}
}

/***************  7S  ***************/

func CreateQueryUSSerial7S(selectQuery *string) {

	/**
	Name	:	CreateQueryUSSerial7S
	Version :	1.0
	Description :
				This function is used to to create the query for the serial extract of a tab for a specicic tablet regarding to the country
				choosen for 7S US Models.

	State	: Used
	*/

	*selectQuery = *selectQuery + " where serial like ('%002_') or serial like ('%003_')"
}

func CreateQueryNLSerial7S(selectQuery *string) {

	/**
	Name	:	CreateQueryNLSerial7S
	Version :	1.0
	Description :
				This function is used to to create the query for the serial extract of a tab for a specicic tablet regarding to the country
				choosen for 7S NL Models.

	State	: Used
	*/

	*selectQuery = *selectQuery + " where serial like ('%025_') or serial like ('%026_')"
}

func CreateQueryDESerial7S(selectQuery *string) {

	/**
	Name	:	CreateQueryDESerial7S
	Version :	1.0
	Description :
				This function is used to to create the query for the serial extract of a tab for a specicic tablet regarding to the country
				choosen for 7S DE Models.

	State	: Used
	*/

	*selectQuery = *selectQuery + " where serial like ('%035_') or serial like ('%036_') or serial like ('%037_')"
}

func CreateQueryDKSerial7S(selectQuery *string) {

	/**
	Name	:	CreateQueryDKSerial7S
	Version :	1.0
	Description :
				This function is used to to create the query for the serial extract of a tab for a specicic tablet regarding to the country
				choosen for 7S DK Models.

	State	: Used
	*/

	*selectQuery = *selectQuery + " where serial like ('%005_')"
}

/***************  4S  ***************/

func CreateQueryUSSerial4S(selectQuery *string) {

	/**
	Name	:	CreateQueryUSSerial4S
	Version :	1.0
	Description :
				This function is used to to create the query for the serial extract of a tab for a specicic tablet regarding to the country
				choosen for 4S US Models.

	State	: Used
	*/

	*selectQuery = *selectQuery + " where serial like ('%505_') or serial like ('%506_') or serial like ('%507_')"
}

func CreateQueryNLSerial4S(selectQuery *string) {

	/**
	Name	:	CreateQueryUSSerial4S
	Version :	1.0
	Description :
				This function is used to to create the query for the serial extract of a tab for a specicic tablet regarding to the country
				choosen for 4S NL Models.

	State	: Used
	*/

	*selectQuery = *selectQuery + " where serial like ('%530_') or serial like ('%531_')"
}

func CreateQueryDKSerial4S(selectQuery *string) {

	/**
	Name	:	CreateQueryDKSerial4S
	Version :	1.0
	Description :
				This function is used to to create the query for the serial extract of a tab for a specicic tablet regarding to the country
				choosen for 4S DK Models.

	State	: Used
	*/

	*selectQuery = *selectQuery + " where serial like ('%520_') "
}

/***************  10S  ***************/

func CreateQueryUSSerial10S(selectQuery *string) {

	/**
	Name	:	CreateQueryUSSerial10S
	Version :	1.0
	Description :
				This function is used to to create the query for the serial extract of a tab for a specicic tablet regarding to the country
				choosen for 10S US Models.

	State	: Used
	*/

	*selectQuery = *selectQuery + " where serial like ('%805_') or serial like ('%807_')"
}

func CreateQueryNLSerial10S(selectQuery *string) {

	/**
	Name	:	CreateQueryNLSerial10S
	Version :	1.0
	Description :
				This function is used to to create the query for the serial extract of a tab for a specicic tablet regarding to the country
				choosen for 10S NL Models.

	State	: Used
	*/

	*selectQuery = *selectQuery + " where serial like ('%817_')"
}

func CreateQueryDESerial10S(selectQuery *string) {

	/**
	Name	:	CreateQueryDESerial10S
	Version :	1.0
	Description :
				This function is used to to create the query for the serial extract of a tab for a specicic tablet regarding to the country
				choosen for 10S DE Models.

	State	: Used
	*/

	*selectQuery = *selectQuery + " where serial like ('%835_')"
}

func CreateQueryDKSerial10S(selectQuery *string) {

	/**
	Name	:	CreateQueryDKSerial10S
	Version :	1.0
	Description :
				This function is used to to create the query for the serial extract of a tab for a specicic tablet regarding to the country
				choosen for 10S DK Models.

	State	: Used
	*/

	*selectQuery = *selectQuery + " where serial like ('%810_') "
}

func CreateProfileAnalyticsFile(w *zip.Writer, dmapExtract *gorp.DbMap, serialToQuery []string) {

	/**
	Name	:	CreateProfileAnalyticsFile
	Version :	1.0
	Description :
				This function is used to to create all data profile for Tab2 Models.

	State	: Used
	*/

	var profiles []ProfileAnalyticsV4
	selectQuery := "select * from Profile_analytics_v4"
	CreateQueryWithRomIdAnalytics(&selectQuery, serialToQuery)

	_, err := dmapExtract.Select(&profiles, selectQuery)
	checkErr(err, "select erreur analytics ")
	f, _ := w.Create("dataProfileAnalytics.xls")

	header := "Serial	Curef	Is_boy	date_birth	Internet_access	is_time_control_on	is_time_slot_activated	"
	header = header + "daily_playtime_week	max_session_duration_week	rest_time_week	"
	header = header + "is_usb_allow	authorize_ads	profile_changed_once	is_web_list_activated	is_profile_deleted_on_tab	\n"
	f.Write([]byte(header))

	for _, profile := range profiles {
		dataOut := profile.Serial + "	"
		if profile.Curef == zero.StringFrom("") {
			dataOut = dataOut + profile.Curef.String + "	"
		} else {
			dataOut = dataOut + "" + "	"
		}
		dataOut = dataOut + profile.IsBoy + "	"
		dataOut = dataOut + profile.Birth.Time.Format("2006/02/01") + "	"
		dataOut = dataOut + profile.InternetAccessMode + "	"
		dataOut = dataOut + profile.TimeControlStatus + "	"
		dataOut = dataOut + profile.IsTimeSlotEnabled + "	"
		dataOut = dataOut + profile.DailyPlayTimeWeekDays + "	"
		dataOut = dataOut + profile.MaxSessionDurationWeekDays + "	"
		dataOut = dataOut + profile.RestTimeWeekDays + "	"
		dataOut = dataOut + profile.AllowUSBStatus + "	"
		dataOut = dataOut + profile.AuthorizeAdsStatus + "	"
		dataOut = dataOut + profile.ProfileChanged + "	"
		dataOut = dataOut + profile.IsWebListActivated + "	"
		dataOut = dataOut + profile.Deleted + "	\n"

		f.Write([]byte(dataOut))
	}
}

func CreateSerialAnalyticsFile(w *zip.Writer, dmapExtract *gorp.DbMap, serialToQuery []string) {

	/**
	Name	:	CreateSerialAnalyticsFile
	Version :	1.0
	Description :
				This function is used to to create all serial profile for Tab2 Models.

	State	: Used
	*/

	var serials []SerialAnalyticsv4
	selectQuery := "select * from Serial_analytics_v4"
	CreateQueryWithRomIdAnalytics(&selectQuery, serialToQuery)

	_, err := dmapExtract.Select(&serials, selectQuery)
	checkErr(err, "select erreur analytics ")
	f, _ := w.Create("dataSerialAnalytics.xls")

	header := "Serial	System_language	date_prod	Email	mailActive	MAC_address	last_parent_log	server_entree	first_activation	setup_date	last_modification	Version	nb_active_profile	nb_child_proifile	nb_app_management_launch	"
	header = header + "nb_faq_launch	nb_contact_us_launch	nb_usermanual_launch	time_control_used_once	\n"
	f.Write([]byte(header))

	for _, serial := range serials {
		dataOut := serial.Serial + "	"
		dataOut = dataOut + serial.SystemLanguage + "	"
		dataOut = dataOut + serial.DateProd.Time.Format("2006/01/02") + "	"
		dataOut = dataOut + serial.ParentEmail + "	"
		dataOut = dataOut + serial.MailActive + "	"
		dataOut = dataOut + serial.MACAddress.String + "	"
		dataOut = dataOut + serial.LastParentLog.Time.Format("2006/01/02") + "	"
		dataOut = dataOut + serial.ServerEntree.Time.Format("2006/01/02") + "	"
		dataOut = dataOut + serial.FirstActivation.Time.Format("2006/01/02") + "	"
		dataOut = dataOut + serial.SetupDate.Time.Format("2006/01/02") + "	"
		dataOut = dataOut + serial.LastModification.Time.Format("2006/01/02") + "	"
		dataOut = dataOut + serial.KurioSystemVersion + "	"
		dataOut = dataOut + strconv.FormatInt(serial.ActiveProfiles.Int64, 10) + "	"
		dataOut = dataOut + strconv.FormatInt(serial.ChildrenProfiles.Int64, 10) + "	"
		dataOut = dataOut + serial.AppManagementCount + "	"
		dataOut = dataOut + serial.FaqCount + "	"
		dataOut = dataOut + serial.ContactUsCount + "	"
		dataOut = dataOut + serial.UserManualCount + "	"
		dataOut = dataOut + serial.TimeControlUsedOnce + "	\n"
		f.Write([]byte(dataOut))

	}
}

func CreateSerialPieCharFilev4(w *zip.Writer, dmapExtract *gorp.DbMap) {

	/**
	Name	:	CreateSerialPieCharFilev4
	Version :	1.0
	Description :
				This function is used to to create all serial profile for Tab2 Models 2015.

	State	: Used
	*/

	var serials []SerialAnalyticsv4
	selectQuery := "select * from Serial_analytics_v4 where serial not like '%40' "

	_, err := dmapExtract.Select(&serials, selectQuery)
	checkErr(err, "select erreur analytics ")
	f, _ := w.Create("Device_Activation_2015.xls")

	header := "Serial	System_language	date_prod	Email	mailActive	MAC_address	last_parent_log	server_entree	first_activation	setup_date	last_modification	Version	nb_active_profile	nb_child_proifile	nb_app_management_launch	"
	header = header + "nb_faq_launch	nb_contact_us_launch	nb_usermanual_launch	time_control_used_once	\n"
	f.Write([]byte(header))

	for _, serial := range serials {
		dataOut := serial.Serial + "	"
		dataOut = dataOut + serial.SystemLanguage + "	"
		dataOut = dataOut + serial.DateProd.Time.Format("2006/01/02") + "	"
		dataOut = dataOut + serial.ParentEmail + "	"
		dataOut = dataOut + serial.MailActive + "	"
		dataOut = dataOut + serial.MACAddress.String + "	"
		dataOut = dataOut + serial.LastParentLog.Time.Format("2006/01/02") + "	"
		dataOut = dataOut + serial.ServerEntree.Time.Format("2006/01/02") + "	"
		dataOut = dataOut + serial.FirstActivation.Time.Format("2006/01/02") + "	"
		dataOut = dataOut + serial.SetupDate.Time.Format("2006/01/02") + "	"
		dataOut = dataOut + serial.LastModification.Time.Format("2006/01/02") + "	"
		dataOut = dataOut + serial.KurioSystemVersion + "	"
		dataOut = dataOut + strconv.FormatInt(serial.ActiveProfiles.Int64, 10) + "	"
		dataOut = dataOut + strconv.FormatInt(serial.ChildrenProfiles.Int64, 10) + "	"
		dataOut = dataOut + serial.AppManagementCount + "	"
		dataOut = dataOut + serial.FaqCount + "	"
		dataOut = dataOut + serial.ContactUsCount + "	"
		dataOut = dataOut + serial.UserManualCount + "	"
		dataOut = dataOut + serial.TimeControlUsedOnce + "	\n"
		f.Write([]byte(dataOut))
	}
}

func CreateContendeliveryTable(w *zip.Writer, dmapExtract *gorp.DbMap) {

	/**
	Name	:	CreateContendeliveryTable
	Version :	1.0
	Description :
				This function is used to to create all content delivery data for Tab2 Models 2016.

	State	: Used
	*/

	selectQuery := ""
	var U40stat []AppAnalytics_v4
	selectQuery = "select distinct package_name, sum(install_count) as install_count, sum( uninstall_count) as uninstall_count from app_analytics_v4 "
	selectQuery = selectQuery + "where package_name in (select package_name from content_delivery_table)"
	selectQuery = selectQuery + "and serial like '%U40'"
	selectQuery = selectQuery + "group by package_name"

	_, err := dmapExtract.Select(&U40stat, selectQuery)
	checkErr(err, "select erreur analytics ")
	f, _ := w.Create("U40stat2016.xls")

	header := "Package Name	InstallCount	UninstallCount" + "	\n"
	f.Write([]byte(header))
	for _, U40stat := range U40stat {
		dataOut := U40stat.PackageName + "	"
		dataOut = dataOut + U40stat.InstallCount + "	"
		dataOut = dataOut + U40stat.UninstallCount + "	\n"
		f.Write([]byte(dataOut))
	}
}

func CreateInstallTableFilev4(w *zip.Writer, romid string, dmapExtract *gorp.DbMap) {

	/**
	Name	:	CreateInstallTableFilev4
	Version :	1.0
	Description :
				This function is used to to create all Install data for Tab2 Models 2015.

	State	: Used
	*/

	selectQuery := ""
	var App_romIdstat []AppAnalytics_v4
	if romid != "" && romid != "all" {
		selectQuery = "select distinct package_name, sum(install_count) as install_count, sum( uninstall_count) as uninstall_count from app_analytics_v4 "
		selectQuery = selectQuery + "where package_name in (select package_name from content_delivery_table)"
		selectQuery = selectQuery + "and serial like '%" + romid + "'"
		selectQuery = selectQuery + "group by package_name"
	} else {
		selectQuery = "select distinct package_name, sum(install_count) as install_count, sum( uninstall_count) as uninstall_count from app_analytics_v4 "
		selectQuery = selectQuery + "where package_name in (select package_name from content_delivery_table)"
		selectQuery = selectQuery + "and serial not like'%40%'"
		selectQuery = selectQuery + "group by package_name"
	}
	_, err := dmapExtract.Select(&App_romIdstat, selectQuery)
	checkErr(err, "select erreur analytics ")
	f, _ := w.Create("Install_count_by_app" + romid + "2016.xls")

	header := "Package Name	Time in Minute	Launch Count" + "	\n"
	f.Write([]byte(header))
	for _, App_romIdstat := range App_romIdstat {
		dataOut := App_romIdstat.PackageName + "	"
		dataOut = dataOut + App_romIdstat.InstallCount + "	"
		dataOut = dataOut + App_romIdstat.UninstallCount + "	\n"
		f.Write([]byte(dataOut))
	}
}
func CreateInstallTableFilev5(w *zip.Writer, romid string, dmapExtract *gorp.DbMap) {

	/**
	Name	:	CreateInstallTableFilev5
	Version :	1.0
	Description :
				This function is used to to create all Install data for Tab2 Models 2016 .

	State	: Used
	*/

	selectQuery := ""
	var App_romIdstat []AppAnalytics_v4
	if romid != "" && romid != "all" {
		selectQuery = "select distinct package_name, sum(install_count) as install_count, sum( uninstall_count) as uninstall_count from app_analytics_v4 "
		selectQuery = selectQuery + "where package_name in (select package_name from content_delivery_table)"
		selectQuery = selectQuery + "and serial like '%" + romid + "'"
		selectQuery = selectQuery + "group by package_name"
	} else {
		selectQuery = "select distinct package_name, sum(install_count) as install_count, sum( uninstall_count) as uninstall_count from app_analytics_v4 "
		selectQuery = selectQuery + "where package_name in (select package_name from content_delivery_table)"
		selectQuery = selectQuery + "and serial  like'%40%'"
		selectQuery = selectQuery + "group by package_name"
	}
	_, err := dmapExtract.Select(&App_romIdstat, selectQuery)
	checkErr(err, "select erreur analytics ")
	f, _ := w.Create("Install_count_by_app" + romid + "2016.xls")

	header := "Package Name	Time in Minute	Launch Count" + "	\n"
	f.Write([]byte(header))
	for _, App_romIdstat := range App_romIdstat {
		dataOut := App_romIdstat.PackageName + "	"
		dataOut = dataOut + App_romIdstat.InstallCount + "	"
		dataOut = dataOut + App_romIdstat.UninstallCount + "	\n"
		f.Write([]byte(dataOut))
	}
}

func CreateInstallTableFilev3(w *zip.Writer, romid string, dmapExtract *gorp.DbMap) {

	/**
	Name	:	CreateInstallTableFilev5
	Version :	1.0
	Description :
				This function is used to to create all Install data for SerieX  .

	State	: Used
	*/

	selectQuery := ""
	var App_romIdstat []AppAnalytics_v4
	if romid != "" && romid != "all" {
		selectQuery = "select distinct package_name, sum(install_count) as install_count, sum( uninstall_count) as uninstall_count from app_analytics "
		selectQuery = selectQuery + "where package_name in (select package_name from content_delivery_table)"
		selectQuery = selectQuery + "and serial like '%" + romid + "'"
		selectQuery = selectQuery + "group by package_name"
	} else {
		selectQuery = "select distinct package_name, sum(install_count) as install_count, sum( uninstall_count) as uninstall_count from app_analytics "
		selectQuery = selectQuery + "where package_name in (select package_name from content_delivery_table)"
		selectQuery = selectQuery + "group by package_name"
	}
	_, err := dmapExtract.Select(&App_romIdstat, selectQuery)
	checkErr(err, "select erreur analytics ")
	f, _ := w.Create("Install_count_by_app" + romid + "2016.xls")

	header := "Package Name	Time in Minute	Launch Count" + "	\n"
	f.Write([]byte(header))
	for _, App_romIdstat := range App_romIdstat {
		dataOut := App_romIdstat.PackageName + "	"
		dataOut = dataOut + App_romIdstat.InstallCount + "	"
		dataOut = dataOut + App_romIdstat.UninstallCount + "	\n"
		f.Write([]byte(dataOut))
	}
}
func CreateActivityStatusTableFilev4(w *zip.Writer, romid string, dmapExtract *gorp.DbMap) {

	/**
	Name	:	CreateActivityStatusTableFilev4
	Version :	1.0
	Description :
				This function is used to to create all Activity Status data for Tab2  2015.

	State	: Used
	*/

	queryWithSerial := "where rom_id = '" + romid + "'"
	selectQuery := ""
	var Activitystatus []TabV4AppPref
	if romid != "" && romid != "all" {
		selectQuery = "select package_name AS package_name, sum(time_spent) AS total_time_spent, sum(launch_count) AS total_launch_count from activity_status_by_romid_v4  " + queryWithSerial + "   group by package_name order by total_time_spent DESC "
	} else {
		selectQuery = "select package_name AS package_name, sum(time_spent) AS total_time_spent, sum(launch_count) AS total_launch_count from activity_status_by_romid_v4   where  rom_id not like '%40' group by package_name order by total_time_spent DESC "
		log.Println("nothing to display v4")
	}
	_, err := dmapExtract.Select(&Activitystatus, selectQuery)
	checkErr(err, "select erreur analytics ")
	f, _ := w.Create("ActivityStatusTable" + romid + "2016.xls")

	header := "Package Name	Time in Minute	Launch Count" + "	\n"
	f.Write([]byte(header))
	for _, Activitystatus := range Activitystatus {
		dataOut := Activitystatus.PackageName + "	"
		dataOut = dataOut + Activitystatus.TotalLaunchCount + "	"
		dataOut = dataOut + Activitystatus.TotalTimeSpent + "	\n"
		f.Write([]byte(dataOut))
	}
}
func CreateActivityStatusTableFilev5(w *zip.Writer, romid string, dmapExtract *gorp.DbMap) {

	/**
	Name	:	CreateActivityStatusTableFilev5
	Version :	1.0
	Description :
				This function is used to to create all Activity Status data for Tab2  2016.

	State	: Used
	*/

	queryWithSerial := "where rom_id = '" + romid + "'"
	selectQuery := ""
	var Activitystatus []TabV4AppPref
	if romid != "" && romid != "all" {
		selectQuery = "select package_name AS package_name, sum(time_spent) AS total_time_spent, sum(launch_count) AS total_launch_count from activity_status_by_romid_v4  " + queryWithSerial + "   group by package_name order by total_time_spent DESC "
	} else {
		selectQuery = "select package_name AS package_name, sum(time_spent) AS total_time_spent, sum(launch_count) AS total_launch_count from activity_status_by_romid_v4   where  rom_id  like '%40' group by package_name order by total_time_spent DESC "
	}
	_, err := dmapExtract.Select(&Activitystatus, selectQuery)
	checkErr(err, "select erreur analytics ")
	f, _ := w.Create("ActivityStatusTable" + romid + "2016.xls")

	header := "Package Name	Time in Minute	Launch Count" + "	\n"
	f.Write([]byte(header))
	for _, Activitystatus := range Activitystatus {
		dataOut := Activitystatus.PackageName + "	"
		dataOut = dataOut + Activitystatus.TotalLaunchCount + "	"
		dataOut = dataOut + Activitystatus.TotalTimeSpent + "	\n"
		f.Write([]byte(dataOut))
	}
}
func CreateActivityStatusTableFilev3(w *zip.Writer, romid string, dmapExtract *gorp.DbMap) {

	/**
	Name	:	CreateActivityStatusTableFilev3
	Version :	1.0
	Description :
				This function is used to to create all Activity Status data for Xseries Models.

	State	: Used
	*/

	queryWithSerial := "where rom_id = '" + romid + "'"
	selectQuery := ""
	var Activitystatus []TabV4AppPref
	if romid != "" && romid != "all" {
		selectQuery = "select package_name AS package_name, sum(time_spent) AS total_time_spent, sum(launch_count) AS total_launch_count from activity_status_by_romid  " + queryWithSerial + "  group by package_name order by total_time_spent DESC "
	} else {
		log.Println("nothing to display")

	}
	_, err := dmapExtract.Select(&Activitystatus, selectQuery)
	checkErr(err, "select erreur analytics ")
	f, _ := w.Create("ActivityStatusTable" + romid + "2016.xls")

	header := "Package Name	Time in Minute	Launch Count" + "	\n"
	f.Write([]byte(header))
	for _, Activitystatus := range Activitystatus {
		dataOut := Activitystatus.PackageName + "	"
		dataOut = dataOut + Activitystatus.TotalLaunchCount + "	"
		dataOut = dataOut + Activitystatus.TotalTimeSpent + "	\n"
		f.Write([]byte(dataOut))
	}
}

func CreateSerialPieCharFilev5(w *zip.Writer, dmapExtract *gorp.DbMap) {

	/**
	Name	:	CreateSerialPieCharFilev5
	Version :	1.0
	Description :
				This function is used to to create all serial data of the pie chart for Tab2 Models 2016.

	State	: Used
	*/

	var serials []SerialAnalyticsv4
	selectQuery := "select * from Serial_analytics_v4 where serial like '%40' and serial not like '%340' "

	_, err := dmapExtract.Select(&serials, selectQuery)
	checkErr(err, "select erreur analytics ")
	f, _ := w.Create("Device_Activation_2016.xls")

	header := "Serial	System_language	date_prod	Email	mailActive	MAC_address	last_parent_log	server_entree	first_activation	setup_date	last_modification	Version	nb_active_profile	nb_child_proifile	nb_app_management_launch	"
	header = header + "nb_faq_launch	nb_contact_us_launch	nb_usermanual_launch	time_control_used_once	\n"
	f.Write([]byte(header))

	for _, serial := range serials {
		dataOut := serial.Serial + "	"
		dataOut = dataOut + serial.SystemLanguage + "	"
		dataOut = dataOut + serial.DateProd.Time.Format("2006/01/02") + "	"
		dataOut = dataOut + serial.ParentEmail + "	"
		dataOut = dataOut + serial.MailActive + "	"
		dataOut = dataOut + serial.MACAddress.String + "	"
		dataOut = dataOut + serial.LastParentLog.Time.Format("2006/01/02") + "	"
		dataOut = dataOut + serial.ServerEntree.Time.Format("2006/01/02") + "	"
		dataOut = dataOut + serial.FirstActivation.Time.Format("2006/01/02") + "	"
		dataOut = dataOut + serial.SetupDate.Time.Format("2006/01/02") + "	"
		dataOut = dataOut + serial.LastModification.Time.Format("2006/01/02") + "	"
		dataOut = dataOut + serial.KurioSystemVersion + "	"
		dataOut = dataOut + strconv.FormatInt(serial.ActiveProfiles.Int64, 10) + "	"
		dataOut = dataOut + strconv.FormatInt(serial.ChildrenProfiles.Int64, 10) + "	"
		dataOut = dataOut + serial.AppManagementCount + "	"
		dataOut = dataOut + serial.FaqCount + "	"
		dataOut = dataOut + serial.ContactUsCount + "	"
		dataOut = dataOut + serial.UserManualCount + "	"
		dataOut = dataOut + serial.TimeControlUsedOnce + "	\n"
		f.Write([]byte(dataOut))

	}
}

func CreateSerialPieCharFilev3(w *zip.Writer, dmapExtract *gorp.DbMap) {

	/**
	Name	:	CreateSerialPieCharFilev3
	Version :	1.0
	Description :
				This function is used to to create all serial data of the pie chart for Xseries Models.

	State	: Used
	*/

	var serials []SerialAnalytics
	selectQuery := "select * from Serial_analytics"

	_, err := dmapExtract.Select(&serials, selectQuery)
	checkErr(err, "select erreur analytics ")
	f, _ := w.Create("Device_Activation_2014.xls")

	header := "Serial	System_language	date_prod	Email	mailActive	MAC_address	last_parent_log	server_entree	first_activation	setup_date	last_modification	Version	nb_active_profile	nb_child_proifile	nb_app_management_launch	"
	header = header + "nb_faq_launch	nb_contact_us_launch	nb_usermanual_launch	time_control_used_once	\n"
	f.Write([]byte(header))

	for _, serial := range serials {
		dataOut := serial.Serial + "	"
		dataOut = dataOut + serial.SystemLanguage + "	"
		dataOut = dataOut + serial.DateProd.Time.Format("2006/01/02") + "	"
		dataOut = dataOut + serial.ParentEmail + "	"
		dataOut = dataOut + serial.MailActive + "	"
		dataOut = dataOut + serial.MACAddress.String + "	"
		dataOut = dataOut + serial.LastParentLog.Time.Format("2006/01/02") + "	"
		dataOut = dataOut + serial.ServerEntree.Time.Format("2006/01/02") + "	"
		dataOut = dataOut + serial.FirstActivation.Time.Format("2006/01/02") + "	"
		dataOut = dataOut + serial.SetupDate.Time.Format("2006/01/02") + "	"
		dataOut = dataOut + serial.LastModification.Time.Format("2006/01/02") + "	"
		dataOut = dataOut + serial.KurioSystemVersion + "	"
		dataOut = dataOut + strconv.FormatInt(serial.ActiveProfiles.Int64, 10) + "	"
		dataOut = dataOut + strconv.FormatInt(serial.ChildrenProfiles.Int64, 10) + "	"
		dataOut = dataOut + serial.AppManagementCount + "	"
		dataOut = dataOut + serial.FaqCount + "	"
		dataOut = dataOut + serial.ContactUsCount + "	"
		dataOut = dataOut + serial.UserManualCount + "	"
		dataOut = dataOut + serial.TimeControlUsedOnce + "	\n"
		f.Write([]byte(dataOut))

	}
}

func CreateActivityStatusAnalyticsFile(w *zip.Writer, dmapExtract *gorp.DbMap, serialToQuery []string) {

	/**
	Name	:	CreateActivityStatusAnalyticsFile
	Version :	1.0
	Description :
				This function is used to to create all launch activity Status data for a ROM ID Tab2 Models 2016.

	State	: Used
	*/

	var activities []ActivityStatusAnalyticsByRomId
	selectQuery := "select * from activity_status_by_romid_v4"
	CreateQueryWithRomIdAnalytics(&selectQuery, serialToQuery)

	_, err := dmapExtract.Select(&activities, selectQuery)
	checkErr(err, "select erreur analytics ")

	f, _ := w.Create("dataActivityStatusAnalytics.xls")

	header := "RomId	package_name	activity_name	launch_count	time_spent_in_second	updated_at\n"
	f.Write([]byte(header))

	for _, activity := range activities {
		dataOut := activity.RomId + "	"
		dataOut = dataOut + activity.PackageName + "	"
		dataOut = dataOut + activity.ActivityName + "	"
		dataOut = dataOut + strconv.FormatInt(activity.LaunchCount, 10) + "	"
		dataOut = dataOut + strconv.FormatInt(activity.TimeSpent, 10) + "	"
		dataOut = dataOut + activity.UpdatedAt.Time.Format("2006/01/02") + "	\n"

		f.Write([]byte(dataOut))

	}
}

func CreateActivityStatusAnalyticsBasicFile(w *zip.Writer, dmapExtract *gorp.DbMap, serialToQuery []string) {

	/**
	Name	:	CreateActivityStatusAnalyticsBasicFile
	Version :	1.0
	Description :
				This function is used to to create all launch activity Status  for a package_name for Tab2 Models.

	State	: Used
	*/

	var appPref []TabV4AppPref
	_, err := dmapExtract.Select(&appPref, "SELECT * FROM KurioS_.activity_status_view_v4")
	checkErr(err, "Erreur select from view activity_status_view_v4")

	f, _ := w.Create("dataActivityStatusAnalyticsBasic.xls")

	header := "package_name	time_spent_in_second	launch_count	\n"
	f.Write([]byte(header))

	for _, app := range appPref {
		dataOut := app.PackageName + "	"
		dataOut = dataOut + app.TotalTimeSpent + "	"
		dataOut = dataOut + app.TotalLaunchCount + "	\n"

		f.Write([]byte(dataOut))

	}
}

func CreateAppAnalyticsFile(w *zip.Writer, dmapExtract *gorp.DbMap, serialToQuery []string) {

	/**
	Name	:	CreateAppAnalyticsFile
	Version :	1.0
	Description :
				This function is used to to create all install data of app  for a serial on Tab2 Models.

	State	: Used
	*/

	var appAnalytics []AppAnalytics_v4
	selectQuery := "select * from app_analytics_v4"
	CreateQueryWithRomIdAnalytics(&selectQuery, serialToQuery)

	_, err := dmapExtract.Select(&appAnalytics, selectQuery)
	checkErr(err, "select erreur analytics_v4 ")

	f, _ := w.Create("dataAppAnalytics.xls")

	header := "Serial	package_name	total_install_count	total_uninstall_count	\n"

	f.Write([]byte(header))

	for _, appAnalytic := range appAnalytics {
		dataOut := appAnalytic.Serial + "	"
		dataOut = dataOut + appAnalytic.PackageName + "	"
		dataOut = dataOut + appAnalytic.InstallCount + "	"
		dataOut = dataOut + appAnalytic.UninstallCount + "	\n"

		f.Write([]byte(dataOut))
	}
}

func CreateAppAnalyticsBasicFile(w *zip.Writer, dmapExtract *gorp.DbMap, serialToQuery []string) {

	/**
	Name	:	CreateAppAnalyticsBasicFile
	Version :	1.0
	Description :
				This function is used to to create all install of app  for a package_name  on Tab2 Models.

	State	: Used
	*/

	var appAnalytics []AppAnalytics_v4
	selectQuery := "select * from KurioS_.app_static_view_v4"

	_, err := dmapExtract.Select(&appAnalytics, selectQuery)
	checkErr(err, "select erreur analytics ")

	f, _ := w.Create("dataAppAnalyticsBasic.xls")

	header := "package_name	total_install_count	total_uninstall_count	\n"
	f.Write([]byte(header))

	for _, appAnalytic := range appAnalytics {
		dataOut := appAnalytic.PackageName + "	"
		dataOut = dataOut + appAnalytic.InstallCount + "	"
		dataOut = dataOut + appAnalytic.UninstallCount + "	\n"

		f.Write([]byte(dataOut))
	}
}

func clearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

func AnalyticsViewerConnect(res http.ResponseWriter, req *http.Request, r render.Render, session sessions.Session) {

	/**
	Name	:	AnalyticsViewerConnect
	Version :	1.0
	Description :
				This function is dedicated to launch the html page where user enter information before connecting.

	State	: Used
	*/
	clearSession(res)
	idsession = ""
	session.Set("kurioPwd", "123")
	newmap := map[string]interface{}{"metatitle": "Analytics Application"}
	r.HTML(200, "analytics_viewer_connect", newmap)
}

func AnalyticsBadPassword(res http.ResponseWriter, req *http.Request, r render.Render, session sessions.Session) {

	/**
	Name	:	AnalyticsBadPassword
	Version :	1.0
	Description :
				This function is dedicated to launch the html page when user enter the bad password.

	State	: Used
	*/

	session.Set("kurioPwd", "123")
	newmap := map[string]interface{}{"metatitle": "Analytics Application"}
	r.HTML(200, "analytics_bad_password", newmap)
}

func AnalyticsBadUser(res http.ResponseWriter, req *http.Request, r render.Render, session sessions.Session) {

	/**
	Name	:	AnalyticsBadUser
	Version :	1.0
	Description :
				This function is dedicated to launch the html page when user try to display a page not allowed for his account.

	State	: Used
	*/

	session.Set("kurioPwd", "123")
	newmap := map[string]interface{}{"metatitle": "Analytics Application"}
	r.HTML(200, "analytics_bad_user", newmap)
}

func hasSymbol(str string) bool {
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ@."
	alphabetSplit := strings.Split(alphabet, "")
	inputLetters := strings.Split(str, "")
	for _, value := range inputLetters {
		found := false
		for _, char := range alphabetSplit {
			if char == value {
				found = true
				break
			}
		}
		if !found {
			return true
		}
	}
	return false
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	return b, err
}

func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func AnalyticsViewerLogin(res http.ResponseWriter, req *http.Request, r render.Render, session sessions.Session) {

	/**
	Name	:	AnalyticsViewerLogin
	Version :	1.0
	Description :
				 this function is dedicated to manage first page launched after registration for each users.
				 Global Variable will be used soon
				 code to use for new password and licenced acces
				 password := req.FormValue("password")
				 session.Set("kurioPwd", password)
				 http.Redirect(res, req, "/analytics_viewer_choice", 302)

	State	: Used
	*/

	login := req.FormValue("Login")
	password := req.FormValue("password")
	token, err := GenerateRandomString(32)

	if hasSymbol(login) {
		session.Set("kurioPwd", "bad")
		session.Options(sessions.Options{
			MaxAge: 0,
		})
		//set session duration to 30 minutes
		http.Redirect(res, req, "/analytics_hacking", 302)
	} else {
		if login == "cidesei3@gmail.com" && password == "cidetoys" {
			if err != nil {
				panic(err)
			}
			idsession = "extract"
			session.Set(token, token)
			session.Options(sessions.Options{
				MaxAge: 60 * 30,
			})
			http.Redirect(res, req, "/analytics_viewer_choice", 302)
		} else if login == "Thomas" && password == "Arronis" {
			if err != nil {
				panic(err)
			}
			idsession = "extract"
			session.Set(token, token)
			session.Options(sessions.Options{
				MaxAge: 0,
			})
			//set session duration to 30 minutes
			http.Redirect(res, req, "/analytics_viewer_julien", 302)
		} else if login == "Julien" && password == "Mer" {
			if err != nil {
				panic(err)
			}
			idsession = "extract"
			session.Set(token, token)
			session.Options(sessions.Options{
				MaxAge: 0,
			})
			//set session duration to 30 minutes
			http.Redirect(res, req, "/analytics_viewer_julien", 302)
		} else if login == "Jeremy" && password == "Fontaine" {
			if err != nil {
				panic(err)
			}
			idsession = "Jeremy"
			session.Set(token, token)
			session.Options(sessions.Options{
				MaxAge: 0,
			})
			//set session duration to 30 minutes
			http.Redirect(res, req, "/analytics_viewer_jeremie", 302)
		} else if login == "Caroline" && password == "Wu" {
			if err != nil {
				panic(err)
			}
			idsession = "Caroline"
			session.Set(token, token)
			session.Options(sessions.Options{
				MaxAge: 0,
			})
			//set session duration to 30 minutes
			http.Redirect(res, req, "/analytics_viewer_jeremie", 302)
		} else {
			session.Set("kurioPwd", "bad")
			session.Options(sessions.Options{
				MaxAge: 0,
			})
			//set session duration to 30 minutes
			http.Redirect(res, req, "/analytics_bad_password", 302)
		}
	}

}

func AnalyticsViewerChoice(res http.ResponseWriter, req *http.Request, r render.Render, session sessions.Session) {

	/**
	Name	:	AnalyticsViewerChoice
	Version :	1.0
	Description :
				 this function is dedicated to launch the html page for the choice of action between extract analytics data and the information
				 for QA

	State	: Used
	*/

	if (idsession == "cidesei3") || (idsession == "Jeremy") || (idsession == "Caroline") {
		newmap := map[string]interface{}{"metatitle": "Analytics Application"}
		r.HTML(200, "analytics_viewer_choice", newmap)

	} else {
		http.Redirect(res, req, "/analytics_bad_user", 302)
	}

}

func AnalyticsViewerJulien(res http.ResponseWriter, req *http.Request, r render.Render, session sessions.Session) {

	/**
	Name	:	AnalyticsViewerOlivier
	Version :	1.0
	Description :
				 this function is dedicated to launch the html page for QA account
				 for QA

	State	: Used
	*/

	if idsession == "extract" {
		newmap := map[string]interface{}{"metatitle": "Analytics Application"}
		r.HTML(200, "analytics_viewer_julien", newmap)
	} else {
		http.Redirect(res, req, "/analytics_bad_user", 302)
	}
}
func AnalyticsViewerjeremie(res http.ResponseWriter, req *http.Request, r render.Render, session sessions.Session) {

	/**
	Name	:	AnalyticsViewerjeremie
	Version :	1.0
	Description :
				 this function is dedicated to launch the html page for Non-QA account
				 for QA

	State	: Used
	*/

	if (idsession == "Jeremy") || (idsession == "Caroline") {
		newmap := map[string]interface{}{"metatitle": "Analytics Application"}
		r.HTML(200, "analytics_viewer", newmap)
	} else {
		http.Redirect(res, req, "/analytics_bad_user", 302)
	}
}

func AnalyticsGetRomIdStatus(params martini.Params, res http.ResponseWriter, req *http.Request, r render.Render, dmap *gorp.DbMap, session sessions.Session) {

	/**
	Name	:	AnalyticsGetRomIdStatus
	Version :	1.0
	Description :
				 this function is dedicated to display the information contained on the piechart for the number of item for each ROMID

	State	: Used
	*/

	var tab7SRomId []Tab7SByRomId
	var tab7XRomId []Tab7XByRomId
	var tabV4RomId []TabV4ByRomId

	if params["version"] == "3" {
		AnalyticsGetRomIdStatus_v3(dmap, &tab7XRomId)
		r.JSON(200, tab7XRomId)
	} else if params["version"] == "4" {
		AnalyticsGetRomIdStatus_v4(dmap, &tabV4RomId)
		r.JSON(200, tabV4RomId)
	} else if params["version"] == "2" {
		AnalyticsGetRomIdStatus_v2(dmap, &tab7SRomId)
		r.JSON(200, tab7SRomId)
	} else if params["version"] == "5" {
		AnalyticsGetRomIdStatus_v6(dmap, &tabV4RomId)
		r.JSON(200, tabV4RomId)
	} else if params["version"] == "6" {
		AnalyticsGetRomIdStatus_v8(dmap, &tabV4RomId)
		r.JSON(200, tabV4RomId)
	}
}

func AnalyticsGetRomIdStatus_v2(dmap *gorp.DbMap, tab7SRomId *[]Tab7SByRomId) {

	/**
	Name	:	AnalyticsGetRomIdStatus_v2
	Version :	1.0
	Description :
				 this function is dedicated to display the information contained on the piechart for the number of item for each ROMID specific for S series

	State	: Used
	*/
	_, err := dmap.Select(tab7SRomId, "select * from 7s_by_rom_id where rom_id !=''")
	checkErr(err, "Erreur select from view 7s_by_rom_id")
	SetCountry_v2(*tab7SRomId)
}

func AnalyticsGetRomIdStatus_v3(dmap *gorp.DbMap, tab7XRomId *[]Tab7XByRomId) {

	/**
	Name	:	AnalyticsGetRomIdStatus_v3
	Version :	1.0
	Description :
				 this function is dedicated to display the information contained on the piechart for the number of item for each ROMID specific for X series

	State	: Used
	*/

	_, err := dmap.Select(tab7XRomId, "SELECT * FROM KurioS_.7x_by_rom_id where rom_id !='' ")
	checkErr(err, "Erreur select from view 7x_by_rom_id")
	SetCountry_v3(*tab7XRomId)
}

func AnalyticsGetRomIdStatus_v4(dmap *gorp.DbMap, tabV4RomId *[]TabV4ByRomId) {

	/**
	Name	:	AnalyticsGetRomIdStatus_v4
	Version :	1.0
	Description :
				 this function is dedicated to display the information contained on the piechart for the number of item for each ROMID specific for Tab2 series 2015

	State	: Used
	*/

	log.Println("on commence v4 ")
	_, err := dmap.Select(tabV4RomId, "SELECT * FROM KurioS_.by_rom_id_view_v4 where rom_id not like '%40' and rom_id not like '%60'")
	//uncomment to make curef work correctly and not whith tmpCuref
	_, err = dmap.Select(tabV4RomId, "SELECT * FROM KurioS_.by_rom_id_view_v4_curef where rom_id not like '%40'  and rom_id not like '%60' and LENGTH(rom_id) != 4")
	checkErr(err, "Erreur select from view by_rom_id_view_v4")
	SetCountry_v4(*tabV4RomId)
}

func AnalyticsGetRomIdStatus_v6(dmap *gorp.DbMap, tabV4RomId *[]TabV4ByRomId) {

	/**
	Name	:	AnalyticsGetRomIdStatus_v6
	Version :	1.0
	Description :
				 this function is dedicated to display the information contained on the piechart for the number of item for each ROMID specific for Tab2 series 2016

	State	: Used
	*/
	log.Println("on commence v6 ")
	_, err := dmap.Select(tabV4RomId, "SELECT * FROM KurioS_.by_rom_id_view_v4 where rom_id like '%40' ")
	//uncomment to make curef work correctly and not whith tmpCuref
	_, err = dmap.Select(tabV4RomId, "SELECT * FROM KurioS_.by_rom_id_view_v4_curef where rom_id like '%40'  and LENGTH(rom_id) != 4")
	checkErr(err, "Erreur select from view by_rom_id_view_v4")
	SetCountry_v6(*tabV4RomId)
}

func AnalyticsGetRomIdStatus_v8(dmap *gorp.DbMap, tabV4RomId *[]TabV4ByRomId) {

	/**
	Name	:	AnalyticsGetRomIdStatus_v6
	Version :	1.0
	Description :
				 this function is dedicated to display the information contained on the piechart for the number of item for each ROMID specific for Tab2 series 2016

	State	: Used
	*/

	_, err := dmap.Select(tabV4RomId, "SELECT * FROM KurioS_.by_rom_id_view_v4 where rom_id like '%60' or rom_id like '%61' ")
	//uncomment to make curef work correctly and not whith tmpCuref
	_, err = dmap.Select(tabV4RomId, "SELECT * FROM KurioS_.by_rom_id_view_v4_curef where rom_id like '%60' or rom_id like '%61' and LENGTH(rom_id) != 4")
	checkErr(err, "Erreur select from view by_rom_id_view_v4")
	SetCountry_v8(*tabV4RomId)
}

func AnalyticsGetProfiles(params martini.Params, res http.ResponseWriter, req *http.Request, r render.Render, dmap *gorp.DbMap) {

	/**
	Name	:	AnalyticsGetProfiles
	Version :	1.0
	Description :
				 this function is dedicated to display profile information from database

	State	: Used
	*/

	var profiles AllProfileInfos
	if params["version"] == "3" {
		AnalyticsGetProfiles_v3(dmap, &profiles)
	} else if params["version"] == "4" {
		AnalyticsGetProfiles_v4(dmap, &profiles)
	}
	r.JSON(200, profiles)
}

func AnalyticsGetProfiles_v3(dmap *gorp.DbMap, profiles *AllProfileInfos) {

	/**
	Name	:	AnalyticsGetProfiles_v3
	Version :	1.0
	Description :
				 this function is dedicated to X series display profile information from database

	State	: Used
	*/

	err := dmap.SelectOne(profiles, "SELECT * FROM KurioS_.7x_profile_status")
	checkErr(err, "Erreur select from view 7x_profile_status")
}

func AnalyticsGetProfiles_v4(dmap *gorp.DbMap, profiles *AllProfileInfos) {

	/**
	Name	:	AnalyticsGetProfiles_v4
	Version :	1.0
	Description :
				 this function is dedicated to Tab2 series display profile information from database

	State	: Used
	*/

	err := dmap.SelectOne(profiles, "SELECT * FROM KurioS_.profile_status_view_v4")
	checkErr(err, "Erreur select from view profile_status_view_v4")
}

func AnalyticsGetActivation(params martini.Params, res http.ResponseWriter, req *http.Request, r render.Render, dmap *gorp.DbMap, session sessions.Session) {

	/**
	Name	:	AnalyticsGetActivation
	Version :	1.0
	Description :
				 this function is dedicated to display activation information from database

	State	: Used
	*/

	var tab7XActivation []Tab7XActivation
	var tabV4Activation []TabV4Activation
	var tab7SActivation []Tab7SActivation

	log.Println("AnalyticsGetActivation")
	// romid := session.Get("romid").(string)

	serialSelected := params["serial"]
	versionSelected := params["version"]

	if versionSelected == "3" {
		log.Println("AnalyticsGetActivation V3")
		AnalyticsGetActivation_v3(dmap, serialSelected, &tab7XActivation)
		r.JSON(200, tab7XActivation)
	} else if versionSelected == "4" {
		log.Println("AnalyticsGetActivation V4")
		AnalyticsGetActivation_v4(dmap, serialSelected, &tabV4Activation)
		r.JSON(200, tabV4Activation)
	} else if versionSelected == "2" {
		log.Println("AnalyticsGetActivation V2")
		AnalyticsGetActivation_v2(dmap, serialSelected, &tab7SActivation)
		r.JSON(200, tab7SActivation)
	} else if versionSelected == "5" {
		log.Println("AnalyticsGetActivation V6")
		AnalyticsGetActivation_v6(dmap, serialSelected, &tabV4Activation)
		r.JSON(200, tabV4Activation)
	} else if versionSelected == "6" {
		log.Println("AnalyticsGetActivation V8")
		AnalyticsGetActivation_v8(dmap, serialSelected, &tabV4Activation)
		r.JSON(200, tabV4Activation)
	}
}

func AnalyticsGetActivation_v3(dmap *gorp.DbMap, serialSelected string, tab7XActivation *[]Tab7XActivation) {

	/**
	Name	:	AnalyticsGetActivation_v3
	Version :	1.0
	Description :
				 this function is dedicated to display X series activation information from database

	State	: Used
	*/

	log.Println("AnalyticsGetActivation_v3 function")
	if serialSelected != "" && serialSelected != "all" {
		queryWithSerial := "where serial like '%" + serialSelected + "'"
		_, err := dmap.Select(tab7XActivation, "SELECT first_activation as activation_date, count(first_activation) as count from Serial_analytics "+queryWithSerial+" group by first_activation")
		checkErr(err, "Erreur select from view 7x_activation_date")
	} else {
		_, err := dmap.Select(tab7XActivation, "SELECT * FROM 7x_activatation_date")
		checkErr(err, "Erreur select from view 7x_activation_date")
	}
}

func AnalyticsGetActivation_v2(dmap *gorp.DbMap, serialSelected string, tab7SActivation *[]Tab7SActivation) {

	/**
	Name	:	AnalyticsGetActivation_v2
	Version :	1.0
	Description :
				 this function is dedicated to display S series activation information from database

	State	: Used
	*/

	log.Println("AnalyticsGetActivation_v2 function")
	if serialSelected != "" && serialSelected != "all" {
		queryWithSerial := "where serial like '%" + serialSelected + "0'"
		_, err := dmap.Select(tab7SActivation, "SELECT first_activation as activation_date, count(first_activation) as count from Serial "+queryWithSerial+" group by first_activation")
		checkErr(err, "Erreur select from view 7s_activation_date")
	} else {
		_, err := dmap.Select(tab7SActivation, "SELECT * FROM 7s_activation_date")
		checkErr(err, "Erreur select from view 7s_activation_date")
	}
}

func AnalyticsGetActivation_v4(dmap *gorp.DbMap, serialSelected string, tabV4Activation *[]TabV4Activation) {

	/**
	Name	:	AnalyticsGetActivation_v4
	Version :	1.0
	Description :
				 this function is dedicated to display Tab 2 2015 activation information from database

	State	: Used
	*/

	log.Println("AnalyticsGetActivation_v4 function")
	log.Println("serial selected : ", serialSelected)
	if serialSelected != "" && serialSelected != "all" {
		if serialSelected == "tmpCuref" {
			queryWithSerial := "where LENGTH(serial) < 20"
			log.Println("Query 1 : ", queryWithSerial)
			_, err := dmap.Select(tabV4Activation, "SELECT first_activation as activation_date, count(first_activation) as count from Serial_analytics_v4 "+queryWithSerial+" group by first_activation")
			checkErr(err, "Erreur select from view activation_date_view_v4 tmpcuref")
		} else if len(serialSelected) == 4 {
			queryWithSerial := "where curef like '%" + serialSelected + "'"
			log.Println("Query 1 : ", queryWithSerial)
			_, err := dmap.Select(tabV4Activation, "SELECT first_activation as activation_date, count(first_activation) as count from Serial_analytics_v4 "+queryWithSerial+" group by first_activation")
			checkErr(err, "Erreur select from view activation_date_view_v4 LA VUE v4")
		} else {
			queryWithSerial := "where serial like '%" + serialSelected + "'"
			log.Println("Query 1 : ", queryWithSerial)
			_, err := dmap.Select(tabV4Activation, "SELECT first_activation as activation_date, count(first_activation) as count from Serial_analytics_v4 "+queryWithSerial+" group by first_activation")
			checkErr(err, "Erreur select from view activation_date_view_v4 special curef")
		}

	} else {
		log.Println("Query 2 : ")
		_, err := dmap.Select(tabV4Activation, "SELECT * FROM KurioS_.activation_date_view_v4 ")
		checkErr(err, "Erreur select from view activation_date_view_v4")
	}
}

func AnalyticsGetActivation_v6(dmap *gorp.DbMap, serialSelected string, tabV4Activation *[]TabV4Activation) {

	/**
	Name	:	AnalyticsGetActivation_v6
	Version :	1.0
	Description :
				 this function is dedicated to display Tab 2 2016 activation information from database

	State	: Used
	*/

	log.Println("AnalyticsGetActivation_v6 function")
	log.Println("serial selected : ", serialSelected)
	if serialSelected != "" && serialSelected != "all" {
		if serialSelected == "tmpCuref" {
			queryWithSerial := "where LENGTH(serial) < 20"
			log.Println("Query 1 : ", queryWithSerial)
			_, err := dmap.Select(tabV4Activation, "SELECT first_activation as activation_date, count(first_activation) as count from Serial_analytics_v4 "+queryWithSerial+" group by first_activation")
			checkErr(err, "Erreur select from view activation_date_view_v4")
		} else if len(serialSelected) == 4 {
			queryWithSerial := "where curef like '%" + serialSelected + "'"
			log.Println("Query 1 : ", queryWithSerial)
			_, err := dmap.Select(tabV4Activation, "SELECT first_activation as activation_date, count(first_activation) as count from Serial_analytics_v4 "+queryWithSerial+" group by first_activation")
			checkErr(err, "Erreur select from view activation_date_view_v4")
		} else {
			queryWithSerial := "where serial like '%" + serialSelected + "'"
			log.Println("Query 1 : ", queryWithSerial)
			_, err := dmap.Select(tabV4Activation, "SELECT first_activation as activation_date, count(first_activation) as count from Serial_analytics_v4 "+queryWithSerial+" group by first_activation")
			checkErr(err, "Erreur select from view activation_date_view_v4 special curef")
		}

	} else {
		log.Println("Query 2 : ")
		_, err := dmap.Select(tabV4Activation, "SELECT * FROM KurioS_.activation_date_2016_view_v4")
		checkErr(err, "Erreur select from view activation_date_view_v4")
	}
}

func AnalyticsGetActivation_v8(dmap *gorp.DbMap, serialSelected string, tabV4Activation *[]TabV4Activation) {

	/**
	Name	:	AnalyticsGetActivation_v8
	Version :	1.0
	Description :
				 this function is dedicated to display Tab 2 2016 activation information from database

	State	: Used
	*/

	log.Println("AnalyticsGetActivation_v8 function")
	log.Println("serial selected : ", serialSelected)
	if serialSelected != "" && serialSelected != "all" {
		if serialSelected == "tmpCuref" {
			queryWithSerial := "where LENGTH(serial) < 20"
			log.Println("Query 1 : ", queryWithSerial)
			_, err := dmap.Select(tabV4Activation, "SELECT first_activation as activation_date, count(first_activation) as count from Serial_analytics_v4 "+queryWithSerial+" group by first_activation")
			checkErr(err, "Erreur select from view activation_date_view_v4")
		} else if len(serialSelected) == 4 {
			queryWithSerial := "where curef like '%" + serialSelected + "'"
			log.Println("Query 1 : ", queryWithSerial)
			_, err := dmap.Select(tabV4Activation, "SELECT first_activation as activation_date, count(first_activation) as count from Serial_analytics_v4 "+queryWithSerial+" group by first_activation")
			checkErr(err, "Erreur select from view activation_date_view_v4")
		} else {
			queryWithSerial := "where serial like '%" + serialSelected + "'"
			log.Println("Query 1 : ", queryWithSerial)
			_, err := dmap.Select(tabV4Activation, "SELECT first_activation as activation_date, count(first_activation) as count from Serial_analytics_v4 "+queryWithSerial+" group by first_activation")
			checkErr(err, "Erreur select from view activation_date_view_v4 special curef")
		}

	} else {
		log.Println("Query 2 : ")
		_, err := dmap.Select(tabV4Activation, "SELECT * FROM KurioS_.activation_date_2017_view_v4")
		checkErr(err, "Erreur select from view activation_date_view_v4")
	}
}

func AnalyticsGetActivityStatus(params martini.Params, res http.ResponseWriter, req *http.Request, r render.Render, dmap *gorp.DbMap, session sessions.Session) {

	/**
	Name	:	AnalyticsGetActivityStatus
	Version :	1.0
	Description :
				 this function is dedicated to display activity_status information from database

	State	: Used
	*/

	var appPref7X []Tab7XAppPref
	var appPrefV4 []TabV4AppPref
	var appPref7S []Tab7SAppPref

	log.Println("AnalyticsGetActivityStatus function")

	versionSelected := params["version"]
	serialSelected := params["serial"]

	if versionSelected == "3" {
		AnalyticsGetActivityStatus_v3(dmap, serialSelected, &appPref7X)
		r.JSON(200, appPref7X)
	} else if versionSelected == "4" {
		AnalyticsGetActivityStatus_v4(dmap, serialSelected, &appPrefV4)
		r.JSON(200, appPrefV4)
	} else if versionSelected == "2" {
		AnalyticsGetActivityStatus_v2(dmap, serialSelected, &appPref7S)
		r.JSON(200, appPref7S)
	} else if versionSelected == "5" {
		AnalyticsGetActivityStatus_v6(dmap, serialSelected, &appPrefV4)
		r.JSON(200, appPrefV4)
	} else if versionSelected == "6" {
		AnalyticsGetActivityStatus_v8(dmap, serialSelected, &appPrefV4)
		r.JSON(200, appPrefV4)
	}

}

func AnalyticsGetActivityStatus_v3(dmap *gorp.DbMap, serialSelected string, appPref *[]Tab7XAppPref) {

	/**
	Name	:	AnalyticsGetActivityStatus_v3
	Version :	1.0
	Description :
				 this function is dedicated to display X series activity_status information from database

	State	: Used
	*/

	if serialSelected != "" && serialSelected != "all" {
		queryWithSerial := "where serial like '%" + serialSelected + "'"
		_, err := dmap.Select(appPref, "select package_name AS package_name, sum(time_spent) AS total_time_spent, sum(launch_count) AS total_launch_count from activity_status "+queryWithSerial+"  and package_name not like 'com.cide.interactive.kurioLauncher'  group by package_name order by total_time_spent DESC LIMIT 50")
		checkErr(err, "Erreur select from view 7x_activity_status")
	} else {
		_, err := dmap.Select(appPref, "SELECT * FROM KurioS_.7x_activity_status  where package_name not like 'com.cide.interactive.kurioLauncher' order by total_time_spent DESC LIMIT 50")
		checkErr(err, "Erreur select from view 7x_activity_status")
	}
}

func AnalyticsGetActivityStatus_v4(dmap *gorp.DbMap, serialSelected string, appPref *[]TabV4AppPref) {

	/**
	Name	:	AnalyticsGetActivityStatus_v4
	Version :	1.0
	Description :
				 this function is dedicated to display Tab 2 2015 activity_status information from database

	State	: Used
	*/

	//serial2016 := " and rom_id not like '%40' "
	if serialSelected != "" && serialSelected != "all" {
		queryWithSerial := "where rom_id = '" + serialSelected + "'"
		_, err := dmap.Select(appPref, "select package_name AS package_name, sum(time_spent) AS total_time_spent, sum(launch_count) AS total_launch_count from activity_status_by_romid_v4  "+queryWithSerial+"  and package_name not like 'com.cide.interactive.kurioLauncher'   group by package_name order by total_time_spent DESC LIMIT 50")
		checkErr(err, "Erreur select from activity_status_v4")
	} else {
		_, err := dmap.Select(appPref, "SELECT * FROM KurioS_.activity_status_view_v4  where package_name not like 'com.cide.interactive.kurioLauncher' order by total_time_spent DESC LIMIT 50")
		checkErr(err, "Erreur select from view activity_status_view_v4")
	}
}

func AnalyticsGetActivityStatus_v6(dmap *gorp.DbMap, serialSelected string, appPref *[]TabV4AppPref) {

	/**
	Name	:	AnalyticsGetActivityStatus_v6
	Version :	1.0
	Description :
				 this function is dedicated to display Tab 2 2016 activity_status information from database

	State	: Used
	*/

	//serial2016 := " and rom_id like '%40' "
	if serialSelected != "" && serialSelected != "all" {
		queryWithSerial := "where rom_id = '" + serialSelected + "'"
		_, err := dmap.Select(appPref, "select package_name AS package_name, sum(time_spent) AS total_time_spent, sum(launch_count) AS total_launch_count from activity_status_by_romid_v4  "+queryWithSerial+"  and package_name not like 'com.cide.interactive.kurioLauncher'  group by package_name order by total_time_spent DESC LIMIT 50")
		checkErr(err, "Erreur select from activity_status_v4")
	} else {
		_, err := dmap.Select(appPref, "SELECT * FROM KurioS_.activity_status_2016_view_v4   where package_name not like 'com.cide.interactive.kurioLauncher' order by total_time_spent DESC LIMIT 50")
		checkErr(err, "Erreur select from view activity_status_view_v4")
	}
}

func AnalyticsGetActivityStatus_v8(dmap *gorp.DbMap, serialSelected string, appPref *[]TabV4AppPref) {

	/**
	Name	:	AnalyticsGetActivityStatus_v8
	Version :	1.0
	Description :
				 this function is dedicated to display Tab 2 2016 activity_status information from database

	State	: Used
	*/

	//	serial2017 := " and rom_id like '%60' or rom_id like '%61' "
	if serialSelected != "" && serialSelected != "all" {
		queryWithSerial := "where rom_id = '" + serialSelected + "'"
		_, err := dmap.Select(appPref, "select package_name AS package_name, sum(time_spent) AS total_time_spent, sum(launch_count) AS total_launch_count from activity_status_by_romid_v4  "+queryWithSerial+"   and package_name not like 'com.cide.interactive.kurioLauncher'  group by package_name order by total_time_spent DESC LIMIT 50")
		checkErr(err, "Erreur select from activity_status_v4")
	} else {
		_, err := dmap.Select(appPref, "SELECT * FROM KurioS_.activity_status_2017_view_v4  where package_name not like 'com.cide.interactive.kurioLauncher' order by total_time_spent DESC LIMIT 50")
		checkErr(err, "Erreur select from view activity_status_view_v4")
	}
}

func AnalyticsGetActivityStatus_v2(dmap *gorp.DbMap, serialSelected string, appPref *[]Tab7SAppPref) {

	/**
	Name	:	AnalyticsGetActivityStatus_v2
	Version :	1.0
	Description :
				 this function is dedicated to display S series activity_status information from database

	State	: Used
	*/

	if serialSelected != "" && serialSelected != "all" {
		queryWithSerial := "where rom_id = '" + serialSelected + "'"
		_, err := dmap.Select(appPref, "select package_name AS package_name, sum(time_spent) AS total_time_spent, sum(launch_count) AS total_launch_count from activity_status_by_romid_v4  "+queryWithSerial+"  and package_name not like 'com.cide.interactive.kurioLauncher' group by package_name order by total_time_spent DESC LIMIT 50")
		checkErr(err, "Erreur select from activity_status_v4")
	} else {
		_, err := dmap.Select(appPref, "SELECT * FROM KurioS_.activity_status_view_v4 where package_name not like 'com.cide.interactive.kurioLauncher' order by total_time_spent DESC LIMIT 50")
		checkErr(err, "Erreur select from view activity_status_view_v4")
	}
}

func AnalyticsGetAppStatic(params martini.Params, res http.ResponseWriter, req *http.Request, r render.Render, dmap *gorp.DbMap, session sessions.Session) {

	/**
	Name	:	AnalyticsGetAppStatic
	Version :	1.0
	Description :
				this function is dedicated to display  app statistics information from database

	State	: Used
	*/

	var appStatic_v3 []AppAnalytics_v3
	var appStatic_v4 []AppAnalytics_v4

	serialSelected := params["serial"]
	versionSelected := params["version"]
	typeSelected := params["type"]

	if typeSelected == "" {
		r.Error(404)
		return
	}

	if versionSelected == "3" {
		AnalyticsGetAppStatic_v3(dmap, serialSelected, typeSelected, &appStatic_v3)
		r.JSON(200, appStatic_v3)
	} else if versionSelected == "4" {
		AnalyticsGetAppStatic_v4(dmap, serialSelected, typeSelected, &appStatic_v4)
		r.JSON(200, appStatic_v4)
	}
}

func AnalyticsGetAppStatic_v3(dmap *gorp.DbMap, serialSelected, typeSelected string, appStatic *[]AppAnalytics_v3) {

	/**
	Name	:	AnalyticsGetAppStatic_v3
	Version :	1.0
	Description :
				 this function is dedicated to display  X series app statistics information from database

	State	: Used
	*/

	if serialSelected != "" && serialSelected != "all" {
		queryWithSerial := "where serial like '%" + serialSelected + "'"
		staticQuery := `select 
		serial as serial,
        package_name AS package_name,
        sum(install_count) AS install_count,
        sum(uninstall_count) AS uninstall_count,
        sum(restored_from_preload) AS restored_from_preload,
        sum(not_restored_from_preload) AS not_restored_from_preload,
        sum(installed_from_parental) AS installed_from_parental,
        sum(uninstalled_from_parental) AS uninstalled_from_parental,
        sum(new_app_installed_from_parental) AS new_app_installed_from_parental,
        sum(updated_from_parental) AS updated_from_parental
    	from app_analytics ` + queryWithSerial + ` group by package_name `
		_, err := dmap.Select(appStatic, staticQuery+" order by "+typeSelected+" DESC LIMIT 50")
		checkErr(err, "Erreur select from view 7x_app_static")
	} else {
		_, err := dmap.Select(appStatic, "SELECT * FROM KurioS_.7x_app_static order by "+typeSelected+" DESC LIMIT 50")
		checkErr(err, "Erreur select from view 7x_app_static")
	}
}

func AnalyticsGetAppStatic_v4(dmap *gorp.DbMap, serialSelected, typeSelected string, appStatic *[]AppAnalytics_v4) {

	/**
	Name	:	AnalyticsGetAppStatic_v4
	Version :	1.0
	Description :
				 this function is dedicated to display  Tab 2 app statistics information from database

	State	: Used
	*/

	if serialSelected != "" && serialSelected != "all" {
		queryWithSerial := "where serial like '%" + serialSelected + "'"
		staticQuery := `select 
		serial as serial,
        package_name AS package_name,
        sum(install_count) AS install_count,
        sum(uninstall_count) AS uninstall_count
    	from app_analytics_v4 ` + queryWithSerial + ` group by package_name `
		_, err := dmap.Select(appStatic, staticQuery+" order by "+typeSelected+" DESC LIMIT 50")
		checkErr(err, "Erreur select from app_analytics_v4")
	} else {
		_, err := dmap.Select(appStatic, "SELECT * FROM KurioS_.app_static_view_v4 order by "+typeSelected+" DESC LIMIT 50")
		checkErr(err, "Erreur select from view app_static_view_v4")
	}
}

func AnalyticsSearchFromPackageName(res http.ResponseWriter, req *http.Request, r render.Render, dmap *gorp.DbMap) {

	
}

func AnalyticsSearchFromPackageNameDate(res http.ResponseWriter, req *http.Request, r render.Render, dmap *gorp.DbMap) {
	
}

func AnalyticsSearchFromSerialMailMac(res http.ResponseWriter, req *http.Request, r render.Render, dmap *gorp.DbMap) {

	/**
	Name	:	AnalyticsSearchFromSerialMailMac
	Version :	1.0
	Description :
				 this function manage the search on the find device page ( extract_data_olivier) with information contained on the form
				 for X,S and Tab devices.

	State	: Used
	*/

	var serialResult []SerialAnalyticsv4
	curefSearch := req.FormValue("curefToSearch")
	serialSearch := req.FormValue("serialToSearch")
	emailSearch := req.FormValue("emailToSearch")
	macAdressSearch := req.FormValue("macAdressToSearch")
	tableToSearch := req.FormValue("tabToSearch")
	searchLike := req.FormValue("searchLike")

	if tableToSearch != "register" {
		query := "select * from " + tableToSearch + " where"

		if curefSearch != "" {
			switch searchLike {
			case "begin":
				log.Println(" curef begin")
				query = query + " curef like '" + curefSearch + "%'"
				break
			case "contain":
				log.Println(" curef contain")
				query = query + " curef like '%" + curefSearch + "%'"
				break
			default:
				query = query + " curef like '" + curefSearch + "%'"
				break
			}
		} else if serialSearch != "" {
			switch searchLike {
			case "begin":
				log.Println(" curef begin")
				query = query + " serial like '" + serialSearch + "%'"
				break
			case "contain":
				query = query + " serial like '%" + serialSearch + "%'"
				break
			default:
				query = query + " serial like '" + serialSearch + "%'"
				break
			}

		} else if emailSearch != "" {
			switch searchLike {
			case "begin":
				query = query + " parent_email like '" + emailSearch + "%'"
				break
			case "contain":
				query = query + " parent_email like '%" + emailSearch + "%'"
				break
			default:
				query = query + " parent_email like '" + emailSearch + "%'"
				break
			}

		} else if macAdressSearch != "" {
			switch searchLike {
			case "begin":
				query = query + " MAC_address like '" + macAdressSearch + "%'"
				break
			case "contain":
				query = query + " MAC_address = '%" + macAdressSearch + "%'"
				break
			default:
				query = query + " MAC_address like '" + macAdressSearch + "%'"
				break
			}

		}

		_, err := dmap.Select(&serialResult, query)
		checkErr(err, "erreur AnalyticsSearchFromSerialMailMac")

		r.JSON(200, serialResult)
		
	} else {
		log.Println(r)
		AnalyticsSearchFromSerialMailMacForOld7Tab(r, serialSearch, emailSearch, macAdressSearch, searchLike)
	}

}

func AnalyticsSearchFromSerialMailMacForOld7Tab(r render.Render, serialSearch string, emailSearch string, macAdressSearch string, searchLike string) {

	/**
	Name	:	AnalyticsSearchFromSerialMailMacForOld7Tab
	Version :	1.0
	Description :
		 this function manage the search on the find device page ( extract_data_olivier) with information contained on the form
		 for Old Tab.

	State	: Used
	*/

	var serialResult []SerialOld7Tab

	dmap := InitOldDb()

	query := "select * from register where"

	if serialSearch != "" {
		switch searchLike {
		case "begin":
			query = query + " serial_number like '" + serialSearch + "%'"
			break
		case "contain":
			query = query + " serial_number like '%" + serialSearch + "%'"
			break
		default:
			query = query + " serial_number like '" + serialSearch + "%'"
			break
		}

	} else if emailSearch != "" {
		switch searchLike {
		case "begin":
			query = query + " email like '" + emailSearch + "%'"
			break
		case "contain":
			query = query + " email like '%" + emailSearch + "%'"
			break
		default:
			query = query + " email like '" + emailSearch + "%'"
			break
		}

	} else if macAdressSearch != "" {
		switch searchLike {
		case "begin":
			query = query + " mac_address like '" + macAdressSearch + "%'"
			break
		case "contain":
			query = query + " mac_address like '%" + macAdressSearch + "%'"
			break
		default:
			query = query + " mac_address like '" + macAdressSearch + "%'"
			break
		}
	}
	log.Println(query)
	_, err := dmap.Select(&serialResult, query)
	checkErr(err, "erreur AnalyticsSearchFromSerialMailMacForOld7Tab")

	r.JSON(200, serialResult)
}

func InitOldDb() *gorp.DbMap {

	/**
	Name	:	InitOldDb
	Version :	1.0
	Description :
				This function is dedicated to establish the connection between the web sever and the database register which is used for old tab.

	State	: Used
	*/

	//db, err := sql.Open("mysql", "register:tG48R573jxuX3rVR@tcp(localhost:3306)/kdtablets_com_register")
	log.Println("initialisation OLD DB")
	db, err := sql.Open("mysql", "register:"+passwdRegister+"@tcp("+myip+":3306)/kdtablets_com_register")
	checkErrFatal(err, "sql.Open failed")

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	checkErr(err, "Create tables failed")
	return dbmap
}

func AnalyticsViewer(res http.ResponseWriter, req *http.Request, r render.Render, dmap *gorp.DbMap) {

	/**
	Name	:	AnalyticsViewer
	Version :	1.0
	Description :
				This function is dedicated to display the dashboard view.

	State	: Used
	*/

	if (idsession == "cidesei3") || (idsession == "Jeremy") || (idsession == "Caroline") {
		r.HTML(200, "analytics_viewer", nil)
	} else {
		http.Redirect(res, req, "/analytics_bad_user", 302)
	}
}



func SetCountry_v2(tab7SRomId []Tab7SByRomId) {

	/**
	Name	:	SetCountry_v2
	Version :	1.0
	Description :
				This function is dedicated to initialize the information of countries for S series.

	State	: Used
	*/

	for count, serial := range tab7SRomId {
		if serial.RomId == "001" {
			tab7SRomId[count].Country = "7S / UK-KDGB / 001"
		} else if serial.RomId == "002" {
			tab7SRomId[count].Country = "7S / US-KDUS / 002"
		} else if serial.RomId == "003" {
			tab7SRomId[count].Country = "7S / US-QVC-TSV / 003"
		} else if serial.RomId == "004" {
			tab7SRomId[count].Country = "7S / CA / 004"
		} else if serial.RomId == "005" {
			tab7SRomId[count].Country = "7S / DK-INTOY / 005"
		} else if serial.RomId == "010" {
			tab7SRomId[count].Country = "7S / FR Gulli / 010"
		} else if serial.RomId == "015" {
			tab7SRomId[count].Country = "7S / AU / 015"
		} else if serial.RomId == "020" {
			tab7SRomId[count].Country = "7S / CH-WALDMEIER / 020"
		} else if serial.RomId == "025" {
			tab7SRomId[count].Country = "7S / NL-ANGRYBIRD / 025"
		} else if serial.RomId == "026" {
			tab7SRomId[count].Country = "7S / NL-TELEKIDS / 026"
		} else if serial.RomId == "030" {
			tab7SRomId[count].Country = "7S / ES-CEFA / 030"
		} else if serial.RomId == "035" {
			tab7SRomId[count].Country = "7S / DE-TOGGOLINO / 035"
		} else if serial.RomId == "036" {
			tab7SRomId[count].Country = "7S / DE-TOGGO / 036"
		} else if serial.RomId == "037" {
			tab7SRomId[count].Country = "7S / DE-TOGGO-V2 / 036"
		} else if serial.RomId == "040" {
			tab7SRomId[count].Country = "7S / RUSSIAN / 040"
		} else if serial.RomId == "045" {
			tab7SRomId[count].Country = "7S / NL-KZOOM / 045"
		} else if serial.RomId == "046" {
			tab7SRomId[count].Country = "7S / MT / 046"
		} else if serial.RomId == "050" {
			tab7SRomId[count].Country = "7S / JP / 050"
		} else if serial.RomId == "055" {
			tab7SRomId[count].Country = "7S / AE-CartoonNetwork / 055"
		} else if serial.RomId == "501" {
			tab7SRomId[count].Country = "4S / UK / 501"
		} else if serial.RomId == "502" {
			tab7SRomId[count].Country = "4S / UK KURIO POCKET/ 502"
		} else if serial.RomId == "505" {
			tab7SRomId[count].Country = "4S / US / 505"
		} else if serial.RomId == "506" {
			tab7SRomId[count].Country = "4S / US-QVC-TSV / 506"
		} else if serial.RomId == "507" {
			tab7SRomId[count].Country = "4S / CA / 507"
		} else if serial.RomId == "508" {
			tab7SRomId[count].Country = "4S / CA SPECIAL EDITION / 508"
		} else if serial.RomId == "510" {
			tab7SRomId[count].Country = "4S / FR Gulli / 510"
		} else if serial.RomId == "515" {
			tab7SRomId[count].Country = "4S / AU / 515"
		} else if serial.RomId == "520" {
			tab7SRomId[count].Country = "4S / DK / 520"
		} else if serial.RomId == "530" {
			tab7SRomId[count].Country = "4S / NL / 530"
		} else if serial.RomId == "531" {
			tab7SRomId[count].Country = "4S / NL-TELEKIDS / 531"
		} else if serial.RomId == "540" {
			tab7SRomId[count].Country = "4S / ES-CEFA / 540"
		} else if serial.RomId == "555" {
			tab7SRomId[count].Country = "4S / AE / 555"
		} else if serial.RomId == "560" {
			tab7SRomId[count].Country = "4S / JP / 560"
		} else if serial.RomId == "561" {
			tab7SRomId[count].Country = "4S / JP-MEGAHOUSE / 501"
		} else if serial.RomId == "801" {
			tab7SRomId[count].Country = "10S / UK / 801"
		} else if serial.RomId == "802" {
			tab7SRomId[count].Country = "10S / UK-TAB XL / 802"
		} else if serial.RomId == "805" {
			tab7SRomId[count].Country = "10S / US / 805"
		} else if serial.RomId == "807" {
			tab7SRomId[count].Country = "10S / CA / 807"
		} else if serial.RomId == "810" {
			tab7SRomId[count].Country = "10S / DK / 810"
		} else if serial.RomId == "817" {
			tab7SRomId[count].Country = "10S / NL-RTL / 817"
		} else if serial.RomId == "820" {
			tab7SRomId[count].Country = "10S / FR-Gulli / 820"
		} else if serial.RomId == "825" {
			tab7SRomId[count].Country = "10S / ES-CEFA / 825"
		} else if serial.RomId == "835" {
			tab7SRomId[count].Country = "10S / DE-TAB XL / 835"
		}
	}
}

func SetCountry_v3(tab7XRomId []Tab7XByRomId) {

	/**
	Name	:	SetCountry_v3
	Version :	1.0
	Description :
				This function is dedicated to initialize the information of countries for X series.

	State	: Used
	*/

	for count, serial := range tab7XRomId {
		if serial.RomId == "101" {
			tab7XRomId[count].Country = "C14100 / UK / 101"
		} else if serial.RomId == "102" {
			tab7XRomId[count].Country = "C14100 / US / 102"
		} else if serial.RomId == "104" {
			tab7XRomId[count].Country = "C14100 / CA / 104"
		} else if serial.RomId == "105" {
			tab7XRomId[count].Country = "C14100 / DK / 105"
		} else if serial.RomId == "110" {
			tab7XRomId[count].Country = "C14100 / FR Gulli / 110"
		} else if serial.RomId == "111" {
			tab7XRomId[count].Country = "C14100 / FR Fnac / 111"
		} else if serial.RomId == "121" {
			tab7XRomId[count].Country = "C14100 / CH / 121"
		} else if serial.RomId == "126" {
			tab7XRomId[count].Country = "C14100 / NL / 126"
		} else if serial.RomId == "130" {
			tab7XRomId[count].Country = "C14100 / ES / 130"
		} else if serial.RomId == "135" {
			tab7XRomId[count].Country = "C14100 / DE / 135"
		} else if serial.RomId == "147" {
			tab7XRomId[count].Country = "C14100 / ZA / 147"
		} else if serial.RomId == "148" {
			tab7XRomId[count].Country = "C14100 / IN / 148"
		} else if serial.RomId == "149" {
			tab7XRomId[count].Country = "C14100 / NZ / 149"
		}
	}
}

func SetCountry_v4(tabV4RomId []TabV4ByRomId) {

	/**
	Name	:	SetCountry_v4
	Version :	1.0
	Description :
				This function is dedicated to initialize the information of countries for Tab 2 2015.

	State	: Used
	*/

	for count, serial := range tabV4RomId {
		if serial.RomId == "301" {
			tabV4RomId[count].Country = "C15100 / UK-GENERIC / 301"
		} else if serial.RomId == "302" {
			tabV4RomId[count].Country = "C15150 / US-GENERIC / 302"
		} else if serial.RomId == "304" {
			tabV4RomId[count].Country = "C15150 / CA / 304"
		} else if serial.RomId == "305" {
			tabV4RomId[count].Country = "C15100 / DK - INTOYS / 305"
		} else if serial.RomId == "310" {
			tabV4RomId[count].Country = "C15100 / FR-GULLI  / 310"
		} else if serial.RomId == "321" {
			tabV4RomId[count].Country = "C15100 / CH-WALDMEIER / 321"
		} else if serial.RomId == "326" {
			tabV4RomId[count].Country = "C15100 / NL-TELEKIDS / 326"
		} else if serial.RomId == "330" {
			tabV4RomId[count].Country = "C15100 / CEFA / 330"
		} else if serial.RomId == "335" {
			tabV4RomId[count].Country = "C15100 / TOGGO / 335"
		} else if serial.RomId == "336" {
			tabV4RomId[count].Country = "C15100 / TOGGOLINO / 336"
		} else if serial.RomId == "347" {
			tabV4RomId[count].Country = "C15100 / ZA / 347"
		} else if serial.RomId == "348" {
			tabV4RomId[count].Country = "C15100 / IN / 348"
		} else if serial.RomId == "355" {
			tabV4RomId[count].Country = "C15100 / AE CN / 355"
		} else if serial.RomId == "403" {
			tabV4RomId[count].Country = "C15150 / US-QVC / 403"
		} else if serial.RomId == "422" {
			tabV4RomId[count].Country = "C15150 / US-WALMART / 422"
		} else if serial.RomId == "E8GB" {
			tabV4RomId[count].Country = "C15450 / EE-KDGB / E8GB"
		} else if serial.RomId == "ALEU" {
			tabV4RomId[count].Country = "EE TEST / ALEU"
		} else if serial.RomId == "" {
			tabV4RomId[count].Country = "C15450 / EE-KDGB / E8GB"
		}
	}
}

func SetCountry_v6(tabV4RomId []TabV4ByRomId) {

	/**
	Name	:	SetCountry_v6
	Version :	1.0
	Description :
				This function is dedicated to initialize the information of countries for Tab 2 2016.

	State	: Used
	*/

	for count, serial := range tabV4RomId {
		if serial.RomId == "G40" {
			tabV4RomId[count].Country = "C15100 / UK-KDGB / G40"
		} else if serial.RomId == "U40" {
			tabV4RomId[count].Country = "C15100 / US-KDUS / U40"
		} else if serial.RomId == "S40" {
			tabV4RomId[count].Country = "C15100 / DK-INTOY / S40"
		} else if serial.RomId == "F40" {
			tabV4RomId[count].Country = "C15100 / FR-GULLI / F40"
		} else if serial.RomId == "W40" {
			tabV4RomId[count].Country = "C15100 / CH-WALDMEIER / W40"
		} else if serial.RomId == "N40" {
			tabV4RomId[count].Country = "C15100 / NL-TELEKIDS / N40"
		} else if serial.RomId == "E40" {
			tabV4RomId[count].Country = "C15100 / ES-CEFA / E40"
		} else if serial.RomId == "D40" {
			tabV4RomId[count].Country = "C15100 / DE-TOGGO / D40"
		} else if serial.RomId == "B40" {
			tabV4RomId[count].Country = "C15100 / BE-KDBNL / B40"
		}
	}
}

func SetCountry_v8(tabV4RomId []TabV4ByRomId) {

	/**
	Name	:	SetCountry_v8
	Version :	1.0
	Description :
				This function is dedicated to initialize the information of countries for Tab 2 2016.

	State	: Used
	*/

	for count, serial := range tabV4RomId {
		if serial.RomId == "G60" {
			tabV4RomId[count].Country = "C17150 / UK-KDGB 16Go/ G60"
		} else if serial.RomId == "G61" {
			tabV4RomId[count].Country = "C17110 / UK-KDGB 8Go / G61"
		} else if serial.RomId == "L60" {
			tabV4RomId[count].Country = "C17150 / ZA-LACEY / L60"
		} else if serial.RomId == "U60" {
			tabV4RomId[count].Country = "C17150 / US-KDUS / U60"
		} else if serial.RomId == "S60" {
			tabV4RomId[count].Country = "C17150 / DK-INTOY / S60"
		} else if serial.RomId == "F60" {
			tabV4RomId[count].Country = "C17150 / FR-GULLI / F60"
		} else if serial.RomId == "W60" {
			tabV4RomId[count].Country = "C17150 / CH-WALDMEIER / W60"
		} else if serial.RomId == "N60" {
			tabV4RomId[count].Country = "C17150 / NL-TELEKIDS / N60"
		} else if serial.RomId == "E60" {
			tabV4RomId[count].Country = "C17150 / ES-CEFA / E60"
		} else if serial.RomId == "D60" {
			tabV4RomId[count].Country = "C17150 / DE-TOGGO / D60"
		} else if serial.RomId == "B60" {
			tabV4RomId[count].Country = "C17150 / BE-KDBNL / B60"
		}
	}
}
