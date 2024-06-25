package model

import (
	"gorm.io/gorm"
)

// Role represents a role entity with specific permissions.
type Role struct {
	gorm.Model
	Name              string
	OrganizationID    uint
	Permissions       []Permission `gorm:"many2many:role_permissions;"`
	UserOrganizations []UserOrganization
}

// Permission represents specific access rights within the system.
type Permission struct {
	gorm.Model
	Name string
}

// RolePermission is the join table for many-to-many relationship between Role and Permission.
type RolePermission struct {
	RoleID       uint
	PermissionID uint
}

// RoleRepository defines methods to interact with the Role storage.
type RoleRepository interface {
	Create(role *Role) error
	RetrieveByID(id uint) (*Role, error)
	Update(role *Role) error
	Delete(id uint) error
	RetrieveByName(name string) (*Role, error)
}
