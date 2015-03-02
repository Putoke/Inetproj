package main

import (
	"Inetproj/controllers"
	"Inetproj/models"
	"log"
	"net/http"
)

func main() {
	router := controllers.NewRouter()
	models.InitDB()

	log.Fatal(http.ListenAndServe(":8000", router))
	models.CloseDB()
}
