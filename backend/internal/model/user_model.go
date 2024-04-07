package model

// User represents the user for the system.
type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"uniqueIndex"`
	// Add additional fields as required
	// Password    string `json:"-"` // Use json:"-" to not expose sensitive information
	// CreatedAt   time.Time `json:"created_at"`
	// UpdatedAt   time.Time `json:"updated_at"`
	// ...
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
