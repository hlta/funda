package api

import (
	"funda/internal/model"
	"funda/internal/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// UserHandler holds the methods to handle user-related HTTP requests.
type UserHandler struct {
	UserService *service.UserService
}

// NewUserHandler creates a new instance of UserHandler.
func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

// Register registers the user routes to the given Echo instance.
func (h *UserHandler) Register(e *echo.Echo) {
	e.POST("/users", h.CreateUser)
	e.GET("/users/:id", h.GetUserByID)
	e.PUT("/users/:id", h.UpdateUser)
	e.DELETE("/users/:id", h.DeleteUser)
}

// CreateUser is the handler for the POST /users endpoint.
func (h *UserHandler) CreateUser(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := h.UserService.CreateUser(u); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, u)
}

// GetUserByID is the handler for the GET /users/:id endpoint.
func (h *UserHandler) GetUserByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	user, err := h.UserService.GetUserByID(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}
	return c.JSON(http.StatusOK, user)
}

// UpdateUser is the handler for the PUT /users/:id endpoint.
func (h *UserHandler) UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	u.ID = uint(id)
	if err := h.UserService.UpdateUser(u); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, u)
}

// DeleteUser is the handler for the DELETE /users/:id endpoint.
func (h *UserHandler) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	if err := h.UserService.DeleteUser(uint(id)); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
