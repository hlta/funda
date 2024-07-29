package middleware

import (
	"time"

	"funda/internal/logger"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func RequestLogger(logger logger.Logger) echo.MiddlewareFunc {
	return echoMiddleware.RequestLoggerWithConfig(echoMiddleware.RequestLoggerConfig{
		LogURI:       true,
		LogStatus:    true,
		LogMethod:    true,
		LogLatency:   true,
		LogRemoteIP:  true,
		LogUserAgent: true,
		LogValuesFunc: func(c echo.Context, values echoMiddleware.RequestLoggerValues) error {
			latency := time.Since(values.StartTime).String()

			logEntry := logger.WithFields(logrus.Fields{
				"URI":        values.URI,
				"status":     values.Status,
				"method":     values.Method,
				"latency":    latency,
				"client_ip":  values.RemoteIP,
				"user_agent": values.UserAgent,
			})

			// Log at different levels based on status code
			if values.Status >= 500 {
				logEntry.Error("Server error")
			} else if values.Status >= 400 {
				logEntry.Warn("Client error")
			} else {
				logEntry.Info("request")
			}

			return nil
		},
	})
}
