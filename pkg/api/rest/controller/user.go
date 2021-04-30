package controller

import (
	"encoding/json"
	"github.com/fahdjamy/standard-structure-layout/pkg/types"
	"log"
	"net/http"
)

type userController struct {
	service types.DBService
	logger *log.Logger
}

func (c *userController) GetUsers(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	users, err := c.service.GetAll("user")
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(writer).Encode(err)
		c.logger.Fatal(err)
		return
	} else {
		writer.WriteHeader(http.StatusOK)
	}
	err = json.NewEncoder(writer).Encode(users)
	c.logger.Print(err)
	return
}

func NewUserController(service types.DBService, logger *log.Logger) types.UserController {
	return &userController{service, logger}
}
