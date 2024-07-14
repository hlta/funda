package main

import (
	"funda/configs"
	"funda/internal/api"
	"funda/internal/auth"
	"funda/internal/db"
	"funda/internal/logger"
	"funda/internal/model"
	"funda/internal/service"
	"funda/internal/store"
	"path/filepath"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// initializeConfig loads and returns the application configuration
func initializeConfig(appLogger logger.Logger) configs.Config {
	appLogger.Info("Loading configuration")
	config, err := configs.LoadConfig(".")
	if err != nil {
		appLogger.WithField("error", err).Fatal("Error loading configuration")
	}
	return config
}

// initializeDatabase sets up and returns the database connection
func initializeDatabase(config configs.Config, dbLogger logger.Logger) *gorm.DB {
	dbLogger.Info("Setting up database")
	database, err := db.SetupDatabase(config.Database, dbLogger)
	if err != nil {
		dbLogger.WithField("error", err).Fatal("Failed to setup database")
	}

	dbLogger.Info("Auto-migrating database models")
	if err := database.AutoMigrate(&model.User{}, &model.Organization{}, &model.UserOrganization{}); err != nil {
		dbLogger.WithField("error", err).Fatal("Failed to auto-migrate")
	}

	return database
}

// initializeCasbin sets up and returns the Casbin enforcer
func initializeCasbin(database *gorm.DB, appLogger logger.Logger) *casbin.Enforcer {
	a, err := gormadapter.NewAdapterByDB(database)
	if err != nil {
		appLogger.WithField("error", err).Fatal("Failed to create Casbin adapter")
	}

	modelPath := filepath.Join("internal", "casbin", "model.conf")
	enforcer, err := casbin.NewEnforcer(modelPath, a)
	if err != nil {
		appLogger.WithField("error", err).Fatal("Failed to create Casbin enforcer")
	}

	if err = enforcer.LoadPolicy(); err != nil {
		appLogger.WithField("error", err).Fatal("Failed to load policy")
	}

	return enforcer
}

// initializeServices initializes and returns all the necessary services
func initializeServices(database *gorm.DB) (*service.UserService, *service.OrganizationService, *service.AuthService) {
	userLogger := logger.NewLogger("userService")
	userRepository := store.NewGormUserRepository(database)
	userService := service.NewUserService(userRepository, userLogger)

	orgRepository := store.NewGormOrganizationRepository(database)
	userOrgRepository := store.NewGormUserOrganizationRepository(database)

	orgLogger := logger.NewLogger("orgService")
	orgService := service.NewOrganizationService(orgRepository, userOrgRepository, orgLogger, database)

	authLogger := logger.NewLogger("authService")
	authService := service.NewAuthService(userService, orgService, authLogger, database)

	return userService, orgService, authService
}

func main() {
	appLogger := logger.NewLogger("default")
	config := initializeConfig(appLogger)

	appLogger.Info("Initialize the auth package with the JWT secret from the config")
	auth.SetupAuth(config.OAuth)

	dbLogger := logger.NewLogger("database")
	database := initializeDatabase(config, dbLogger)

	enforcer := initializeCasbin(database, appLogger)

	_, orgService, authService := initializeServices(database)

	e := echo.New()

	// Setup dependencies
	deps := &api.Dependencies{
		Config:      config,
		Logger:      appLogger,
		AuthService: authService,
		OrgService:  orgService,
		Enforcer:    enforcer,
	}

	api.SetupRoutes(e, deps)

	appLogger.WithField("port", config.Server.Port).Info("Starting server")
	e.Logger.Fatal(e.Start(":" + config.Server.Port))
}
