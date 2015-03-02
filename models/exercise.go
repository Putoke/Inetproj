package models

import (
	"log"
)

type Exercise struct {
	Userid      string `json:"userid"`
	Name        string `json:"name"`
	Musclegroup string `json:"musclegroup"`
}

const GetExercisesQuery = "SELECT name, musclegroup FROM exercises WHERE userid = ?"

func GetExercises(id string) []*Exercise {

	stmt, err := db.Prepare(GetExercisesQuery)
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
		exercises = append(exercises, &Exercise{id, name, musclegroup})
	}

	return exercises
}
