package utils

import (
	"errors"
	"funda/internal/model"
)

func ValidateOrganization(org *model.Organization) error {
	if org.Name == "" {
		return errors.New("organization name is required")
	}
	return nil
}
