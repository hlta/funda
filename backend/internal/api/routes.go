package api

import (
	"funda/internal/service"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, userService *service.UserService, authService *service.AuthService, orgService *service.OrganizationService) {
	userHandler := NewUserHandler(userService)
	authHandler := NewAuthHandler(authService)
	orgHandler := NewOrganizationHandler(orgService)

	authHandler.Register(e)
	userHandler.Register(e)
	orgHandler.Register(e)
}
