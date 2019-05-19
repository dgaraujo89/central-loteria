package server

import "github.com/diegogomesaraujo/central-loteria/controller"

var userController = &controller.UserController{}

type routes []Route

var routesList = routes{
	Route{
		Name:        "List all User",
		Method:      "GET",
		Path:        "/users",
		HandlerFunc: userController.Find,
	},
	Route{
		Name:        "Get User",
		Method:      "GET",
		Path:        "/users/{id}",
		HandlerFunc: userController.FindByID,
	},
	Route{
		Name:        "Add User",
		Method:      "POST",
		Path:        "/users",
		HandlerFunc: userController.Add,
	},
	Route{
		Name:        "Update User",
		Method:      "PUT",
		Path:        "/users",
		HandlerFunc: userController.Update,
	},
	Route{
		Name:        "Update User's Password",
		Method:      "PUT",
		Path:        "/users/password",
		HandlerFunc: userController.UpdatePassword,
	},
	Route{
		Name:        "Delete User",
		Method:      "DELETE",
		Path:        "/users/{id}",
		HandlerFunc: userController.Delete,
	},
}
