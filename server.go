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


    http.HandleFunc("/kill",  func(w http.ResponseWriter, r *http.Request) {
        log.Println("Received shutdown request from " + r.RemoteAddr + ". Exiting")
    })
    go http.ListenAndServe(config.Values.ShutdownAddress+":"+config.Values.ShutdownPort, nil)

    log.Println("Server started listening on "+ config.Values.ListeningAddress+":"+config.Values.ListeningPort)
	log.Fatal(http.ListenAndServe(config.Values.ListeningAddress+":"+config.Values.ListeningPort, router))
	models.CloseDB()
}
