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

func (repository *OrderRepositoryImpl) UpdateOrder(db *gorm.DB, request models.Order, id int) (models.Order, error) {
	order := models.Order{
		CustomerName: request.CustomerName,
	}

	err := db.Model(&order).Where("order_id = ?", id).Updates(order).Error

	if err != nil {
		return order, errors.New(err.Error())
	}

	newItems := []models.Item{}
	for _, tempItem := range request.Items {
		item := models.Item{
			ItemCode:    tempItem.ItemCode,
			Description: tempItem.Description,
			Quantity:    tempItem.Quantity,
		}

		db.Model(&item).Where("order_id = ? and item_id = ?", id, tempItem.ItemID).Updates(item)

		item.ItemID = tempItem.ItemID
		newItems = append(newItems, item)
	}

	order.Items = newItems

	return order, nil
}

func (repository *OrderRepositoryImpl) GetOrderByOrderID(db *gorm.DB, id int) (models.Order, error) {
	orders := models.Order{}

	result := db.Where("order_id = ?", id).Preload("Items").Find(&orders)

	if result.RowsAffected == 0 {
		return orders, errors.New("order not found")
	}

	return orders, nil
}
