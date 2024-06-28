package mapper

import (
	"funda/internal/model"
	"funda/internal/response"
)

// ToOrganizationResponse maps a model.Organization to a response.OrganizationResponse.
func ToOrganizationResponse(org model.Organization) response.OrganizationResponse {
	return response.OrganizationResponse{
		ID:            org.ID,
		Name:          org.Name,
		Industry:      *org.Industry,
		GSTRegistered: *org.GSTRegistered,
		OwnerID:       org.OwnerID,
	}
}
