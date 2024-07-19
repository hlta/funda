package model

import (
	"errors"

	"gorm.io/gorm"
)

// Organization represents an organization entity.
type Organization struct {
	gorm.Model
	Name              string `gorm:"uniqueIndex"`
	Industry          *string
	GSTRegistered     *bool
	OwnerID           uint
	UserOrganizations []UserOrganization
}

// OrganizationRepository defines methods to interact with the Organization storage.
type OrganizationRepository interface {
	CreateWithTx(tx *gorm.DB, org *Organization) error
	RetrieveByID(id uint) (*Organization, error)
	Update(org *Organization) error
	Delete(id uint) error
}

// Predefined errors to handle specific scenarios
var ErrOrgExists = errors.New("organization already exists")
