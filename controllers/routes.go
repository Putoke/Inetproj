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

var routes = Routes{
	Route{
		"exercises",
		"GET",
		"/exercises/{id}",
		Use(Exercises, RequireLogin),
	},

    Route{
        "ex",
        "GET",
        "/ex",
        Use(ExercisesTest, RequireLogin),
    },

	Route{
		"workouts",
		"GET",
		"/workouts/{id}",
		Workouts,
	},

	Route{
		"schedules",
		"GET",
		"/schedules/{id}",
		Schedules,
	},

	Route{
		"registeruser",
		"GET",
		"/user/register/{name}/{lastname}/{email}/{password}",
		RegisterUser,
	},

	Route{
		"login",
		"GET",
		"/login/{email}/{password}",
		Logintest,
	},
}
