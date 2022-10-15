package utils

import (
	"clean2/domain/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandleSucces(ctx *gin.Context, data interface{}) {
	responData := model.Respon{
		Status:  "200",
		Message: "Success",
		Data:    data,
	}
	ctx.JSON(http.StatusOK, responData)
}

func HandleError(ctx *gin.Context, status int, message string) {
	responData := model.Respon{
		Status:  strconv.Itoa(status),
		Message: message,
	}
	ctx.JSON(status, responData)
}
