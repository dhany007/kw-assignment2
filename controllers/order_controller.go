package controllers

import "github.com/gin-gonic/gin"

type OrderController interface {
	GetAllOrders(ctx *gin.Context)
	CreateOrder(ctx *gin.Context)
}
