package types

import "net/http"

type User struct {
	Email string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserService interface {
	DeleteUser(uid string) error
	FindAll() ([]User, error)
	Validate(user *User) error
	Create(user *User) (*User, error)
}

type UserController interface {
	GetUsers(response http.ResponseWriter, request *http.Request)
}
