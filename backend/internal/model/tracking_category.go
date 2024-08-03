package model

import (
	"gorm.io/gorm"
)

// TrackingCategory represents a category for tracking transactions.
type TrackingCategory struct {
	gorm.Model
	Name    string           `json:"name"`
	Options []TrackingOption `json:"options"`
}
