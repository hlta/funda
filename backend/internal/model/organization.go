package model

import (
	"gorm.io/gorm"
)

// Organization represents an organization entity.
type Organization struct {
	gorm.Model
	Name              string
	Industry          *string // Optional field
	GSTRegistered     *bool   // Optional field
	OwnerID           uint
	UserOrganizations []UserOrganization
}
