package controllers

import (
	"assignment2/models"
	"assignment2/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderControllerImpl struct {
	OrderService services.OrderService
}

func NewOrderController(orderService services.OrderService) OrderController {
	return &OrderControllerImpl{
		OrderService: orderService,
	}
}

func (controller *OrderControllerImpl) GetAllOrders(ctx *gin.Context) {
	orders := controller.OrderService.GetAllOrdersItems()

	ctx.JSON(http.StatusOK, models.ResponseJSON{
		Code:    200,
		Status:  "OK",
		Payload: orders,
	})
}
