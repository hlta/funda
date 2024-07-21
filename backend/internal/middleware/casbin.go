package middleware

import (
	"funda/internal/auth"
	"funda/internal/constants"
	"funda/internal/logger"
	"funda/internal/utils"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/labstack/echo/v4"
)

// CasbinMiddleware returns an Echo middleware that enforces Casbin authorization.
func CasbinMiddleware(enforcer *casbin.Enforcer, log logger.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Extract user and other context information
			user, ok := c.Get(constants.UserClaimsKey).(*auth.Claims)
			if !ok {
				log.WithField("path", c.Path()).Warn(constants.UserNotFoundInContext)
				return c.JSON(http.StatusForbidden, map[string]interface{}{
					"message": constants.UserNotFoundInContext,
				})
			}

			org := utils.UintToString(user.OrgID)
			if org == "0" {
				log.WithField("path", c.Path()).Warn(constants.OrganizationNotFoundInRequest)
				return c.JSON(http.StatusForbidden, map[string]interface{}{
					"message": constants.OrganizationNotFoundInRequest,
				})
			}

			path := c.Request().URL.Path
			method := c.Request().Method

			// Get roles for the user from Casbin
			roles, err := enforcer.GetRolesForUser(utils.UintToString(user.UserID), org)
			if err != nil {
				log.WithFields(map[string]interface{}{
					"user":  user.UserID,
					"error": err,
				}).Error(constants.ErrorRetrievingRolesForUser)
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"message": constants.ErrorRetrievingRolesForUser,
				})
			}
			// Check permissions for each role
			allowed := false
			for _, role := range roles {
				allowed, err = enforcer.Enforce(role, org, path, method)
				if err != nil {
					log.WithFields(map[string]interface{}{
						"user":  user.UserID,
						"error": err,
					}).Error(constants.ErrorDuringAuthorization)
					return c.JSON(http.StatusInternalServerError, map[string]interface{}{
						"message": constants.ErrorDuringAuthorization,
					})
				}
				if allowed {
					break
				}
			}

			if !allowed {
				log.WithFields(map[string]interface{}{
					"user":   user.UserID,
					"path":   c.Path(),
					"method": method,
				}).Warn(constants.Forbidden)
				return c.JSON(http.StatusForbidden, map[string]interface{}{
					"message": constants.Forbidden,
				})
			}

			log.WithFields(map[string]interface{}{
				"user":   user.UserID,
				"path":   c.Path(),
				"method": method,
			}).Info("access granted")
			return next(c)
		}
	}
}
