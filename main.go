package main

import (
	"log"
	"os"

	"github.com/diegogomesaraujo/central-loteria/server"
)

func main() {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port: %v\n", port)

	allowedOrigins := []string{"*"}
	allowedMethods := []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}

	server.StartServer(":"+port, allowedOrigins, allowedMethods)
}
