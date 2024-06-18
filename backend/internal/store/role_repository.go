package store

import (
	"funda/internal/model"

	"gorm.io/gorm"
)

// GormRoleRepository is the GORM implementation of the RoleRepository interface.
type GormRoleRepository struct {
	DB *gorm.DB
}

func NewGormRoleRepository(db *gorm.DB) *GormRoleRepository {
	return &GormRoleRepository{DB: db}
}

func (r *GormRoleRepository) Create(role *model.Role) error {
	return r.DB.Create(role).Error
}

func (r *GormRoleRepository) RetrieveByID(id uint) (*model.Role, error) {
	role := &model.Role{}
	result := r.DB.Preload("Permissions").First(role, id)
	return role, result.Error
}

func (r *GormRoleRepository) RetrieveByName(name string) (*model.Role, error) {
	role := &model.Role{}
	result := r.DB.Preload("Permissions").Where("name = ?", name).First(role)
	return role, result.Error
}

func (r *GormRoleRepository) Update(role *model.Role) error {
	return r.DB.Save(role).Error
}

func (r *GormRoleRepository) Delete(id uint) error {
	return r.DB.Delete(&model.Role{}, id).Error
}
