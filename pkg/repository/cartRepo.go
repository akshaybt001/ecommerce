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
	// var proitems domain.Model
	// findPrice := `SELECT * FROM models WHERE id=$1`
	// err = tx.Raw(findPrice, productId).Scan(&proitems).Error
	// if err != nil {
	// 	tx.Rollback()
	// 	return err
	// }

	var productdetails domain.Model
	findProductDetails := `SELECT * FROM models WHERE id = ?`
	err = tx.Raw(findProductDetails, productId).Scan(&productdetails).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	// fmt.Printf("brand id is %d", productdetails.Brand_id)

	// type Amount struct {
	// 	Subtotal int
	// 	Total    int
	// }

	//Updating the subtotal in cart table
	var subtotal int
	updateSubTotal := `UPDATE carts SET sub_total=carts.sub_total+$1 WHERE user_id=$2 RETURNING sub_total`
	err = tx.Raw(updateSubTotal, productdetails.Price, userId).Scan(&subtotal).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// getting the subtotal and total
	var amount domain.Carts
	findAmount := `SELECT sub_total,total FROM carts WHERE user_id=$1`
	err = tx.Raw(findAmount, userId).Scan(&amount).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	fmt.Println("cart subtotal is ", amount.SubTotal)

	fmt.Println("cart total is ", amount.Total)

	var discount domain.Discount
	finddiscount := `SELECT * FROM discounts WHERE brand_id = ? AND expiration_date>NOW()`
	err = tx.Raw(finddiscount, productdetails.Brand_id).Scan(&discount).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	fmt.Printf("discount id is %d", discount.Id)

	if discount.Id == 0 {
		total := 0
		if amount.SubTotal == int(productdetails.Price) {
			total = amount.SubTotal
		} else {
			total = amount.Total + int(productdetails.Price)
		}
		updateTotal := `UPDATE carts SET total=$1 WHERE id=$2`
		err = tx.Exec(updateTotal, total, cartId).Error
		if err != nil {
			tx.Rollback()
			return err
		}

	} else {
		disc := 0
		total := 0
		if int(productdetails.Price) >= discount.MinimumPurchaseAmount {
			disc = ((amount.SubTotal * int(discount.DiscountPercent)) / 100)
			if disc > discount.DiscountMaximumAmount {
				disc = discount.DiscountMaximumAmount
			}
		}

		if amount.SubTotal == int(productdetails.Price) {
			total = amount.SubTotal - disc
		} else {
			total = amount.Total + int(productdetails.Price) - disc
		}

		fmt.Println("discouted price is ", total)

		updateTotal := `UPDATE carts SET total=$1 WHERE id=$2`
		err = tx.Exec(updateTotal, total, cartId).Error
		if err != nil {
			tx.Rollback()
			return err
		}
		if disc != 0 {
			updateIsDiscounted := `UPDATE cart_items SET is_discounted=true WHERE carts_id=$1`
			err = tx.Exec(updateIsDiscounted, cartId).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		}

		fmt.Println("cart id is ", cartItemId)

	}

	// Update the subtotal in cart table with discount check
	// var total int
	// updateSubTotal := `
	// 	UPDATE carts
	// 	SET sub_total = CASE WHEN sub_total+$1 < $2::int THEN sub_total
	// 						ELSE sub_total - ((sub_total * $3::int) / 100)
	// 				   END
	// 	WHERE user_id = $4 RETURNING sub_total
	//  `
	// err = tx.Raw(updateSubTotal, proitems.Price, int(discount.MinimumPurchaseAmount), int(discount.DiscountPercent), userId).Scan(&total).Error
	// if err != nil {
	// 	tx.Rollback()
	// 	return err
	// }

	// if total < int(discount.MinimumPurchaseAmount) {
	// 	// No need for explicit rollback as updateSubTotal handles non-discounted price
	// 	return nil
	// }
	// //Updating the subtotal in cart table
	// var total int
	// updateSubTotal := `UPDATE carts SET sub_total=carts.sub_total+$1 WHERE user_id=$2 RETURNING sub_total`
	// err = tx.Raw(updateSubTotal, proitems.Price, userId).Scan(&total).Error
	// if err != nil {
	// 	tx.Rollback()
	// 	return err
	// }

	// // if discount.MinimumPurchaseAmount > 0 && total < int(discount.MinimumPurchaseAmount) {
	// // 	tx.Rollback()
	// // 	return fmt.Errorf("the minimum purchase amount condition is not net for the offer")
	// // }

	// // Calculating the discounted price

	// discountedPrice := total - ((total * int(discount.DiscountPercent)) / 100)
	// if discountedPrice < 0 {
	// 	discountedPrice = 0
	// }

	// subtotal := discountedPrice

	// updateTotal := `UPDATE carts SET sub_total=$1 WHERE id=$2`
	// err = tx.Exec(updateTotal, total, cartId).Error
	// if err != nil {
	// 	tx.Rollback()
	// 	return err
	// }

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
	var productdetails domain.Model
	findproductdetails := `SELECT * FROM models WHERE id=$1`
	err = tx.Raw(findproductdetails, productId).Scan(&productdetails).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// Retrieve the discount amount for the product and user
	var discount domain.Discount
	findDiscountAmount := `SELECT * FROM discounts WHERE brand_id =$1`
	err = tx.Raw(findDiscountAmount, productdetails.Brand_id).Scan(&discount).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	// Calculate the total discount amount
	// var totalDiscountAmount = (discountAmount * productdetails.Price) / 100
	// fmt.Printf("the total discount amount %d", totalDiscountAmount)

	// disc := ((productdetails.Price * discount.DiscountPercent) / 100)
	// if disc > discount.DiscountMaximumAmount {
	// 	disc = discount.DiscountMaximumAmount
	// }
	disc := 0
	if int(productdetails.Price) >= discount.MinimumPurchaseAmount {
		disc = ((int(productdetails.Price) * int(discount.DiscountPercent)) / 100)
		if disc > discount.DiscountMaximumAmount {
			disc = discount.DiscountMaximumAmount
		}
	}
	fmt.Printf("the total discount amount %d", disc)

	// updatediscountamount := `UPDATE carts SET total=total+$1 WHERE user_id=$2`
	// err = tx.Raw(updatediscountamount, disc, userId).Error
	// if err != nil {
	// 	tx.Rollback()
	// 	return err
	// }

	//Update the subtotal reduce the price of the cart total with price of the product
	// var subTotal int

	// updateSubTotal := `UPDATE carts SET sub_total=sub_total-
	// $1 WHERE user_id=$2 RETURNING sub_total`
	// err = tx.Raw(updateSubTotal, productdetails.Price, userId).Scan(&subTotal).Error
	// if err != nil {
	// 	tx.Rollback()
	// 	return err
	// }
	updateSubTotal := `UPDATE carts SET sub_total = sub_total - $1, total = total - $1 + $2 WHERE user_id = $3`
	err = tx.Exec(updateSubTotal, productdetails.Price, disc, userId).Error
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
