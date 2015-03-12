package controllers

import (
	"github.com/dchest/authcookie"
	"net/http"
)

func RequireLogin(handler http.Handler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		cookie, _ := r.Cookie("session")

		secret := []byte("my secret key")
		login := authcookie.Login(cookie.Value, secret)

		if login != "" {
			handler.ServeHTTP(w, r)
		} else {
			http.Redirect(w, r, "/", http.StatusForbidden)
		}

	}

}
