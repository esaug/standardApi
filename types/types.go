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
	FirstName string `json: "fistname" validate:"required"`
	LastName  string `json: "lastname" validate:"required"`
	Email     string `json: "email" validate:"required,email"`
	Password  string `json: "password" validate:"required,min=3,max=130"`
}
