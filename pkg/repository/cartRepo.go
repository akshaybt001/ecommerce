package repository

import (
	"fmt"

	"gorm.io/gorm"
	"main.go/pkg/common/response"
	"main.go/pkg/domain"
	interfaces "main.go/pkg/repository/interface"
)

type CartDatabase struct {
	DB *gorm.DB
}

func NewCartRepository(DB *gorm.DB) interfaces.CartRepository {
	return &CartDatabase{DB}
}

//----------------------Create-Cart -----------------------

func (c *CartDatabase) CreateCart(id int) error {
	query := `INSERT INTO carts (user_id,sub_total,total) VALUES($1,0,0)`
	err := c.DB.Exec(query, id).Error
	return err
}

// -------------------------- ADD-To-Cart --------------------------//

func (c *CartDatabase) AddToCart(productId, userId int) error {
	tx := c.DB.Begin()
	//finding cart id coresponding to the user
	var cartId int
	findCartId := `SELECT id FROM carts WHERE user_id=? `
	err := tx.Raw(findCartId, userId).Scan(&cartId).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	//Check whether the product exists in the cart_items
	var cartItemId int
	cartItemCheck := `SELECT id FROM cart_items WHERE carts_id = $1 AND model_id = $2 LIMIT 1`
	err = tx.Raw(cartItemCheck, cartId, productId).Scan(&cartItemId).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	if cartItemId == 0 {
		addToCart := `INSERT INTO cart_items (carts_id,model_id,quantity)VALUES($1,$2,1)`
		err = tx.Exec(addToCart, cartId, productId).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	} else {
		updatCart := `UPDATE cart_items SET quantity = cart_items.quantity+1 WHERE id = $1 `
		err = tx.Exec(updatCart, cartItemId).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	//finding the price of the product
	var price int
	findPrice := `SELECT price FROM models WHERE id=$1`
	err = tx.Raw(findPrice, productId).Scan(&price).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	//Updating the subtotal in cart table
	var subtotal int
	updateSubTotal := `UPDATE carts SET sub_total=carts.sub_total+$1 WHERE user_id=$2 RETURNING sub_total`
	err = tx.Raw(updateSubTotal, price, userId).Scan(&subtotal).Error
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

// -------------------------- Remove-From-Cart --------------------------//

func (c *CartDatabase) RemoveFromCart(userId, productId int) error {
	tx := c.DB.Begin()

	//Find cart id
	var cartId int
	findCartId := `SELECT id FROM carts WHERE user_id=? `
	err := tx.Raw(findCartId, userId).Scan(&cartId).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	//Find the qty of the product in cart
	var qty int
	findQty := `SELECT quantity FROM cart_items WHERE carts_id=$1 AND model_id=$2`
	err = tx.Raw(findQty, cartId, productId).Scan(&qty).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	if qty == 0 {
		tx.Rollback()
		return fmt.Errorf("no items in cart to reomve")
	}

	//If the qty is 1 dlt the product from the cart
	if qty == 1 {
		dltItem := `DELETE FROM cart_items WHERE carts_id=$1 AND model_id=$2`
		err := tx.Exec(dltItem, cartId, productId).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	} else { // If there is  more than one product reduce the qty by 1
		updateQty := `UPDATE cart_items SET quantity=cart_items.quantity-1 WHERE carts_id=$1 AND model_id=$2`
		err = tx.Exec(updateQty, cartId, productId).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	//Find the price of the product item
	var price int
	productPrice := `SELECT price FROM models WHERE id=$1`
	err = tx.Raw(productPrice, productId).Scan(&price).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	//Update the subtotal reduce the price of the cart total with price of the product
	var subTotal int
	updateSubTotal := `UPDATE carts SET sub_total=sub_total-$1 WHERE user_id=$2 RETURNING sub_total`
	err = tx.Raw(updateSubTotal, price, userId).Scan(&subTotal).Error
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

// -------------------------- List-Cart --------------------------//

func (c *CartDatabase) ListCart(userId int) (response.ViewCart, error) {
	tx := c.DB.Begin()
	//get cart details
	type cartDetails struct {
		Id       int
		SubTotal float64
		Total    float64
	}
	var cart cartDetails
	getCartDetails := `SELECT
		c.id,
		c.sub_total,
		c.total
		FROM carts c WHERE c.user_id=$1`
	err := tx.Raw(getCartDetails, userId).Scan(&cart).Error

	if err != nil {
		tx.Rollback()
		return response.ViewCart{}, err
	}
	//get cart_items details
	var cartItems domain.CartItem
	getCartItemsDetails := `SELECT * FROM cart_items WHERE carts_id=$1`
	err = tx.Raw(getCartItemsDetails, cart.Id).Scan(&cartItems).Error
	if err != nil {
		tx.Rollback()
		return response.ViewCart{}, err
	}
	//get the product details
	var details []response.DisplayCart
	getDetails := `SELECT p.brand, pi.sku AS productname, 
		pi.color,
		pi.ram,
		pi.battery,
		pi.storage,
		pi.camera,
		ci.quantity,
		pi.price AS price_per_unit,
		(pi.price*ci.quantity) AS total
		FROM cart_items ci JOIN models pi  ON ci.model_id = pi.id
		JOIN brands p ON pi.brand_id = p.id WHERE ci.carts_id=$1`
	err = tx.Raw(getDetails, cart.Id).Scan(&details).Error
	if err != nil {
		tx.Rollback()
		return response.ViewCart{}, err
	}

	var carts response.ViewCart
	carts.CartTotal = cart.Total
	carts.SubTotal = cart.SubTotal
	carts.CartItems = details
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return response.ViewCart{}, err
	}
	return carts, nil
}
