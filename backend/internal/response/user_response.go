package response

import "encoding/json"

type OrganizationResponse struct {
	ID   uint   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type UserResponse struct {
	FirstName    string               `json:"firstName"`
	LastName     string               `json:"lastName"`
	Email        string               `json:"email"`
	Token        string               `json:"token,omitempty"`
	Organization OrganizationResponse `json:"organization,omitempty"`
	Roles        []string             `json:"roles"`
	Permissions  []string             `json:"permissions"`
	SelectedOrg  uint                 `json:"selectedOrg"`
}

func (o OrganizationResponse) MarshalJSON() ([]byte, error) {
	type Alias OrganizationResponse
	aux := &struct {
		*Alias
		ID *uint `json:"id,omitempty"`
	}{
		Alias: (*Alias)(&o),
		ID:    nil,
	}

	if o.ID != 0 {
		aux.ID = &o.ID
	}

	return json.Marshal(aux)
}
