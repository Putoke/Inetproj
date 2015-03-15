package controllers

import (
	"Inetproj/config"
	"Inetproj/models"
	//"fmt"
	"github.com/dchest/authcookie"
	ctx "github.com/gorilla/context"
	"net/http"
)

func RequireLogin(handler http.Handler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

        if cookie, ok := requestHasValidSession(r); ok {
            email, _, _ := authcookie.Parse(cookie.Value, []byte(config.Values.AuthSecret))
            user := models.GetUserByEmail(email)
            ctx.Set(r, "email", user.Email)
            ctx.Set(r, "id", user.Id)
            handler.ServeHTTP(w, r)
        } else {
            ctx.Set(r, "email", nil)
            ctx.Set(r, "id", nil)
            SendStatusJSON(w, "", http.StatusForbidden)
        }

	}

}
