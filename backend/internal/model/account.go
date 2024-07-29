package model

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Name         string
	Type         string // e.g., Asset, Liability, Equity, Revenue, Expense
	Balance      float64
	OrgID        uint
	Organization Organization `gorm:"foreignKey:OrgID"`
}
