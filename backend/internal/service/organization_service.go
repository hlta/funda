package service

import (
	"funda/internal/logger"
	"funda/internal/mapper"
	"funda/internal/model"
	"funda/internal/response"
	"funda/internal/utils"

	"gorm.io/gorm"
)

// OrganizationService provides organization management functionalities.
type OrganizationService struct {
	repo        model.OrganizationRepository
	userOrgRepo model.UserOrganizationRepository
	log         logger.Logger
	db          *gorm.DB
}

// NewOrganizationService creates a new instance of OrganizationService.
func NewOrganizationService(repo model.OrganizationRepository, userOrgRepo model.UserOrganizationRepository, log logger.Logger, db *gorm.DB) *OrganizationService {
	return &OrganizationService{
		repo:        repo,
		userOrgRepo: userOrgRepo,
		log:         log,
		db:          db,
	}
}

func (s *OrganizationService) CreateOrganizationWithTx(tx *gorm.DB, org *model.Organization) error {
	if err := utils.ValidateOrganization(org); err != nil {
		return err
	}

	// Create the organization within the transaction
	if err := s.repo.CreateWithTx(tx, org); err != nil {
		utils.LogError(s.log, "creating organization", err)
		return err
	}

	// Add the owner to the organization within the transaction
	userOrg := model.UserOrganization{
		UserID:         org.OwnerID,
		OrganizationID: org.ID,
	}
	if err := s.userOrgRepo.AddUserToOrganizationWithTx(tx, &userOrg); err != nil {
		utils.LogError(s.log, "adding user to organization", err)
		return err
	}

	utils.LogSuccess(s.log, "organization created", "Organization successfully created", org.OwnerID)
	return nil
}

// CreateOrganization handles the creation of a new organization.
func (s *OrganizationService) CreateOrganization(org *model.Organization) error {
	if err := utils.ValidateOrganization(org); err != nil {
		return err
	}

	// Start a transaction
	tx := s.db.Begin()
	if tx.Error != nil {
		utils.LogError(s.log, "starting transaction", tx.Error)
		return tx.Error
	}

	// Create the organization within the transaction
	if err := s.repo.CreateWithTx(tx, org); err != nil {
		tx.Rollback()
		utils.LogError(s.log, "creating organization", err)
		return err
	}

	// Add the owner to the organization within the transaction
	userOrg := model.UserOrganization{
		UserID:         org.OwnerID,
		OrganizationID: org.ID,
	}
	if err := s.userOrgRepo.AddUserToOrganizationWithTx(tx, &userOrg); err != nil {
		tx.Rollback()
		utils.LogError(s.log, "adding user to organization", err)
		return err
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		utils.LogError(s.log, "committing transaction", err)
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
func (s *OrganizationService) DeleteOrganization(userID uint, id uint) error {
	if err := s.userOrgRepo.RemoveUserFromOrganization(userID, id); err != nil {
		utils.LogError(s.log, "removing user from organization", err)
		return err
	}
	if err := s.repo.Delete(id); err != nil {
		utils.LogError(s.log, "deleting organization", err)
		return err
	}
	utils.LogSuccess(s.log, "organization deleted", "Organization successfully deleted", id)
	return nil
}

// GetUserOrganizations retrieves the organizations a user belongs to.
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

// AddUserToOrganization associates a user with an organization.
func (s *OrganizationService) AddUserToOrganization(userID, orgID uint) error {
	userOrg := model.UserOrganization{
		UserID:         userID,
		OrganizationID: orgID,
	}

	if err := s.userOrgRepo.AddUserToOrganizationWithTx(s.db, &userOrg); err != nil {
		utils.LogError(s.log, "adding user to organization", err)
		return err
	}
	utils.LogSuccess(s.log, "user added to organization", "User successfully added to organization", userID)
	return nil
}
