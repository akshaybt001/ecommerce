package interfaces

import (
	"main.go/pkg/common/helper"
	"main.go/pkg/common/response"
	"main.go/pkg/domain"
)

type OrderUseCase interface {
	OrderAll(id, paymentTypeId int) (domain.Orders, error)
	UserCancelOrder(orderId, userId int) error
	ListOrder(userId, orderId int) (response.OrderResponse, error)
	ListAllOrders(userId int) ([]domain.Orders, error)
	ListAllOrderForAdmin() ([]response.AdminOrder,error)
	UpdateOrder(UpdateOrder helper.UpdateOrder) error
	ReturnOrder(userId, orderId int) (int, error)
	// AddOrderStatus(orderStatus helper.OrderStatus) (response.OrderStatus, error)

}
