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
type AdminData struct {
	Id    int
	Name  string
	Email string
}
type AdminDetails struct {
	Id                uint
	Name              string
	Email             string
	IsBlocked         bool
	BlockedAt         string `json:",omitempty"`
	BlockedBy         uint   `json:",omitempty"`
	ReasonForBlocking string `json:",omitempty"`
}
