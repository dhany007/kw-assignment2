package controllers

import (
	"assignment2/helper"
	"assignment2/models"
	"assignment2/params"
	"assignment2/services"
	"net/http"
	"strconv"

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

	if helper.NotFoundError(ctx, err) {
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

	if helper.PanicIfError(ctx, err) {
		return
	}

	ctx.JSON(http.StatusOK, models.ResponseJSON{
		Code:    200,
		Status:  "OK",
		Payload: response,
	})
}

func (controller *OrderControllerImpl) UpdateOrder(ctx *gin.Context) {
	requestOrder := params.RequestCreateOrder{}
	helper.ReadFromRequestBody(ctx, &requestOrder)

	orderID := ctx.Param("orderID")
	id, err := strconv.Atoi(orderID)
	if helper.PanicIfError(ctx, err) {
		return
	}

	_, err = controller.OrderService.GetOrderByOrderID(id)
	if helper.NotFoundError(ctx, err) {
		return
	}

	response, err := controller.OrderService.UpdateOrderItems(requestOrder, id)

	if helper.PanicIfError(ctx, err) {
		return
	}

	ctx.JSON(http.StatusOK, models.ResponseJSON{
		Code:    200,
		Status:  "OK",
		Payload: response,
	})
}

func (controller *OrderControllerImpl) GetOrderByOrderID(ctx *gin.Context) {
	orderID := ctx.Param("orderID")
	id, err := strconv.Atoi(orderID)

	if helper.PanicIfError(ctx, err) {
		return
	}

	response, err := controller.OrderService.GetOrderByOrderID(id)

	if helper.NotFoundError(ctx, err) {
		return
	}

	ctx.JSON(http.StatusOK, models.ResponseJSON{
		Code:    200,
		Status:  "OK",
		Payload: response,
	})
}
