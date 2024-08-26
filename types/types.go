package types

type RegisterUserPayload struct {
	FirstName string `json: "fistname"`
	LastName  string `json: "lastname"`
	Email     string `json: "email"`
	Password  string `json: "password"`
}
