package model

import (
	"errors"

	"gorm.io/gorm"
)

// User represents the user entity.
type User struct {
	gorm.Model
	FirstName             string
	LastName              string `json:"LastName,omitempty"`
	Email                 string `gorm:"uniqueIndex"`
	Token                 string `json:"Token,omitempty" gorm:"-"`
	Password              string
	UserOrganizations     []UserOrganization
	DefaultOrganizationID uint
	DefaultOrganization   Organization `gorm:"foreignKey:DefaultOrganizationID"`
}

// UserRepository is the interface that defines methods to interact with the User storage.
type UserRepository interface {
	Create(user *User) error                     // Create a new user
	RetrieveByID(id uint) (*User, error)         // Retrieve a user by ID
	RetrieveByEmail(email string) (*User, error) // Retrieve a user by email
	Update(user *User) error                     // Update a user
	Delete(id uint) error                        // Delete a user by ID
}

// Predefined errors to handle specific scenarios
var ErrEmailExists = errors.New("email already exists")
