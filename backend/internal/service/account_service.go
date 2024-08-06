package service

import (
	"funda/internal/logger"
	"funda/internal/mapper"
	"funda/internal/model"
	"funda/internal/response"
	"funda/internal/store"
	"time"
)

type AccountService interface {
	CreateAccount(account *model.Account) error
	UpdateAccount(account *model.Account) error
	DeleteAccount(account *model.Account) error
	GetAccountById(id uint) (*model.Account, error)
	GetAllAccounts() ([]model.Account, error)
	GetAccountResponseById(id uint) (*response.AccountResponse, error)
	GetAllAccountResponses() ([]response.AccountResponse, error)
}

type accountService struct {
	accountRepo store.AccountRepository
	log         logger.Logger
}

func NewAccountService(accountRepo store.AccountRepository, log logger.Logger) AccountService {
	return &accountService{accountRepo: accountRepo, log: log}
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

func (s *accountService) GetAccountResponseById(id uint) (*response.AccountResponse, error) {
	account, err := s.accountRepo.FindById(id)
	if err != nil {
		return nil, err
	}
	ytd := s.calculateYTD(account.Transactions)
	accountResponse := mapper.ToAccountResponse(*account, ytd)
	return &accountResponse, nil
}

func (s *accountService) GetAllAccountResponses() ([]response.AccountResponse, error) {
	accounts, err := s.accountRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var accountResponses []response.AccountResponse
	for _, account := range accounts {
		ytd := s.calculateYTD(account.Transactions)
		accountResponse := mapper.ToAccountResponse(account, ytd)
		accountResponses = append(accountResponses, accountResponse)
	}
	return accountResponses, nil
}

func (s *accountService) calculateYTD(transactions []model.Transaction) float64 {
	var ytd float64
	currentYear := time.Now().Year()
	startOfFinancialYear := time.Date(currentYear, time.July, 1, 0, 0, 0, 0, time.Local)
	endOfFinancialYear := time.Date(currentYear+1, time.June, 30, 23, 59, 59, 999999999, time.Local)
	for _, txn := range transactions {
		txnDate, err := time.Parse("2006-01-02", txn.Date)
		if err != nil {
			continue
		}
		if txnDate.After(startOfFinancialYear) && txnDate.Before(endOfFinancialYear) {
			ytd += txn.Amount
		}
	}
	return ytd
}
