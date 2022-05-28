package helper

import (
	"assignment2/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadFromRequestBody(ctx *gin.Context, result interface{}) {
	err := ctx.ShouldBindJSON(result)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ResponseJSON{
			Code:  http.StatusBadRequest,
			Error: err.Error(),
		})
		return
	}
}

func PanicIfError(ctx *gin.Context, err error) {
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ResponseJSON{
			Code:  http.StatusBadRequest,
			Error: err.Error(),
		})
		return
	}
}
