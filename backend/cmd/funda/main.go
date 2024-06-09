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
	log := logger.NewLogger()
	log.Info("Loading configuration")
	config, err := configs.LoadConfig(".")
	if err != nil {
		log.WithField("error", err).Fatal("Error loading configuration")
	}

	log.Info("Initialize the auth package with the JWT secret from the config")
	auth.SetupAuth(config.OAuth)

	log.Info("Setting up database")
	database, err := db.SetupDatabase(config.Database, log)
	if err != nil {
		log.WithField("error", err).Fatal("Failed to setup database")
	}

	log.Info("Auto-migrating database models")
	if err := database.AutoMigrate(&model.User{}); err != nil {
		log.WithField("error", err).Fatal("Failed to auto-migrate")
	}

	e := echo.New()
	middleware.SetupMiddlewares(e, log, config)

	userRepository := store.NewGormUserRepository(database)
	userService := service.NewUserService(userRepository, log)
	authService := service.NewAuthService(userRepository, log)

	api.SetupRoutes(e, userService, authService)

	log.WithField("port", config.Server.Port).Info("Starting server")
	e.Logger.Fatal(e.Start(":" + config.Server.Port))
}
