package repository

import (
	"fmt"

	"gorm.io/gorm"
	"main.go/pkg/common/response"
	"main.go/pkg/domain"
	interfaces "main.go/pkg/repository/interface"
)

type PaymentDatabase struct {
	DB *gorm.DB
}

func NewPaymentRepository(DB *gorm.DB) interfaces.PaymentRepository {
	return &PaymentDatabase{DB}
}

func (c *PaymentDatabase) ViewPaymentDetails(orderID int) (domain.PaymentDetails, error) {
	var paymentDetails domain.PaymentDetails
	fetchPaymentDetailsQuery := `SELECT * FROM payment_details WHERE orders_id = $1;`
	err := c.DB.Raw(fetchPaymentDetailsQuery, orderID).Scan(&paymentDetails).Error
	fmt.Println("2", paymentDetails)
	return paymentDetails, err
}

func (c *PaymentDatabase) UpdatePaymentDetails(orderID int, paymentRef string) (domain.PaymentDetails, error) {
	var updatedPayment domain.PaymentDetails
	updatePaymentQuery := `	UPDATE payment_details SET payment_type_id = 2, payment_status_id = 3, payment_ref = $1, updated_at = NOW()
							WHERE orders_id = $2 RETURNING *;`
	tx := c.DB.Begin()
	err := c.DB.Raw(updatePaymentQuery, paymentRef, orderID).Scan(&updatedPayment).Error
	if err != nil {
		tx.Rollback()
		return updatedPayment, err
	}
	updateOrderTable := `UPDATE orders SET payment_status_id=3 WHERE id=$1`
	err = tx.Exec(updateOrderTable, orderID).Error
	if err != nil {
		tx.Rollback()
		return domain.PaymentDetails{}, err
	}
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return domain.PaymentDetails{}, err
	}
	return updatedPayment, err
}

// -------------------------- Wallet --------------------------//

func (c *PaymentDatabase) DisplayWallet(userId int) (response.Wallet, error) {
	var wallet response.Wallet
	query := `SELECT * FROM user_wallets WHERE users_id=?`
	err := c.DB.Raw(query, userId).Scan(&wallet).Error
	return wallet, err
}
