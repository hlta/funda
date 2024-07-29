package model

import (
	"gorm.io/gorm"
)

type Expense struct {
	gorm.Model
	AccountID uint
	Account   Account `gorm:"foreignKey:AccountID"`
	Amount    float64
	Date      string
}
