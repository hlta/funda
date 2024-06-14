package middleware

import (
	"funda/configs"
	"funda/internal/logger"

	"github.com/labstack/echo/v4"
)

// SetupMiddlewares configures all the middleware for the Echo server.
func SetupMiddlewares(e *echo.Echo, log logger.Logger, config configs.Config) {
	e.Use(logger.EchoLogger(logger.EchoLoggerConfig{Logger: log}))
	e.Use(CORSMiddleware(config.CORS))
	e.Use(OAuthMiddleware(config.OAuth))
}
