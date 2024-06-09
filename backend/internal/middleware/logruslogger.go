package middleware

import (
	"funda/internal/logger" // Adjust the path based on your actual package structure

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// LogrusLoggerConfig defines the config for LogrusLogger middleware.
type LogrusLoggerConfig struct {
	Logger logger.Logger
}

// LogrusLogger returns an Echo middleware for logging requests using the custom Logger.
func LogrusLogger(config LogrusLoggerConfig) echo.MiddlewareFunc {
	logWriter := config.Logger.Writer()

	logFormat := `{"time":"${time_rfc3339}", "remote_ip":"${remote_ip}", "method":"${method}", "uri":"${uri}", "status":${status}, "error":"${error}"}` + "\n"

	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: logWriter,
		Format: logFormat,
	})
}
