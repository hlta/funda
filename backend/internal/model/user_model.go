package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string `gorm:"uniqueIndex"`
	Password  string // This should be a hashed password
}

// UserRepository is the interface that defines methods to interact with the User storage.
type UserRepository interface {
	Create(user *User) error
	RetrieveByID(id uint) (*User, error)
	RetrieveByEmail(email string) (*User, error)
	Update(user *User) error
	Delete(id uint) error
	// Additional methods can be added as needed
	// List(offset, limit int) ([]*User, error)
	// ...
}
