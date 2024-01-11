package request

// Запрос на получения параметров
type GetParameterValuesRequest struct {
	Path string `json:"path,omitempty"` // Путь для которого запрашиваются значения параметров
}
