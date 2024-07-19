package service

import (
	"funda/internal/logger"
	"funda/internal/model"

	"gorm.io/gorm"
)

// UserService provides user management functionalities.
type UserService struct {
	repo model.UserRepository
	log  logger.Logger
}

// NewUserService creates a new instance of UserService.
func NewUserService(repo model.UserRepository, log logger.Logger) *UserService {
	return &UserService{
		repo: repo,
		log:  log,
	}
}

// CreateUser handles the creation of a new user.
func (s *UserService) CreateUser(user *model.User) error {
	s.log.WithField("action", "creating user").Info("Attempting to create user")
	if err := s.repo.Create(user); err != nil {
		s.log.WithField("action", "creating user").Error(err.Error())
		return err
	}
	s.log.WithField("action", "user created").Info("User successfully created")
	return nil
}

// CreateUserWithTx handles the creation of a new user within a transaction.
func (s *UserService) CreateUserWithTx(tx *gorm.DB, user *model.User) error {
	s.log.WithField("action", "creating user").Info("Attempting to create user")
	if err := s.repo.CreateWithTx(tx, user); err != nil {
		s.log.WithField("action", "creating user").Error(err.Error())
		return err
	}
	s.log.WithField("action", "user created").Info("User successfully created")
	return nil
}

// GetUserByID retrieves a user by their ID.
func (s *UserService) GetUserByID(id uint) (*model.User, error) {
	s.log.WithField("action", "retrieving user by ID").Info("Attempting to retrieve user")
	user, err := s.repo.RetrieveByID(id)
	if err != nil {
		s.log.WithField("action", "retrieving user by ID").Error(err.Error())
		return nil, err
	}
	s.log.WithField("action", "user retrieved").Info("User successfully retrieved")
	return user, nil
}

// GetUserByEmail retrieves a user by their email.
func (s *UserService) GetUserByEmail(email string) (*model.User, error) {
	s.log.WithField("action", "retrieving user by email").Info("Attempting to retrieve user")
	user, err := s.repo.RetrieveByEmail(email)
	if err != nil {
		s.log.WithField("action", "retrieving user by email").Error(err.Error())
		return nil, err
	}
	s.log.WithField("action", "user retrieved").Info("User successfully retrieved")
	return user, nil
}

// UpdateUser handles updates to an existing user.
func (s *UserService) UpdateUser(user *model.User) error {
	s.log.WithField("action", "updating user").Info("Attempting to update user")
	if err := s.repo.Update(user); err != nil {
		s.log.WithField("action", "updating user").Error(err.Error())
		return err
	}
	s.log.WithField("action", "user updated").Info("User successfully updated")
	return nil
}

// UpdateUserWithTx handles updates to an existing user within a transaction.
func (s *UserService) UpdateUserWithTx(tx *gorm.DB, user *model.User) error {
	s.log.WithField("action", "updating user").Info("Attempting to update user")
	if err := s.repo.UpdateWithTx(tx, user); err != nil {
		s.log.WithField("action", "updating user").Error(err.Error())
		return err
	}
	s.log.WithField("action", "user updated").Info("User successfully updated")
	return nil
}

// DeleteUser removes a user by their ID.
func (s *UserService) DeleteUser(id uint) error {
	s.log.WithField("action", "deleting user").Info("Attempting to delete user")
	if err := s.repo.Delete(id); err != nil {
		s.log.WithField("action", "deleting user").Error(err.Error())
		return err
	}
	s.log.WithField("action", "user deleted").Info("User successfully deleted")
	return nil
}
