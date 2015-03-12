package controllers

import (
	"github.com/dchest/authcookie"
    ctx "github.com/gorilla/context"
	"net/http"
    "log"
    "Inetproj/models"
)

func RequireLogin(handler http.Handler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		cookie, _ := r.Cookie("session")

        if cookie != nil {
            secret := []byte("my secret key")
            login := authcookie.Login(cookie.Value, secret)
            email,_,_ := authcookie.Parse(cookie.Value, secret)

            log.Println("login = " + email)


            if login != ""  {

                user := models.GetUserByEmail(email)
                session, _ := Store.Get(r, "inet")

                ctx.Set(r, "user", user)
                session.Values["id"] = user.Id
                handler.ServeHTTP(w, r)
            } else {
                http.Redirect(w, r, "/", http.StatusForbidden)
                ctx.Set(r, "user", nil)
            }
        } else {
            // not permitted
        }



	}

}
