package middleware

import (
	"errors"
	"funda/configs"
	"funda/internal/auth"
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
				return echo.NewHTTPError(http.StatusUnauthorized, "Authorization header is required")
			}

			tokenString, err := extractToken(authHeader)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}

			claims, err := auth.ValidateToken(tokenString)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or expired JWT token")
			}

			// Set user context using extracted claims
			c.Set("userClaims", claims)

			return next(c)
		}
	}
}

// isAuthOptional checks if the request path is excluded from authentication.
func isAuthOptional(path string) bool {
	switch path {
	case "/login", "/signup", "/auth/check", "/logout", "/auth/orgs", "/auth/switch-org":
		return true
	default:
		return false
	}
}

// extractToken retrieves the JWT token from the Authorization header.
func extractToken(authHeader string) (string, error) {
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("authorization header format must be 'Bearer {token}'")
	}
	return parts[1], nil
}
