package types

import (
	"time"
)

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserById(id int) (*User, error)
	CreateUser(User) error
}

type User struct {
	ID        int       `json: "id"`
	FirstName string    `json: "firstname"`
	LastName  string    `json: "lastname"`
	Email     string    `json: "email"`
	Password  string    `json: "-"`
	CreateAt  time.Time `json: "createat"`
}

type RegisterUserPayload struct {
	FirstName string `json: "fistname"`
	LastName  string `json: "lastname"`
	Email     string `json: "email"`
	Password  string `json: "password"`
}
