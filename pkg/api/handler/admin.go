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

type AdminHandler struct {
	adminUseCase services.AdminUseCase
}

func NewAdminHandler(usecase services.AdminUseCase) *AdminHandler {
	return &AdminHandler{
		adminUseCase: usecase,
	}
}

//-------------------------- Login --------------------------//

func (cr *AdminHandler) AdminLogin(c *gin.Context) {
	var admin helper.LoginReq
	err := c.BindJSON(&admin)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "unable to bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	token, err := cr.adminUseCase.AdminLogin(admin)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "failed to login",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("adminToken", token, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "login succesfully",
		Data:       nil,
		Errors:     nil,
	})
}

//-------------------------- Log-Out --------------------------//

func (cr *AdminHandler) AdminLogout(c *gin.Context) {
	c.SetCookie("adminToken", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "admin logouted",
		Data:       nil,
		Errors:     nil,
	})
}



//-------------------------- Show-Single-User --------------------------//

func (cr *AdminHandler) ShowUser(c *gin.Context) {
	paramID := c.Param("user_id")
	userID, err := strconv.Atoi(paramID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	user, err := cr.adminUseCase.ShowUser(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant find user",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "user details",
		Data:       user,
		Errors:     nil,
	})

}

//-------------------------- Show-All-Users --------------------------//

func (cr *AdminHandler) ShowAllUsers(c *gin.Context) {
	users, err := cr.adminUseCase.ShowAllUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant find user",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "users are",
		Data:       users,
		Errors:     nil,
	})

}

//-------------------------- Block-User --------------------------//

func (cr *AdminHandler) BlockUser(c *gin.Context) {
	var body helper.BlockData
	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	adminID, err := handlerutil.GetAdminIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't find SupAdminId",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.adminUseCase.BlockUser(body, adminID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "Can't Block",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "User Blocked",
		Data:       nil,
		Errors:     nil,
	})
}

//-------------------------- UnBlock-User --------------------------//

func (cr *AdminHandler) UnblockUser(c *gin.Context) {
	paramsId := c.Param("user_id")
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
	err = cr.adminUseCase.UnblockUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant unblock user",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "user unblocked",
		Data:       nil,
		Errors:     nil,
	})
}
