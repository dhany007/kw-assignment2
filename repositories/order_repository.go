package repositories

import (
	"assignment2/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	GetAllOrders(db *gorm.DB) ([]models.Order, error)
}
