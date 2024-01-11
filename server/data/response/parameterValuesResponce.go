package response

type ParameterValuesResponce struct {
	Parameters []ParameterValue `json:"parameters"`
}

type ParameterValue struct {
	Path  string `json:"path"`
	Name  string `json:"name"`
	Value string `json:"value"`
}
