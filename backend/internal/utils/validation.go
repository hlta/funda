package utils

import (
	"errors"
	"funda/internal/model"
	"regexp"
)

func ValidateOrganization(org *model.Organization) error {
	if org.Name == "" {
		return errors.New("organization name cannot be empty")
	}
	if !regexp.MustCompile(`^[a-zA-Z0-9 ]+$`).MatchString(org.Name) {
		return errors.New("organization name contains invalid characters")
	}
	return nil
}
