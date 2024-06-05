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

// Register registers the user routes to the given Echo instance.
func (h *AuthHandler) Register(e *echo.Echo) {
	e.POST("/signup", h.Signup)
	e.POST("/login", h.Login)
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Signup(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}

	// Attempt to create the user
	if err := h.authService.Signup(&user); err != nil {
		if strings.Contains(err.Error(), "unique constraint") {
			return echo.NewHTTPError(http.StatusConflict, "Email already in use")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to register user")
	}

	return c.NoContent(http.StatusCreated)
}

func (h *AuthHandler) Login(c echo.Context) error {
	var loginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.Bind(&loginReq); err != nil {
		return err
	}

	token, err := h.authService.Login(loginReq.Email, loginReq.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
