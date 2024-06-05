package middleware

import (
	"funda/configs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// CORSMiddleware creates a configurable CORS middleware based on the given settings.
func CORSMiddleware(config configs.CORSConfig) echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     config.AllowOrigins,     // Origins allowed to access the server
		AllowMethods:     config.AllowMethods,     // Methods allowed for the CORS requests
		AllowHeaders:     config.AllowHeaders,     // Headers allowed in requests
		AllowCredentials: config.AllowCredentials, // Specifies whether credentials can be shared
		ExposeHeaders:    config.ExposeHeaders,    // Headers that browsers are allowed to access
		MaxAge:           config.MaxAge,           // Indicates how long the results of a preflight request can be cached
	})
}
