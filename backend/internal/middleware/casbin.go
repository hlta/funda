package middleware

import (
	"funda/internal/auth"
	"funda/internal/constants"
	"net/http"
	"strconv"

	"github.com/casbin/casbin/v2"
	"github.com/labstack/echo/v4"
)

// CasbinMiddleware returns an Echo middleware that enforces Casbin authorization.
func CasbinMiddleware(enforcer *casbin.Enforcer) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Extract user and other context information
			user, ok := c.Get(constants.UserClaimsKey).(auth.Claims)
			if !ok {
				return c.JSON(http.StatusForbidden, map[string]interface{}{
					"message": constants.UserNotFoundInContext,
				})
			}

			org := strconv.FormatUint(uint64(user.OrgID), 10)
			if org == "0" {
				return c.JSON(http.StatusForbidden, map[string]interface{}{
					"message": constants.OrganizationNotFoundInRequest,
				})
			}

			path := c.Request().URL.Path
			method := c.Request().Method

			// Get roles for the user from Casbin
			roles, err := enforcer.GetRolesForUser(user.ID, org)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"message": constants.ErrorRetrievingRolesForUser,
				})
			}

			// Check permissions for each role
			allowed := false
			for _, role := range roles {
				allowed, err = enforcer.Enforce(role, org, path, method)
				if err != nil {
					return c.JSON(http.StatusInternalServerError, map[string]interface{}{
						"message": constants.ErrorDuringAuthorization,
					})
				}
				if allowed {
					break
				}
			}

			if !allowed {
				return c.JSON(http.StatusForbidden, map[string]interface{}{
					"message": constants.Forbidden,
				})
			}

			return next(c)
		}
	}
}
