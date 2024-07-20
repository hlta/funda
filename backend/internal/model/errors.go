package model

var (
	ErrEmailExists = FieldError{Field: "email", Message: "Email already exists"}
	ErrOrgExists   = FieldError{Field: "organizationName", Message: "Organization name already exists"}
)

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e FieldError) Error() string {
	return e.Message
}
