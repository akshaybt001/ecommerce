package response

type Category struct {
	Id       int
	Category string
}

type Brands struct {
	Id           int `json:",omitempty"`
	Name         string
	Description  string
	CategoryId   int
	CategoryName string
}

type Model struct {
	Id           uint
	ModelName    string
	Brand        string
	Description  string
	CategoryName string
	Sku          string
	QtyInStock   int
	Color        string
	Ram          int
	Battery      int
	ScreenSize   float64
	Storage      int
	Camera       int
	Price        int
	Image        []string
}
