package db

import (
	"fmt"
	"funda/configs"
	"funda/internal/logger"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

func SetupDatabase(cfg configs.DatabaseConfig, log logger.Logger) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	// Initialize the Gorm logger with your custom logger
	gormLog := logger.NewGormLogger(log, gormLogger.LogLevel(log.GetLevel()))

	switch cfg.Type {
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(cfg.Name), &gorm.Config{
			Logger: gormLog,
		})
	case "postgres":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.Host, cfg.Username, cfg.Password, cfg.Name, cfg.Port)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: gormLog,
		})
	default:
		err = fmt.Errorf("unsupported database type: %s", cfg.Type)
	}
	if err != nil {
		// Using formatted strings to log the error
		log.Error(fmt.Sprintf("Database setup failed: type=%s, error=%v", cfg.Type, err))
		return nil, err
	}
	return db, nil
}
