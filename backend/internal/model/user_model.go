package model

import (
	"errors"

	"gorm.io/gorm"
)

// User represents the user entity.
type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string `gorm:"uniqueIndex"` // Email must be unique
	Password  string // This should be a hashed password
}

// UserRepository is the interface that defines methods to interact with the User storage.
type UserRepository interface {
	Create(user *User) error                     // Create a new user
	RetrieveByID(id uint) (*User, error)         // Retrieve a user by ID
	RetrieveByEmail(email string) (*User, error) // Retrieve a user by email
	Update(user *User) error                     // Update a user
	Delete(id uint) error                        // Delete a user by ID
	// List can be added if listing users is required
	// List(offset, limit int) ([]*User, error)
}

// Predefined errors to handle specific scenarios
var ErrEmailExists = errors.New("email already exists")
