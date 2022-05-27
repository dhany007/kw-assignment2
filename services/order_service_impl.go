package services

import (
	"assignment2/models"
	"assignment2/repositories"
	"fmt"

	"gorm.io/gorm"
)

type OrderServiceImpl struct {
	OrderRepository repositories.OrderRepository
	DB              *gorm.DB
}

func NewOrderService(db *gorm.DB, orderRepository repositories.OrderRepository) OrderService {
	return &OrderServiceImpl{
		OrderRepository: orderRepository,
		DB:              db,
	}
}

func (service *OrderServiceImpl) GetAllOrdersItems() []models.Order {
	orders, err := service.OrderRepository.GetAllOrders(service.DB)

	fmt.Println("order", orders)
	fmt.Println("err", err)
	return orders
}
