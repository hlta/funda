package mapper

import "funda/internal/response"

func ToAuthResponse(user response.UserResponse, token string, roles []string, permissions [][]string) response.AuthResponse {
	return response.AuthResponse{
		User:        user,
		Token:       token,
		Roles:       roles,
		Permissions: permissions,
	}
}
