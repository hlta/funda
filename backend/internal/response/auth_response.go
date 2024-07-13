package response

// SwitchOrganizationResponse represents the response structure for switching organizations.
type AuthResponse struct {
	User        UserResponse `json:"user"`
	Token       string       `json:"token"`
	Roles       []string     `json:"roles"`
	Permissions [][]string   `json:"permissions"`
}
