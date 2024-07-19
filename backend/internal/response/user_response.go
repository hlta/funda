package response

// UserResponse represents the response structure for a user.
type UserResponse struct {
	ID          uint   `json:"id,omitempty"`
	FirstName   string `json:"firstName,omitempty"`
	LastName    string `json:"lastName,omitempty"`
	Email       string `json:"email,omitempty"`
	SelectedOrg uint   `json:"selectedOrg,omitempty"`
}
