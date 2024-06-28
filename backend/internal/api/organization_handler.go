package api

import (
	"funda/internal/auth"
	"funda/internal/middleware"
	"funda/internal/model"
	"funda/internal/response"
	"funda/internal/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type OrganizationHandler struct {
	orgService *service.OrganizationService
}

func NewOrganizationHandler(orgService *service.OrganizationService) *OrganizationHandler {
	return &OrganizationHandler{
		orgService: orgService,
	}
}

func (h *OrganizationHandler) Register(e *echo.Echo) {
	e.POST("/organizations", h.CreateOrganization)
	e.GET("/organizations/:id", h.GetOrganization, middleware.OrganizationOwnerMiddleware(h.orgService))
	e.PUT("/organizations/:id", h.UpdateOrganization, middleware.OrganizationOwnerMiddleware(h.orgService))
}

func (h *OrganizationHandler) CreateOrganization(c echo.Context) error {
	userClaims := c.Get("userClaims").(auth.Claims)

	var orgReq struct {
		Name          string `json:"name" validate:"required"`
		Industry      string `json:"industry"`
		GSTRegistered bool   `json:"gst_registered"`
	}

	if err := c.Bind(&orgReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, middleware.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request payload",
		})
	}

	org := &model.Organization{
		Name:          orgReq.Name,
		Industry:      &orgReq.Industry,
		GSTRegistered: &orgReq.GSTRegistered,
		OwnerID:       userClaims.UserID,
	}

	if err := h.orgService.CreateOrganization(org); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, middleware.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create organization",
		})
	}

	return c.JSON(http.StatusOK, response.GenericResponse{
		Message: "Organization created successfully",
		Data:    org,
	})
}

func (h *OrganizationHandler) GetOrganization(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, middleware.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid organization ID",
		})
	}

	org, err := h.orgService.GetOrganizationByID(uint(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, middleware.ErrorResponse{
			Code:    http.StatusNotFound,
			Message: "Organization not found",
		})
	}

	return c.JSON(http.StatusOK, response.GenericResponse{
		Message: "Organization retrieved successfully",
		Data:    org,
	})
}

func (h *OrganizationHandler) UpdateOrganization(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, middleware.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid organization ID",
		})
	}

	var orgReq struct {
		Name          string `json:"name" validate:"required"`
		Industry      string `json:"industry"`
		GSTRegistered bool   `json:"gst_registered"`
	}

	if err := c.Bind(&orgReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, middleware.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request payload",
		})
	}

	org := &model.Organization{
		Model:         gorm.Model{ID: uint(id)},
		Name:          orgReq.Name,
		Industry:      &orgReq.Industry,
		GSTRegistered: &orgReq.GSTRegistered,
	}

	if err := h.orgService.UpdateOrganization(org); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, middleware.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to update organization",
		})
	}

	return c.JSON(http.StatusOK, response.GenericResponse{
		Message: "Organization updated successfully",
		Data:    org,
	})
}
