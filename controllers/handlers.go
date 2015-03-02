package models

import (
	"encoding/json"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	exercises := GetExercises("0")
	json.NewEncoder(w).Encode(exercises)
}
