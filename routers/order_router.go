package routers

import (
	"assignment2/controllers"
	"assignment2/repositories"
	"assignment2/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	orderRepository := repositories.NewOrderRepository()
	orderService := services.NewOrderService(db, orderRepository)
	orderController := controllers.NewOrderController(orderService)

	orders := router.Group("/orders")
	{
		orders.GET("/", orderController.GetAllOrders)
	}

	return router
}
