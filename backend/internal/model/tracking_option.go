package model

import (
	"gorm.io/gorm"
)

// TrackingOption represents an option within a tracking category.
type TrackingOption struct {
	gorm.Model
	TrackingCategoryID uint   `json:"tracking_category_id"`
	Name               string `json:"name"`
}
