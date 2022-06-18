package repositories

import (
	"assignment2/models"
	"errors"

	"gorm.io/gorm"
)

type OrderRepositoryImpl struct {
}

func NewOrderRepository() OrderRepository {
	return &OrderRepositoryImpl{}
}

func (r *OrderRepositoryImpl) GetAllOrders(db *gorm.DB) ([]models.Order, error) {
	orders := []models.Order{}

	result := db.Preload("Items").Find(&orders)
	if result.RowsAffected == 0 {
		return nil, errors.New("not found")
	}

	return orders, nil
}

func (r *OrderRepositoryImpl) CreateOrder(db *gorm.DB, request models.Order) (models.Order, error) {
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

func (r *OrderRepositoryImpl) UpdateOrder(db *gorm.DB, request models.Order, id int) (models.Order, error) {
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

func (r *OrderRepositoryImpl) GetOrderByOrderID(db *gorm.DB, id int) (models.Order, error) {
	orders := models.Order{}

	result := db.Where("order_id = ?", id).Preload("Items").Find(&orders)

	if result.RowsAffected == 0 {
		return orders, errors.New("order not found")
	}

	return orders, nil
}

func (r *OrderRepositoryImpl) DeleteOrderByOrderID(db *gorm.DB, request models.Order) error {
	err := db.Delete(&request.Items).Error
	if err != nil {
		return errors.New(err.Error())
	}

	err = db.Delete(&request).Error
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}
