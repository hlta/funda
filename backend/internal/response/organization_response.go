package response

import "funda/internal/utils"

// OrganizationResponse represents the response structure for an organization.
type OrganizationResponse struct {
	ID   uint   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// MarshalJSON is a custom marshaler for OrganizationResponse to omit the ID field if it is 0.
func (o OrganizationResponse) MarshalJSON() ([]byte, error) {
	return utils.MarshalUintOmitZero(o, "id", o.ID)
}
