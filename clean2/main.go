package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"clean2/infrastructure/database"
	"clean2/infrastructure/router"
	"clean2/registry"
)

const (
	port = ":8080"
)

func main() {

	db := database.PostgresInstance()
	defer db.Close()

	registry := registry.NewRegistry(db)

	ginrouter := gin.Default()
	router.NewRouter(ginrouter, registry.NewAppController())

	// Run the http server
	if err := ginrouter.Run(port); err != nil {
		log.Fatalln("could not run server: ", err.Error())
	} else {
		log.Println("Server listening on port: ", port)
	}
}
