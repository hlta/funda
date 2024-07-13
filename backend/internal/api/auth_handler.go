package api

import (
	"net/http"
	"time"

	"funda/internal/constants"
	"funda/internal/mapper"
	"funda/internal/middleware"
	"funda/internal/model"
	"funda/internal/response"
	"funda/internal/service"
	"funda/internal/utils"

	"github.com/casbin/casbin/v2"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authService *service.AuthService
	enforcer    *casbin.Enforcer
}

func NewAuthHandler(authService *service.AuthService, enforcer *casbin.Enforcer) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		enforcer:    enforcer,
	}
}

func (h *AuthHandler) Register(e *echo.Echo) {
	e.POST(constants.SignupRoute, h.Signup)
	e.POST(constants.LoginRoute, h.Login)
	e.POST(constants.LogoutRoute, h.Logout)
	e.GET(constants.CheckAuthRoute, h.CheckAuth)
	e.GET(constants.GetUserOrganizationsRoute, h.GetUserOrganizations)
	e.POST(constants.SwitchOrgRoute, h.SwitchOrganization)
}

func (h *AuthHandler) Signup(c echo.Context) error {
	var signupReq struct {
		FirstName        string `json:"firstName"`
		LastName         string `json:"lastName"`
		Email            string `json:"email"`
		Password         string `json:"password"`
		OrganizationName string `json:"organizationName"`
	}
	if err := c.Bind(&signupReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, middleware.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: constants.InvalidRequestDetails,
		})
	}

	user := &model.User{
		FirstName: signupReq.FirstName,
		LastName:  signupReq.LastName,
		Email:     signupReq.Email,
		Password:  signupReq.Password,
	}

	if err := h.authService.Signup(user, signupReq.OrganizationName); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, middleware.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: constants.FailedCreateUserAndOrg,
		})
	}

	// Assign specified role
	if _, err := h.enforcer.AddGroupingPolicy(user.Email, constants.AdminRoleName); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, middleware.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: constants.FailedAssignRole,
		})
	}

	return c.JSON(http.StatusOK, response.GenericResponse{
		Message: constants.SignupSuccessful,
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
			Message: constants.InvalidRequestDetails,
		})
	}

	userResp, token, err := h.authService.Login(loginReq.Email, loginReq.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, middleware.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: constants.InvalidCredentials,
		})
	}

	roles, err := h.enforcer.GetRolesForUser(userResp.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, middleware.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: constants.FailedRetrieveRoles,
		})
	}

	permissions, err := h.enforcer.GetImplicitPermissionsForUser(userResp.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, middleware.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: constants.FailedRetrievePermissions,
		})
	}

	utils.SetCookie(c, *token, 24*time.Hour)

	return c.JSON(http.StatusOK, response.GenericResponse{
		Message: constants.LoginSuccessful,
		Data:    mapper.ToAuthResponse(*userResp, *token, roles, permissions),
	})
}

func (h *AuthHandler) Logout(c echo.Context) error {
	utils.ClearCookie(c)
	return c.JSON(http.StatusOK, response.GenericResponse{
		Message: constants.LogoutSuccessful,
	})
}

func (h *AuthHandler) CheckAuth(c echo.Context) error {
	cookie, err := c.Cookie(constants.TokenCookieName)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, middleware.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: constants.NotAuthenticated,
		})
	}

	token := cookie.Value
	userResp, err := h.authService.VerifyToken(token)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, middleware.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: constants.InvalidToken,
		})
	}
	utils.SetCookie(c, token, 24*time.Hour)

	roles, err := h.enforcer.GetRolesForUser(userResp.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, middleware.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: constants.FailedRetrieveRoles,
		})
	}

	permissions, err := h.enforcer.GetImplicitPermissionsForUser(userResp.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, middleware.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: constants.FailedRetrievePermissions,
		})
	}

	return c.JSON(http.StatusOK, response.GenericResponse{
		Message: constants.Authenticated,
		Data:    mapper.ToAuthResponse(*userResp, token, roles, permissions),
	})
}

func (h *AuthHandler) GetUserOrganizations(c echo.Context) error {
	cookie, err := c.Cookie(constants.TokenCookieName)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, middleware.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: constants.NotAuthenticated,
		})
	}

	token := cookie.Value
	userResp, err := h.authService.VerifyToken(token)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, middleware.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: constants.InvalidToken,
		})
	}

	orgs, err := h.authService.GetUserOrganizations(userResp.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, middleware.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: constants.FailedRetrieveOrganizations,
		})
	}

	return c.JSON(http.StatusOK, response.GenericResponse{
		Message: constants.OrganizationsRetrieved,
		Data:    orgs,
	})
}

func (h *AuthHandler) SwitchOrganization(c echo.Context) error {
	cookie, err := c.Cookie(constants.TokenCookieName)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, middleware.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: constants.NotAuthenticated,
		})
	}

	token := cookie.Value
	userResp, err := h.authService.VerifyToken(token)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, middleware.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: constants.InvalidToken,
		})
	}

	var switchOrgReq struct {
		OrgID uint `json:"org_id"`
	}
	if err := c.Bind(&switchOrgReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, middleware.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: constants.InvalidRequestDetails,
		})
	}

	// Generate a new token with the new organization context
	newToken, err := h.authService.SwitchOrganization(userResp.ID, switchOrgReq.OrgID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, middleware.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: constants.FailedGenerateToken,
		})
	}

	utils.SetCookie(c, newToken, 24*time.Hour)

	return c.JSON(http.StatusOK, response.GenericResponse{
		Message: constants.OrganizationSwitched,
		Data:    newToken,
	})
}
