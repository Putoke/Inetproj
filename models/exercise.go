package models

import (
	"log"
    "encoding/json"
)

type Exercise struct {
	//Userid      string `json:"userid"`
	Name        string `json:"name"`
	Musclegroup string `json:"musclegroup"`
}


const getExercisesQuery = "SELECT name, musclegroup FROM exercises WHERE userid = ?"
const addExerciseQuery = "INSERT INTO exercises (userid, name, musclegroup) VALUES (?, ?, ?)"
const removeExerciseQuery = "DELETE FROM exercises WHERE userid=? AND name=?"

func GetExercises(id string) []*Exercise {

	stmt, err := db.Prepare(getExercisesQuery)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := stmt.Query(id)

	exercises := make([]*Exercise, 0, 10)

	var name, musclegroup string

	for rows.Next() {
		err = rows.Scan(&name, &musclegroup)
		if err != nil {
			log.Fatal(err)
		}
		exercises = append(exercises, &Exercise{name, musclegroup})
	}

	return exercises
}

func AddExercise(id string, data string) {
    stmt, err := db.Prepare(addExerciseQuery);
    if err != nil {
        log.Fatal(err)
    }
    var ex Exercise
    err = json.Unmarshal([]byte(data), &ex)

    if len(ex.Name) > 0 {
        stmt.Exec(id, ex.Name, ex.Musclegroup)
    }


}

func RemoveExercise(id string, data string) {

    stmt, err := db.Prepare(removeExerciseQuery);
    if err != nil {
        log.Fatal(err)
    }

    var ex Exercise
    err = json.Unmarshal([]byte(data), &ex)

    stmt.Exec(id, ex.Name)
}