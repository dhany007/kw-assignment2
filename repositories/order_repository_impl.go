package repositories

import (
	"assignment2/models"
	"errors"
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

	fmt.Println("kesini")
	result := db.Preload("Items").Find(&orders)
	// result := db.Model(models.Order{}).Joins("JOIN items on items.order_id = orders.order_id").Scan(&orders)
	fmt.Println(result)
	if result.RowsAffected == 0 {
		return nil, errors.New("not found")
	}

	return orders, nil
}

func (repository *OrderRepositoryImpl) CreateOrder(db *gorm.DB, request models.Order) (models.Order, error) {
	orders := request
	items := request.Items

	err := db.Create(&orders).Error

	if err != nil {
		return orders, errors.New(err.Error())
	}

	err = db.Create(&items).Error

	if err != nil {
		return orders, errors.New(err.Error())
	}

	return orders, nil
}
