package service

import (
	"errors"
	"funda/internal/model"
)

// UserService provides methods to deal with users in the system.
type UserService struct {
	repo model.UserRepository
}

// NewUserService creates a new instance of UserService.
func NewUserService(repo model.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// CreateUser creates a new user.
func (s *UserService) CreateUser(user *model.User) error {
	// Here you can add business logic before creating the user.
	// For example, you might want to hash the password, validate the input, etc.

	return s.repo.Create(user)
}

// GetUserByID retrieves a user by their ID.
func (s *UserService) GetUserByID(id uint) (*model.User, error) {
	return s.repo.RetrieveByID(id)
}

// GetUserByEmail retrieves a user by their email.
func (s *UserService) GetUserByEmail(email string) (*model.User, error) {
	return s.repo.RetrieveByEmail(email)
}

// UpdateUser updates an existing user.
func (s *UserService) UpdateUser(user *model.User) error {
	// Perform business logic checks before updating the user.
	// For instance, check if the email is already taken, if certain fields should not be updated, etc.

	return s.repo.Update(user)
}

// DeleteUser deletes a user by their ID.
func (s *UserService) DeleteUser(id uint) error {
	// You can include business logic here, like cleaning up related resources.

	return s.repo.Delete(id)
}

// ListUsers retrieves a list of users with pagination.
func (s *UserService) ListUsers(offset, limit int) ([]*model.User, error) {
	// You can include additional business logic here if needed, like filtering based on criteria.

	// This is a placeholder for a method that would be implemented in your UserRepository.
	// return s.repo.List(offset, limit)
	return nil, errors.New("not implemented")
}

// Additional methods for user service can be added as needed.
