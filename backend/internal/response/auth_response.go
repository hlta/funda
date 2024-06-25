package response

// SwitchOrganizationResponse represents the response structure for switching organizations.
type SwitchOrganizationResponse struct {
	Token       string   `json:"token"`
	Roles       []string `json:"roles"`
	Permissions []string `json:"permissions"`
}
