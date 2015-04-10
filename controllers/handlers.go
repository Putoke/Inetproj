package controllers

import (
	"Inetproj/models"
	"encoding/json"
    ctx "github.com/gorilla/context"
	"log"
	"net/http"
    "runtime"
    "github.com/gorilla/mux"
)



func Exercises(w http.ResponseWriter, r *http.Request) {
    id, email := getIDAndEmail(r)

    printHandlerLog(id, email, r)
    exercises := models.GetExercises(id)
    json.NewEncoder(w).Encode(exercises)
}

func ExercisesDefault(w http.ResponseWriter, r * http.Request) {
    id, email := getIDAndEmail(r)

    printHandlerLog(id, email, r)
    exercises := models.GetExercises("0")
    json.NewEncoder(w).Encode(exercises)
}

func ExercisesAdd(w http.ResponseWriter, r * http.Request) {
    id, email := getIDAndEmail(r);
    printHandlerLog(id, email, r );
    json := mux.Vars(r)["json"]
    models.AddExercise(id, json)
}

func ExerciseRemove(w http.ResponseWriter, r * http.Request) {
    id, email := getIDAndEmail(r);
    printHandlerLog(id, email, r );
    json := mux.Vars(r)["json"]
    models.RemoveExercise(id, json)

}

func Workouts(w http.ResponseWriter, r *http.Request) {
    id, email := getIDAndEmail(r)

    printHandlerLog(id, email, r)
    workouts := models.GetWorkouts(id)
    json.NewEncoder(w).Encode(workouts)
}

func WorkoutsDefault(w http.ResponseWriter, r * http.Request) {
    id, email := getIDAndEmail(r)

    printHandlerLog(id, email, r)
    workouts := models.GetWorkouts("0")
    json.NewEncoder(w).Encode(workouts)
}


func Schedules(w http.ResponseWriter, r *http.Request) {
	id, email := getIDAndEmail(r)

    printHandlerLog(id, email, r)
	schedules := models.GetSchedules(id)
	json.NewEncoder(w).Encode(schedules)
}

func SchedulesDefault(w http.ResponseWriter, r * http.Request) {
    id, email := getIDAndEmail(r)

    printHandlerLog(id, email, r)
    schedules := models.GetSchedules("0")
    json.NewEncoder(w).Encode(schedules)

}

func UserInfo(w http.ResponseWriter, r * http.Request) {
    id, email := getIDAndEmail(r)
    user := models.GetUserByEmail(email)

    printHandlerLog(id, email, r)
    json.NewEncoder(w).Encode(user)
}

func getIDAndEmail(r * http.Request) (id, email string) {
    id = ctx.Get(r, "id").(string)
    email = ctx.Get(r, "email").(string)
    return id, email
}

func printHandlerLog(id string, email string, r * http.Request) {

    pc, _, _, _ := runtime.Caller(1)
    fname := runtime.FuncForPC(pc).Name()
    //fname := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
    log.Println(fname + " called by email=" + email + ", id=" + id + ", client=" + r.RemoteAddr)
}

