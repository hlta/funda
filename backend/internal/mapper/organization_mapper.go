package mapper

import (
	"funda/internal/model"
	"funda/internal/response"
)

// ToOrganizationResponse maps a model.Organization to a response.OrganizationResponse.
func ToOrganizationResponse(org model.Organization) response.OrganizationResponse {
	var industry string
	if org.Industry != nil {
		industry = *org.Industry
	}

	var gstRegistered bool
	if org.GSTRegistered != nil {
		gstRegistered = *org.GSTRegistered
	}

	return response.OrganizationResponse{
		ID:            org.ID,
		Name:          org.Name,
		Industry:      industry,
		GSTRegistered: gstRegistered,
		OwnerID:       org.OwnerID,
	}
}
