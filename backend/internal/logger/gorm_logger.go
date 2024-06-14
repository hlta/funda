package logger

import (
	"context"
	"time"

	"gorm.io/gorm/logger"
)

type gormLogger struct {
	log      Logger
	logLevel logger.LogLevel
}

// NewGormLogger creates a new instance of gormLogger with the specified logger and log level.
func NewGormLogger(log Logger, level logger.LogLevel) logger.Interface {
	return &gormLogger{
		log:      log,
		logLevel: level,
	}
}

// LogMode sets the log level for the logger.
func (g *gormLogger) LogMode(level logger.LogLevel) logger.Interface {
	g.logLevel = level
	return g
}

// Info logs informational messages.
func (g *gormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if g.logLevel >= logger.Info {
		g.log.WithField("context", ctx).Info(append([]interface{}{msg}, data...)...)
	}
}

// Warn logs warning messages.
func (g *gormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if g.logLevel >= logger.Warn {
		g.log.WithField("context", ctx).Warn(append([]interface{}{msg}, data...)...)
	}
}

// Error logs error messages.
func (g *gormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if g.logLevel >= logger.Error {
		g.log.WithField("context", ctx).Error(append([]interface{}{msg}, data...)...)
	}
}

// Trace logs SQL queries and their execution time.
func (g *gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if g.logLevel <= 0 {
		return
	}
	sql, rows := fc()
	elapsed := time.Since(begin)
	entry := g.log.WithFields(map[string]interface{}{
		"sql":           sql,
		"rows_affected": rows,
		"elapsed_time":  elapsed,
	})
	if err != nil {
		entry.WithField("error", err).Error("Trace")
	} else if elapsed > 200*time.Millisecond { // adjust the threshold as needed
		entry.Warn("Trace")
	} else {
		entry.Info("Trace")
	}
}
