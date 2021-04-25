package controller

import (
	"github.com/fahdjamy/standard-structure-layout/pkg/types"
	"net/http"
)

type userController struct {}

func (c *userController) GetUsers(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
}

func NewUserController() types.UserController {
	return &userController{}
}
