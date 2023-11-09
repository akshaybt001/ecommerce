package helper

type UserReq struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Mobile   string `json:"mobile" binding:"required"`
	Password string `json:"password"`
	OTP      string `json:"otp"`
}
type LoginReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password"`
}
