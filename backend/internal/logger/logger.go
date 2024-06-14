package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Logger defines a standard logging interface that all services will use.
type Logger interface {
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	WithField(key string, value interface{}) *logrus.Entry
	WithFields(fields map[string]interface{}) *logrus.Entry
	Writer() io.Writer
	SetLevel(level logrus.Level)
	SetFormatter(formatter logrus.Formatter)
	GetLevel() logrus.Level
}

// logrusLogger wraps around logrus to implement the Logger interface.
type logrusLogger struct {
	*logrus.Logger
}

// loggers holds the loggers for different components.
var loggers = make(map[string]*logrusLogger)

// NewLogger creates and configures a new instance of Logger.
func NewLogger(component string) Logger {
	if logger, exists := loggers[component]; exists {
		return logger
	}

	logger := logrus.New()
	configureLogger(logger, component)
	loggers[component] = &logrusLogger{logger}
	return loggers[component]
}

func configureLogger(logger *logrus.Logger, component string) {
	logger.SetOutput(os.Stdout)

	level := viper.GetString("log." + component + ".level")
	if parsedLevel, err := logrus.ParseLevel(level); err == nil {
		logger.SetLevel(parsedLevel)
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}

	format := viper.GetString("log." + component + ".format")
	if format == "json" {
		logger.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logger.SetFormatter(&logrus.TextFormatter{})
	}
}

// Writer returns the output destination of the logger.
func (l *logrusLogger) Writer() io.Writer {
	return l.Out
}

func (l *logrusLogger) WithField(key string, value interface{}) *logrus.Entry {
	return l.Logger.WithField(key, value)
}

// WithFields adds multiple fields to the log entry.
func (l *logrusLogger) WithFields(fields map[string]interface{}) *logrus.Entry {
	return l.Logger.WithFields(fields)
}

// GetLevel returns the current log level of the logger.
func (l *logrusLogger) GetLevel() logrus.Level {
	return l.Logger.GetLevel()
}
