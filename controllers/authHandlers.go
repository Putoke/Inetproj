package controllers

import (
	"Inetproj/config"
	"Inetproj/models"
	"encoding/json"
	"github.com/dchest/authcookie"
	ctx "github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
	"time"
    "io/ioutil"
    "Inetproj/util"
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

    if _, ok:= requestHasValidSession(r); ok {
        status := "Already logged in"
        log.Println("Login failed from " + r.RemoteAddr + " as email=" + email + " (" + status + ")")
        SendHTTPStatusJSON(w, http.StatusForbidden)
    } else {
        if util.EqualHashes([]byte(util.StringToMD5(password)), []byte(user.Password)) {
            duration := time.Duration(config.Values.SessionExpirationTimeMinutes) * time.Minute
            expiration := time.Now().Add(duration)

            c := authcookie.NewSinceNow(email, duration, []byte(config.Values.AuthSecret))
            cookie := http.Cookie{Name: config.Values.SessionCookieName, Value: c, Path: "/", Expires: expiration, HttpOnly: true}
            http.SetCookie(w, &cookie)

            ctx.Set(r, "email", user.Email)
            ctx.Set(r, "id", user.Id)

            log.Println("Login successful from " + r.RemoteAddr + " as email=" + user.Email)
            SendHTTPStatusJSON(w, http.StatusOK)
        } else {
            ctx.Set(r, "email", nil)
            ctx.Set(r, "id", nil)
            log.Println("Login failed from " + r.RemoteAddr + " as email=" + email)
            SendHTTPStatusJSON(w, http.StatusUnauthorized)
        }
    }
}

func LoginPost( w http.ResponseWriter, r * http.Request) {

    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        log.Fatal(err)
    }

    user := new(models.User)
    err = json.Unmarshal(body, &user)
    if err != nil {
        log.Fatal(err)
    }

    email := user.Email
    password := user.Password

    user = models.GetUserByEmail(email)

    if _, ok:= requestHasValidSession(r); ok {
        status := "Already logged in"
        log.Println("Login failed from " + r.RemoteAddr + " as email=" + email + " (" + status + ")")
        SendHTTPStatusJSON(w, http.StatusForbidden)
    } else {
        if util.EqualHashes([]byte(util.StringToMD5(password)), []byte(user.Password)) {
            duration := time.Duration(config.Values.SessionExpirationTimeMinutes) * time.Minute
            expiration := time.Now().Add(duration)

            c := authcookie.NewSinceNow(email, duration, []byte(config.Values.AuthSecret))
            cookie := http.Cookie{Name: config.Values.SessionCookieName, Value: c, Path: "/", Expires: expiration, HttpOnly: true}
            http.SetCookie(w, &cookie)

            ctx.Set(r, "email", user.Email)
            ctx.Set(r, "id", user.Id)

            log.Println("Login successful from " + r.RemoteAddr + " as email=" + user.Email)
            SendHTTPStatusJSON(w, http.StatusOK)
        } else {
            ctx.Set(r, "email", nil)
            ctx.Set(r, "id", nil)
            log.Println("Login failed from " + r.RemoteAddr + " as email=" + email)
            SendHTTPStatusJSON(w, http.StatusUnauthorized)
        }
    }


}

func RegisterUser(w http.ResponseWriter, r *http.Request) {

    //id, email := getIDAndEmail(r)
    //printHandlerLog(id, email, r)

    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        log.Fatal(err)
    }

    success := models.RegisterUser(string(body[:]))

    if success {
        SendHTTPStatusJSON(w, http.StatusOK)
        log.Println("User registered successful")
    } else {
        SendHTTPStatusJSON(w, http.StatusForbidden)
        log.Println("User registered failed")
    }

}



func SendStatusJSON(w http.ResponseWriter, status string, code int) {
	s := &Status{Status: status, Code: code}
	json.NewEncoder(w).Encode(s)
}

func SendHTTPStatusJSON(w http.ResponseWriter, code int) {
    SendStatusJSON(w, http.StatusText(code), code)
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
