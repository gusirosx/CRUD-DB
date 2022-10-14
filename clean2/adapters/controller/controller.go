package controller

import (
	"clean2/domain/model"
	interactor "clean2/usecase"
	"net/http"
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
	GetUsers(c Context) error
}

func NewUserController(us interactor.UserInteractor) UserController {
	return &userController{us}
}

func (uc *userController) GetUsers(c Context) error {
	var u []*model.User

	u, err := uc.userInteractor.Get(u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}
