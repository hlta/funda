package model

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	VendorID uint
	Vendor   Vendor `gorm:"foreignKey:VendorID"`
	Date     string
	Amount   float64
	Status   string // e.g., Pending, Completed
}
