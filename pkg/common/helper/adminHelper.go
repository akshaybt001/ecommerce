package helper

type BlockData struct {
	UserId uint   ` json:"userid" validate:"required"`
	Reason string ` json:"reason" validate:"required"`
}