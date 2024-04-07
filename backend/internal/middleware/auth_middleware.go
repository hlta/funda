package middleware

import (
	"errors"
	"funda/configs"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func OAuthMiddleware(config configs.OAuthConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "authorization header is required")
			}

			tokenString, err := getTokenFromHeader(authHeader)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}

			token, err := validateToken(tokenString, config.JWTSecret)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid or expired JWT token")
			}

			// Optional: Inject token claims into context if needed for further processing
			c.Set("user", token.Claims)

			return next(c)
		}
	}
}

// getTokenFromHeader extracts the token from the Authorization header.
func getTokenFromHeader(authHeader string) (string, error) {
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("authorization header format must be Bearer {token}")
	}
	return parts[1], nil
}

// validateToken validates the JWT token using the provided secret or public key.
func validateToken(tokenString, secretKey string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}
