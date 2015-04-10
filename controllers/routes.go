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

var routes = []Route{

    Route{
        "exercises",
        "GET",
        "/exercises",
        Use(Exercises, RequireLogin),
    },

	Route{
		"exercisesDefault",
		"GET",
		"/exercises/default",
		Use(ExercisesDefault, RequireLogin),
	},

    Route{
        "schedulesAdd",
        "GET",
        "/exercises/add/{json}",
        Use(ExercisesAdd, RequireLogin),
    },

    Route{
        "workouts",
        "GET",
        "/workouts",
        Use(Workouts, RequireLogin),
    },

    Route {
        "workoutsDefault",
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
        "schedulesDefault",
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

    Route {
        "login",
        "POST",
        "/login",
        LoginPost,
    },
}


