package repositories

import (
	"assignment2/models"
	"fmt"

	"gorm.io/gorm"
)

type OrderRepositoryImpl struct {
}

func NewOrderRepository() OrderRepository {
	return &OrderRepositoryImpl{}
}

func (repository *OrderRepositoryImpl) GetAllOrders(db *gorm.DB) ([]models.Order, error) {
	orders := []models.Order{}

	err := db.Find(&orders).Error

	test := db.Preload("Items").Find(&orders)
	fmt.Printf("%v\n", test)

	return orders, err
}
