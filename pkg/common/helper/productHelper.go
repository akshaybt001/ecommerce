package helper

type Category struct {
	Name string `json:"name" validate:"required"`
}
