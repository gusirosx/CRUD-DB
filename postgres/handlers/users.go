package handlers

import (
	"crudAPI/entity"
	"crudAPI/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// type MongoController struct {
// 	session *mongo.Client
// }

// func NewMongoController(session *mongo.Client) *MongoController {
// 	return &MongoController{session}
// }

func GetUsers(ctx *gin.Context) {
	response, err := models.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing user items"})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func GetUser(ctx *gin.Context) {
	userID := ctx.Param("id")

	user, err := models.GetUser(userID)
	if err != nil {
		// Arrumar isso ==================================================================
		if err.Error() == "mongo: no documents in result" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func CreateUser(ctx *gin.Context) {
	var user entity.User

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.CreateUser(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"success": "User was successfully created"})
}

func UpdateUser(ctx *gin.Context) {
	var user entity.User

	user.Id = ctx.Param("id")
	if user.Id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no user ID was provided"})
		return
	}

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.UpdateUser(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"success": "user was successfully updated"})
}

func DeleteUser(ctx *gin.Context) {

	userID := ctx.Param("id")
	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no user ID was provided"})
		return
	}

	if err := models.DeleteUser(userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"success": "User was successfully deleted"})
}
