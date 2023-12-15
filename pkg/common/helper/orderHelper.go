package helper

type Cart struct {
	Id    int
	Total int
}

type CartItems struct {
	ModelId    int
	Quantity   int
	Price      int
	QtyInStock int
}

type UpdateOrder struct {
	OrderId       uint `json:"orderid"`
	OrderStatusID uint `json:"orderstatusid"`
}
type OrderStatus struct {
	Id     uint   `json:"id,omitempty"`
	Status string `json:"status"`
}