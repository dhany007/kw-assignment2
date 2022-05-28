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

	orderRouter := router.Group("/orders")
	{
		orderRouter.GET("/", orderController.GetAllOrders)
		orderRouter.GET("/:orderID", orderController.GetOrderByOrderID)
		orderRouter.POST("/", orderController.CreateOrder)
		orderRouter.PUT("/:orderID", orderController.UpdateOrder)
		orderRouter.DELETE("/:orderID", orderController.DeleteOrderByOrderID)
	}

	router.Use(gin.Recovery())

	return router
}
