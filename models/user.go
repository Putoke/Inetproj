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

const RegisterUserQuery = "INSERT INTO users (name, lastname, email, password) values (?, ?, ?, ?);"

func RegisterUser(name, lastname, email, password string) bool {

	stmt, err := db.Prepare(RegisterUserQuery)
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Query(name, lastname, email, password)

	if err != nil {
		log.Fatal(err)
	}

	return err == nil
}
