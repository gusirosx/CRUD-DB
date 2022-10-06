package main

import (
	"CRUD-DB/clean/adapters"
	"CRUD-DB/clean/infrastructure"
	"CRUD-DB/clean/usecases"
	"log"

	"github.com/gin-gonic/gin"
)

const (
	port = ":8080"
)

func main() {

	db := infrastructure.PostgresInstance()
	defer db.Close()

	router := gin.Default()

	userRepo := usecases.NewUserRepoImpl(db)
	userUsecase := adapters.NewUserUsecase(userRepo)
	// personHandler.CreatePersonHandler(router, personUsecase)
	adapters.NewUserHandler(router, userUsecase)
	// Run the http server
	if err := router.Run(port); err != nil {
		log.Fatalln("could not run server: ", err.Error())
	} else {
		log.Println("Server listening on port: ", port)
	}

}
