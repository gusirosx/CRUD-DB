package adapters

import (
	"CRUD-DB/clean/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInteractor interface {
	GetUsers() ([]domain.User, error)
	//Add(userId, orderId, itemId int) error
}

type ServiceHandler struct {
	UserInteractor UserInteractor
}

type UserHandler struct {
	personUsecase UserUsecase
}

func NewUserHandler(r *gin.Engine, userUsecase UserUsecase) {
	userHandler := UserHandler{userUsecase}

	r.GET("/users", userHandler.GetUsers)
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
