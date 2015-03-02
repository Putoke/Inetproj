package models

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

var db * sql.DB

var dbuser string = "frebernuser"
var dbpwd string = "itn6Fy9E"
var dbhost string = "mysql-vt2015.csc.kth.se"
var dbport string = "3306"
var dbname string = "frebern"
var dbprot string = "tcp"



func InitDB() {
    var err error
    db, err = sql.Open("mysql", dbuser + ":" +dbpwd + "@" + dbprot + "(" + dbhost + ":" +dbport + ")/" + dbname)

    if err != nil {
        log.Fatal(err)
    }

    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

}

func CloseDB() {
    defer db.Close()
}