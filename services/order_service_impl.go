package services

import (
	"assignment2/models"
	"assignment2/params"
	"assignment2/repositories"
	"errors"

	"gorm.io/gorm"
)

type OrderServiceImpl struct {
	OrderRepository repositories.OrderRepository
	DB              *gorm.DB
}

func NewOrderService(db *gorm.DB, orderRepository repositories.OrderRepository) OrderService {
	return &OrderServiceImpl{
		OrderRepository: orderRepository,
		DB:              db,
	}
}

func (service *OrderServiceImpl) GetAllOrdersItems() ([]models.Order, error) {
	orders, err := service.OrderRepository.GetAllOrders(service.DB)

	if err != nil {
		return nil, errors.New("orders of items not found")
	}

	return orders, nil
}

func (service *OrderServiceImpl) CreateOrderItems(request params.RequestCreateOrder) (models.Order, error) {
	items := []models.Item{}

	for _, item := range request.Items {
		tempItem := models.Item{
			ItemCode:    item.ItemCode,
			Quantity:    item.Quantity,
			Description: item.Description,
		}
		items = append(items, tempItem)
	}

	order := models.Order{
		CustomerName: request.CustomerName,
		Items:        items,
	}

	response, _ := service.OrderRepository.CreateOrder(service.DB, order)

	return response, nil
}
