package middleware

import (
	"funda/internal/auth"
	"funda/internal/service"
	"funda/internal/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// OrganizationOwnerMiddleware checks if the logged-in user is the owner of the organization.
func OrganizationOwnerMiddleware(orgService *service.OrganizationService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userClaims, ok := c.Get("userClaims").(auth.Claims)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "User claims are not available")
			}

			orgID, err := utils.ParseUint(c.Param("id"))
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, "Invalid organization ID format")
			}

			orgResp, err := orgService.GetOrganizationByID(uint(orgID))
			if err != nil {
				return echo.NewHTTPError(http.StatusNotFound, "Organization not found")
			}

			// Check if the current user is the owner of the organization
			if orgResp.OwnerID != userClaims.UserID {
				return echo.NewHTTPError(http.StatusForbidden, "Access denied")
			}

			return next(c)
		}
	}
}
