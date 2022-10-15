package adapters

import (
	"clean2/domain/model"
	interactor "clean2/usecase"
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
