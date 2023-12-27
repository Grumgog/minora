package response

type LoginInfo struct {
	Id       uint   `json:"id" validate:"required"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Token    string `json:"token" validate:"required"`
}
