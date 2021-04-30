package rest

import (
	"github.com/fahdjamy/standard-structure-layout/pkg/api/rest/controller"
	"github.com/fahdjamy/standard-structure-layout/pkg/types"
	"log"
)

func Routes(router types.Router, port string, service types.DBService, logger *log.Logger) {
	userController := controller.NewUserController(service, logger)

	router.Get("/users", userController.GetUsers)

	router.Serve(port)
}
