package interfaces

import (
	"main.go/pkg/common/helper"
	"main.go/pkg/common/response"
)

type PaymentUseCase interface {
	CreateRazorpayPayment(orderId int) (response.OrderResponse, string, int, error)
	UpdatePaymentDetails(paymentVerifier helper.PaymentVerification) error
	DisplayWallet(userId int) (response.Wallet, error)
}
