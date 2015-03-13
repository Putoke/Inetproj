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
        "/exercises",
        Use(Exercises, RequireLogin),
    },

	Route{
		"exercises_default",
		"GET",
		"/exercises/default",
		Use(ExercisesDefault, RequireLogin),
	},

    Route{
        "workouts",
        "GET",
        "/workouts",
        Use(Workouts, RequireLogin),
    },

    Route {
        "workouts_default",
        "GET",
        "/workouts/default",
        Use(WorkoutsDefault, RequireLogin),
    },

	Route{
		"schedules",
		"GET",
		"/schedules",
		Use(Schedules, RequireLogin),
	},

    Route{
        "schedules_default",
        "GET",
        "/schedules/default",
        Use(SchedulesDefault, RequireLogin),
    },

    Route{
        "userinfo",
        "GET",
        "/user/info",
        Use(UserInfo, RequireLogin),
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
		Login,
	},
}


