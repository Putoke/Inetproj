package models

import (
    "log"
    "encoding/json"
)

type Workout struct {
    Name string `json:"name,omitempty"`
	Exercise string `json:"exercise"`
	Sets     string `json:"sets"`
	Reps     string `json:"reps"`
}

type workoutResult struct {
    Name string `json:"name"`
    Exercises []*Workout `json:"exercises"`
}

const getWorkoutsQuery = "SELECT name, exercise, sets, reps FROM workouts WHERE userid = ?"
const addWorkoutQuery = "INSERT INTO workouts (userid, name, exercise, sets, reps) VALUES (?, ?, ?, ?, ?)"
const removeWorkoutQuery = "DELETE FROM workouts WHERE userid=? AND name=?"

func GetWorkouts(id string) []*workoutResult {

	stmt, err := db.Prepare(getWorkoutsQuery)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := stmt.Query(id)

    m := make(map[string][]*Workout)

	var name, workout, sets, reps string

	for rows.Next() {
		err = rows.Scan(&name, &workout, &sets, &reps)
		if err != nil {
			log.Fatal(err)
		}

        m[name] = append(m[name], &Workout{"", workout, sets, reps})

	}

    result := make([]*workoutResult,0, 10)

    for k, _ := range m {
        result = append(result, &workoutResult{k, m[k]})
    }

	return result
}


func AddWorkout(id string, data string) {
    stmt, err := db.Prepare(addWorkoutQuery);
    if err != nil {
        log.Fatal(err)
    }
    var ex Workout
    err = json.Unmarshal([]byte(data), &ex)

    if len(ex.Name) > 0 && len(ex.Exercise) > 0 {
        stmt.Exec(id, ex.Name, ex.Exercise, ex.Reps, ex.Sets)
    }
}

func RemoveWorkout(id string, data string) {

    stmt, err := db.Prepare(removeWorkoutQuery);
    if err != nil {
        log.Fatal(err)
    }

    var ex Workout
    err = json.Unmarshal([]byte(data), &ex)

    stmt.Exec(id, ex.Name)
}
