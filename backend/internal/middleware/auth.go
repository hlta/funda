package middleware

import (
	"errors"
	"funda/configs"
	"funda/internal/auth"
	"funda/internal/constants"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// OAuthMiddleware creates a middleware for JWT validation.
func OAuthMiddleware(config configs.OAuthConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Allow unauthenticated access to login and signup endpoints.
			if isAuthOptional(c.Path()) {
				return next(c)
			}

			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, constants.AuthorizationHeaderRequired)
			}

			tokenString, err := extractToken(authHeader)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}

			claims, err := auth.ValidateToken(tokenString)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, constants.InvalidOrExpiredToken)
			}

			// Set user context using extracted claims
			c.Set(constants.UserClaimsKey, claims)

			return next(c)
		}
	}
}

// isAuthOptional checks if the request path is excluded from authentication.
func isAuthOptional(path string) bool {
	switch path {
	case constants.LoginRoute, constants.SignupRoute, constants.CheckAuthRoute, constants.LogoutRoute, constants.GetUserOrganizationsRoute, constants.SwitchOrgRoute:
		return true
	default:
		return false
	}
}

// extractToken retrieves the JWT token from the Authorization header.
func extractToken(authHeader string) (string, error) {
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New(constants.InvalidAuthorizationHeader)
	}
	return parts[1], nil
}
