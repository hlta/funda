package model

import (
	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	CustomerID uint
	Customer   Customer `gorm:"foreignKey:CustomerID"`
	Date       string
	DueDate    string
	Amount     float64
	Status     string // e.g., Pending, Paid
}
