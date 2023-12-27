package request

type LoginRequest struct {
	Login      string `json:"login" binding:"required"`
	Password   string `json:"password" binding:"required"`
	IsRemember bool   `json:"isRemember" binding:"required"`
}
