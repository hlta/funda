package utils

import (
	"funda/internal/logger"
)

func LogError(log logger.Logger, action string, err error) {
	log.WithField("action", action).WithError(err).Error("Failed to " + action)
}

func LogSuccess(log logger.Logger, action, message string, userID uint) {
	log.WithField("action", action).WithField("userID", userID).Info(message)
}
