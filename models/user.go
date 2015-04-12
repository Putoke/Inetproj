package models

import (
	"log"
    "encoding/json"
    "Inetproj/util"
)

type User struct {
	Id       string `json:"user"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

const registerUserQuery = "INSERT INTO users (name, lastname, email, password) values (?, ?, ?, ?);"
const getUserQuery = "SELECT id, name, lastname, email, password FROM users where email = ?"

func GetUserByEmail(email string) * User {
    stmt, err := db.Prepare(getUserQuery)
    if err != nil {
        log.Fatal(err)
    }

    user := &User{}
    err = stmt.QueryRow(email).Scan(&user.Id, &user.Name, &user.Lastname, &user.Email, &user.Password)

    return user
}

func RegisterUser(data string) bool {

	stmt, err := db.Prepare(registerUserQuery)

    if err != nil {
        log.Fatal(err)
    }
    var ex User
    _ = json.Unmarshal([]byte(data), &ex)

	_, err = stmt.Exec(ex.Name, ex.Lastname, ex.Email, util.StringToMD5(ex.Password))

	if err != nil {
		log.Println(err)
	}

	return err == nil
}


