package types

import (
	"errors"
	"net/http"
)

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserService interface {
	FindAll() ([]User, error)
	Validate(user *User) error
	DeleteUser(uid string) error
	Create(user *User) (*User, error)
}

type UserController interface {
	GetUsers(response http.ResponseWriter, request *http.Request)
}

func (u User) ValidateData() error {
	if u.Email == "" {
		return errors.New("please provide a valid email address")
	}
	if u.Password == "" {
		return errors.New("please provide a valid password")
	}
	if u.Username == "" {
		return errors.New("please provide a username")
	}
	return nil
}
