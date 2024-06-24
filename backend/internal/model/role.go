package model

import (
	"gorm.io/gorm"
)

// Role represents a role entity with specific permissions.
type Role struct {
	gorm.Model
	Name              string             `json:"name"`
	OrganizationID    uint               `json:"organizationId"`
	Permissions       []Permission       `json:"permissions" gorm:"many2many:role_permissions;"`
	UserOrganizations []UserOrganization `json:"userOrganizations"`
}

// Permission represents specific access rights within the system.
type Permission struct {
	gorm.Model
	Name string `json:"name"`
}

// RolePermission is the join table for many-to-many relationship between Role and Permission.
type RolePermission struct {
	RoleID       uint `json:"roleId"`
	PermissionID uint `json:"permissionId"`
}

// RoleRepository defines methods to interact with the Role storage.
type RoleRepository interface {
	Create(role *Role) error
	RetrieveByID(id uint) (*Role, error)
	Update(role *Role) error
	Delete(id uint) error
	RetrieveByName(name string) (*Role, error)
}
