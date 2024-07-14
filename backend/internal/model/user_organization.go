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
	User           User         `gorm:"foreignKey:UserID"`
	Organization   Organization `gorm:"foreignKey:OrganizationID"`
}

// UserOrganizationRepository defines methods to interact with the UserOrganization storage.
type UserOrganizationRepository interface {
	AddUserToOrganizationWithTx(tx *gorm.DB, userOrg *UserOrganization) error
	RemoveUserFromOrganization(userID, orgID uint) error
	GetUserOrganizations(userID uint) ([]UserOrganization, error)
}
