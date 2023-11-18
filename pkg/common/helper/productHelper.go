package helper

type Category struct {
	Name string `json:"name" validate:"required"`
}

type Brands struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	CategoryId  string `json:"categoryid" validate:"required"`
}

type Model struct {
	Brand_id    uint    `json:"brandid"`
	Model_name  string  `json:"modelname"`
	Sku         string  `json:"sku"`
	Qty         int     `json:"quantity"`
	Color       string  `json:"colour"`
	Ram         int     `json:"ram"`
	Battery     int     `json:"battery"`
	Screen_size float64 `json:"screensize"`
	Storage     int     `json:"storage"`
	Camera      int     `json:"camera"`
	Price       int     `json:"price"`
	Image       string  `json:"image"`
}

type QueryParams struct {
	Page     int    `json:"page"`
	Limit    int    `json:"limit"`
	Query    string `json:"query"`   //search key word
	Filter   string `json:"filter"`  //to specify the column name
	SortBy   string `json:"sort_by"` //to specify column to set the sorting
	SortDesc bool   `json:"sort_desc"`
}