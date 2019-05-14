package server

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Route specification
type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

func configureRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routesList {
		router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

// StartServer to listen connections
func StartServer(addressListen string, allowedOrigins []string, allowedMethods []string) {
	router := configureRoutes()

	allowedOriginsCors := handlers.AllowedOrigins(allowedOrigins)
	allowedMethodsCors := handlers.AllowedMethods(allowedMethods)

	err := http.ListenAndServe(addressListen, handlers.CORS(allowedOriginsCors, allowedMethodsCors)(router))

	log.Fatal(err)
}
