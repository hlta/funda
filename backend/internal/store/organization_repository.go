package store

import (
	"funda/internal/model"

	"gorm.io/gorm"
)

// GormOrganizationRepository is the GORM implementation of the OrganizationRepository interface.
type GormOrganizationRepository struct {
	DB *gorm.DB
}

func NewGormOrganizationRepository(db *gorm.DB) *GormOrganizationRepository {
	return &GormOrganizationRepository{DB: db}
}

func (r *GormOrganizationRepository) Create(org *model.Organization) error {
	return r.DB.Create(org).Error
}

func (r *GormOrganizationRepository) RetrieveByID(id uint) (*model.Organization, error) {
	org := &model.Organization{}
	result := r.DB.First(org, id)
	return org, result.Error
}

func (r *GormOrganizationRepository) Update(org *model.Organization) error {
	return r.DB.Save(org).Error
}

func (r *GormOrganizationRepository) Delete(id uint) error {
	return r.DB.Delete(&model.Organization{}, id).Error
}
