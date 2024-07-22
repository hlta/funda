package mapper

import (
	"funda/internal/model"
	"funda/internal/response"
)

// ToUserResponse maps a model.User to a response.UserResponse.
func ToUserResponse(user model.User) response.UserResponse {
	return response.UserResponse{
		ID:                    user.ID,
		FirstName:             user.FirstName,
		LastName:              user.LastName,
		DefaultOrganizationID: user.DefaultOrganizationID,
		Email:                 user.Email,
	}
}
