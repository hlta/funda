package response

// OrganizationResponse represents the response structure for an organization.
type OrganizationResponse struct {
	ID            uint   `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	Industry      string `json:"industry,omitempty"`
	OwnerID       uint   `json:"owner_id,omitempty"`
	GSTRegistered bool   `json:"gst_registered,omitempty"`
}
