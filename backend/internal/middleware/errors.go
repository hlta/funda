package middleware

import "funda/internal/model"

// ErrorResponse for structuring error output
type ErrorResponse struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Errors  []model.FieldError `json:"errors,omitempty"`
}
