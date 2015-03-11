package controllers

import (
	"Inetproj/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Exercises(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	log.Println("Get exercises called by userid: " + id + " (" + r.RemoteAddr + ")")
	exercises := models.GetExercises(id)
	json.NewEncoder(w).Encode(exercises)
}

func Workouts(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	log.Println("Get workouts called by userid: " + id + " (" + r.RemoteAddr + ")")
	workouts := models.GetWorkouts(id)
	json.NewEncoder(w).Encode(workouts)
}

func Schedules(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	log.Println("Get schedules called by userid: " + id + " (" + r.RemoteAddr + ")")
	schedules := models.GetSchedules(id)
	json.NewEncoder(w).Encode(schedules)
}



