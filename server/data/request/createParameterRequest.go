package request

// Запрос на создание параметра
type CreateParameterRequest struct {
	Name          string `json:"name" binding:"required"`                // Имя параметра
	ParameterType string `json:"parameterType" binding:"required"`       // Тип параметра
	IsDefault     bool   `json:"isDefault,omitempty" binding:"required"` // Параметр является стандартным и доступен для всех шаблонов
}
