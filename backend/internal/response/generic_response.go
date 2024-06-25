package response

// GenericResponse represents a generic API response.
type GenericResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
