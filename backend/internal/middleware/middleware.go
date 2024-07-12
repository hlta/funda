package middleware

import (
	"funda/configs"
	"funda/internal/logger"

	"github.com/casbin/casbin/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// SetupMiddlewares configures all the middleware for the Echo server.
func SetupMiddlewares(e *echo.Echo, log logger.Logger, config configs.Config, enforcer *casbin.Enforcer) {
	e.Use(logger.EchoLogger(logger.EchoLoggerConfig{Logger: log}))
	e.Use(middleware.Recover())
	e.Use(CORSMiddleware(config.CORS))
	e.Use(OAuthMiddleware(config.OAuth))
	e.Use(CasbinMiddleware(enforcer))

}
