package controllers

import (
	"assignment2/helper"
	"assignment2/models"
	"assignment2/params"
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
	orders, err := controller.OrderService.GetAllOrdersItems()

	if err != nil {
		ctx.JSON(http.StatusNotFound, models.ResponseJSON{
			Code:  http.StatusNotFound,
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.ResponseJSON{
		Code:    200,
		Status:  "OK",
		Payload: orders,
	})
}

func (controller *OrderControllerImpl) CreateOrder(ctx *gin.Context) {
	requestOrder := params.RequestCreateOrder{}
	helper.ReadFromRequestBody(ctx, &requestOrder)

	response, err := controller.OrderService.CreateOrderItems(requestOrder)
	helper.PanicIfError(ctx, err)

	ctx.JSON(http.StatusOK, models.ResponseJSON{
		Code:    200,
		Status:  "OK",
		Payload: response,
	})
}
