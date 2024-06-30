package service

import (
	"funda/internal/logger"
	"funda/internal/mapper"
	"funda/internal/model"
	"funda/internal/response"
	"funda/internal/utils"
)

// OrganizationService provides organization management functionalities.
type OrganizationService struct {
	repo        model.OrganizationRepository
	roleRepo    model.RoleRepository
	userOrgRepo model.UserOrganizationRepository
	log         logger.Logger
}

// NewOrganizationService creates a new instance of OrganizationService.
func NewOrganizationService(repo model.OrganizationRepository, roleRepo model.RoleRepository, userOrgRepo model.UserOrganizationRepository, log logger.Logger) *OrganizationService {
	return &OrganizationService{
		repo:        repo,
		roleRepo:    roleRepo,
		userOrgRepo: userOrgRepo,
		log:         log,
	}
}

// CreateOrganization handles the creation of a new organization.
func (s *OrganizationService) CreateOrganization(org *model.Organization) error {
	if err := utils.ValidateOrganization(org); err != nil {
		return err
	}
	if err := s.repo.Create(org); err != nil {
		utils.LogError(s.log, "creating organization", err)
		return err
	}
	if err := s.assignAdminRole(org.OwnerID, org.ID); err != nil {
		return err
	}
	utils.LogSuccess(s.log, "organization created", "Organization successfully created", org.OwnerID)
	return nil
}

// GetOrganizationByID retrieves an organization by its ID.
func (s *OrganizationService) GetOrganizationByID(id uint) (*response.OrganizationResponse, error) {
	org, err := s.repo.RetrieveByID(id)
	if err != nil {
		utils.LogError(s.log, "retrieving organization by ID", err)
		return nil, err
	}
	utils.LogSuccess(s.log, "organization retrieved", "Organization successfully retrieved", org.OwnerID)
	orgResp := mapper.ToOrganizationResponse(*org)
	return &orgResp, nil
}

// UpdateOrganization handles updates to an existing organization.
func (s *OrganizationService) UpdateOrganization(org *model.Organization) error {
	if err := utils.ValidateOrganization(org); err != nil {
		return err
	}
	if err := s.repo.Update(org); err != nil {
		utils.LogError(s.log, "updating organization", err)
		return err
	}
	utils.LogSuccess(s.log, "organization updated", "Organization successfully updated", org.OwnerID)
	return nil
}

// DeleteOrganization removes an organization by its ID.
func (s *OrganizationService) DeleteOrganization(id uint) error {
	if err := s.repo.Delete(id); err != nil {
		utils.LogError(s.log, "deleting organization", err)
		return err
	}
	utils.LogSuccess(s.log, "organization deleted", "Organization successfully deleted", id)
	return nil
}

func (s *OrganizationService) GetUserOrganizations(userID uint) ([]response.OrganizationResponse, error) {

	userOrgs, err := s.userOrgRepo.GetUserOrganizations(userID)
	if err != nil {
		return nil, err
	}

	var orgs []response.OrganizationResponse
	for _, userOrg := range userOrgs {
		org, err := s.repo.RetrieveByID(userOrg.OrganizationID)
		if err != nil {
			return nil, err
		}
		orgs = append(orgs, mapper.ToOrganizationResponse(*org))
	}
	return orgs, nil
}

// Helper Methods

func (s *OrganizationService) assignAdminRole(userID, orgID uint) error {
	role, err := s.roleRepo.RetrieveByName("Admin")
	if err != nil {
		utils.LogError(s.log, "retrieving role", err)
		return err
	}
	userOrg := &model.UserOrganization{UserID: userID, OrganizationID: orgID, RoleID: role.ID}
	if err := s.userOrgRepo.AddUserToOrganization(userOrg); err != nil {
		utils.LogError(s.log, "assigning user to organization", err)
		return err
	}
	return nil
}
