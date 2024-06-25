package response

// OrganizationResponse represents the response structure for an organization.
type OrganizationResponse struct {
	ID   uint   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
