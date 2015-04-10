package models

import (
    "log"
)

type Workout struct {
	//Name     string `json:"name"`
	Exercise string `json:"exercise"`
	Sets     string `json:"sets"`
	Reps     string `json:"reps"`
}

type workoutResult struct {
    Name string `json:"name"`
    Exercises []*Workout `json:"exercises"`
}

const getWorkoutsQuery = "SELECT name, exercise, sets, reps FROM workouts WHERE userid = ?"

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

        m[name] = append(m[name], &Workout{workout, sets, reps})

	}

    result := make([]*workoutResult,0, 10)

    for k, _ := range m {
        result = append(result, &workoutResult{k, m[k]})
    }




	return result
}
