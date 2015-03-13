package controllers

import (
	"Inetproj/models"
	"crypto/md5"
	"encoding/hex"
	"github.com/dchest/authcookie"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
    ctx "github.com/gorilla/context"
	"log"
	"net/http"
	"time"

	"bytes"
	"errors"
	"fmt"
    "Inetproj/config"
)

var Store = sessions.NewCookieStore(
	[]byte(securecookie.GenerateRandomKey(64)), // signing key
	[]byte(securecookie.GenerateRandomKey(32)))

var ErrInvalidLogin = errors.New("Hash mismatch")

func Login(w http.ResponseWriter, r *http.Request) {



	email := mux.Vars(r)["email"]
	password := mux.Vars(r)["password"]

	user := models.GetUserByEmail(email)

	if equalHashes(stringToMD5(password), user.Password) {
        duration := time.Duration(config.Values.SessionExpirationTimeMinutes) * time.Minute
        expiration := time.Now().Add(duration)

		c := authcookie.NewSinceNow(email, duration, []byte(config.Values.AuthSecret))

		cookie := http.Cookie{Name: config.Values.SessionCookieName, Value: c, Path: "/", Expires: expiration}
		http.SetCookie(w, &cookie)

        ctx.Set(r, "email", user.Email)
        ctx.Set(r, "id", user.Id)

        log.Println("Login successful from " +r.RemoteAddr + " as email=" + user.Email)
		fmt.Fprintln(w, http.StatusOK)

	} else {
        ctx.Set(r, "email", nil)
        ctx.Set(r, "id", nil)
        log.Println("Login failed from " +r.RemoteAddr + " as email=" + user.Email)
		fmt.Fprintln(w, http.StatusForbidden)
	}

}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	lastname := vars["lastname"]
	email := vars["email"]
	pwhash := stringToMD5(vars["password"])

	success := models.RegisterUser(name, lastname, email, pwhash)
	if success {
		log.Println("New user registered (" + name + ", " + lastname + ", " + email + ")")
		w.Write([]byte("User registered"))
	}
}

func stringToMD5(s string) string {
	hasher := md5.New()
	hasher.Write([]byte(s))
	return hex.EncodeToString(hasher.Sum(nil))
}

func equalHashes(hash1, hash2 string) bool {

	return hash1 == hash2
}

func compareHashAndPassword(hash []byte, passwd []byte) error {
	if bytes.Compare(hash[:], passwd[:]) != 0 {
		return ErrInvalidLogin
	}
	return nil
}
