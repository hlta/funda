package middleware

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func SetCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("token").(string)
		cookie := new(http.Cookie)
		cookie.Name = "token"
		cookie.Value = token
		cookie.Expires = time.Now().Add(24 * time.Hour)
		cookie.HttpOnly = true
		cookie.Secure = true
		cookie.SameSite = http.SameSiteStrictMode
		cookie.Path = "/"
		c.SetCookie(cookie)
		return next(c)
	}
}

func ClearCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie := new(http.Cookie)
		cookie.Name = "token"
		cookie.Value = ""
		cookie.Expires = time.Now().Add(-time.Hour)
		cookie.HttpOnly = true
		cookie.Secure = true
		cookie.SameSite = http.SameSiteStrictMode
		cookie.Path = "/"
		c.SetCookie(cookie)
		return next(c)
	}
}
