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
		return err
	}

	if err := h.authService.Signup(&user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
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
