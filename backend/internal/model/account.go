package model

import (
	"gorm.io/gorm"
)

// Account represents an account in the accounting system.
type Account struct {
	gorm.Model
	Code               int                `json:"code" gorm:"unique;not null"`
	Name               string             `json:"name" gorm:"not null"`
	Type               string             `json:"type" gorm:"not null"` // e.g., Asset, Liability, Equity, Revenue, Expense
	TaxRate            string             `json:"tax_rate"`
	Balance            float64            `json:"balance" gorm:"default:0"`
	OrgID              uint               `json:"org_id"`
	Organization       Organization       `gorm:"foreignKey:OrgID"`
	Transactions       []Transaction      `json:"transactions"`
	TrackingCategories []TrackingCategory `gorm:"many2many:account_tracking_categories;"`
}
