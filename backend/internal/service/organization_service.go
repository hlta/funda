package service

import (
	"errors"
	"funda/internal/logger"
	"funda/internal/mapper"
	"funda/internal/model"
	"funda/internal/response"
)

// OrganizationService provides organization management functionalities.
type OrganizationService struct {
	repo model.OrganizationRepository
	log  logger.Logger
}

// NewOrganizationService creates a new instance of OrganizationService.
func NewOrganizationService(repo model.OrganizationRepository, log logger.Logger) *OrganizationService {
	return &OrganizationService{
		repo: repo,
		log:  log,
	}
}

// CreateOrganization handles the creation of a new organization.
func (s *OrganizationService) CreateOrganization(org *model.Organization) error {
	if org.Name == "" {
		return errors.New("organization name is required")
	}
	if err := s.repo.Create(org); err != nil {
		s.log.WithField("action", "creating organization").Error(err.Error())
		return err
	}
	s.log.WithField("action", "organization created").Info("Organization successfully created")
	return nil
}

// GetOrganizationByID retrieves an organization by its ID.
func (s *OrganizationService) GetOrganizationByID(id uint) (*response.OrganizationResponse, error) {
	org, err := s.repo.RetrieveByID(id)
	if err != nil {
		s.log.WithField("action", "retrieving organization by ID").Error(err.Error())
		return nil, err
	}
	s.log.WithField("action", "organization retrieved").Info("Organization successfully retrieved")
	orgRespResp := mapper.ToOrganizationResponse(*org)
	return &orgRespResp, nil
}

// UpdateOrganization handles updates to an existing organization.
func (s *OrganizationService) UpdateOrganization(org *model.Organization) error {
	if org.Name == "" {
		return errors.New("organization name is required")
	}
	if err := s.repo.Update(org); err != nil {
		s.log.WithField("action", "updating organization").Error(err.Error())
		return err
	}
	s.log.WithField("action", "organization updated").Info("Organization successfully updated")
	return nil
}

// DeleteOrganization removes an organization by its ID.
func (s *OrganizationService) DeleteOrganization(id uint) error {
	if err := s.repo.Delete(id); err != nil {
		s.log.WithField("action", "deleting organization").Error(err.Error())
		return err
	}
	s.log.WithField("action", "organization deleted").Info("Organization successfully deleted")
	return nil
}
