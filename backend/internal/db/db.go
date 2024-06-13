package db

import (
	"fmt"
	"funda/configs"
	"funda/internal/logger"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDatabase(cfg configs.DatabaseConfig, log logger.Logger) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	switch cfg.Type {
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(cfg.Name), &gorm.Config{TranslateError: true})
	case "postgres":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.Host, cfg.Username, cfg.Password, cfg.Name, cfg.Port)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
