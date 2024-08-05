package utils

import (
	"funda/internal/constants"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// ClearCookie clears the token cookie.
func ClearCookie(c echo.Context) {
	SetCookie(c, "", -time.Hour)
}

// SetCookie sets a cookie with the specified token and duration.
func SetCookie(c echo.Context, token string, duration time.Duration) {
	cookie := &http.Cookie{
		Name:     constants.TokenCookieName,
		Value:    token,
		Expires:  time.Now().Add(duration),
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteNoneMode,
		Path:     "/",
	}
	c.SetCookie(cookie)
}
