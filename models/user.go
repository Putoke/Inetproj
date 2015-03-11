package models

import (
	"log"
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

func RegisterUser(name, lastname, email, password string) bool {

	stmt, err := db.Prepare(registerUserQuery)
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Query(name, lastname, email, password)

	if err != nil {
		log.Fatal(err)
	}

	return err == nil
}

func GetUserByEmail(email string) * User {
    stmt, err := db.Prepare(getUserQuery)
    if err != nil {
        log.Fatal(err)
    }


    user := &User{}
    err = stmt.QueryRow(email).Scan(&user.Id, &user.Name, &user.Lastname, &user.Email, &user.Password)

    return user
}
