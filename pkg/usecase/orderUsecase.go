package usecase

import (
	"main.go/pkg/common/helper"
	"main.go/pkg/common/response"
	"main.go/pkg/domain"
	interfaces "main.go/pkg/repository/interface"
	services "main.go/pkg/usecase/interface"
)

type OrderUseCase struct {
	orderRepo interfaces.OrderRepository
}



func NewOrderUseCase(orderRepo interfaces.OrderRepository) services.OrderUseCase {
	return &OrderUseCase{
		orderRepo: orderRepo,
	}
}

// -------------------------- Order-All --------------------------//

func (c *OrderUseCase) OrderAll(id, paymentTypeId int) (domain.Orders, error) {
	order, err := c.orderRepo.OrderAll(id, paymentTypeId)
	return order, err
}

// -------------------------- Cancel-Order --------------------------//

func (c *OrderUseCase) UserCancelOrder(orderId, userId int) error {
	err := c.orderRepo.UserCancelOrder(orderId, userId)
	return err
}

// -------------------------- List-Order --------------------------//

func (c *OrderUseCase) ListOrder(userId, orderId int) (response.OrderResponse, error) {
	order, err := c.orderRepo.ListOrder(userId, orderId)
	return order, err
}

// -------------------------- List-All-Order --------------------------//

func (c *OrderUseCase) ListAllOrders(userId int) ([]domain.Orders, error) {
	orders, err := c.orderRepo.ListAllOrders(userId)
	return orders, err
}

// -------------------------- Update-Order --------------------------//

func (c *OrderUseCase) UpdateOrder(updateOrder helper.UpdateOrder) error {
	err := c.orderRepo.UpdateOrder(updateOrder)
	return err
}

// -------------------------- Return-Order --------------------------//

func (c *OrderUseCase) ReturnOrder(userId, orderId int) (int, error) {
	returnAmount, err := c.orderRepo.ReturnOrder(userId, orderId)
	return returnAmount, err
}

// ListAllOrderForAdmin implements interfaces.OrderUseCase.
func (c *OrderUseCase) ListAllOrderForAdmin() ([]response.AdminOrder, error) {
	orders, err := c.orderRepo.ListAllOrdersForAdmin()
	return orders, err
}

// AddOrderStatus implements interfaces.OrderUseCase.
// func (c *OrderUseCase) AddOrderStatus(orderStatus helper.OrderStatus) (response.OrderStatus, error) {
// 	newOrderStatus, err := c.orderRepo.AddOrderStatus(orderStatus)
// 	return newOrderStatus, err
// }