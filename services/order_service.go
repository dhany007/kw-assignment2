package services

import (
	"assignment2/models"
)

type OrderService interface {
	GetAllOrdersItems() []models.Order
}
