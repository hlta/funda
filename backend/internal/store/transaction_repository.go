package store

import (
	"funda/internal/model"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(transaction *model.Transaction) error
	Update(transaction *model.Transaction) error
	Delete(transaction *model.Transaction) error
	FindById(id uint) (*model.Transaction, error)
	FindAll() ([]model.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) Create(transaction *model.Transaction) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(transaction).Error; err != nil {
			return err
		}
		return r.updateAccountBalance(tx, transaction, "create")
	})
	return err
}

func (r *transactionRepository) Update(transaction *model.Transaction) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		var oldTransaction model.Transaction
		if err := tx.First(&oldTransaction, transaction.ID).Error; err != nil {
			return err
		}
		if err := tx.Save(transaction).Error; err != nil {
			return err
		}
		return r.updateAccountBalance(tx, transaction, "update", &oldTransaction)
	})
	return err
}

func (r *transactionRepository) Delete(transaction *model.Transaction) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(transaction).Error; err != nil {
			return err
		}
		return r.updateAccountBalance(tx, transaction, "delete")
	})
	return err
}

func (r *transactionRepository) FindById(id uint) (*model.Transaction, error) {
	var transaction model.Transaction
	err := r.db.First(&transaction, id).Error
	return &transaction, err
}

func (r *transactionRepository) FindAll() ([]model.Transaction, error) {
	var transactions []model.Transaction
	err := r.db.Find(&transactions).Error
	return transactions, err
}

func (r *transactionRepository) updateAccountBalance(tx *gorm.DB, transaction *model.Transaction, operation string, oldTransaction ...*model.Transaction) error {
	var account model.Account
	if err := tx.First(&account, transaction.AccountID).Error; err != nil {
		return err
	}

	switch operation {
	case "create":
		if transaction.Type == "Debit" {
			account.Balance += transaction.Amount
		} else if transaction.Type == "Credit" {
			account.Balance -= transaction.Amount
		}
	case "update":
		if len(oldTransaction) > 0 {
			oldTx := oldTransaction[0]
			if oldTx.Type == "Debit" {
				account.Balance -= oldTx.Amount
			} else if oldTx.Type == "Credit" {
				account.Balance += oldTx.Amount
			}
		}
		if transaction.Type == "Debit" {
			account.Balance += transaction.Amount
		} else if transaction.Type == "Credit" {
			account.Balance -= transaction.Amount
		}
	case "delete":
		if transaction.Type == "Debit" {
			account.Balance -= transaction.Amount
		} else if transaction.Type == "Credit" {
			account.Balance += transaction.Amount
		}
	}

	return tx.Save(&account).Error
}
