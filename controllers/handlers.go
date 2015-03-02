package controllers

import (
	"encoding/json"
	"net/http"
    "Inetproj/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	exercises := models.GetExercises("0")
	json.NewEncoder(w).Encode(exercises)
}
