package response

// UserResponse represents the response structure for a user.
type UserResponse struct {
	ID           uint                 `json:"id,omitempty"`
	FirstName    string               `json:"firstName,omitempty"`
	LastName     string               `json:"lastName,omitempty"`
	Email        string               `json:"email,omitempty"`
	Token        string               `json:"token,omitempty"`
	Organization OrganizationResponse `json:"organization,omitempty"`
	SelectedOrg  uint                 `json:"selectedOrg,omitempty"`
	Roles        []string             `json:"roles,omitempty"`
	Permissions  []string             `json:"permissions,omitempty"`
}
