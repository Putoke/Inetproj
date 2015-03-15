package controllers

import (
	"Inetproj/config"
	"Inetproj/models"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/dchest/authcookie"
	ctx "github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
	"time"
)

var Store = sessions.NewCookieStore(
	[]byte(securecookie.GenerateRandomKey(64)), // signing key
	[]byte(securecookie.GenerateRandomKey(32)))

type Status struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

func Login(w http.ResponseWriter, r *http.Request) {

	email := mux.Vars(r)["email"]
	password := mux.Vars(r)["password"]

	user := models.GetUserByEmail(email)

    if _, ok:= requestHasValidUserSession(r, email); ok {
        status := "Already logged in"
        log.Println("Login failed from " + r.RemoteAddr + " as email=" + email + " (" + status + ")")
        SendStatusJSON(w, status, http.StatusPreconditionFailed)
    } else {
        if equalHashes([]byte(stringToMD5(password)), []byte(user.Password)) {
            duration := time.Duration(config.Values.SessionExpirationTimeMinutes) * time.Minute
            expiration := time.Now().Add(duration)

            c := authcookie.NewSinceNow(email, duration, []byte(config.Values.AuthSecret))

            cookie := http.Cookie{Name: config.Values.SessionCookieName, Value: c, Path: "/", Expires: expiration}
            http.SetCookie(w, &cookie)

            ctx.Set(r, "email", user.Email)
            ctx.Set(r, "id", user.Id)

            log.Println("Login successful from " + r.RemoteAddr + " as email=" + user.Email)
            SendStatusJSON(w, "", http.StatusOK)

        } else {
            ctx.Set(r, "email", nil)
            ctx.Set(r, "id", nil)
            log.Println("Login failed from " + r.RemoteAddr + " as email=" + email)
            SendStatusJSON(w, "", http.StatusForbidden)
        }
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

func equalHashes(hash1 []byte, hash2 []byte) bool {
	return bytes.Compare(hash1[:], hash2[:]) == 0
}

func SendStatusJSON(w http.ResponseWriter, status string, code int) {
	s := &Status{Status: status, Code: code}
	json.NewEncoder(w).Encode(s)
}
func requestHasValidUserSession(r * http.Request, email string) (* http.Cookie, bool) {
    cookie, _ := r.Cookie("session")

    if cookie != nil {
        login := authcookie.Login(cookie.Value, []byte(config.Values.AuthSecret))
        cookieEmail, _,_ := authcookie.Parse(cookie.Value, []byte(config.Values.AuthSecret))

        if login != "" && email == cookieEmail {
            return cookie, true
        } else {
            return nil, false
        }

    } else {
        return nil, false
    }
}

func requestHasValidSession( r * http.Request) (* http.Cookie, bool) {
    cookie, _ := r.Cookie("session")

    if cookie != nil {
        login := authcookie.Login(cookie.Value, []byte(config.Values.AuthSecret))

        if login != "" {
            return cookie, true
        } else {
            return nil, false
        }

    } else {
        return nil, false
    }
}
