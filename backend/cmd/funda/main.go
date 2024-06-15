package main

import (
	"funda/configs"
	"funda/internal/api"
	"funda/internal/auth"
	"funda/internal/db"
	"funda/internal/logger"
	"funda/internal/middleware"
	"funda/internal/model"
	"funda/internal/service"
	"funda/internal/store"

	"github.com/labstack/echo/v4"
)

func main() {
	appLogger := logger.NewLogger("default")
	appLogger.Info("Loading configuration")

	config, err := configs.LoadConfig(".")
	if err != nil {
		appLogger.WithField("error", err).Fatal("Error loading configuration")
	}

	appLogger.Info("Initialize the auth package with the JWT secret from the config")
	auth.SetupAuth(config.OAuth)

	dbLogger := logger.NewLogger("database")
	appLogger.Info("Setting up database")
	database, err := db.SetupDatabase(config.Database, dbLogger)
	if err != nil {
		dbLogger.WithField("error", err).Fatal("Failed to setup database")
	}

	appLogger.Info("Auto-migrating database models")
	if err := database.AutoMigrate(&model.User{}); err != nil {
		dbLogger.WithField("error", err).Fatal("Failed to auto-migrate")
	}

	e := echo.New()
	appLogger.Info("Setting up middlewares")
	middleware.SetupMiddlewares(e, appLogger, config)

	userLogger := logger.NewLogger("userService")
	userRepository := store.NewGormUserRepository(database)
	userService := service.NewUserService(userRepository, userLogger)

	authLogger := logger.NewLogger("authService")
	authService := service.NewAuthService(userRepository, authLogger)

	api.SetupRoutes(e, userService, authService)

	appLogger.WithField("port", config.Server.Port).Info("Starting server")
	e.Logger.Fatal(e.Start(":" + config.Server.Port))
}
