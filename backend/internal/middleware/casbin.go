package middleware

import (
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/labstack/echo/v4"
)

// CasbinMiddleware returns an Echo middleware that enforces Casbin authorization.
func CasbinMiddleware(enforcer *casbin.Enforcer) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Extract user and other context information
			user, ok := c.Get("user").(string)
			if !ok {
				return c.JSON(http.StatusForbidden, map[string]interface{}{
					"message": "user not found in context",
				})
			}

			org := c.Param("org")
			if org == "" {
				return c.JSON(http.StatusForbidden, map[string]interface{}{
					"message": "organization not found in request",
				})
			}

			path := c.Request().URL.Path
			method := c.Request().Method

			// Enforce Casbin policies
			allowed, err := enforcer.Enforce(user, org, path, method)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"message": "error occurred during authorization",
				})
			}

			if !allowed {
				return c.JSON(http.StatusForbidden, map[string]interface{}{
					"message": "forbidden",
				})
			}

			return next(c)
		}
	}
}
