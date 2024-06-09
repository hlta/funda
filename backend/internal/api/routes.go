package api

import (
	"funda/internal/service"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, userService *service.UserService, authService *service.AuthService) {
	userHandler := NewUserHandler(userService)
	authHandler := NewAuthHandler(authService)

	authHandler.Register(e)
	userHandler.Register(e)
}
