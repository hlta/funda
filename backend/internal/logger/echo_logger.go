package logger

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// EchoLoggerConfig defines the config for LogrusLogger middleware.
type EchoLoggerConfig struct {
	Logger Logger
}

// LogrusLogger returns an Echo middleware for logging requests using the custom Logger.
func EchoLogger(config EchoLoggerConfig) echo.MiddlewareFunc {
	logWriter := config.Logger.Writer()

	logFormat := `{"time":"${time_rfc3339}", "remote_ip":"${remote_ip}", "method":"${method}", "uri":"${uri}", "status":${status}, "error":"${error}"}` + "\n"

	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: logWriter,
		Format: logFormat,
	})
}
