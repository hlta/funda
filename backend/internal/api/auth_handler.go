package api

import (
	"errors"
	"funda/internal/middleware"
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
		return echo.NewHTTPError(http.StatusBadRequest, middleware.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request details",
		})
	}

	if err := h.authService.Signup(&user); err != nil {
		if errors.Is(err, model.ErrEmailExists) {
			return echo.NewHTTPError(http.StatusConflict, middleware.ErrorResponse{
				Code:    http.StatusConflict,
				Message: "This email is already registered",
				Errors: []middleware.FieldError{
					{
						Field:   "email",
						Message: "This email is already in use",
					},
				},
			})
		}
		return echo.NewHTTPError(http.StatusInternalServerError, middleware.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "An unexpected error occurred. Please try again later.",
		})
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
		return echo.NewHTTPError(http.StatusBadRequest, middleware.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request details",
		})
	}

	token, err := h.authService.Login(loginReq.Email, loginReq.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, middleware.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: "Invalid credentials",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}
