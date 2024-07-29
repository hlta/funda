package api

import (
	"funda/configs"
	"funda/internal/logger"
	"funda/internal/middleware"
	"funda/internal/service"

	"github.com/casbin/casbin/v2"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

// Dependencies holds all the dependencies for setting up routes
type Dependencies struct {
	Config      configs.Config
	Logger      logger.Logger
	UserService *service.UserService
	AuthService *service.AuthService
	OrgService  *service.OrganizationService
	Enforcer    *casbin.Enforcer
}

// SetupRoutes initializes all the routes and their handlers
func SetupRoutes(e *echo.Echo, deps *Dependencies) {
	// Global Middleware
	e.Use(middleware.RequestLogger(deps.Logger))
	e.Use(echoMiddleware.Recover())
	e.Use(middleware.CORSMiddleware(deps.Config.CORS))
	e.Use(middleware.ErrorHandlingMiddleware)

	// Initialize Handlers
	handlers := NewHandlers(deps.AuthService, deps.OrgService, deps.Enforcer)

	// Register Routes
	registerPublicRoutes(e, handlers.AuthHandler)
	registerProtectedRoutes(e, deps, handlers)
}

// registerPublicRoutes registers routes that do not require authentication
func registerPublicRoutes(e *echo.Echo, authHandler *AuthHandler) {
	authHandler.Register(e)
}

// registerProtectedRoutes registers routes that require authentication
func registerProtectedRoutes(e *echo.Echo, deps *Dependencies, handlers *Handlers) {
	protectedRoutes := e.Group("/api")
	protectedRoutes.Use(middleware.OAuthMiddleware(deps.Config.OAuth, deps.Logger))
	protectedRoutes.Use(middleware.CasbinMiddleware(deps.Enforcer, deps.Logger))

	handlers.OrganizationHandler.Register(protectedRoutes)
}
