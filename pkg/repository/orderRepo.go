package repository

import (
	"fmt"

	"gorm.io/gorm"
	"main.go/pkg/common/helper"
	"main.go/pkg/common/response"
	"main.go/pkg/domain"
	interfaces "main.go/pkg/repository/interface"
)

type OrderDatabase struct {
	DB *gorm.DB
}

func NewOrderRepository(DB *gorm.DB) interfaces.OrderRepository {
	return &OrderDatabase{DB}
}

//-------------------------- Order-All --------------------------//

func (c *OrderDatabase) OrderAll(id, paymentTypeId int) (domain.Orders, error) {

	paymentStatus := 1
	orderStatus := 1

	tx := c.DB.Begin()

	//Find the cart id and tottal of the cart
	var cart domain.Carts
	findCart := `SELECT * FROM carts WHERE user_id=? `
	err := tx.Raw(findCart, id).Scan(&cart).Error
	if err != nil {
		tx.Rollback()
		return domain.Orders{}, err
	}
	if cart.Total == 0 {
		setTotal := `UPDATE carts SET total=carts.sub_total`
		err = tx.Exec(setTotal).Error
		if err != nil {
			tx.Rollback()
			return domain.Orders{}, err
		}
		cart.Total = cart.SubTotal
	}
	if cart.SubTotal == 0 {
		tx.Rollback()
		return domain.Orders{}, fmt.Errorf("no items in cart")
	}
	// var modelIds []int
	// findModelIds := `SELECT model_id FROM cart_items WHERE carts_id = $1;`
	// err = tx.Raw(findModelIds, cart.Id).Scan(&modelIds).Error
	// if err != nil {
	// 	tx.Rollback()
	// 	return domain.Orders{}, err
	// }
	// var brandIds []int
	// for i, modelId := range modelIds {
	// 	var productdetails domain.Model
	// 	findproductdetails := `SELECT models.*,discounts.*  FROM models LEFT JOIN discounts ON discounts.brand_id=models.brand_id AND expiration_date WHERE id=$1`
	// 	err = tx.Raw(findproductdetails, modelId).Scan(&productdetails).Error
	// 	if err != nil {
	// 		tx.Rollback()
	// 		return domain.Orders{}, err
	// 	}

	// 	brandIds = append(brandIds, int(productdetails.Brand_id))
	// 		 // Check discount for each item
	// 		 findDiscount := `SELECT * FROM discounts WHERE brand_id = ? AND expiration_date > NOW();`
	// 		 var discount domain.Discount
	// 		 err = tx.Raw(findDiscount, brandIds[i]).Scan(&discount).Error
	// 		 if err != nil {
	// 			 tx.Rollback()
	// 			 return domain.Orders{}, err
	// 		 }
	//   // Add error message for expired discount
	// if discount.Id == 0 || discount.ExpirationDate.Before(time.Now()) {
	// 	return domain.Orders{}, fmt.Errorf("Discount expired! Please add item again to benefit from the discount.")
	// }

	// }

	var expiredCartItems []domain.CartItem
	findExpiredCartItems := `SELECT cart_items.*
	FROM cart_items
	JOIN models ON models.id = cart_items.model_id
	JOIN discounts ON models.brand_id = discounts.brand_id
	WHERE cart_items.is_discounted = true
		AND models.price > discounts.minimum_purchase_amount
		AND discounts.expiration_date < NOW()
		AND cart_items.carts_id = ?`
	err = tx.Raw(findExpiredCartItems, cart.Id).Scan(&expiredCartItems).Error
	if err != nil {
		tx.Rollback()
		return domain.Orders{}, fmt.Errorf("there seems to be items in your cart whose offers have expired, please remove them from your cart and add them again")
	}
	if len(expiredCartItems) != 0 {
		return domain.Orders{}, fmt.Errorf("there seems to be items in your cart whose offers have expired, please remove them from your cart and add them again")
	}

	//Find the default address of the user
	var addressId int
	address := `SELECT id FROM addresses WHERE users_id=$1 AND is_default=true`
	err = tx.Raw(address, id).Scan(&addressId).Error
	if err != nil {
		tx.Rollback()
		return domain.Orders{}, err
	}
	if addressId == 0 {
		tx.Rollback()
		return domain.Orders{}, fmt.Errorf("add address pls")
	}
	if paymentTypeId == 3 {
		var userWallet domain.UserWallet

		wallet := `SELECT * FROM user_wallets WHERE users_id=$1`
		err := c.DB.Raw(wallet, id).Scan(&userWallet).Error
		if err != nil {
			return domain.Orders{}, err
		}

		if cart.Total <= userWallet.Amount {
			userWallet.Amount = userWallet.Amount - cart.Total
			updateWallet := `UPDATE user_wallets SET amount=$1 WHERE users_id=$2`
			err := c.DB.Raw(updateWallet, userWallet.Amount, id).Scan(&userWallet).Error
			if err != nil {
				return domain.Orders{}, err
			}

			// cart.Total = 0
			paymentStatus = 3
			orderStatus = 2
		} else {
			return domain.Orders{}, fmt.Errorf("Wallet does not have enough money")
		}

	}

	//Add the details to the orders and return the orderid
	var order domain.Orders
	insetOrder := `INSERT INTO orders (user_id,order_date,payment_type_id,shipping_address,order_total,order_status_id,payment_status_id)
		VALUES($1,NOW(),$2,$3,$4,$5,$6) RETURNING *`
	err = tx.Raw(insetOrder, id, paymentTypeId, addressId, cart.Total, orderStatus, paymentStatus).Scan(&order).Error
	if err != nil {
		tx.Rollback()
		return domain.Orders{}, err
	}

	//Get the cart item details of the user
	var cartItmes []helper.CartItems
	cartDetail := `select ci.model_id,ci.quantity,pi.price,pi.qty_in_stock  from cart_items ci join models pi on ci.model_id = pi.id where ci.carts_id=$1`
	err = tx.Raw(cartDetail, cart.Id).Scan(&cartItmes).Error
	if err != nil {
		tx.Rollback()
		return domain.Orders{}, err
	}

	//Add the items in the cart into the orderitems one by one
	for _, items := range cartItmes {
		//check whether the item is available
		if items.Quantity > items.QtyInStock {
			return domain.Orders{}, fmt.Errorf("out of stock")
		}
		insetOrderItems := `INSERT INTO order_items (orders_id,model_id,quantity,price) VALUES($1,$2,$3,$4)`
		err = tx.Exec(insetOrderItems, order.Id, items.ModelId, items.Quantity, items.Price).Error

		if err != nil {
			tx.Rollback()
			return domain.Orders{}, err
		}
	}

	//Update the cart total
	updateCart := `UPDATE carts SET total=0,sub_total=0 WHERE user_id=?`
	err = tx.Exec(updateCart, id).Error
	if err != nil {
		tx.Rollback()
		return domain.Orders{}, err
	}

	//Remove the items from the cart_items
	for _, items := range cartItmes {
		removeCartItems := `DELETE FROM cart_items WHERE carts_id =$1 AND model_id=$2`
		err = tx.Exec(removeCartItems, cart.Id, items.ModelId).Error
		if err != nil {
			tx.Rollback()
			return domain.Orders{}, err
		}
	}

	//Reduce the product qty in stock details
	for _, items := range cartItmes {
		updateQty := `UPDATE models SET qty_in_stock=models.qty_in_stock-$1 WHERE id=$2`
		err = tx.Exec(updateQty, items.Quantity, items.ModelId).Error
		if err != nil {
			tx.Rollback()
			return domain.Orders{}, err
		}
	}

	//update the PaymentDetails table with OrdersID, OrderTotal, PaymentTypeID, PaymentStatusID
	createPaymentDetails := `INSERT INTO payment_details
			(orders_id,
			order_total,
			payment_type_id,
			payment_status_id,
			updated_at)
			VALUES($1,$2,$3,$4,NOW())`
	if err = tx.Exec(createPaymentDetails, order.Id, order.OrderTotal, paymentTypeId, paymentStatus).Error; err != nil {
		tx.Rollback()
		return domain.Orders{}, err
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return domain.Orders{}, err
	}
	return order, nil
}

//-------------------------- Cancel-Order --------------------------//

func (c *OrderDatabase) UserCancelOrder(orderId, userId int) error {

	tx := c.DB.Begin()

	//find the orderd product and qty and update the product_items with those
	var items []helper.CartItems
	findProducts := `SELECT model_id,quantity FROM order_items WHERE orders_id=?`
	err := tx.Raw(findProducts, orderId).Scan(&items).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	if len(items) == 0 {
		return fmt.Errorf("no order found with this id")
	}
	for _, item := range items {
		updateProductItem := `UPDATE models SET qty_in_stock=qty_in_stock+$1 WHERE id=$2`
		err = tx.Exec(updateProductItem, item.Quantity, item.ModelId).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	var orders domain.Orders
	var payment domain.PaymentDetails

	findOrders := `SELECT * FROM orders WHERE id=?`
	err = tx.Raw(findOrders, orderId).Scan(&orders).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	findPayment := `SELECT * FROM orders WHERE id=?`
	err = tx.Raw(findPayment, orderId).Scan(&payment).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	if orders.PaymentTypeId == 2 || orders.PaymentTypeId == 3 && payment.PaymentStatusID == 3 {
		var userWallet domain.UserWallet

		wallet := `SELECT * FROM user_wallets WHERE users_id=$1`
		err := c.DB.Raw(wallet, userId).Scan(&userWallet).Error
		if err != nil {
			return err
		}

		updateWallet := `UPDATE user_wallets SET amount=amount + ? WHERE users_id=?`
		err = c.DB.Exec(updateWallet, orders.OrderTotal, userId).Error
		if err != nil {
			return err
		}

	}

	//Remove the items from order_items
	removeItems := `DELETE FROM order_items WHERE orders_id=$1`
	err = tx.Exec(removeItems, orderId).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	//update the order status as canceled
	cancelOrder := `UPDATE orders SET order_status_id=$1,payment_status_id=$2 WHERE id=$3 AND user_id=$4`
	err = tx.Exec(cancelOrder, 5, 2, orderId, userId).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil

}

//-------------------------- List-Order --------------------------//

func (c *OrderDatabase) ListOrder(userId, orderId int) (domain.Orders, error) {
	var order domain.Orders
	findOrder := `SELECT * FROM orders WHERE user_id=$1 AND id=$2`
	err := c.DB.Raw(findOrder, userId, orderId).Scan(&order).Error
	return order, err
}

//-------------------------- List-All-Order --------------------------//

func (c *OrderDatabase) ListAllOrders(userId int) ([]domain.Orders, error) {
	var orders []domain.Orders

	findOrders := `SELECT * FROM orders WHERE user_id=?`
	err := c.DB.Raw(findOrders, userId).Scan(&orders).Error
	return orders, err
}

//-------------------------- Update-Order --------------------------//

func (c *OrderDatabase) UpdateOrder(updateOrder helper.UpdateOrder) error {
	//check whether there is a order with this order number

	var isExists bool
	query1 := `select exists(select 1 from orders where id=?)`
	err := c.DB.Raw(query1, updateOrder.OrderId).Scan(&isExists).Error
	if err != nil {
		return err
	}
	if !isExists {
		return fmt.Errorf("no such order")
	}
	updateOrderQry := `UPDATE orders SET order_status_id=$1 WHERE id=$2`
	err = c.DB.Exec(updateOrderQry, updateOrder.OrderStatusID, updateOrder.OrderId).Error
	if err != nil {
		return err
	}
	return nil
}

// ListAllOrdersForAdmin implements interfaces.OrderRepository.
func (c *OrderDatabase) ListAllOrdersForAdmin() ([]response.AdminOrder, error) {
	var orders []response.AdminOrder
	findOrders := `SELECT orders.id AS order_id,orders.payment_type_id,order_statuses.status AS order_status,payment_types.type AS payment_type
	FROM orders JOIN order_statuses ON orders.order_status_id=order_statuses.id
	JOIN payment_types ON orders.payment_type_id=payment_types.id`
	err := c.DB.Raw(findOrders).Scan(&orders).Error
	return orders, err
}

//-------------------------- Return-Order --------------------------//

func (c *OrderDatabase) ReturnOrder(userId, orderId int) (int, error) {
	var orders domain.Orders
	var userWallet domain.UserWallet

	wallet := `SELECT * FROM user_wallets WHERE users_id=$1`
	err := c.DB.Raw(wallet, userId).Scan(&userWallet).Error
	if err != nil {
		return 0, err
	}

	getOrderDetails := `SELECT * FROM orders WHERE user_id=$1 AND id=$2`
	err = c.DB.Raw(getOrderDetails, userId, orderId).Scan(&orders).Error
	if err != nil {
		return 0, err
	}
	if orders.OrderStatusID != 4 {
		return 0, fmt.Errorf("the order is not deleverd")
	}

	updateWallet := `UPDATE user_wallets SET amount=amount + ? WHERE users_id=?`
	err = c.DB.Exec(updateWallet, orders.OrderTotal, userId).Error
	if err != nil {
		return 0, err
	}

	returnOder := `UPDATE orders SET order_status_id=$1 WHERE id=$2`
	err = c.DB.Exec(returnOder, 6, orderId).Error
	if err != nil {
		return 0, err
	}
	return orders.OrderTotal, nil

}

func (c *OrderDatabase) UserIdFromOrder(id int) (int, error) {
	var userId int
	err := c.DB.Raw(`SELECT user_id FROM orders WHERE id=?`, id).Scan(&userId).Error
	return userId, err
}
