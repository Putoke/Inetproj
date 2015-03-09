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

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	lastname := vars["lastname"]
	email := vars["email"]
	password := vars["password"]
	success := models.RegisterUser(name, lastname, email, password)
	if success {
		log.Println("New user registered (" + name + ", " + lastname + ")")
		w.Write([]byte("User registered"))
	}
}

func Login(w http.ResponseWriter, r *http.Request) {

	// https://github.com/iamjem/go-passwordless-demo
}
