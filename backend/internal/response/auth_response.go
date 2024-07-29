package response

// SwitchOrganizationResponse represents the response structure for switching organizations.
type AuthResponse struct {
	User        UserResponse `json:"user"`
	SelectedOrg uint         `json:"selectedOrg"`
	Token       string       `json:"token"`
	Roles       []string     `json:"roles"`
	Permissions [][]string   `json:"permissions"`
}
