package interfaces

import (
	"main.go/pkg/common/response"
	"main.go/pkg/domain"
)

type PaymentRepository interface {
	ViewPaymentDetails(orderID int) (domain.PaymentDetails, error)
	UpdatePaymentDetails(orderID int, paymentRef string) (domain.PaymentDetails, error)
	DisplayWallet(userId int)(response.Wallet,error)
}