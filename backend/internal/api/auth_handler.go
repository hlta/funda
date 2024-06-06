package api

import (
	"funda/internal/model"
	"funda/internal/service"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Register(e *echo.Echo) {
	e.POST("/signup", h.Signup)
	e.POST("/login", h.Login)
}

func (h *AuthHandler) Signup(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request details"})
	}

	if err := h.authService.Signup(&user); err != nil {
		if strings.Contains(err.Error(), "unique constraint") {
			return c.JSON(http.StatusConflict, map[string]string{"error": "Email already in use"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to register user"})
	}
	return c.JSON(http.StatusCreated, map[string]string{"message": "User successfully registered"})
}

func (h *AuthHandler) Login(c echo.Context) error {
	var loginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.Bind(&loginReq); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request details"})
	}

	token, err := h.authService.Login(loginReq.Email, loginReq.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}
