package model

import (
	"gorm.io/gorm"
)

// UserOrganization represents the many-to-many relationship between users and organizations with roles.
type UserOrganization struct {
	gorm.Model
	UserID         uint
	OrganizationID uint
	RoleID         uint
	User           User
	Organization   Organization
	Role           Role
}

// UserOrganizationRepository defines methods to interact with the UserOrganization storage.
type UserOrganizationRepository interface {
	AddUserToOrganization(userOrg *UserOrganization) error
	RemoveUserFromOrganization(userID uint, orgID uint) error
	GetUserOrganizations(userID uint) ([]UserOrganization, error)
}
