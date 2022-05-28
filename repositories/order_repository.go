package repositories

import (
	"assignment2/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	GetAllOrders(db *gorm.DB) ([]models.Order, error)
	CreateOrder(db *gorm.DB, request models.Order) (models.Order, error)
}
