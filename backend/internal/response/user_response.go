package response

type OrganizationResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type UserResponse struct {
	FirstName    string               `json:"firstName"`
	LastName     string               `json:"lastName"`
	Email        string               `json:"email"`
	Token        string               `json:"token,omitempty"`
	Organization OrganizationResponse `json:"organization"`
	Roles        []string             `json:"roles"`
	Permissions  []string             `json:"permissions"`
}
