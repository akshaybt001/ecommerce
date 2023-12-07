package interfaces

import (
	"main.go/pkg/common/helper"
	"main.go/pkg/common/response"
	"main.go/pkg/domain"
)

type OrderRepository interface {
	OrderAll(id, paymentTypeId int) (domain.Orders, error)
	UserCancelOrder(orderId, userId int) error
	ListOrder(userId, orderId int) (response.OrderResponse,error)
	ListAllOrders(userId int) ([]domain.Orders, error)
	ListAllOrdersForAdmin()([]response.AdminOrder,error)
	UpdateOrder(updateOrder helper.UpdateOrder) error
	ReturnOrder(userId, orderId int) (int, error)
	UserIdFromOrder(id int) (int, error)
}