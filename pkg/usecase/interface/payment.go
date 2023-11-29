package interfaces

import (
	"main.go/pkg/common/helper"
	"main.go/pkg/common/response"
	"main.go/pkg/domain"
)

type PaymentUseCase interface {
	CreateRazorpayPayment(orderId int) (domain.Orders, string, int, error)
	UpdatePaymentDetails(paymentVerifier helper.PaymentVerification) error
	DisplayWallet(userId int)(response.Wallet,error)
}