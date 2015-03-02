package main

import (
	"Inetproj/controllers"
	"Inetproj/models"
	"log"
	"net/http"
    "os"
)

func main() {
	router := controllers.NewRouter()
    models.InitDB()
	log.Println("Server started")
    router.HandleFunc("/kill", func(w http.ResponseWriter, r * http.Request) {os.Exit(0)})
 
	log.Fatal(http.ListenAndServe(":8001", router))
	models.CloseDB()
}

