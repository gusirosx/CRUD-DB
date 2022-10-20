package main

import (
	"crudAPI/config"
	"crudAPI/routes"
)

func main() {
	config.Init()
	// Setup GinGonic Routes
	routes.RoutesSetup()
}
