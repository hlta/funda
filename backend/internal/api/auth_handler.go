package api

import (
	"funda/internal/model"
	"funda/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authService *service.AuthService
}

// Constructor for AuthHandler, injecting dependencies
func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// Register routes for authentication
func (h *AuthHandler) Register(e *echo.Echo) {
	e.POST("/signup", h.Signup)
	e.POST("/login", h.Login)
}

// Handler for user signup
func (h *AuthHandler) Signup(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		c.Logger().Errorf("Signup: Failed to bind user data: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request details"})
	}

	if err := h.authService.Signup(&user); err != nil {
		c.Logger().Errorf("Signup: Failed: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Signup failure"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "User successfully registered"})
}

// Handler for user login
func (h *AuthHandler) Login(c echo.Context) error {
	var loginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.Bind(&loginReq); err != nil {
		c.Logger().Errorf("Login: Failed to bind login data: %v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request details"})
	}

	token, err := h.authService.Login(loginReq.Email, loginReq.Password)
	if err != nil {
		c.Logger().Errorf("Login: Failed for %s: %v", loginReq.Email, err)
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}
