package interfaces

import (
	"main.go/pkg/common/helper"
	"main.go/pkg/domain"
)

type OrderUseCase interface {
	OrderAll(id, paymentTypeId int) (domain.Orders, error)
	UserCancelOrder(orderId, userId int) error
	ListOrder(userId, orderId int) (domain.Orders, error)
	ListAllOrders(userId int) ([]domain.Orders, error)
	UpdateOrder(UpdateOrder helper.UpdateOrder) error

}