package controllers

import (
	"Inetproj/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func Exercises(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	exercises := models.GetExercises(id)
	json.NewEncoder(w).Encode(exercises)
}
