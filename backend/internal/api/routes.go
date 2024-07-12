package api

import (
	"funda/internal/service"

	"github.com/casbin/casbin/v2"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, userService *service.UserService, authService *service.AuthService, orgService *service.OrganizationService, enforcer *casbin.Enforcer) {
	authHandler := NewAuthHandler(authService, enforcer)
	orgHandler := NewOrganizationHandler(orgService, enforcer)

	authHandler.Register(e)
	orgHandler.Register(e)
}
