package main

import (
	"fmt"
	"funda/configs"
	"funda/internal/api"
	"funda/internal/auth"
	"funda/internal/middleware"
	"funda/internal/model"
	"funda/internal/service"
	"funda/internal/store"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var log = setupLogger()

func setupLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	level, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logger.SetLevel(logrus.InfoLevel)
	} else {
		logger.SetLevel(level)
	}
	return logger
}

func main() {
	log.Info("Loading configuration")
	config, err := configs.LoadConfig(".")
	if err != nil {
		log.WithField("error", err).Fatal("Error loading configuration")
	}

	log.Info("Initialize the auth package with the JWT secret from the config")
	auth.SetupAuth(config.OAuth)

	log.Info("Setting up database")
	db, err := setupDatabase(config.Database)
	if err != nil {
		log.WithField("error", err).Fatal("Failed to setup database")
	}

	log.Info("Auto-migrating database models")
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.WithField("error", err).Fatal("Failed to auto-migrate")
	}

	e := echo.New()
	e.Use(middleware.LogrusLogger(middleware.LogrusLoggerConfig{Logger: log}))
	e.Use(middleware.CORSMiddleware(config.CORS))
	e.Use(middleware.OAuthMiddleware(config.OAuth))
	userRepository := store.NewGormUserRepository(db)
	userService := service.NewUserService(userRepository)
	authService := service.NewAuthService(userRepository)

	userHandler := api.NewUserHandler(userService)
	authHandler := api.NewAuthHandler(authService)

	authHandler.Register(e)
	userHandler.Register(e)

	log.WithField("port", config.Server.Port).Info("Starting server")
	e.Logger.Fatal(e.Start(":" + config.Server.Port))
}

func setupDatabase(cfg configs.DatabaseConfig) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	switch cfg.Type {
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(cfg.Name), &gorm.Config{})
	case "postgres":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.Host, cfg.Username, cfg.Password, cfg.Name, cfg.Port)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	default:
		err = fmt.Errorf("unsupported database type: %s", cfg.Type)
	}
	if err != nil {
		log.WithFields(logrus.Fields{
			"type":  cfg.Type,
			"error": err,
		}).Error("Database setup failed")
	}
	return db, err
}
