package store

import (
	"funda/internal/model"

	"gorm.io/gorm"
)

// GormUserOrganizationRepository is the GORM implementation of the UserOrganizationRepository interface.
type GormUserOrganizationRepository struct {
	DB *gorm.DB
}

func NewGormUserOrganizationRepository(db *gorm.DB) *GormUserOrganizationRepository {
	return &GormUserOrganizationRepository{DB: db}
}

func (r *GormUserOrganizationRepository) AddUserToOrganizationWithTx(tx *gorm.DB, userOrg *model.UserOrganization) error {
	return tx.Create(userOrg).Error
}

func (r *GormUserOrganizationRepository) RemoveUserFromOrganization(userID uint, orgID uint) error {
	return r.DB.Where("user_id = ? AND organization_id = ?", userID, orgID).Delete(&model.UserOrganization{}).Error
}

func (r *GormUserOrganizationRepository) GetUserOrganizations(userID uint) ([]model.UserOrganization, error) {
	var userOrgs []model.UserOrganization
	err := r.DB.Where("user_id = ?", userID).Find(&userOrgs).Error
	return userOrgs, err
}
