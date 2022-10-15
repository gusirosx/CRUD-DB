package router

import (
	"clean2/adapters"
	"net/http"

	"github.com/gin-gonic/gin"
)

// func NewRouter(e *echo.Echo, c adapters.AppController) *echo.Echo {
// 	e.Use(middleware.Logger())
// 	e.Use(middleware.Recover())

// 	e.GET("/users", func(context echo.Context) error { return c.User.GetUsers(context) })

// 	return e
// }

func NewRouter(router *gin.Engine, app adapters.AppController) {
	//userHandler := UserHandler{userUsecase}

	// Handle the index route
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"success": "Up and running..."})
	})
	// Handle the no route case
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Page not found"})
	})

	router.GET("/users", app.User.GetUsers)

	// //Group user related routes together
	// userRoutes := router.Group("/users")
	// {
	// 	// Read users at /users
	// 	userRoutes.GET("", userHandler.GetUsers)
	// 	// Read users at /users/ID
	// 	userRoutes.GET("/:id", userHandler.GetUser)
	// }

	//router.GET("/users", userHandler.GetUsers)
	// r.GET("/person", personHandler.viewPersons)
	// r.GET("person/:id", personHandler.viewPersonId)
	// r.PUT("/person/:id", personHandler.editPerson)
	// r.DELETE("/person/:id", personHandler.deletePerson)
}
