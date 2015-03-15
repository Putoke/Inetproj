package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
    "Inetproj/config"
)

var db *sql.DB

var tmp, _ = ioutil.ReadFile("database/user")
var dbuser string = string(tmp)
var tmp2, _ = ioutil.ReadFile("database/userpassword")
var dbpwd string = string(tmp2)
var dbhost string = "mysql-vt2015.csc.kth.se"
var dbport string = "3306"
var dbname string = "jedlu"
var dbprot string = "tcp"

func InitDB() {

	var err error
	db, err = sql.Open("mysql", config.Values.User+":"+config.Values.Userpwd+"@"+dbprot+"("+dbhost+":"+dbport+")/"+dbname)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

    log.Println("Successfully connected to database")

}

func CloseDB() {
	db.Close()
}
