package services

import (
	"assignment2/models"
	"assignment2/params"
)

type OrderService interface {
	GetAllOrdersItems() ([]models.Order, error)
	CreateOrderItems(request params.RequestCreateOrder) (models.Order, error)
	UpdateOrderItems(request params.RequestCreateOrder, id int) (models.Order, error)
	GetOrderByOrderID(id int) (models.Order, error)
	DeleteOrderByOrderID(id int) error
}
