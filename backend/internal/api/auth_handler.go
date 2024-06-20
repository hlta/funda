package api

import (
	"errors"
	"funda/internal/middleware"
	"funda/internal/model"
	"funda/internal/response"
	"funda/internal/service"
	"net/http"
	"time"

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
	e.POST("/logout", h.Logout)
	e.GET("/auth/check", h.CheckAuth)
	e.GET("/auth/orgs", h.GetUserOrganizations)
	e.POST("/auth/switch-org", h.SwitchOrganization)
}

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

	return c.JSON(http.StatusCreated, response.GenericResponse{
		Message: "User successfully registered",
	})
}

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

	user, err := h.authService.Login(loginReq.Email, loginReq.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, middleware.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: "Invalid credentials",
		})
	}

	token, err := h.authService.GenerateToken(user, 0) // Initial login without org context
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, middleware.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to generate token",
		})
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HttpOnly = true
	cookie.Secure = true
	cookie.SameSite = http.SameSiteStrictMode
	cookie.Path = "/"
	c.SetCookie(cookie)

	userResp := &response.UserResponse{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Token:     token,
	}

	return c.JSON(http.StatusOK, response.GenericResponse{
		Message: "Login successful",
		Data:    userResp,
	})
}

func (h *AuthHandler) Logout(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now().Add(-time.Hour)
	cookie.HttpOnly = true
	cookie.Secure = true
	cookie.SameSite = http.SameSiteStrictMode
	cookie.Path = "/"
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, response.GenericResponse{
		Message: "Logout successful",
	})
}

func (h *AuthHandler) CheckAuth(c echo.Context) error {
	cookie, err := c.Cookie("token")
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, middleware.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: "Not authenticated",
		})
	}

	token := cookie.Value
	user, err := h.authService.VerifyToken(token)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, middleware.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: "Invalid token",
		})
	}

	userResp := &response.UserResponse{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Token:     token,
	}

	return c.JSON(http.StatusOK, response.GenericResponse{
		Message: "Authenticated",
		Data:    userResp,
	})
}

func (h *AuthHandler) GetUserOrganizations(c echo.Context) error {
	cookie, err := c.Cookie("token")
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, middleware.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: "Not authenticated",
		})
	}

	token := cookie.Value
	user, err := h.authService.VerifyToken(token)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, middleware.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: "Invalid token",
		})
	}

	orgs, err := h.authService.GetUserOrganizations(user.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, middleware.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to retrieve organizations",
		})
	}

	return c.JSON(http.StatusOK, response.GenericResponse{
		Message: "Organizations retrieved successfully",
		Data:    orgs,
	})
}

func (h *AuthHandler) SwitchOrganization(c echo.Context) error {
	cookie, err := c.Cookie("token")
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, middleware.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: "Not authenticated",
		})
	}

	token := cookie.Value
	user, err := h.authService.VerifyToken(token)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, middleware.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: "Invalid token",
		})
	}

	var switchOrgReq struct {
		OrgID uint `json:"org_id"`
	}
	if err := c.Bind(&switchOrgReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, middleware.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request details",
		})
	}

	// Generate a new token with the new organization context
	newToken, roles, permissions, err := h.authService.SwitchOrganization(user, switchOrgReq.OrgID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, middleware.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to generate token",
		})
	}

	newCookie := new(http.Cookie)
	newCookie.Name = "token"
	newCookie.Value = newToken
	newCookie.Expires = time.Now().Add(24 * time.Hour)
	newCookie.HttpOnly = true
	newCookie.Secure = true
	newCookie.SameSite = http.SameSiteStrictMode
	newCookie.Path = "/"
	c.SetCookie(newCookie)

	return c.JSON(http.StatusOK, response.GenericResponse{
		Message: "Organization switched successfully",
		Data: map[string]interface{}{
			"token":       newToken,
			"roles":       roles,
			"permissions": permissions,
		},
	})
}
