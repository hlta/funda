package main

import (
	"funda/configs"
	"funda/internal/api"
	"funda/internal/auth"
	"funda/internal/db"
	"funda/internal/initializer"
	"funda/internal/logger"
	"funda/internal/middleware"
	"funda/internal/model"
	"funda/internal/service"
	"funda/internal/store"
	"log"
	"path/filepath"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
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
	if err := database.AutoMigrate(&model.User{}, &model.Organization{}, &model.UserOrganization{}); err != nil {
		dbLogger.WithField("error", err).Fatal("Failed to auto-migrate")
	}

	// Initialize the Casbin adapter
	a, err := gormadapter.NewAdapterByDB(database)
	if err != nil {
		log.Fatalf("Failed to create Casbin adapter: %v", err)
	}

	// Initialize the Casbin enforcer
	modelPath := filepath.Join("internal", "casbin", "model.conf")
	enforcer, err := casbin.NewEnforcer(modelPath, a)
	if err != nil {
		log.Fatalf("Failed to create Casbin enforcer: %v", err)
	}

	// Load the policy from DB
	if err = enforcer.LoadPolicy(); err != nil {
		log.Fatalf("Failed to load policy: %v", err)
	}

	// Load predefined roles and permissions
	policyPath := filepath.Join("internal", "casbin", "policy.csv")
	if err := initializer.LoadPoliciesFromCSV(enforcer, policyPath, appLogger); err != nil {
		log.Fatalf("Failed to load policies from CSV: %v", err)
	}

	e := echo.New()
	appLogger.Info("Setting up middlewares")
	middleware.SetupMiddlewares(e, appLogger, config, enforcer)

	userLogger := logger.NewLogger("userService")
	userRepository := store.NewGormUserRepository(database)
	userService := service.NewUserService(userRepository, userLogger)

	orgRepository := store.NewGormOrganizationRepository(database)
	userOrgRepository := store.NewGormUserOrganizationRepository(database)

	orgLogger := logger.NewLogger("orgService")
	orgService := service.NewOrganizationService(orgRepository, userOrgRepository, orgLogger)

	authLogger := logger.NewLogger("authService")
	authService := service.NewAuthService(userService, orgService, authLogger)

	api.SetupRoutes(e, userService, authService, orgService, enforcer)

	appLogger.WithField("port", config.Server.Port).Info("Starting server")
	e.Logger.Fatal(e.Start(":" + config.Server.Port))
}
