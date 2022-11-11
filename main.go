package main

import (
	"github.com/baguseka01/golang_echo_RESTAPI/config"
	"github.com/baguseka01/golang_echo_RESTAPI/routes"
)

func main() {
	// Initialize database
	config.Connect()

	// Initialize router
	e := routes.Route()

	// Initialize app port
	e.Logger.Fatal(e.Start(":9999"))

}
