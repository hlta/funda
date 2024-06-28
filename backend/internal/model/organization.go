package model

import (
	"gorm.io/gorm"
)

// Organization represents an organization entity.
type Organization struct {
	gorm.Model
	Name              string
	Industry          *string
	GSTRegistered     *bool
	OwnerID           uint
	UserOrganizations []UserOrganization
}

// OrganizationRepository defines methods to interact with the Organization storage.
type OrganizationRepository interface {
	Create(org *Organization) error
	RetrieveByID(id uint) (*Organization, error)
	Update(org *Organization) error
	Delete(id uint) error
}
