package controllers

import (
	"Inetproj/models"
	"encoding/json"
	"github.com/gorilla/mux"
    ctx "github.com/gorilla/context"
	"log"
	"net/http"
    "fmt"
)

func Exercises(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	log.Println("Get exercises called by userid: " + id + " (" + r.RemoteAddr + ")")
	exercises := models.GetExercises(id)
	json.NewEncoder(w).Encode(exercises)
}

func ExercisesTest(w http.ResponseWriter, r *http.Request) {
    id := ctx.Get(r, "id").(string)
    email := ctx.Get(r, "email").(string)

    log.Println("Get exercises called by userid: " + id + " (" + r.RemoteAddr + ")")
    exercises := models.GetExercises(id)
    json.NewEncoder(w).Encode(exercises)
    fmt.Fprintln(w, email + " " + id)
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



