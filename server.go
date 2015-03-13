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
	router.HandleFunc("/kill", func(w http.ResponseWriter, r *http.Request) { os.Exit(0) })

	log.Fatal(http.ListenAndServe(":8000", router))
	models.CloseDB()
}
