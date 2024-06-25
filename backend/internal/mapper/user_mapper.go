package mapper

import (
	"funda/internal/model"
	"funda/internal/response"
)

// ToUserResponse maps a model.User to a response.UserResponse.
func ToUserResponse(user model.User, roles []string, permissions []string, token string) response.UserResponse {
	return response.UserResponse{
		ID:          user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		Token:       token,
		SelectedOrg: user.DefaultOrganizationID,
		Roles:       roles,
		Permissions: permissions,
		Organization: response.OrganizationResponse{
			ID:   user.DefaultOrganization.ID,
			Name: user.DefaultOrganization.Name,
		},
	}
}
