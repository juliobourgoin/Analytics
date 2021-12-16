package main

import (
	"database/sql"
	"github.com/lioonel/gorp"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"

	"crypto/aes"
	"crypto/cipher"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
)

var idsession string = ""
var root_directory string = ""

func CheckSession(w http.ResponseWriter, res *http.Request, session sessions.Session, r render.Render) {

	/**
	Name	:	CheckSession
	Version :	1.0
	Description :
				This function help to check if a session ring is on place.

	State	: Used
	*/

	value := session.Get("kurioPwd")
	if res.RequestURI != "/login" {
		log.Println(value)
		log.Println(res.RequestURI)
		//30 min
		session.Options(sessions.Options{
			MaxAge: 60 * 30,
		})
		if value == "1u05DX155!>S17H" { // All
			session.Set("romid", "")
		} else if value == "EY#~74kx)^R74:R" {
			session.Set("romid", "101")
		} else if value == "{(</<K54}q8o2[i" {
			session.Set("romid", "102")
		} else if value == "e99c9+oRu0]82je" {
			session.Set("romid", "104")
		} else if value == "JU000-[xKrx7]!i" {
			session.Set("romid", "105")
		} else if value == "7-5P9GQLiae>ISd" {
			session.Set("romid", "110")
		} else if value == "1kw:+)?4c1@&:=O" {
			session.Set("romid", "111")
		} else if value == "x7dLvo_p&SmD2jd" {
			session.Set("romid", "121")
		} else if value == "eHe(wXy0138b.Wz" {
			session.Set("romid", "126")
		} else if value == "6782RV47^0G176h" {
			session.Set("romid", "130")
		} else if value == "6BOBCPj)ggMZ<fm" {
			session.Set("romid", "135")
		}
	}
}
func getKurioSPassword() string {

	/**
	Name	:	getKurioSPassword
	Version :	1.0
	Description :
				This function retrieve the encrypted password of the KurioS_ Database.

	State	: Used
	*/

	key := "opensesame123456opensesame123456" // 32 bytes!

	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		panic(err)
	}
	ciphertext := []byte("abcdef1234567890")
	iv := ciphertext[:aes.BlockSize] // const BlockSize = 32
	b, err := ioutil.ReadFile("templates/kurio.pwd")
	if err != nil {
		panic(err)
	}

	decrypter := cipher.NewCFBDecrypter(block, iv) // simple!
	decrypted := make([]byte, 16)
	decrypter.XORKeyStream(decrypted, b)

	return string(decrypted)
}
func thelocalip() string {

	msg := ""
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			msg = ipv4.String()
		}
	}
	log.Println("c est l adresse" + msg)
	return msg
}

func getRegisterpassword() string {

	/**
	Name	:	getRegisterpassword
	Version :	1.0
	Description :
				This function retrieve the encrypted password of the register Database.

	State	: Used
	*/

	key := "opensesame123456opensesame123456" // 16 bytes!

	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		panic(err)
	}
	ciphertext := []byte("abcdef1234567890")
	iv := ciphertext[:aes.BlockSize] // const BlockSize = 16
	b, err := ioutil.ReadFile("templates/register.pwd")
	if err != nil {
		panic(err)
	}

	decrypter := cipher.NewCFBDecrypter(block, iv) // simple!
	decrypted := make([]byte, 16)
	decrypter.XORKeyStream(decrypted, b)

	return string(decrypted)
}

func initDb() *gorp.DbMap {

	/**
	Name	:	initDb
	Version :	1.0
	Description :
				This function is dedicated to establish the connection between the web sever and the database.

	State	: Used
	*/

	db, err := sql.Open("mysql", "KurioS_:"+passwdKurioS+"@tcp("+myip+":3306)/KurioS_")
	checkErrFatal(err, "sql.Open failed")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	checkErr(err, "Create tables failed")
	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Println(msg, err)
	}
}

func checkErrFatal(err error, msg string) {
	if err != nil {
		log.Panicln(msg, err)
	}
}

func activateDBTrace(dbmap *gorp.DbMap) {
	dbmap.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds))
}

func deactivateDBTrace(dbmap *gorp.DbMap) {
	dbmap.TraceOff()
}
