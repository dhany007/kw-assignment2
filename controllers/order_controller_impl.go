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

func (c *OrderControllerImpl) GetAllOrders(ctx *gin.Context) {
	orders, err := c.OrderService.GetAllOrdersItems()

	if helper.NotFoundError(ctx, err) {
		return
	}

	ctx.JSON(http.StatusOK, models.ResponseJSON{
		Code:    200,
		Status:  "OK",
		Payload: orders,
	})
}

func (c *OrderControllerImpl) CreateOrder(ctx *gin.Context) {
	requestOrder := params.RequestCreateOrder{}
	helper.ReadFromRequestBody(ctx, &requestOrder)

	response, err := c.OrderService.CreateOrderItems(requestOrder)

	if helper.PanicIfError(ctx, err) {
		return
	}

	ctx.JSON(http.StatusOK, models.ResponseJSON{
		Code:    200,
		Status:  "OK",
		Payload: response,
	})
}

func (c *OrderControllerImpl) UpdateOrder(ctx *gin.Context) {
	requestOrder := params.RequestCreateOrder{}
	helper.ReadFromRequestBody(ctx, &requestOrder)

	orderID := ctx.Param("orderID")
	id, err := strconv.Atoi(orderID)
	if helper.PanicIfError(ctx, err) {
		return
	}

	_, err = c.OrderService.GetOrderByOrderID(id)
	if helper.NotFoundError(ctx, err) {
		return
	}

	response, err := c.OrderService.UpdateOrderItems(requestOrder, id)

	if helper.PanicIfError(ctx, err) {
		return
	}

	ctx.JSON(http.StatusOK, models.ResponseJSON{
		Code:    200,
		Status:  "OK",
		Payload: response,
	})
}

func (c *OrderControllerImpl) GetOrderByOrderID(ctx *gin.Context) {
	orderID := ctx.Param("orderID")
	id, err := strconv.Atoi(orderID)

	if helper.PanicIfError(ctx, err) {
		return
	}

	response, err := c.OrderService.GetOrderByOrderID(id)

	if helper.NotFoundError(ctx, err) {
		return
	}

	ctx.JSON(http.StatusOK, models.ResponseJSON{
		Code:    200,
		Status:  "OK",
		Payload: response,
	})
}

func (c *OrderControllerImpl) DeleteOrderByOrderID(ctx *gin.Context) {
	orderID := ctx.Param("orderID")
	id, err := strconv.Atoi(orderID)

	if helper.PanicIfError(ctx, err) {
		return
	}

	err = c.OrderService.DeleteOrderByOrderID(id)

	if helper.NotFoundError(ctx, err) {
		return
	}

	ctx.JSON(http.StatusOK, models.ResponseJSON{
		Code:   200,
		Status: "OK",
	})
}
