package adapters

import (
	"clean2/domain/model"
	interactor "clean2/usecase"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppController struct {
	User interface{ UserController }
}

type Context interface {
	JSON(code int, i interface{}) error
}

type userController struct {
	userInteractor interactor.UserInteractor
}

type UserController interface {
	GetUsers(ctx *gin.Context)
	GetUser(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

func NewUserController(us interactor.UserInteractor) UserController {
	return &userController{us}
}

func (uc *userController) GetUsers(ctx *gin.Context) {
	var u []*model.User

	u, err := uc.userInteractor.Get(u)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing user items"})
		return
	}

	ctx.JSON(http.StatusOK, u)
}

// GetUser will return a specific user
func (uc *userController) GetUser(ctx *gin.Context) {
	// get the userID from the ctx params, key is "id"
	userID := ctx.Param("id")
	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no user ID was provided"})
		return
	}

	// call GetUser to get the user
	response, err := uc.userInteractor.GetByID(userID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing user items"})
		}
		return
	}
	// send the response message
	ctx.JSON(http.StatusOK, response)
}

// CreateUser create a user in the postgres database
func (uc *userController) CreateUser(ctx *gin.Context) {
	// create an empty user of type model.User
	var user model.User

	// decode the json request to user
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// call CreateUser to create the user
	_, err := uc.userInteractor.Create(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//response
	// send the response message
	ctx.JSON(http.StatusCreated, gin.H{"success": "User was successfully created"})
}

// UpdateUser update a user in the postgres database
func (uc *userController) UpdateUser(ctx *gin.Context) {
	// create an empty user of type model.User
	var user model.User

	// get the userID from the ctx params, key is "id"
	user.Id = ctx.Param("id")
	if user.Id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no user ID was provided"})
		return
	}

	// decode the json request to user
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// call UpdateUser to update the user
	_, err := uc.userInteractor.Update(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// send the response message
	ctx.JSON(http.StatusOK, gin.H{"success": "user was successfully updated"})
}

// DeleteUser delete a user in the postgres database
func (uc *userController) DeleteUser(ctx *gin.Context) {
	// get the userID from the ctx params, key is "id"
	userID := ctx.Param("id")
	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no user ID was provided"})
		return
	}

	// call DeleteUser to update the user
	_, err := uc.userInteractor.Delete(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// send the response message
	ctx.JSON(http.StatusOK, gin.H{"success": "User was successfully deleted"})
}
