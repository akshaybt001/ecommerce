package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"main.go/pkg/common/helper"
	"main.go/pkg/common/response"
	services "main.go/pkg/usecase/interface"
)

type ProductHandler struct {
	productUsecase services.ProductUsecase
}

func NewProductHandler(productUsecase services.ProductUsecase) *ProductHandler {
	return &ProductHandler{
		productUsecase: productUsecase,
	}
}

// -------------------------- Create-Category --------------------------//

func (cr *ProductHandler) CreateCategory(c *gin.Context) {
	var category helper.Category
	err := c.Bind(&category)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind failed",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	NewCategoery, err := cr.productUsecase.CreateCategory(category)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't create category",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Category Created",
		Data:       NewCategoery,
		Errors:     nil,
	})
}

func (cr *ProductHandler) UpdateCategory(c *gin.Context) {
	var category helper.Category
	err := c.Bind(&category)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind failed",
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
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	updatedCategory, err := cr.productUsecase.UpdateCategory(category, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't update category",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Category Updated",
		Data:       updatedCategory,
		Errors:     nil,
	})
}

// -------------------------- Delete-Category --------------------------//

func (cr *ProductHandler) DeleteCategory(c *gin.Context) {
	parmasId := c.Param("id")
	id, err := strconv.Atoi(parmasId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't bind data",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	err = cr.productUsecase.DeleteCategory(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't dlete category",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Category deleted",
		Data:       nil,
		Errors:     nil,
	})

}

func (cr *ProductHandler) ListAllCategories(c *gin.Context) {
	categories, err := cr.productUsecase.ListAllCategories()
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find category",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Categories are",
		Data:       categories,
		Errors:     nil,
	})
}

// -------------------------- List-Single-Category --------------------------//

func (cr *ProductHandler) ListCategory(c *gin.Context) {
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
	category, err := cr.productUsecase.ListCategory(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find category",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Category is",
		Data:       category,
		Errors:     nil,
	})
}

// -------------------------- Create-Brand --------------------------//

func (cr *ProductHandler) AddBrand(c *gin.Context) {
	var Brand helper.Brands
	err := c.Bind(&Brand)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	newBrand, err := cr.productUsecase.AddBrand(Brand)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't add Brand",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Brand Added",
		Data:       newBrand,
		Errors:     nil,
	})

}

// -------------------------- Update-Brand --------------------------//

func (cr *ProductHandler) UpdateBrand(c *gin.Context) {
	var Brand helper.Brands
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Cant find id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	if err := c.Bind(&Brand); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Cant bind body",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	updatedBrand, err := cr.productUsecase.UpdateBrand(id, Brand)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Cant update brand",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusBadRequest, response.Response{
		StatusCode: 200,
		Message:    "Brand updated",
		Data:       updatedBrand,
		Errors:     nil,
	})

}

func (cr *ProductHandler) DeleteBrand(c *gin.Context) {

	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find Brandid",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.productUsecase.DeleteBrand(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't delete brand",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "brand deleted",
		Data:       nil,
		Errors:     nil,
	})
}

// -------------------------- List-All-Brand --------------------------//

func (cr *ProductHandler) ListAllBrand(c *gin.Context) {

	brands, err := cr.productUsecase.ListAllBrand()
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find Brands",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Brands",
		Data:       brands,
		Errors:     nil,
	})
}

// -------------------------- List-Single-Brand --------------------------//

func (cr *ProductHandler) ListBrand(c *gin.Context) {

	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find brandid",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	brand, err := cr.productUsecase.ListBrand(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find brand",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "brand",
		Data:       brand,
		Errors:     nil,
	})
}

// -------------------------- Add-Model --------------------------//

func (cr *ProductHandler) AddModel(c *gin.Context) {
	var model helper.Model
	err := c.Bind(&model)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Cant bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	newModel, err := cr.productUsecase.AddModel(model)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Cant create",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "product created",
		Data:       newModel,
		Errors:     nil,
	})
}

// -------------------------- Add-Model --------------------------//

func (cr *ProductHandler) UpdateModel(c *gin.Context) {
	var model helper.Model
	err := c.Bind(&model)
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

	updatedItem, err := cr.productUsecase.UpdateModel(id, model)
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
		Message:    "productitem updated",
		Data:       updatedItem,
		Errors:     nil,
	})

}

// -------------------------- Delete-Model --------------------------//

func (cr *ProductHandler) DeleteModel(c *gin.Context) {
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
	err = cr.productUsecase.DeleteModel(id)
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

func (cr *ProductHandler) ListAllModel(c *gin.Context) {

	var viewProductaItem helper.QueryParams

	viewProductaItem.Page, _ = strconv.Atoi(c.Query("page"))
	viewProductaItem.Limit, _ = strconv.Atoi(c.Query("limit"))
	viewProductaItem.Query = c.Query("query")
	viewProductaItem.Filter = c.Query("filter")
	viewProductaItem.SortBy = c.Query("sort_by")
	viewProductaItem.SortDesc, _ = strconv.ParseBool(c.Query("sort_desc"))

	model, err := cr.productUsecase.ListAllModel(viewProductaItem)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't disaply items",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "product items are",
		Data:       model,
		Errors:     nil,
	})
}

// -------------------------- List-Single-Model --------------------------//

func (cr *ProductHandler) ListModel(c *gin.Context) {
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
	productItem, err := cr.productUsecase.ListModel(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "can't find product",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "product",
		Data:       productItem,
		Errors:     nil,
	})

}

// -------------------------- Upload-Image --------------------------//

func (cr *ProductHandler) UploadImage(c *gin.Context) {

	id := c.Param("id")
	productId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant find product id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	// Multipart form
	form, _ := c.MultipartForm()

	files := form.File["images"]

	for _, file := range files {
		// Upload the file to specific dst.
		c.SaveUploadedFile(file, "../asset/uploads/"+file.Filename)

		err := cr.productUsecase.UploadImage(file.Filename, productId)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Response{
				StatusCode: 400,
				Message:    "cant upload images",
				Data:       nil,
				Errors:     err.Error(),
			})
			return
		}

	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "image uploaded",
		Data:       nil,
		Errors:     nil,
	})
	return
}
