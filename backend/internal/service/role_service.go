package service

import (
	"funda/internal/logger"
	"funda/internal/model"
)

// RoleService provides role management functionalities.
type RoleService struct {
	repo model.RoleRepository
	log  logger.Logger
}

// NewRoleService creates a new instance of RoleService.
func NewRoleService(repo model.RoleRepository, log logger.Logger) *RoleService {
	return &RoleService{
		repo: repo,
		log:  log,
	}
}

// CreateRole handles the creation of a new role with permissions.
func (s *RoleService) CreateRole(role *model.Role) error {
	if err := s.repo.Create(role); err != nil {
		s.log.WithField("action", "creating role").Error(err.Error())
		return err
	}
	s.log.WithField("action", "role created").Info("Role successfully created")
	return nil
}

// GetRoleByID retrieves a role by its ID including its permissions.
func (s *RoleService) GetRoleByID(id uint) (*model.Role, error) {
	role, err := s.repo.RetrieveByID(id)
	if err != nil {
		s.log.WithField("action", "retrieving role by ID").Error(err.Error())
		return nil, err
	}
	s.log.WithField("action", "role retrieved").Info("Role successfully retrieved")
	return role, nil
}

// GetRoleByName retrieves a role by its name.
func (s *RoleService) GetRoleByName(name string) (*model.Role, error) {
	role, err := s.repo.RetrieveByName(name)
	if err != nil {
		s.log.WithField("action", "retrieving role by name").Error(err.Error())
		return nil, err
	}
	s.log.WithField("action", "role retrieved").Info("Role successfully retrieved")
	return role, nil
}

// UpdateRole handles updates to an existing role including its permissions.
func (s *RoleService) UpdateRole(role *model.Role) error {
	if err := s.repo.Update(role); err != nil {
		s.log.WithField("action", "updating role").Error(err.Error())
		return err
	}
	s.log.WithField("action", "role updated").Info("Role successfully updated")
	return nil
}

// DeleteRole removes a role by its ID.
func (s *RoleService) DeleteRole(id uint) error {
	if err := s.repo.Delete(id); err != nil {
		s.log.WithField("action", "deleting role").Error(err.Error())
		return err
	}
	s.log.WithField("action", "role deleted").Info("Role successfully deleted")
	return nil
}
