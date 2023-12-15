package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
	handlerutil "main.go/pkg/api/handlerUtil"
	"main.go/pkg/common/helper"
	"main.go/pkg/common/response"
	services "main.go/pkg/usecase/interface"
)

type OrderHandler struct {
	OrderUseCase services.OrderUseCase
}

func NewOrderHandler(orderUseCase services.OrderUseCase) *OrderHandler {
	return &OrderHandler{
		OrderUseCase: orderUseCase,
	}
}

//-------------------------- Order-All --------------------------//

func (cr *OrderHandler) OrderAll(c *gin.Context) {
	paramsId := c.Param("payment_id")
	paymentTypeId, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	userId, err := handlerutil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant find userid",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	order, err := cr.OrderUseCase.OrderAll(userId, paymentTypeId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant place order",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "orderplaced",
		Data:       order,
		Errors:     nil,
	})
}

//-------------------------- Cancel-Order --------------------------//

func (cr *OrderHandler) UserCancelOrder(c *gin.Context) {
	userId, err := handlerutil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant find userid",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	paramsId := c.Param("orderId")
	orderId, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.OrderUseCase.UserCancelOrder(orderId, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't cancel order",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "order canceled",
		Data:       nil,
		Errors:     nil,
	})
}

//-------------------------- List-Order --------------------------//

func (cr *OrderHandler) ListOrder(c *gin.Context) {
	userId, err := handlerutil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant find userid",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	paramsId := c.Param("orderId")
	orderId, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	order, err := cr.OrderUseCase.ListOrder(userId, orderId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find order",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "order ",
		Data:       order,
		Errors:     nil,
	})
}

//-------------------------- List-All-Order --------------------------//

func (cr *OrderHandler) ListAllOrders(c *gin.Context) {
	Id, err := handlerutil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't find Id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	orders, err := cr.OrderUseCase.ListAllOrders(Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find order",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "order ",
		Data:       orders,
		Errors:     nil,
	})
}

//-------------------------- Update-Order --------------------------//

func (cr *OrderHandler) UpdateOrder(c *gin.Context) {
	var UpdateOrder helper.UpdateOrder
	err := c.BindJSON(&UpdateOrder)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't bind ",
			Data:       nil,
			Errors:     err.Error(),
		})
	}
	err = cr.OrderUseCase.UpdateOrder(UpdateOrder)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't update order",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "order updated ",
		Data:       nil,
		Errors:     nil,
	})
}

//-------------------------- Return-Order --------------------------//

func (cr *OrderHandler) ReturnOrder(c *gin.Context) {
	userId, err := handlerutil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant find userid",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	paramsId := c.Param("orderId")
	orderId, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	returnAmount, err := cr.OrderUseCase.ReturnOrder(userId, orderId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't return order",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "order returnd ",
		Data:       returnAmount,
		Errors:     nil,
	})
}
func (cr *OrderHandler) ListAllOrderForAdmin(c *gin.Context) {
	order, err := cr.OrderUseCase.ListAllOrderForAdmin()
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "error listing all orders",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "orders listed successfully",
		Data:       order,
		Errors:     nil,
	})
}

// //------------------- Add Order Status --------------------------------

// func (cr *OrderHandler) AddOrderStatus(c *gin.Context) {
// 	var orderStatus helper.OrderStatus
// 	err := c.BindJSON(&orderStatus)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, response.Response{
// 			StatusCode: 400,
// 			Message:    "error binding json",
// 			Data:       nil,
// 			Errors:     err.Error(),
// 		})
// 		return
// 	}
// 	newOrderStatus, err := cr.OrderUseCase.AddOrderStatus(orderStatus)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, response.Response{
// 			StatusCode: 400,
// 			Message:    "error adding new orderStatus",
// 			Data:       nil,
// 			Errors:     err.Error(),
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, response.Response{
// 		StatusCode: 200,
// 		Message:    "orderstatus added successfully",
// 		Data:       newOrderStatus,
// 		Errors:     nil,
// 	})
// }


func (cr *OrderHandler) InvoiceDownload(c *gin.Context) {
	paramId := c.Param("orderId")
	orderId, err := strconv.Atoi(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "error converting ordeId to integer",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	userId, err := handlerutil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "error retrieving userId",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	order, err := cr.OrderUseCase.ListOrder(userId, orderId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "error retrieving orderInformation",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	if order.PaymentStatus != "completed" {
		err = fmt.Errorf("please complete the payment to download the invoice")
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Cannot download invoice unless you have completed the payment",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	// Generate PDF from JSON data
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 18)
	pdf.Cell(40, 10, "Invoice")
	pdf.Ln(10)

	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Techify")
	pdf.Ln(8)
	pdf.Cell(40, 10, "Contact : 8592817810")
	pdf.Ln(10)

	// Add order information
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Order Information")
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, fmt.Sprintf("Order ID: %d", order.Id))
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("Customer Name: %s", order.UserName))
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("Order Date: %s", order.OrderDate.Format("2006-01-02 15:04:05")))
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("Payment Type: %s", order.PaymentType))
	pdf.Ln(8)

	// Add shipping address
	pdf.Ln(10)
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(40, 10, "Shipping Address")
	pdf.Ln(8)

	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, fmt.Sprintf("House Number: %s", order.House_number))
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("Street: %s", order.Street))
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("City: %s", order.City))
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("District: %s", order.District))
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("Landmark: %s", order.Landmark))
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("Pincode: %d", order.Pincode))
	pdf.Ln(10)

	// Add product details
	pdf.Ln(10)
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(40, 10, "Product Details")
	pdf.Ln(8)

	pdf.SetFont("Arial", "", 12)

	pdf.Cell(40, 10, fmt.Sprintf("Product: %s", order.ModelName))
	pdf.Ln(8)
	pdf.Cell(40, 10, fmt.Sprintf("Price: Rs.%d", order.OrderTotal))
	pdf.Ln(8)

	
	pdf.Ln(10)
	pdf.SetFont("Arial", "", 10)
	pdf.Cell(40, 10, "Thank you for shopping with us")
	pdf.Ln(10)

	// Set headers for file download
	c.Header("Content-Disposition", "attachment; filename=order.pdf")
	c.Header("Content-Type", "application/pdf")

	// Output the PDF to the response writer
	err = pdf.Output(c.Writer)
	if err != nil {
		// Handle error appropriately
		fmt.Println("Error generating PDF:", err)
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "downloaded response successfully",
		Data:       nil,
		Errors:     nil,
	})
}