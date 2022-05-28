package helper

import (
	"assignment2/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadFromRequestBody(ctx *gin.Context, result interface{}) bool {
	err := ctx.ShouldBindJSON(result)
	errRequest := false

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.ResponseJSON{
			Code:  http.StatusBadRequest,
			Error: err.Error(),
		})
		errRequest = true
	}

	return errRequest
}

func PanicIfError(ctx *gin.Context, err error) bool {
	result := false

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.ResponseJSON{
			Code:  http.StatusBadRequest,
			Error: err.Error(),
		})
		result = true
	}

	return result
}

func NotFoundError(ctx *gin.Context, err error) bool {
	result := false

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, models.ResponseJSON{
			Code:  http.StatusNotFound,
			Error: err.Error(),
		})
		result = true
	}

	return result
}
