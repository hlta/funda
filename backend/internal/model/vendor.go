package model

import (
	"gorm.io/gorm"
)

type Vendor struct {
	gorm.Model
	Name  string
	Email string
}
