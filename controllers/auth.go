package controllers
import (
    "net/http"
    "github.com/gorilla/mux"
    "github.com/gorilla/sessions"
    "github.com/gorilla/securecookie"
    "Inetproj/models"
    "log"
    "crypto/md5"
    "encoding/hex"

    "bytes"
    "fmt"
    "errors"
)

var Store = sessions.NewCookieStore(
    []byte(securecookie.GenerateRandomKey(64)), // signing key
    []byte(securecookie.GenerateRandomKey(32)))

var ErrInvalidLogin = errors.New("Hash mismatch")

func Login(w http.ResponseWriter, r *http.Request) {

    email := mux.Vars(r)["email"]
    password := mux.Vars(r)["password"]
//    session, _ := Store.Get(r, "inetproj")

    user := models.GetUserByEmail(email)

    err := compareHashAndPassword([]byte(user.Password), []byte(stringToMD5(password)))


    log.Println("Login attempt: email="+ email +", password="+password + "\nUser=" + user.Email +", password="+user.Password + "\nhashcompute="+ stringToMD5(password))
    fmt.Fprintln(w, err)



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

func compareHashAndPassword( hash []byte , passwd []byte ) error {
    if (bytes.Compare(hash[:], passwd[:]) != 0) {
        return ErrInvalidLogin
    }
    return nil
}