package router

import (
	"clean2/adapters"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

	//Group user related routes together
	userRoutes := router.Group("/users")
	{
		// Read users at /users
		userRoutes.GET("", app.User.GetUsers)
		// Read users at /users/ID
		userRoutes.GET("/:id", app.User.GetUser)
		// Create user at /users
		userRoutes.POST("", app.User.CreateUser)
		// Update users at /users
		userRoutes.PUT("/:id", app.User.UpdateUser)
		// Delete users at /users
		userRoutes.DELETE("/:id", app.User.DeleteUser)
	}
}
