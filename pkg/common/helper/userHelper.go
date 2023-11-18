package helper

type UserReq struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Mobile   string `json:"mobile" binding:"required"`
	Password string `json:"password"`
	OTP      string `json:"otp"`
}
type ForgotPassword struct{
	Email       string `json:"email" binding:"required"`
	NewPassword string `json:"newpassword" binding:"required"`
	OTP         string `json:"otp"`
}

type LoginReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password"`
}

type UpdatePassword struct {
	OldPassword string `json:"oldpassword" `
	NewPasswoed string `json:"newpassword" `
}

type Userprofile struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Mobile  string `json:"mobile"`
	Address `gorm:"embedded" json:"address"`
}
type Address struct {
	House_number string `json:"house_number" `
	Street       string `json:"street" `
	City         string `json:"city" `
	District     string `json:"district" `
	Landmark     string `json:"landmark" `
	Pincode      int    `json:"pincode" `
	IsDefault    bool   `json:"isdefault" `
}
