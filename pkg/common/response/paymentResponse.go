package response

type Wallet struct {
	Amount int `gorm:"default:0;check:Amount>=0" sql:"CHECK(Amount >= 0)"`
}
