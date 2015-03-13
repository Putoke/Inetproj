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
)

var Store = sessions.NewCookieStore(
	[]byte(securecookie.GenerateRandomKey(64)), // signing key
	[]byte(securecookie.GenerateRandomKey(32)))

var ErrInvalidLogin = errors.New("Hash mismatch")

func Logintest(w http.ResponseWriter, r *http.Request) {

	secret := []byte("my secret key")

	email := mux.Vars(r)["email"]
	password := mux.Vars(r)["password"]

	user := models.GetUserByEmail(email)

	if equalHashes(stringToMD5(password), user.Password) {
		c := authcookie.NewSinceNow(email, 24*time.Hour, secret)
		expiration := time.Now().Add(60 * time.Second)
		cookie := http.Cookie{Name: "session", Value: c, Path: "/", Expires: expiration}
		http.SetCookie(w, &cookie)

        ctx.Set(r, "email", user.Email)
        ctx.Set(r, "id", user.Id)

		fmt.Fprintln(w, http.StatusOK)

	} else {
        ctx.Set(r, "email", nil)
        ctx.Set(r, "id", nil)
		fmt.Fprintln(w, http.StatusForbidden)
	}

}

func Login(w http.ResponseWriter, r *http.Request) {

	email := mux.Vars(r)["email"]
	password := mux.Vars(r)["password"]
	//    session, _ := Store.Get(r, "inetproj")

	user := models.GetUserByEmail(email)

	err := compareHashAndPassword([]byte(user.Password), []byte(stringToMD5(password)))

	log.Println("Login attempt: email=" + email + ", password=" + password + "\nUser=" + user.Email + ", password=" + user.Password + "\nhashcompute=" + stringToMD5(password))
	if err != nil {
		fmt.Fprintln(w, err)
	} else {
		fmt.Fprint(w, "Login successfull")
	}

	// https://github.com/iamjem/go-passwordless-demo
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
