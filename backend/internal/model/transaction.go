package model

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	AccountID uint
	Account   Account `gorm:"foreignKey:AccountID"`
	Type      string  // e.g., Debit, Credit
	Amount    float64
	Date      string
}
