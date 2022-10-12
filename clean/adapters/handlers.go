package adapters

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewUserHandler(router *gin.Engine, userUsecase UserUsecase) {
	userHandler := UserHandler{userUsecase}

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
		userRoutes.GET("", userHandler.GetUsers)
		// Read users at /users/ID
		userRoutes.GET("/:id", userHandler.GetUser)
	}

	//router.GET("/users", userHandler.GetUsers)
	// r.GET("/person", personHandler.viewPersons)
	// r.GET("person/:id", personHandler.viewPersonId)
	// r.PUT("/person/:id", personHandler.editPerson)
	// r.DELETE("/person/:id", personHandler.deletePerson)
}

// GetUsers will return all the users
func (userHandler *UserHandler) GetUsers(ctx *gin.Context) {
	response, err := userHandler.personUsecase.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing user items"})
		return
	}
	// send the response message
	ctx.JSON(http.StatusOK, response)
}

// GetUser will return a specific user
func (userHandler *UserHandler) GetUser(ctx *gin.Context) {
	// get the userID from the ctx params, key is "id"
	userID := ctx.Param("id")
	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no user ID was provided"})
		return
	}

	// call GetUser to get the user
	response, err := userHandler.personUsecase.GetUser(userID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	// send the response message
	ctx.JSON(http.StatusOK, response)
}
