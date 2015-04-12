package models

import (
	"encoding/json"
	"log"
)

type Schedule struct {
	Name    string `json:"name,omitempty"`
	Workout string `json:"workout"`
	Day     string `json:"day"`
}

type scheduleResult struct {
	Name      string      `json:"name"`
	Schedules []*Schedule `json:"schedules"`
}

const getSchedulesQuery = "SELECT name, workout, day FROM schedules WHERE userid = ?"
const addSchedulesQuery = "INSERT INTO schedules (userid, name, workout, day) VALUES (?, ?, ?, ?)"
const removeSchedulesQuery = "DELETE FROM schedules WHERE userid=? AND name=?"

func GetSchedules(id string) []*scheduleResult {

	stmt, err := db.Prepare(getSchedulesQuery)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := stmt.Query(id)

	m := make(map[string][]*Schedule)

	var name, workout, day string

	for rows.Next() {
		err = rows.Scan(&name, &workout, &day)
		if err != nil {
			log.Fatal(err)
		}
		m[name] = append(m[name], &Schedule{"", workout, day})
	}

	result := make([]*scheduleResult, 0, 10)

	for k, _ := range m {
		result = append(result, &scheduleResult{k, m[k]})
	}

	return result
}

func AddSchedule(id string, data string) {
	stmt, err := db.Prepare(addSchedulesQuery)
	if err != nil {
		log.Fatal(err)
	}
	var ex Schedule
	err = json.Unmarshal([]byte(data), &ex)

	if len(ex.Name) > 0 && len(ex.Workout) > 0 && len(ex.Day) > 0 {
		stmt.Exec(id, ex.Name, ex.Workout, ex.Day)
	}
}
