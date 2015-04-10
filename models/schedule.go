package models

import "log"

type Schedule struct {
    Name string `json:"name,omitempty"`
	Workout string `json:"workout"`
	Day     string `json:"day"`
}

type scheduleResult struct {
    Name string `json:"name"`
    Schedules []*Schedule `json:"schedules"`
}

const getSchedulesQuery = "SELECT name, workout, day FROM schedules WHERE userid = ?"

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

    result := make([]*scheduleResult,0, 10)

    for k, _ := range m {
        result = append(result, &scheduleResult{k, m[k]})
    }

	return result
}
