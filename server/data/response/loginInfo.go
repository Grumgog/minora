package response

type LoginInfo struct {
	Login string `json:"username" validate:"required"`
	Token string `json:"token" validate:"required"`
}
