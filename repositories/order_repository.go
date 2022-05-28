package repositories

import (
	"assignment2/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	GetAllOrders(db *gorm.DB) ([]models.Order, error)
	CreateOrder(db *gorm.DB, request models.Order) (models.Order, error)
	UpdateOrder(db *gorm.DB, request models.Order, id int) (models.Order, error)
	GetOrderByOrderID(db *gorm.DB, id int) (models.Order, error)
	DeleteOrderByOrderID(db *gorm.DB, request models.Order) error
}
