package handler

import (
	"encoding/csv"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

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

func (cr *AdminHandler) AdminDashBoard(c *gin.Context) {
	var filterDash helper.ReportParams

	filterDash.Status, _ = strconv.Atoi(c.Query("status"))
	filterDash.Day, _ = strconv.Atoi(c.Query("day"))
	filterDash.Week, _ = strconv.Atoi(c.Query("week"))
	filterDash.Month, _ = strconv.Atoi(c.Query("month"))
	filterDash.Year, _ = strconv.Atoi(c.Query("year"))
	filterDash.StartDate = c.Query("startdate")
	filterDash.EndDate = c.Query("enddate")

	dashBoard, err := cr.adminUseCase.GetDashBoard(filterDash)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant get dashboard",
			Data:       nil,
			Errors:     err.Error(),
		})
		return

	}
	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Dash board",
		Data:       dashBoard,
		Errors:     nil,
	})
}
func (cr *AdminHandler) ViewSalesReport(c *gin.Context) {

	var filterReport helper.ReportParams

	filterReport.Status, _ = strconv.Atoi(c.Query("status"))
	filterReport.Day, _ = strconv.Atoi(c.Query("day"))
	filterReport.Week, _ = strconv.Atoi(c.Query("week"))
	filterReport.Month, _ = strconv.Atoi(c.Query("month"))
	filterReport.Year, _ = strconv.Atoi(c.Query("year"))
	filterReport.StartDate = c.Query("startdate")
	filterReport.EndDate = c.Query("enddate")

	sales, err := cr.adminUseCase.ViewSalesReport(filterReport)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant get sales report",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Response{
		StatusCode: 200,
		Message:    "Sales report",
		Data:       sales,
		Errors:     nil,
	})

}
func (cr *AdminHandler) DownloadSalesReport(c *gin.Context) {

	var filterReport helper.ReportParams

	filterReport.Status, _ = strconv.Atoi(c.Query("status"))
	filterReport.Day, _ = strconv.Atoi(c.Query("day"))
	filterReport.Week, _ = strconv.Atoi(c.Query("week"))
	filterReport.Month, _ = strconv.Atoi(c.Query("month"))
	filterReport.Year, _ = strconv.Atoi(c.Query("year"))
	filterReport.StartDate = c.Query("startdate")
	filterReport.EndDate = c.Query("endday")

	sales, err := cr.adminUseCase.ViewSalesReport(filterReport)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "cant get sales report",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	// Set headers so browser will download the file
	c.Header("Content-Type", "text/csv")
	c.Header("Content-Disposition", "attachment;filename=sales.csv")

	// c.Header("Content-Type", "text/csv")
	// c.Header("Content-Disposition", "attachment;filename=sales.csv")

	// Create a CSV writer using our response writer as our io.Writer
	wr := csv.NewWriter(c.Writer)

	// Write CSV header row
	headers := []string{"Name", "PaymentType", "OrderDate", "OrderTotal"}
	if err := wr.Write(headers); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Write data rows
	for _, sale := range sales {
		row := []string{sale.Name, sale.PaymentType, sale.OrderDate.Format("2006-01-02 15:04:05"), strconv.Itoa(sale.OrderTotal)}
		if err := wr.Write(row); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	// Flush the writer's buffer to ensure all data is written to the client
	wr.Flush()
	if err := wr.Error(); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

}
