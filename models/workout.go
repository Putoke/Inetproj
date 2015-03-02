package models
import "log"

type Workout struct {
    Userid      string `json:"userid"`
    Name        string `json:"name"`
    Exercise string `json:"exercise"`
    Sets string `json:"sets"`
    Reps string `json:"reps"`
}

const getWorkoutsQuery = "SELECT name, exercise, sets, reps FROM workouts WHERE userid = ?"

func GetWorkouts(id string) []*Workout {

    stmt, err := db.Prepare(getWorkoutsQuery)
    if err != nil {
        log.Fatal(err)
    }
    rows, err := stmt.Query(id)

    workouts := make([]*Workout, 0, 10)

    var name, workout, sets, reps string

    for rows.Next() {
        err = rows.Scan(&name, &workout, &sets, &reps)
        if err != nil {
            log.Fatal(err)
        }
        workouts = append(workouts, &Workout{id, name, workout, sets, reps})
    }

    return workouts
}