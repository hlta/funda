package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

// Logger defines a standard logging interface that all services will use.
type Logger interface {
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	WithField(key string, value interface{}) *logrus.Entry
	Writer() io.Writer
}

// logrusLogger wraps around logrus to implement the Logger interface.
type logrusLogger struct {
	*logrus.Logger
}

// NewLogger creates and configures a new instance of Logger.
func NewLogger() Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	level, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logger.SetLevel(logrus.InfoLevel)
	} else {
		logger.SetLevel(level)
	}
	return &logrusLogger{logger}
}

// Writer returns the output destination of the logger.
func (l *logrusLogger) Writer() io.Writer {
	return l.Out
}
