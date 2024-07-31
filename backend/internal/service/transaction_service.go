package service

import (
	"funda/internal/model"
	"funda/internal/store"
)

type TransactionService interface {
	CreateTransaction(transaction *model.Transaction) error
	UpdateTransaction(transaction *model.Transaction) error
	DeleteTransaction(transaction *model.Transaction) error
	GetTransactionById(id uint) (*model.Transaction, error)
	GetAllTransactions() ([]model.Transaction, error)
}

type transactionService struct {
	transactionRepo store.TransactionRepository
}

func NewTransactionService(transactionRepo store.TransactionRepository) TransactionService {
	return &transactionService{transactionRepo: transactionRepo}
}

func (s *transactionService) CreateTransaction(transaction *model.Transaction) error {
	return s.transactionRepo.Create(transaction)
}

func (s *transactionService) UpdateTransaction(transaction *model.Transaction) error {
	return s.transactionRepo.Update(transaction)
}

func (s *transactionService) DeleteTransaction(transaction *model.Transaction) error {
	return s.transactionRepo.Delete(transaction)
}

func (s *transactionService) GetTransactionById(id uint) (*model.Transaction, error) {
	return s.transactionRepo.FindById(id)
}

func (s *transactionService) GetAllTransactions() ([]model.Transaction, error) {
	return s.transactionRepo.FindAll()
}
