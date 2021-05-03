package rest

import (
	"github.com/fahdjamy/standard-structure-layout/pkg/api/rest/controller"
	"github.com/fahdjamy/standard-structure-layout/pkg/api/rest/middlewares"
	"github.com/fahdjamy/standard-structure-layout/pkg/types"
	"log"
)

func Routes(router types.Router, port string, service types.DBService, logger *log.Logger) {
	customController := controller.NewCustomController(logger)
	userController := controller.NewUserController(service, logger)

	router.Get("/users", middlewares.Log(userController.GetUsers))

	router.NotFound(customController.PageNotFoundHandler)
	router.Serve(port)
}
