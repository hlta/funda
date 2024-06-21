package service

import (
	"errors"
	"funda/internal/auth"
	"funda/internal/logger"
	"funda/internal/model"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userService *UserService
	orgRepo     model.OrganizationRepository
	roleRepo    model.RoleRepository
	userOrgRepo model.UserOrganizationRepository
	log         logger.Logger
}

func NewAuthService(userService *UserService, orgRepo model.OrganizationRepository, roleRepo model.RoleRepository, userOrgRepo model.UserOrganizationRepository, log logger.Logger) *AuthService {
	return &AuthService{
		userService: userService,
		orgRepo:     orgRepo,
		roleRepo:    roleRepo,
		userOrgRepo: userOrgRepo,
		log:         log,
	}
}

func (s *AuthService) Signup(user *model.User, orgName string) error {
	// Hash the password
	if err := s.hashPassword(user); err != nil {
		return err
	}

	// Create the user
	if err := s.userService.CreateUser(user); err != nil {
		s.logError("creating user", err)
		return err
	}

	// Create the organization provided by the user
	if err := s.createOrganization(user, orgName); err != nil {
		return err
	}

	// Log success
	s.log.WithField("action", "user signed up").WithField("userID", user.ID).Info("User successfully registered")
	return nil
}

func (s *AuthService) Login(email, password string) (*model.User, error) {
	user, err := s.userService.GetUserByEmail(email)
	if err != nil {
		s.logError("retrieving user", err)
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		s.log.WithField("action", "password verification").Error("Invalid credentials")
		return nil, err
	}

	// Load the default organization
	if err := s.userService.LoadDefaultOrganization(user); err != nil {
		s.logError("loading default organization", err)
		return nil, err
	}

	token, err := auth.GenerateToken(user, user.DefaultOrganizationID)
	if err != nil {
		s.logError("generating token", err)
		return nil, err
	}

	user.Token = token
	s.log.WithField("action", "user logged in").Info("Token successfully generated")
	return user, nil
}

func (s *AuthService) VerifyToken(tokenString string) (*model.User, []string, []string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &auth.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return auth.GetJWTKey(), nil
	})

	if err != nil || !token.Valid {
		return nil, nil, nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*auth.Claims)
	if !ok {
		return nil, nil, nil, errors.New("invalid claims")
	}

	user, err := s.userService.GetUserByID(claims.UserID)
	if err != nil {
		return nil, nil, nil, err
	}

	roles, permissions, err := s.GetRolesAndPermissions(claims.UserID, claims.OrgID)
	if err != nil {
		return nil, nil, nil, err
	}

	return user, roles, permissions, nil
}

func (s *AuthService) GetUserOrganizations(userID uint) ([]model.UserOrganization, error) {
	return s.userOrgRepo.GetUserOrganizations(userID)
}

func (s *AuthService) GetRolesAndPermissions(userID uint, orgID uint) ([]string, []string, error) {
	var roles []string
	var permissions []string

	// Query the specific UserOrganization
	userOrg, err := s.userOrgRepo.GetUserOrganization(userID, orgID)
	if err != nil {
		return roles, permissions, err
	}

	role, err := s.roleRepo.RetrieveByID(userOrg.RoleID)
	if err != nil {
		return roles, permissions, err
	}

	roles = append(roles, role.Name)
	for _, perm := range role.Permissions {
		permissions = append(permissions, perm.Name)
	}

	return roles, permissions, nil
}

func (s *AuthService) SwitchOrganization(user *model.User, orgID uint) (string, []string, []string, error) {
	newToken, err := auth.GenerateToken(user, orgID)
	if err != nil {
		return "", nil, nil, err
	}

	token, err := jwt.ParseWithClaims(newToken, &auth.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return auth.GetJWTKey(), nil
	})

	if err != nil || !token.Valid {
		return "", nil, nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*auth.Claims)
	if !ok {
		return "", nil, nil, errors.New("invalid claims")
	}
	roles, permissions, err := s.GetRolesAndPermissions(claims.UserID, claims.OrgID)
	if err != nil {
		return "", nil, nil, err
	}

	return newToken, roles, permissions, nil
}

// Helper Methods

func (s *AuthService) hashPassword(user *model.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		s.logError("hashing password", err)
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

func (s *AuthService) createOrganization(user *model.User, orgName string) error {
	org := &model.Organization{Name: orgName, OwnerID: user.ID}
	if err := s.orgRepo.Create(org); err != nil {
		s.logError("creating organization", err)
		return err
	}

	role, err := s.roleRepo.RetrieveByName("Admin")
	if err != nil {
		s.logError("retrieving role", err)
		return err
	}

	userOrg := &model.UserOrganization{UserID: user.ID, OrganizationID: org.ID, RoleID: role.ID}
	if err := s.userOrgRepo.AddUserToOrganization(userOrg); err != nil {
		s.logError("assigning user to organization", err)
		return err
	}

	user.DefaultOrganizationID = org.ID
	if err := s.userService.UpdateUser(user); err != nil {
		s.logError("updating user", err)
		return err
	}

	return nil
}

func (s *AuthService) logError(action string, err error) {
	s.log.WithField("action", action).WithError(err).Error("Failed to " + action)
}
