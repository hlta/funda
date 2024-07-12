package service

import (
	"errors"
	"funda/internal/auth"
	"funda/internal/logger"
	"funda/internal/mapper"
	"funda/internal/model"
	"funda/internal/response"
	"funda/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userService *UserService
	orgService  *OrganizationService
	log         logger.Logger
}

// NewAuthService creates a new instance of AuthService.
func NewAuthService(userService *UserService, orgService *OrganizationService, log logger.Logger) *AuthService {
	return &AuthService{
		userService: userService,
		orgService:  orgService,
		log:         log,
	}
}

func (s *AuthService) Signup(user *model.User, orgName string) error {
	if err := s.hashPassword(user); err != nil {
		return err
	}

	if err := s.userService.CreateUser(user); err != nil {
		utils.LogError(s.log, "creating user", err)
		return err
	}

	if err := s.createOrganization(user, orgName); err != nil {
		return err
	}

	utils.LogSuccess(s.log, "user signed up", "User successfully registered", user.ID)
	return nil
}

func (s *AuthService) Login(email, password string) (*response.UserResponse, error) {
	user, err := s.userService.GetUserByEmail(email)
	if err != nil {
		utils.LogError(s.log, "retrieving user", err)
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		utils.LogError(s.log, "password verification", err)
		return nil, err
	}

	if err := s.userService.LoadDefaultOrganization(user); err != nil {
		utils.LogError(s.log, "loading default organization", err)
		return nil, err
	}

	token, err := auth.GenerateToken(user, user.DefaultOrganizationID)
	if err != nil {
		utils.LogError(s.log, "generating token", err)
		return nil, err
	}

	roles, permissions, err := s.GetRolesAndPermissions(user.ID, user.DefaultOrganizationID)
	if err != nil {
		return nil, err
	}

	userResp := mapper.ToUserResponse(*user, roles, permissions, user.DefaultOrganizationID, token)
	utils.LogSuccess(s.log, "user logged in", "Token successfully generated", user.ID)
	return &userResp, nil
}

func (s *AuthService) VerifyToken(tokenString string) (*response.UserResponse, error) {
	claims, err := auth.ValidateToken(tokenString)
	if err != nil {
		return nil, errors.New("invalid claims")
	}

	user, err := s.userService.GetUserByID(claims.UserID)
	if err != nil {
		return nil, err
	}

	roles, permissions, err := s.GetRolesAndPermissions(claims.UserID, claims.OrgID)
	if err != nil {
		return nil, err
	}

	userResp := mapper.ToUserResponse(*user, roles, permissions, claims.OrgID, tokenString)
	return &userResp, nil
}

func (s *AuthService) GetUserOrganizations(userID uint) ([]response.OrganizationResponse, error) {
	return s.orgService.GetUserOrganizations(userID)
}

func (s *AuthService) GetRolesAndPermissions(userID, orgID uint) ([]string, []string, error) {
	return nil, nil, nil
}

func (s *AuthService) SwitchOrganization(userID, orgID uint) (*response.SwitchOrganizationResponse, error) {
	user, err := s.userService.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	newToken, err := auth.GenerateToken(user, orgID)
	if err != nil {
		return nil, err
	}
	roles, permissions, err := s.GetRolesAndPermissions(userID, orgID)
	if err != nil {
		return nil, err
	}

	switchOrgResp := mapper.ToSwitchOrganizationResponse(newToken, roles, permissions)
	return &switchOrgResp, nil
}

// Helper Methods

func (s *AuthService) hashPassword(user *model.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.LogError(s.log, "hashing password", err)
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

func (s *AuthService) createOrganization(user *model.User, orgName string) error {
	org := &model.Organization{Name: orgName, OwnerID: user.ID}
	if err := s.orgService.CreateOrganization(org); err != nil {
		utils.LogError(s.log, "creating user organization", err)
		return err
	}

	user.DefaultOrganizationID = org.ID
	if err := s.userService.UpdateUser(user); err != nil {
		utils.LogError(s.log, "updating user", err)
		return err
	}

	return nil
}
