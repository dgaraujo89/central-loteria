package server

import "github.com/diegogomesaraujo/central-loteria/controller"

var userController = &controller.UserController{}

type routes []Route

var routesList = routes{
	Route{
		Name:        "User",
		Method:      "GET",
		Path:        "/users",
		HandlerFunc: userController.Find,
	},
	Route{
		Name:        "User Add",
		Method:      "POST",
		Path:        "/users",
		HandlerFunc: userController.Add,
	},
}
