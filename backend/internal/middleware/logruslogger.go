package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

// LogrusLoggerConfig defines the config for LogrusLogger middleware.
type LogrusLoggerConfig struct {
	Logger *logrus.Logger
}

// LogrusLogger returns an Echo middleware for logging requests using Logrus.
func LogrusLogger(config LogrusLoggerConfig) echo.MiddlewareFunc {
	// Set default output if none provided
	if config.Logger.Out == nil {
		config.Logger.Out = logrus.StandardLogger().Out
	}

	// Configure the Echo log format to use Logrus
	logFormat := `{"time":"${time_rfc3339}", "remote_ip":"${remote_ip}", "method":"${method}", "uri":"${uri}", "status":${status}, "error":"${error}"}` + "\n"

	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: config.Logger.Out,
		Format: logFormat,
	})
}
