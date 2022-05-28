package services

import (
	"assignment2/models"
	"assignment2/params"
)

type OrderService interface {
	GetAllOrdersItems() ([]models.Order, error)
	CreateOrderItems(params.RequestCreateOrder) (models.Order, error)
}
