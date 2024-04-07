package store

import (
	"funda/internal/model"

	"gorm.io/gorm"
)

// GormUserRepository is the GORM implementation of the UserRepository interface.
type GormUserRepository struct {
	DB *gorm.DB
}

// NewGormUserRepository creates a new user repository with a given GORM DB connection.
func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{DB: db}
}

// Create adds a new user to the database.
func (r *GormUserRepository) Create(user *model.User) error {
	result := r.DB.Create(user)
	return result.Error
}

// RetrieveByID finds a user by their ID.
func (r *GormUserRepository) RetrieveByID(id uint) (*model.User, error) {
	user := &model.User{}
	result := r.DB.First(user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

// RetrieveByEmail finds a user by their email.
func (r *GormUserRepository) RetrieveByEmail(email string) (*model.User, error) {
	user := &model.User{}
	result := r.DB.Where("email = ?", email).First(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

// Update modifies an existing user in the database.
func (r *GormUserRepository) Update(user *model.User) error {
	result := r.DB.Save(user)
	return result.Error
}

// Delete removes a user from the database.
func (r *GormUserRepository) Delete(id uint) error {
	result := r.DB.Delete(&model.User{}, id)
	return result.Error
}

// List retrieves users from the database with pagination.
func (r *GormUserRepository) List(offset, limit int) ([]*model.User, error) {
	var users []*model.User
	result := r.DB.Offset(offset).Limit(limit).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
