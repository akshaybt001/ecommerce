package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
