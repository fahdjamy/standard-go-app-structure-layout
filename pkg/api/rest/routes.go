package rest

import (
	"github.com/fahdjamy/standard-structure-layout/pkg/api/rest/controller"
	"github.com/fahdjamy/standard-structure-layout/pkg/types"
)

var (
	userController types.UserController = controller.NewUserController()
)

func Routes (router types.Router, port string)  {
	router.Get("/users", userController.GetUsers)
	router.Serve(port)
}
