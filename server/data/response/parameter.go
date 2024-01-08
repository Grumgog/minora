package response

type Parameter struct {
	ID            uint   `json:"id" validate:"required"`
	Name          string `json:"name" validate:"required"`
	ParameterType string `json:"parameterType" validate:"required"`
	IsDefault     bool   `json:"isDefault" validate:"required"`
	DefaultValue  string `json:"defaultValue" validate:"required"`
	IsSystem      bool   `json:"isSystem" validate:"required"`
}
