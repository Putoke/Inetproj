package controllers

import (
	"github.com/dchest/authcookie"
    ctx "github.com/gorilla/context"
	"net/http"
    "Inetproj/models"
    "fmt"
    "Inetproj/config"
)

func RequireLogin(handler http.Handler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		cookie, _ := r.Cookie("session")

        if cookie != nil {
            login := authcookie.Login(cookie.Value, []byte(config.Values.AuthSecret))
            email,_,_ := authcookie.Parse(cookie.Value, []byte(config.Values.AuthSecret))

            if login != ""  {
                user := models.GetUserByEmail(email)
                ctx.Set(r, "email", user.Email)
                ctx.Set(r, "id", user.Id)
                handler.ServeHTTP(w, r)
            } else {
                ctx.Set(r, "email", nil)
                ctx.Set(r, "id", nil)
                http.Redirect(w, r, "/", http.StatusForbidden)
            }
        } else {
            fmt.Fprintln(w, http.StatusForbidden)
            // not permitted
        }



	}

}
