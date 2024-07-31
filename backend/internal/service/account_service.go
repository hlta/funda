package service

import (
	"funda/internal/model"
	"funda/internal/store"
)

type AccountService interface {
	CreateAccount(account *model.Account) error
	UpdateAccount(account *model.Account) error
	DeleteAccount(account *model.Account) error
	GetAccountById(id uint) (*model.Account, error)
	GetAllAccounts() ([]model.Account, error)
}

type accountService struct {
	accountRepo store.AccountRepository
}

func NewAccountService(accountRepo store.AccountRepository) AccountService {
	return &accountService{accountRepo: accountRepo}
}

func (s *accountService) CreateAccount(account *model.Account) error {
	return s.accountRepo.Create(account)
}

func (s *accountService) UpdateAccount(account *model.Account) error {
	return s.accountRepo.Update(account)
}

func (s *accountService) DeleteAccount(account *model.Account) error {
	return s.accountRepo.Delete(account)
}

func (s *accountService) GetAccountById(id uint) (*model.Account, error) {
	return s.accountRepo.FindById(id)
}

func (s *accountService) GetAllAccounts() ([]model.Account, error) {
	return s.accountRepo.FindAll()
}
