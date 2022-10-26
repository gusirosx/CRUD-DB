package handlers

import (
	"crudAPI/models"
	"crudAPI/services"
	"crudAPI/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUsers will return all the users
func GetUsers(ctx *gin.Context) {
	response, err := services.GetUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing user items"})
		return
	}
	// send the response message
	ctx.JSON(http.StatusOK, response)
}

// GetUser will return a specific user
func GetUser(ctx *gin.Context) {
	// get the userID from the ctx params, key is "id"
	userID := ctx.Param("id")
	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no user ID was provided"})
		return
	}

	// call GetUser to get the user
	user, err := models.GetUser(userID)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	// send the response message
	ctx.JSON(http.StatusOK, user)
}

// CreateUser create a user in the postgres database
func CreateUser(ctx *gin.Context) {
	// create an empty user of type entity.User
	var user types.User

	// decode the json request to user
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// call CreateUser to create the user
	if err := models.CreateUser(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// send the response message
	ctx.JSON(http.StatusCreated, gin.H{"success": "User was successfully created"})
}

// UpdateUser update a user in the postgres database
func UpdateUser(ctx *gin.Context) {
	// create an empty user of type entity.User
	var user types.User

	// get the userID from the ctx params, key is "id"
	userID := ctx.Param("id")
	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no user ID was provided"})
		return
	}

	// decode the json request to user
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// call UpdateUser to update the user
	if err := models.UpdateUser(userID, user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// send the response message
	ctx.JSON(http.StatusOK, gin.H{"success": "user was successfully updated"})
}

// DeleteUser delete a user in the postgres database
func DeleteUser(ctx *gin.Context) {

	// get the userID from the ctx params, key is "id"
	userID := ctx.Param("id")
	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no user ID was provided"})
		return
	}

	// call DeleteUser to delete the user
	if err := models.DeleteUser(userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// send the response message
	ctx.JSON(http.StatusOK, gin.H{"success": "User was successfully deleted"})
}
