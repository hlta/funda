package service

import (
	"funda/internal/logger"
	"funda/internal/model"
)

// UserOrganizationService provides user-organization relationship management functionalities.
type UserOrganizationService struct {
	repo model.UserOrganizationRepository
	log  logger.Logger
}

// NewUserOrganizationService creates a new instance of UserOrganizationService.
func NewUserOrganizationService(repo model.UserOrganizationRepository, log logger.Logger) *UserOrganizationService {
	return &UserOrganizationService{
		repo: repo,
		log:  log,
	}
}

// AddUserToOrganization handles adding a user to an organization with a specific role.
func (s *UserOrganizationService) AddUserToOrganization(userOrg *model.UserOrganization) error {
	if err := s.repo.AddUserToOrganization(userOrg); err != nil {
		s.log.WithField("action", "adding user to organization").Error(err.Error())
		return err
	}
	s.log.WithField("action", "user added to organization").Info("User successfully added to organization")
	return nil
}

// RemoveUserFromOrganization handles removing a user from an organization.
func (s *UserOrganizationService) RemoveUserFromOrganization(userID uint, orgID uint) error {
	if err := s.repo.RemoveUserFromOrganization(userID, orgID); err != nil {
		s.log.WithField("action", "removing user from organization").Error(err.Error())
		return err
	}
	s.log.WithField("action", "user removed from organization").Info("User successfully removed from organization")
	return nil
}

// GetUserOrganizations retrieves the organizations a user belongs to.
func (s *UserOrganizationService) GetUserOrganizations(userID uint) ([]model.UserOrganization, error) {
	userOrgs, err := s.repo.GetUserOrganizations(userID)
	if err != nil {
		s.log.WithField("action", "retrieving user organizations").Error(err.Error())
		return nil, err
	}
	s.log.WithField("action", "user organizations retrieved").Info("User organizations successfully retrieved")
	return userOrgs, nil
}
