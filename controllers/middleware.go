package controllers

import (
	"github.com/dchest/authcookie"
	"net/http"
    "log"
)

func RequireLogin(handler http.Handler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		cookie, _ := r.Cookie("session")

        if cookie != nil {
            secret := []byte("my secret key")
            login := authcookie.Login(cookie.Value, secret)
            s,_,_ := authcookie.Parse(cookie.Value, secret)

            log.Println("login = " + s)


            if login != ""  {
                handler.ServeHTTP(w, r)
            } else {
                http.Redirect(w, r, "/", http.StatusForbidden)
            }
        } else {
            // not permitted
        }



	}

}
