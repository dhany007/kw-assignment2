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

func (service *OrderServiceImpl) UpdateOrderItems(request params.RequestCreateOrder, id int) (models.Order, error) {
	items := []models.Item{}

	for _, item := range request.Items {
		tempItem := models.Item{
			ItemID:      item.ItemID,
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

	response, _ := service.OrderRepository.UpdateOrder(service.DB, order, id)

	return response, nil
}

func (service *OrderServiceImpl) GetOrderByOrderID(id int) (models.Order, error) {
	order, err := service.OrderRepository.GetOrderByOrderID(service.DB, id)

	if err != nil {
		return order, errors.New("orders of items not found")
	}

	return order, nil
}

func (service *OrderServiceImpl) DeleteOrderByOrderID(id int) error {

	order, err := service.OrderRepository.GetOrderByOrderID(service.DB, id)

	if err != nil {
		return errors.New("orders of items not found")
	}

	err = service.OrderRepository.DeleteOrderByOrderID(service.DB, order)

	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}
