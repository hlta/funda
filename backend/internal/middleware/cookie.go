package middleware

import (
	"funda/internal/constants"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// SetCookieMiddleware is a middleware to set the token cookie.
func SetCookieMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get(constants.TokenCookieName).(string)
		if !ok || token == "" {
			return echo.NewHTTPError(http.StatusInternalServerError, constants.TokenNotFound)
		}
		setCookie(c, token, 24*time.Hour)
		return next(c)
	}
}

// ClearCookieMiddleware is a middleware to clear the token cookie.
func ClearCookieMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		setCookie(c, "", -time.Hour)
		return next(c)
	}
}

// setCookie sets a cookie with the specified token and duration.
func setCookie(c echo.Context, token string, duration time.Duration) {
	cookie := &http.Cookie{
		Name:     constants.TokenCookieName,
		Value:    token,
		Expires:  time.Now().Add(duration),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	}
	c.SetCookie(cookie)
}
