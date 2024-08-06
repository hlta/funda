// api/account_handler.go
package api

import (
	"net/http"
	"strconv"

	"funda/internal/model"
	"funda/internal/response"
	"funda/internal/service"

	"github.com/casbin/casbin/v2"
	"github.com/labstack/echo/v4"
)

type AccountHandler struct {
	accountService service.AccountService
	enforcer       *casbin.Enforcer
}

func NewAccountHandler(accountService service.AccountService, enforcer *casbin.Enforcer) *AccountHandler {
	return &AccountHandler{accountService: accountService, enforcer: enforcer}
}

func (h *AccountHandler) Register(e *echo.Group) {
	e.POST("/accounts", h.CreateAccount)
	e.GET("/accounts/:id", h.GetAccountByID)
	e.GET("/accounts", h.GetAllAccounts)
	e.PUT("/accounts/:id", h.UpdateAccount)
	e.DELETE("/accounts/:id", h.DeleteAccount)
}

func (h *AccountHandler) CreateAccount(c echo.Context) error {
	var account model.Account
	if err := c.Bind(&account); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	if err := h.accountService.CreateAccount(&account); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create account")
	}

	accountResponse, _ := h.accountService.GetAccountResponseById(account.ID)
	return c.JSON(http.StatusOK, response.GenericResponse{
		Message: "Account created successfully",
		Data:    accountResponse,
	})
}

func (h *AccountHandler) GetAccountByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid account ID")
	}

	accountResponse, err := h.accountService.GetAccountResponseById(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Account not found")
	}

	return c.JSON(http.StatusOK, response.GenericResponse{
		Message: "Account retrieved successfully",
		Data:    accountResponse,
	})
}

func (h *AccountHandler) GetAllAccounts(c echo.Context) error {
	accountResponses, err := h.accountService.GetAllAccountResponses()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve accounts")
	}

	return c.JSON(http.StatusOK, response.GenericResponse{
		Message: "Accounts retrieved successfully",
		Data:    accountResponses,
	})
}

func (h *AccountHandler) UpdateAccount(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid account ID")
	}

	var account model.Account
	if err := c.Bind(&account); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	account.Model.ID = uint(id) // Set the ID using the embedded gorm.Model
	if err := h.accountService.UpdateAccount(&account); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update account")
	}

	accountResponse, _ := h.accountService.GetAccountResponseById(uint(id))
	return c.JSON(http.StatusOK, response.GenericResponse{
		Message: "Account updated successfully",
		Data:    accountResponse,
	})
}

func (h *AccountHandler) DeleteAccount(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid account ID")
	}

	account := model.Account{}
	account.Model.ID = uint(id) // Set the ID using the embedded gorm.Model
	if err := h.accountService.DeleteAccount(&account); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete account")
	}

	return c.NoContent(http.StatusNoContent)
}
