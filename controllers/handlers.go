package controllers

import (
	"Inetproj/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Exercises(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	log.Println("Get exercises called by userid: " + id + " (" + r.RemoteAddr + ")")
	exercises := models.GetExercises(id)
	json.NewEncoder(w).Encode(exercises)
}
