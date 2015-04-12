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
		"exercisesAdd",
		"POST",
		"/exercises/add",
		Use(ExercisesAdd, RequireLogin),
	},

	Route{
		"exercisesAdd",
		"POST",
		"/exercises/remove",
		Use(ExerciseRemove, RequireLogin),
	},

	Route{
		"workouts",
		"GET",
		"/workouts",
		Use(Workouts, RequireLogin),
	},

	Route{
		"workoutsDefault",
		"GET",
		"/workouts/default",
		Use(WorkoutsDefault, RequireLogin),
	},

	Route{
		"workoutsAdd",
		"POST",
		"/workouts/add",
		Use(WorkoutsAdd, RequireLogin),
	},

	Route{
		"workoutsRemove",
		"POST",
		"/workouts/remove",
		Use(WorkoutsRemove, RequireLogin),
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
        "schedulesAdd",
        "POST",
        "/schedules/add",
        Use(SchedulesAdd, RequireLogin),
    },

    Route{
        "schedulesRemove",
        "POST",
        "/schedules/remove",
        Use(SchedulesRemove, RequireLogin),
    },

	Route{
		"userinfo",
		"GET",
		"/user/info",
		Use(UserInfo, RequireLogin),
	},

	Route{
		"registeruser",
		"POST",
		//"/user/register/{name}/{lastname}/{email}/{password}",
        "/register",
		RegisterUser,
	},

	Route{
		"login",
		"GET",
		"/login/{email}/{password}",
		Login,
	},

	Route{
		"login",
		"POST",
		"/login",
		LoginPost,
	},
}
