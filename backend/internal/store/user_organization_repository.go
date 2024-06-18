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

func (r *GormUserOrganizationRepository) AddUserToOrganization(userOrg *model.UserOrganization) error {
	return r.DB.Create(userOrg).Error
}

func (r *GormUserOrganizationRepository) RemoveUserFromOrganization(userID uint, orgID uint) error {
	return r.DB.Where("user_id = ? AND organization_id = ?", userID, orgID).Delete(&model.UserOrganization{}).Error
}

func (r *GormUserOrganizationRepository) GetUserOrganizations(userID uint) ([]model.UserOrganization, error) {
	var userOrgs []model.UserOrganization
	result := r.DB.Where("user_id = ?", userID).Find(&userOrgs)
	return userOrgs, result.Error
}
