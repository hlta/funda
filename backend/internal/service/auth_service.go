package service

import (
	"errors"
	"funda/internal/auth"
	"funda/internal/logger"
	"funda/internal/model"
	"time"

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

func (s *AuthService) Signup(user *model.User) error {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		s.log.WithField("action", "hashing password").WithError(err).Error("Failed to hash password")
		return err
	}

	// Update the user password to the hashed password
	user.Password = string(hashedPassword)

	// Create the user using the UserService
	if err := s.userService.CreateUser(user); err != nil {
		s.log.WithField("action", "creating user").WithError(err).Error("Failed to create user")
		return err
	}

	// Create a default organization for the user
	org := &model.Organization{Name: "Default Organization", OwnerID: user.ID}
	if err := s.orgRepo.Create(org); err != nil {
		s.log.WithField("action", "creating organization").WithError(err).Error("Failed to create organization")
		return err
	}

	// Assign the user a default role (e.g., "Admin") in the organization
	role, err := s.roleRepo.RetrieveByName("Admin")
	if err != nil {
		s.log.WithField("action", "retrieving role").WithError(err).Error("Failed to retrieve default role")
		return err
	}

	userOrg := &model.UserOrganization{UserID: user.ID, OrganizationID: org.ID, RoleID: role.ID}
	if err := s.userOrgRepo.AddUserToOrganization(userOrg); err != nil {
		s.log.WithField("action", "assigning user to organization").WithError(err).Error("Failed to assign user to organization")
		return err
	}

	// Log success
	s.log.WithField("action", "user signed up").WithField("userID", user.ID).Info("User successfully registered")
	return nil
}

func (s *AuthService) Login(email, password string) (*model.User, error) {
	user, err := s.userService.GetUserByEmail(email)
	if err != nil {
		s.log.WithField("action", "retrieving user").Error(err.Error())
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		s.log.WithField("action", "password verification").Error("Invalid credentials")
		return nil, err
	}

	token, err := s.GenerateToken(user, 0)
	if err != nil {
		s.log.WithField("action", "generating token").Error(err.Error())
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

	if err != nil {
		return nil, nil, nil, err
	}

	if !token.Valid {
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

	return user, claims.Roles, claims.Permissions, nil
}

func (s *AuthService) GetUserOrganizations(userID uint) ([]model.UserOrganization, error) {
	return s.userOrgRepo.GetUserOrganizations(userID)
}

func (s *AuthService) GenerateToken(user *model.User, orgID uint) (string, error) {
	var roles []string
	var permissions []string

	if orgID != 0 {
		userOrgs, err := s.userOrgRepo.GetUserOrganizations(user.ID)
		if err != nil {
			return "", err
		}

		for _, userOrg := range userOrgs {
			if userOrg.OrganizationID == orgID {
				role, err := s.roleRepo.RetrieveByID(userOrg.RoleID)
				if err != nil {
					return "", err
				}
				roles = append(roles, role.Name)
				for _, perm := range role.Permissions {
					permissions = append(permissions, perm.Name)
				}
			}
		}
	}

	claims := auth.Claims{
		UserID:      user.ID,
		Email:       user.Email,
		Roles:       roles,
		Permissions: permissions,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(auth.GetJWTKey()))
}

func (s *AuthService) SwitchOrganization(user *model.User, orgID uint) (string, []string, []string, error) {
	newToken, err := s.GenerateToken(user, orgID)
	if err != nil {
		return "", nil, nil, err
	}

	token, err := jwt.ParseWithClaims(newToken, &auth.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return auth.GetJWTKey(), nil
	})

	if err != nil {
		return "", nil, nil, err
	}

	claims, ok := token.Claims.(*auth.Claims)
	if !ok {
		return "", nil, nil, errors.New("invalid claims")
	}

	return newToken, claims.Roles, claims.Permissions, nil
}
