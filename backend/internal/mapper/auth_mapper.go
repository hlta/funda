package mapper

import "funda/internal/response"

// ToSwitchOrganizationResponse maps the result of switching organizations to a response.SwitchOrganizationResponse.
func ToSwitchOrganizationResponse(token string, roles []string, permissions []string) response.SwitchOrganizationResponse {
	return response.SwitchOrganizationResponse{
		Token:       token,
		Roles:       roles,
		Permissions: permissions,
	}
}
