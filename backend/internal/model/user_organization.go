package model

import (
	"gorm.io/gorm"
)

// UserOrganization represents the many-to-many relationship between users and organizations with roles.
type UserOrganization struct {
	gorm.Model
	UserID         uint         `json:"userId"`
	OrganizationID uint         `json:"organizationId"`
	RoleID         uint         `json:"roleId"`
	User           User         `json:"user" gorm:"foreignKey:UserID"`
	Organization   Organization `json:"organization" gorm:"foreignKey:OrganizationID"`
	Role           Role         `json:"role" gorm:"foreignKey:RoleID"`
}

// UserOrganizationRepository defines methods to interact with the UserOrganization storage.
type UserOrganizationRepository interface {
	AddUserToOrganization(userOrg *UserOrganization) error
	RemoveUserFromOrganization(userID uint, orgID uint) error
	GetUserOrganizations(userID uint) ([]UserOrganization, error)
	GetUserOrganization(userID uint, orgID uint) (*UserOrganization, error)
}
