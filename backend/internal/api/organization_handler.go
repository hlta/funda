package api

import (
	"net/http"
	"strconv"

	"funda/internal/auth"
	"funda/internal/constants"
	"funda/internal/mapper"
	"funda/internal/middleware"
	"funda/internal/model"
	"funda/internal/response"
	"funda/internal/service"

	"github.com/casbin/casbin/v2"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type OrganizationHandler struct {
	orgService *service.OrganizationService
	enforcer   *casbin.Enforcer
}

func NewOrganizationHandler(orgService *service.OrganizationService, enforcer *casbin.Enforcer) *OrganizationHandler {
	return &OrganizationHandler{orgService: orgService, enforcer: enforcer}
}

func (h *OrganizationHandler) Register(e *echo.Group) {
	e.POST(constants.CreateOrganizationRoute, h.CreateOrganization)
	e.GET(constants.GetOrganizationRoute, h.GetOrganization, middleware.OrganizationOwnerMiddleware(h.orgService))
	e.PUT(constants.UpdateOrganizationRoute, h.UpdateOrganization, middleware.OrganizationOwnerMiddleware(h.orgService))
}

func (h *OrganizationHandler) CreateOrganization(c echo.Context) error {
	userClaims, ok := c.Get("userClaims").(*auth.Claims)
	if !ok {
		return newHTTPError(http.StatusBadRequest, constants.InvalidRequestDetails)
	}

	orgReq := new(struct {
		Name          string `json:"name" validate:"required"`
		Industry      string `json:"industry"`
		GSTRegistered bool   `json:"gst_registered"`
	})

	if err := c.Bind(orgReq); err != nil {
		return newHTTPError(http.StatusBadRequest, constants.InvalidRequestDetails)
	}

	org := &model.Organization{
		Name:          orgReq.Name,
		Industry:      &orgReq.Industry,
		GSTRegistered: &orgReq.GSTRegistered,
		OwnerID:       userClaims.UserID,
	}

	if err := h.orgService.CreateOrganization(org); err != nil {
		return newHTTPError(http.StatusInternalServerError, constants.FailedCreateOrganization)
	}

	return c.JSON(http.StatusOK, response.GenericResponse{
		Message: constants.OrganizationCreatedSuccessfully,
		Data:    mapper.ToOrganizationResponse(*org),
	})
}

func (h *OrganizationHandler) GetOrganization(c echo.Context) error {
	id, err := parseID(c.Param("id"))
	if err != nil {
		return newHTTPError(http.StatusBadRequest, constants.InvalidOrganizationID)
	}

	org, err := h.orgService.GetOrganizationByID(id)
	if err != nil {
		return newHTTPError(http.StatusNotFound, constants.OrganizationNotFound)
	}

	return c.JSON(http.StatusOK, response.GenericResponse{
		Message: constants.OrganizationsRetrieved,
		Data:    org,
	})
}

func (h *OrganizationHandler) UpdateOrganization(c echo.Context) error {
	id, err := parseID(c.Param("id"))
	if err != nil {
		return newHTTPError(http.StatusBadRequest, constants.InvalidOrganizationID)
	}

	orgReq := new(struct {
		Name          string `json:"name" validate:"required"`
		Industry      string `json:"industry"`
		GSTRegistered bool   `json:"gst_registered"`
	})

	if err := c.Bind(orgReq); err != nil {
		return newHTTPError(http.StatusBadRequest, constants.InvalidRequestDetails)
	}

	org := &model.Organization{
		Model:         gorm.Model{ID: id},
		Name:          orgReq.Name,
		Industry:      &orgReq.Industry,
		GSTRegistered: &orgReq.GSTRegistered,
	}

	if err := h.orgService.UpdateOrganization(org); err != nil {
		return newHTTPError(http.StatusInternalServerError, constants.FailedUpdateOrganization)
	}

	return c.JSON(http.StatusOK, response.GenericResponse{
		Message: constants.OrganizationUpdatedSuccessfully,
		Data:    org,
	})
}

func parseID(idParam string) (uint, error) {
	id, err := strconv.ParseUint(idParam, 10, 32)
	return uint(id), err
}

func newHTTPError(statusCode int, message string) *echo.HTTPError {
	return echo.NewHTTPError(statusCode, middleware.ErrorResponse{
		Code:    statusCode,
		Message: message,
	})
}
