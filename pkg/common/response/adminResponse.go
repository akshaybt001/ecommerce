package response

import "time"

type DashBoard struct {
	TotalRevenue        int
	TotalOrders         int
	TotalProductsSelled int
	TotalUsers          int
}
type SalesReport struct {
	Name        string
	PaymentType string
	OrderDate   time.Time
	OrderTotal  int
}
