package response

type GenericResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
