package request

// Запрос на авторизацию
type LoginRequest struct {
	Login      string `json:"login" binding:"required"`      // Логин пользователя
	Password   string `json:"password" binding:"required"`   // Пароль пользователя
	IsRemember bool   `json:"isRemember" binding:"required"` // Необходимо ли запоминать пользователя
}
