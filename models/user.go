package models

type User struct {
    Id string `json:"user"`
    Name string `json:"name"`
    Lastname string `json:"lastname"`
    Email string `json:"email"`
    Password string `json:"password"`
}