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
