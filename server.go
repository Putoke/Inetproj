package main
import (
    "Inetproj/controllers"
    "log"
    "net/http"
    "Inetproj/models"
    "github.com/goji/httpauth"
)

func main() {
    router := controllers.NewRouter()
    models.InitDB()

    http.Handle("/", httpauth.SimpleBasicAuth("test", "korv")(router))

    log.Fatal(http.ListenAndServe(":8000", nil))
    models.CloseDB()
}