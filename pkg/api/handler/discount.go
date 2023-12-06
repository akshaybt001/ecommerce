package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"main.go/pkg/common/helper"
	"main.go/pkg/common/response"
	services "main.go/pkg/usecase/interface"
)

type DiscountHandler struct {
	discountUsecase services.DiscountUsecase
}

func NewDiscountHandler(discountusecase services.DiscountUsecase) *DiscountHandler {
	return &DiscountHandler{
		discountUsecase: discountusecase,
	}
}

// -------------------------- Add-Discount --------------------------//

func (cr *DiscountHandler) AddDiscount(c *gin.Context) {
	var discount helper.Discount
	err := c.Bind(&discount)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't bind ",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	err = cr.discountUsecase.AddDiscount(discount)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't create",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "discount added",
		Data:       nil,
		Errors:     nil,
	})

}

// -------------------------- Edit-Discount --------------------------//

func (cr *DiscountHandler) EditDiscount(c *gin.Context) {
	var discount helper.Discount
	err := c.Bind(&discount)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't Bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	updatedDiscount, err := cr.discountUsecase.EditDiscount(id, discount)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't update productitem",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Discount updated",
		Data:       updatedDiscount,
		Errors:     nil,
	})

}

// -------------------------- Delete-Discount --------------------------//

func (cr *DiscountHandler) DeleteDiscount(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.discountUsecase.DeleteDiscount(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't delete item",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "item deleted",
		Data:       nil,
		Errors:     nil,
	})
}

// -------------------------- List-All-Model --------------------------//

func (cr *DiscountHandler) ListAllDiscount(c *gin.Context) {

	discount, err := cr.discountUsecase.ListAllDiscount()

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't discount items",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "discount items are",
		Data:       discount,
		Errors:     nil,
	})
}

// -------------------------- List-Single-Discount --------------------------//

func (cr *DiscountHandler) ListDiscount(c *gin.Context) {

	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find discountid",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	discount, err := cr.discountUsecase.ListDiscount(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find discount",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "discount",
		Data:       discount,
		Errors:     nil,
	})
}
