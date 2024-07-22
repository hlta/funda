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
		return h.respondWithError(c, http.StatusBadRequest, constants.InvalidRequestDetails, err)
	}

	user := &model.User{
		FirstName: signupReq.FirstName,
		LastName:  signupReq.LastName,
		Email:     signupReq.Email,
		Password:  signupReq.Password,
	}

	if err := h.authService.Signup(user, signupReq.OrganizationName); err != nil {
		switch err {
		case model.ErrEmailExists, model.ErrOrgExists:
			return err
		default:
			return h.respondWithError(c, http.StatusInternalServerError, constants.FailedCreateUserAndOrg, err)
		}
	}

	// Add predefined roles and permissions for the new organization
	if err := utils.AddPredefinedRolesAndPermissions(h.enforcer, user.DefaultOrganizationID); err != nil {
		return h.respondWithError(c, http.StatusInternalServerError, constants.FailedToSetPermissions, err)
	}

	// Assign the admin role to the new user in their default organization
	if err := h.assignRole(user.ID, user.DefaultOrganizationID, constants.AdminRoleName); err != nil {
		return h.respondWithError(c, http.StatusInternalServerError, constants.FailedAssignRole, err)
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
		return h.respondWithError(c, http.StatusBadRequest, constants.InvalidRequestDetails, err)
	}

	userResp, token, err := h.authService.Login(loginReq.Email, loginReq.Password)
	if err != nil {
		return h.respondWithError(c, http.StatusUnauthorized, constants.InvalidCredentials, err)
	}

	roles, permissions, err := h.getUserRolesAndPermissions(userResp.ID, userResp.DefaultOrganizationID)
	if err != nil {
		return h.respondWithError(c, http.StatusInternalServerError, constants.FailedRetrieveRoles, err)
	}

	utils.SetCookie(c, *token, 24*time.Hour)

	return c.JSON(http.StatusOK, response.GenericResponse{
		Message: constants.LoginSuccessful,
		Data:    mapper.ToAuthResponse(*userResp, userResp.DefaultOrganizationID, *token, roles, permissions),
	})
}

func (h *AuthHandler) getUserRolesAndPermissions(userID uint, orgID uint) ([]string, [][]string, error) {
	roles := h.enforcer.GetRolesForUserInDomain(utils.UintToString(userID), utils.UintToString(orgID))

	permissions, err := h.enforcer.GetImplicitPermissionsForUser(utils.UintToString(userID), utils.UintToString(orgID))
	if err != nil {
		return nil, nil, err
	}

	return roles, permissions, nil
}

func (h *AuthHandler) assignRole(userID uint, orgID uint, roleName string) error {
	_, err := h.enforcer.AddGroupingPolicy(utils.UintToString(userID), roleName, utils.UintToString(orgID))
	return err
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
		return h.respondWithError(c, http.StatusUnauthorized, constants.NotAuthenticated, err)
	}

	token := cookie.Value
	userResp, selectedOrg, err := h.authService.VerifyToken(token)
	if err != nil {
		return h.respondWithError(c, http.StatusUnauthorized, constants.InvalidToken, err)
	}
	utils.SetCookie(c, token, 24*time.Hour)

	roles, permissions, err := h.getUserRolesAndPermissions(userResp.ID, *selectedOrg)
	if err != nil {
		return h.respondWithError(c, http.StatusInternalServerError, constants.FailedRetrieveRoles, err)
	}

	return c.JSON(http.StatusOK, response.GenericResponse{
		Message: constants.Authenticated,
		Data:    mapper.ToAuthResponse(*userResp, *selectedOrg, token, roles, permissions),
	})
}

func (h *AuthHandler) GetUserOrganizations(c echo.Context) error {
	cookie, err := c.Cookie(constants.TokenCookieName)
	if err != nil {
		return h.respondWithError(c, http.StatusUnauthorized, constants.NotAuthenticated, err)
	}

	token := cookie.Value
	userResp, _, err := h.authService.VerifyToken(token)
	if err != nil {
		return h.respondWithError(c, http.StatusUnauthorized, constants.InvalidToken, err)
	}

	orgs, err := h.authService.GetUserOrganizations(userResp.ID)
	if err != nil {
		return h.respondWithError(c, http.StatusInternalServerError, constants.FailedRetrieveOrganizations, err)
	}

	return c.JSON(http.StatusOK, response.GenericResponse{
		Message: constants.OrganizationsRetrieved,
		Data:    orgs,
	})
}

func (h *AuthHandler) SwitchOrganization(c echo.Context) error {
	cookie, err := c.Cookie(constants.TokenCookieName)
	if err != nil {
		return h.respondWithError(c, http.StatusUnauthorized, constants.NotAuthenticated, err)
	}

	token := cookie.Value
	userResp, _, err := h.authService.VerifyToken(token)
	if err != nil {
		return h.respondWithError(c, http.StatusUnauthorized, constants.InvalidToken, err)
	}

	var switchOrgReq struct {
		OrgID uint `json:"org_id"`
	}
	if err := c.Bind(&switchOrgReq); err != nil {
		return h.respondWithError(c, http.StatusBadRequest, constants.InvalidRequestDetails, err)
	}

	// Generate a new token with the new organization context
	newToken, err := h.authService.SwitchOrganization(userResp.ID, switchOrgReq.OrgID)
	if err != nil {
		return h.respondWithError(c, http.StatusInternalServerError, constants.FailedGenerateToken, err)
	}

	utils.SetCookie(c, newToken, 24*time.Hour)

	roles, permissions, err := h.getUserRolesAndPermissions(userResp.ID, switchOrgReq.OrgID)
	if err != nil {
		return h.respondWithError(c, http.StatusInternalServerError, constants.FailedRetrieveRoles, err)
	}

	return c.JSON(http.StatusOK, response.GenericResponse{
		Message: constants.Authenticated,
		Data:    mapper.ToAuthResponse(*userResp, switchOrgReq.OrgID, newToken, roles, permissions),
	})

}

// Helper method to respond with error and log it
func (h *AuthHandler) respondWithError(c echo.Context, statusCode int, message string, err error) error {
	c.Logger().Error(err)
	return echo.NewHTTPError(statusCode, middleware.ErrorResponse{
		Code:    statusCode,
		Message: message,
	})
}
