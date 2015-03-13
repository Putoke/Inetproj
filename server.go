package main

import (
	"Inetproj/controllers"
	"Inetproj/models"
	"log"
	"net/http"
	"os"
    "Inetproj/config"
)

func main() {

    config.InitConfig()
    models.InitDB()
	router := controllers.NewRouter()
	log.Println("Server started")

    http.HandleFunc("/kill",  func(w http.ResponseWriter, r *http.Request) {
        log.Println("Received shutdown request from " + r.RemoteAddr + ". Exiting")
        os.Exit(0)
    })
    go http.ListenAndServe("127.0.0.1:8001", nil)

	log.Fatal(http.ListenAndServe(":8000", router))
	models.CloseDB()
}
