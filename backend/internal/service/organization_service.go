package service

import (
	"funda/internal/logger"
	"funda/internal/mapper"
	"funda/internal/model"
	"funda/internal/response"
	"funda/internal/seed"
	"funda/internal/utils"

	"gorm.io/gorm"
)

const (
	errCreatingOrganization = "error creating organization"
	errSeedingAccounts      = "error seeding accounts for organization"
	errAddingUser           = "error adding user to organization"
	successCreatingOrg      = "organization created successfully"
	successUpdatingOrg      = "organization updated successfully"
	successDeletingOrg      = "organization deleted successfully"
	successRetrievingOrg    = "organization retrieved successfully"
	successAddingUserToOrg  = "user added to organization successfully"
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

// CreateOrganizationWithTx creates an organization within a transaction.
func (s *OrganizationService) CreateOrganizationWithTx(tx *gorm.DB, org *model.Organization) error {
	if err := utils.ValidateOrganization(org); err != nil {
		return err
	}

	logContext := s.log.WithField("orgID", org.ID)

	if err := s.repo.CreateWithTx(tx, org); err != nil {
		logContext.Error(errCreatingOrganization, err)
		return err
	}

	if err := seed.SeedAccountsForOrg(tx, org.ID, s.log); err != nil {
		logContext.Error(errSeedingAccounts, err)
		return err
	}

	userOrg := model.UserOrganization{
		UserID:         org.OwnerID,
		OrganizationID: org.ID,
	}
	if err := s.userOrgRepo.AddUserToOrganizationWithTx(tx, &userOrg); err != nil {
		logContext.Error(errAddingUser, err)
		return err
	}

	logContext.Info(successCreatingOrg)
	return nil
}

// CreateOrganization handles the creation of a new organization.
func (s *OrganizationService) CreateOrganization(org *model.Organization) error {
	return s.withTransaction(func(tx *gorm.DB) error {
		return s.CreateOrganizationWithTx(tx, org)
	})
}

// withTransaction handles database transaction management.
func (s *OrganizationService) withTransaction(fn func(tx *gorm.DB) error) error {
	tx := s.db.Begin()
	if tx.Error != nil {
		s.log.Error("starting transaction", tx.Error)
		return tx.Error
	}

	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		s.log.Error("committing transaction", err)
		return err
	}

	return nil
}

// GetOrganizationByID retrieves an organization by its ID.
func (s *OrganizationService) GetOrganizationByID(id uint) (*response.OrganizationResponse, error) {
	org, err := s.repo.RetrieveByID(id)
	if err != nil {
		s.log.WithField("orgID", id).Error("retrieving organization by ID", err)
		return nil, err
	}
	s.log.WithField("orgID", id).Info(successRetrievingOrg)
	orgResp := mapper.ToOrganizationResponse(*org)
	return &orgResp, nil
}

// UpdateOrganization handles updates to an existing organization.
func (s *OrganizationService) UpdateOrganization(org *model.Organization) error {
	if err := utils.ValidateOrganization(org); err != nil {
		return err
	}
	if err := s.repo.Update(org); err != nil {
		s.log.WithField("orgID", org.ID).Error("updating organization", err)
		return err
	}
	s.log.WithField("orgID", org.ID).Info(successUpdatingOrg)
	return nil
}

// DeleteOrganization removes an organization by its ID.
func (s *OrganizationService) DeleteOrganization(userID uint, id uint) error {
	if err := s.userOrgRepo.RemoveUserFromOrganization(userID, id); err != nil {
		s.log.WithField("orgID", id).Error("removing user from organization", err)
		return err
	}
	if err := s.repo.Delete(id); err != nil {
		s.log.WithField("orgID", id).Error("deleting organization", err)
		return err
	}
	s.log.WithField("orgID", id).Info(successDeletingOrg)
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
			s.log.WithField("orgID", userOrg.OrganizationID).Error("retrieving organization", err)
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
		s.log.WithField("userID", userID).WithField("orgID", orgID).Error("adding user to organization", err)
		return err
	}
	s.log.WithField("userID", userID).WithField("orgID", orgID).Info(successAddingUserToOrg)
	return nil
}
