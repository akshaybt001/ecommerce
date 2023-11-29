package repository

import (
	"fmt"

	"gorm.io/gorm"
	"main.go/pkg/common/helper"
	"main.go/pkg/common/response"
	"main.go/pkg/domain"
	"main.go/pkg/repository/controllers"
	interfaces "main.go/pkg/repository/interface"
)

type adminDatabase struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) interfaces.AdminRespository {
	return &adminDatabase{DB}
}

//-------------------------- Login --------------------------//

func (c *adminDatabase) AdminLogin(email string) (domain.Admins, error) {
	var adminData domain.Admins
	err := c.DB.Raw("SELECT * FROM admins WHERE email=?", email).Scan(&adminData).Error
	return adminData, err
}

//-------------------------- Show-Single-User --------------------------//

func (c *adminDatabase) ShowUser(userID int) (response.UserDetails, error) {
	var userData response.UserDetails
	qury := `SELECT users.id,
			 users.name,
			 users.email, 
			 users.mobile, 
			 users.is_blocked, 
			 block_infos.blocked_at,
			 block_infos.reason_for_blocking 
			 FROM users as users 
			 FULL OUTER JOIN user_block_infos as block_infos ON users.id = block_infos.users_id
			 WHERE users.id = $1;`

	err := c.DB.Raw(qury, userID).Scan(&userData).Error
	if err != nil {
		return response.UserDetails{}, err
	}
	if userData.Email == "" {
		return response.UserDetails{}, fmt.Errorf("no such user")
	}
	return userData, nil
}

//-------------------------- Show-All-Users --------------------------//

func (c *adminDatabase) ShowAllUser() ([]response.UserDetails, error) {
	var userDatas []response.UserDetails

	getUsers := `SELECT users.id,
				users.name,
				users.email, 
				users.mobile, 
				users.is_blocked,
				block_infos.blocked_at,
				block_infos.reason_for_blocking 
				FROM users as users 
				FULL OUTER JOIN user_block_infos as block_infos ON users.id = block_infos.users_id;`

	err := c.DB.Raw(getUsers).Scan(&userDatas).Error
	return userDatas, err
}

//-------------------------- Dashboard --------------------------//

func (c *adminDatabase) GetDashBoard(reports helper.ReportParams) (response.DashBoard, error) {
	tx := c.DB.Begin()
	var dashBoard response.DashBoard
	getDasheBoard := `SELECT SUM(oi.quantity*oi.price)as Total_Revenue,
			SUM (oi.quantity)as Total_Products_Selled,
			COUNT(DISTINCT o.id)as Total_Orders FROM orders o
			JOIN order_items oi on o.id=oi.orders_id`

	getTotalUsers := `SELECT COUNT(id)AS TotalUsers FROM users`
	
	if reports.Status!=0{
		getDasheBoard=fmt.Sprintf("%s WHERE o.order_id=%d",getDasheBoard,reports.Status)
	} else{
		getDasheBoard = fmt.Sprintf("%s WHERE o.order_status_id is not null", getDasheBoard)

	}
	if reports.Day != 0 && reports.Month != 0 && reports.Year != 0 {
		date := controllers.GetDate(reports.Year, reports.Month, reports.Week, reports.Day)
		getDasheBoard = fmt.Sprintf("%s AND o.order_date::date='%s'", getDasheBoard, date)
		getTotalUsers = fmt.Sprintf("%s WHERE created_at::date='%s'", getTotalUsers, date)
	} else if reports.Week != 0 && reports.Month != 0 && reports.Year != 0 {
		date := controllers.GetDate(reports.Year, reports.Month, reports.Week, reports.Day)
		getDasheBoard = fmt.Sprintf("%s AND o.order_date BETWEEN %s", getDasheBoard, date)
		getTotalUsers = fmt.Sprintf("%s WHERE created_at BETWEEN %s", getTotalUsers, date)
	} else if reports.Month != 0 && reports.Year != 0 {
		getDasheBoard = fmt.Sprintf("%s AND EXTRACT(YEAR FROM order_date) = %d AND EXTRACT(MONTH FROM order_date) = %d", getDasheBoard, reports.Year, reports.Month)
		getTotalUsers = fmt.Sprintf("%s WHERE EXTRACT(YEAR FROM created_at) = %d AND EXTRACT(MONTH FROM created_at) = %d", getTotalUsers, reports.Year, reports.Month)
	} else if reports.StartDate != "" && reports.EndDate != "" {
		getDasheBoard = fmt.Sprintf("%s AND o.order_date BETWEEN '%s 00:00:00'::timestamp AND '%s 23:59:59'::timestamp", getDasheBoard, reports.StartDate, reports.EndDate)
		getTotalUsers = fmt.Sprintf("%s WHERE created_at BETWEEN '%s 00:00:00'::timestamp AND '%s 23:59:59'::timestamp", getTotalUsers, reports.StartDate, reports.EndDate)
	} else if reports.Year != 0 {
		getDasheBoard = fmt.Sprintf("%s AND EXTRACT ( YEAR FROM order_date) = %d", getDasheBoard, reports.Year)
		getTotalUsers = fmt.Sprintf("%s WHERE EXTRACT ( YEAR FROM created_at) = %d", getTotalUsers, reports.Year)
	}

	if err := tx.Raw(getDasheBoard).Scan(&dashBoard).Error; err != nil {
		tx.Rollback()
		return response.DashBoard{}, err
	}
	if err := tx.Raw(getTotalUsers).Scan(&dashBoard.TotalUsers).Error; err != nil {
		tx.Rollback()
		return response.DashBoard{}, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return response.DashBoard{}, err
	}
	return dashBoard, nil
}

func (c *adminDatabase) ViewSalesReport(reports helper.ReportParams) ([]response.SalesReport, error) {

	var sales []response.SalesReport
	getReports := `SELECT u.name,
		pt.type AS payment_type,
		o.order_date,
		o.order_total 
		FROM orders o JOIN users u ON u.id=o.user_id 
		JOIN payment_types pt ON o.payment_type_id= pt.id`

	if reports.Status!=0{
		getReports=fmt.Sprintf("%s WHERE o.order_status_id=%d", getReports,reports.Status)
	}else{
		getReports=fmt.Sprintf("%s WHERE o.order_status_id is not null",getReports)
	}
	if reports.Day != 0 && reports.Month != 0 && reports.Year != 0 {
		date := controllers.GetDate(reports.Year, reports.Month, reports.Week, reports.Day)
		getReports = fmt.Sprintf("%s AND o.order_date::date='%s'", getReports, date)
	} else if reports.Week != 0 && reports.Month != 0 && reports.Year != 0 {
		date := controllers.GetDate(reports.Year, reports.Month, reports.Week, reports.Day)
		getReports = fmt.Sprintf("%s AND o.order_date BETWEEN %s", getReports, date)
	} else if reports.Month != 0 && reports.Year != 0 {
		getReports = fmt.Sprintf("%s AND EXTRACT(YEAR FROM order_date) = %d AND EXTRACT(MONTH FROM order_date) = %d", getReports, reports.Year, reports.Month)
	} else if reports.StartDate != "" && reports.EndDate != "" {
		getReports = fmt.Sprintf("%s AND o.order_date BETWEEN '%s 00:00:00'::timestamp AND '%s 23:59:59'::timestamp", getReports, reports.StartDate, reports.EndDate)
	} else if reports.Year != 0 {
		getReports = fmt.Sprintf("%s AND EXTRACT ( YEAR FROM order_date) = %d", getReports, reports.Year)
	}
	err := c.DB.Raw(getReports).Scan(&sales).Error
	return sales, err

}