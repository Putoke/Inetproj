package models

import "log"

type Schedule struct {
	Userid  string `json:"userid"`
	Name    string `json:"name"`
	Workout string `json:"workout"`
	Day     string `json:"day"`
}

const getSchedulesQuery = "SELECT name, workout, day FROM schedules WHERE userid = ?"

func GetSchedules(id string) []*Schedule {

	stmt, err := db.Prepare(getSchedulesQuery)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := stmt.Query(id)

	schedules := make([]*Schedule, 0, 10)

	var name, workout, day string

	for rows.Next() {
		err = rows.Scan(&name, &workout, &day)
		if err != nil {
			log.Fatal(err)
		}
		schedules = append(schedules, &Schedule{id, name, workout, day})
	}

	return schedules
}
