package controllers

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes {
	Route {
		"exercises",
		"GET",
		"/exercises/{id}",
		Exercises,
	},

    Route {
        "workouts",
        "GET",
        "/workouts/{id}",
        Workouts,
    },

    Route {
        "schedules",
        "GET",
        "/schedules/{id}",
        Schedules,
    },
}
