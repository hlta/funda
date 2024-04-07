package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"funda/configs"
	"funda/internal/api"
	"funda/internal/auth"
	"funda/internal/middleware"
	"funda/internal/model"
	"funda/internal/service"
	"funda/internal/store"
)

func main() {
	// Load configuration
	config, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	// Initialize the auth package with the JWT secret from the config
	auth.SetupAuth(config.OAuth)

	// Setup database connection
	db, err := setupDatabase(config.Database)
	if err != nil {
		log.Fatalf("Failed to setup database: %v", err)
	}

	// Perform auto-migration
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate: %v", err)
	}

	// Initialize Echo and other components as before...
	e := echo.New()
	e.Use(middleware.OAuthMiddleware(config.OAuth)) // Adjusted for your setup

	userRepository := store.NewGormUserRepository(db)
	userService := service.NewUserService(userRepository)
	authService := service.NewAuthService(userRepository)

	// Handler
	userHandler := api.NewUserHandler(userService)
	authHandler := api.NewAuthHandler(authService)

	// Register routes
	authHandler.Register(e)
	userHandler.Register(e)

	// Start server
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

	return db, err
}
