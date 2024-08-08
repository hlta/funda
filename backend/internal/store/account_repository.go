package store

import (
	"funda/internal/model"

	"gorm.io/gorm"
)

type AccountRepository interface {
	Create(account *model.Account) error
	Update(account *model.Account) error
	Delete(account *model.Account) error
	FindById(id uint) (*model.Account, error)
	FindByIdAndOrg(id uint, orgID uint) (*model.Account, error)
	FindAll() ([]model.Account, error)
	FindAllByOrg(orgID uint) ([]model.Account, error)
	FindByCodeAndOrg(code int, orgID uint) (*model.Account, error)
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{db: db}
}

func (r *accountRepository) Create(account *model.Account) error {
	return r.db.Create(account).Error
}

func (r *accountRepository) Update(account *model.Account) error {
	return r.db.Save(account).Error
}

func (r *accountRepository) Delete(account *model.Account) error {
	return r.db.Delete(account).Error
}

func (r *accountRepository) FindById(id uint) (*model.Account, error) {
	var account model.Account
	err := r.db.First(&account, id).Error
	return &account, err
}

func (r *accountRepository) FindByIdAndOrg(id uint, orgID uint) (*model.Account, error) {
	var account model.Account
	err := r.db.Where("id = ? AND org_id = ?", id, orgID).First(&account).Error
	return &account, err
}

func (r *accountRepository) FindAll() ([]model.Account, error) {
	var accounts []model.Account
	err := r.db.Find(&accounts).Error
	return accounts, err
}

func (r *accountRepository) FindAllByOrg(orgID uint) ([]model.Account, error) {
	var accounts []model.Account
	err := r.db.Where("org_id = ?", orgID).Find(&accounts).Error
	return accounts, err
}

func (r *accountRepository) FindByCodeAndOrg(code int, orgID uint) (*model.Account, error) {
	var account model.Account
	err := r.db.Where("code = ? AND org_id = ?", code, orgID).First(&account).Error
	return &account, err
}
