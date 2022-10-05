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

// GetUsers will return all the users
func (service ServiceHandler) GetUsers(ctx *gin.Context) {
	response, err := service.UserInteractor.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing user items"})
		return
	}
	// send the response message
	ctx.JSON(http.StatusOK, response)
}
