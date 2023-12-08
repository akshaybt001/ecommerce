package helper

type BlockData struct {
	UserId uint   ` json:"userid" validate:"required"`
	Reason string ` json:"reason" validate:"required"`
}
type ReportParams struct {
	Status    int   `json:"status"`
	Year      int    `json:"year"`
	Month     int    `json:"month"`
	Week      int    `json:"week"`
	Day       int    `json:"day"`
	StartDate string `json:"startdate"`
	EndDate   string `json:"enddate"`
}
type CreateAdmin struct {
	Name     string ` json:"name" validate:"required"`
	Email    string ` json:"email" validate:"required" binding:"email"`
	Password string ` json:"password" validate:"required"`
}