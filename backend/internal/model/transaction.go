package model

import (
	"gorm.io/gorm"
)

// Transaction represents a financial transaction.
type Transaction struct {
	gorm.Model
	AccountID       uint             `json:"account_id"`
	Account         Account          `gorm:"foreignKey:AccountID"`
	Type            string           `json:"type"` // e.g., Debit, Credit
	Amount          float64          `json:"amount"`
	Date            string           `json:"date"`
	TrackingOptions []TrackingOption `gorm:"many2many:transaction_tracking_options;"`
}
