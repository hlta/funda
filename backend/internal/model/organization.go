package model

import (
	"gorm.io/gorm"
)

// Organization represents an organization entity.
type Organization struct {
	gorm.Model
	Name              string             `json:"name"`
	OwnerID           uint               `json:"ownerId"` // Reference to the User who created the organization
	UserOrganizations []UserOrganization `json:"userOrganizations"`
}

// OrganizationRepository defines methods to interact with the Organization storage.
type OrganizationRepository interface {
	Create(org *Organization) error
	RetrieveByID(id uint) (*Organization, error)
	Update(org *Organization) error
	Delete(id uint) error
}
