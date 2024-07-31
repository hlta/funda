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
	FindAll() ([]model.Account, error)
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

func (r *accountRepository) FindAll() ([]model.Account, error) {
	var accounts []model.Account
	err := r.db.Find(&accounts).Error
	return accounts, err
}
