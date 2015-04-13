package controllers

import (
	"Inetproj/models"
	"encoding/json"
	ctx "github.com/gorilla/context"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
)

/*
*   EXERCISES
 */

func Exercises(w http.ResponseWriter, r *http.Request) {
	id, email := getIDAndEmail(r)

	printHandlerLog(id, email, r)
	exercises := models.GetExercises(id)
	json.NewEncoder(w).Encode(exercises)
}

func ExercisesDefault(w http.ResponseWriter, r *http.Request) {
	id, email := getIDAndEmail(r)

	printHandlerLog(id, email, r)
	exercises := models.GetExercises("0")
	json.NewEncoder(w).Encode(exercises)

}

func ExercisesAdd(w http.ResponseWriter, r *http.Request) {
	id, email := getIDAndEmail(r)
	printHandlerLog(id, email, r)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	models.AddExercise(id, string(body[:]))
	SendHTTPStatusJSON(w, http.StatusOK)
}

func ExerciseRemove(w http.ResponseWriter, r *http.Request) {
	id, email := getIDAndEmail(r)
	printHandlerLog(id, email, r)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	models.RemoveExercise(id, string(body[:]))
	SendHTTPStatusJSON(w, http.StatusOK)

}

/*
*   WORKOUTS
 */

func Workouts(w http.ResponseWriter, r *http.Request) {
	id, email := getIDAndEmail(r)

	printHandlerLog(id, email, r)
	workouts := models.GetWorkouts(id)
	json.NewEncoder(w).Encode(workouts)
}

func WorkoutsDefault(w http.ResponseWriter, r *http.Request) {
	id, email := getIDAndEmail(r)

	printHandlerLog(id, email, r)
	workouts := models.GetWorkouts("0")
	json.NewEncoder(w).Encode(workouts)
}

func WorkoutsAdd(w http.ResponseWriter, r *http.Request) {
	id, email := getIDAndEmail(r)
	printHandlerLog(id, email, r)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	models.AddWorkout(id, string(body[:]))
	SendHTTPStatusJSON(w, http.StatusOK)
}

func WorkoutsRemove(w http.ResponseWriter, r *http.Request) {
	id, email := getIDAndEmail(r)
	printHandlerLog(id, email, r)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	models.RemoveWorkout(id, string(body[:]))
	SendHTTPStatusJSON(w, http.StatusOK)

}

/*
*   SCHEDULES
 */

func Schedules(w http.ResponseWriter, r *http.Request) {
	id, email := getIDAndEmail(r)

	printHandlerLog(id, email, r)
	schedules := models.GetSchedules(id)
	json.NewEncoder(w).Encode(schedules)
}

func SchedulesDefault(w http.ResponseWriter, r *http.Request) {
	id, email := getIDAndEmail(r)

	printHandlerLog(id, email, r)
	schedules := models.GetSchedules("0")
	json.NewEncoder(w).Encode(schedules)

}

func SchedulesAdd(w http.ResponseWriter, r *http.Request) {
	id, email := getIDAndEmail(r)
	printHandlerLog(id, email, r)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	models.AddSchedule(id, string(body[:]))
	SendHTTPStatusJSON(w, http.StatusOK)
}

func SchedulesRemove(w http.ResponseWriter, r * http.Request) {
    id, email := getIDAndEmail(r)
    printHandlerLog(id, email, r)

    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        log.Fatal(err)
    }

    models.RemoveSchedule(id, string(body[:]))
    SendHTTPStatusJSON(w, http.StatusOK)
}

/*
*   MISC
 */

func UserInfo(w http.ResponseWriter, r *http.Request) {
	id, email := getIDAndEmail(r)
	user := models.GetUserByEmail(email)

	printHandlerLog(id, email, r)
	json.NewEncoder(w).Encode(user)
}

func getIDAndEmail(r *http.Request) (id, email string) {

    if r != nil {
        id = ctx.Get(r, "id").(string)
        email = ctx.Get(r, "email").(string)
    } else {
        log.Println("getIDAndEmail - nil request")
    }

	return id, email
}

func printHandlerLog(id string, email string, r *http.Request) {

	pc, _, _, _ := runtime.Caller(1)
	fname := runtime.FuncForPC(pc).Name()
	//fname := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	log.Println(fname + " called by email=" + email + ", id=" + id + ", client=" + r.RemoteAddr)
}
