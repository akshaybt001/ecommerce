package repository

import (
	"fmt"

	"gorm.io/gorm"
	"main.go/pkg/common/helper"
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

	//Add the details to the orders and return the orderid
	var order domain.Orders
	insetOrder := `INSERT INTO orders (user_id,order_date,payment_type_id,shipping_address,order_total,order_status_id)
		VALUES($1,NOW(),$2,$3,$4,1) RETURNING *`
	err = tx.Raw(insetOrder, id, paymentTypeId, addressId, cart.Total).Scan(&order).Error
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
	if err = tx.Exec(createPaymentDetails, order.Id, order.OrderTotal, paymentTypeId, 1).Error; err != nil {
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
	//Remove the items from order_items
	removeItems := `DELETE FROM order_items WHERE orders_id=$1`
	err = tx.Exec(removeItems, orderId).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	//update the order status as canceled
	cancelOrder := `UPDATE orders SET order_status_id=$1 WHERE id=$2 AND user_id=$3`
	err = tx.Exec(cancelOrder, 5, orderId, userId).Error
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

// // ListAllOrdersForAdmin implements interfaces.OrderRepository.
// func (o *OrderDatabase) ListAllOrdersForAdmin() ([]response.AdminOrder, error) {
// 	var orders []response.AdminOrder
// 	findOrders := `SELECT orders.id AS order_id,orders.payment_type_id,order_statuses.status AS order_status,payment_types.type AS payment_type,payment_statuses.status AS payment_status
// 	FROM orders JOIN order_statuses ON orders.order_status_id=order_statuses.id
// 	JOIN payment_types ON orders.payment_type_id=payment_types.id 
// 	JOIN payment_statuses ON orders.payment_status_id=payment_statuses.id
//      `
// 	err := o.DB.Raw(findOrders).Scan(&orders).Error
// 	return orders,err
// }