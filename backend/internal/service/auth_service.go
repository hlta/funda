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
	"gorm.io/gorm"
)

type AuthService struct {
	userService *UserService
	orgService  *OrganizationService
	log         logger.Logger
	db          *gorm.DB
}

// NewAuthService creates a new instance of AuthService.
func NewAuthService(userService *UserService, orgService *OrganizationService, log logger.Logger, db *gorm.DB) *AuthService {
	return &AuthService{
		userService: userService,
		orgService:  orgService,
		log:         log,
		db:          db,
	}
}

func (s *AuthService) Signup(user *model.User, orgName string) error {
	// Start a transaction
	tx := s.db.Begin()
	if tx.Error != nil {
		utils.LogError(s.log, "starting transaction", tx.Error)
		return tx.Error
	}

	if err := s.hashPassword(user); err != nil {
		tx.Rollback()
		return err
	}

	if err := s.userService.CreateUserWithTx(tx, user); err != nil {
		tx.Rollback()
		utils.LogError(s.log, "creating user", err)
		return err
	}

	if err := s.createOrganizationWithTx(tx, user, orgName); err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		utils.LogError(s.log, "committing transaction", err)
		return err
	}

	utils.LogSuccess(s.log, "user signed up", "User successfully registered", user.ID)
	return nil
}

func (s *AuthService) Login(email, password string) (*response.UserResponse, *string, error) {
	user, err := s.userService.GetUserByEmail(email)
	if err != nil {
		utils.LogError(s.log, "retrieving user", err)
		return nil, nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		utils.LogError(s.log, "password verification", err)
		return nil, nil, err
	}

	token, err := auth.GenerateToken(user, user.DefaultOrganizationID)
	if err != nil {
		utils.LogError(s.log, "generating token", err)
		return nil, nil, err
	}

	userResp := mapper.ToUserResponse(*user)
	utils.LogSuccess(s.log, "user logged in", "Token successfully generated", user.ID)
	return &userResp, &token, nil
}

func (s *AuthService) VerifyToken(tokenString string) (*response.UserResponse, *uint, error) {
	claims, err := auth.ValidateToken(tokenString)
	if err != nil {
		return nil, nil, errors.New("invalid claims")
	}

	user, err := s.userService.GetUserByID(claims.UserID)
	if err != nil {
		return nil, nil, err
	}

	userResp := mapper.ToUserResponse(*user)
	return &userResp, &claims.OrgID, nil
}

func (s *AuthService) GetUserOrganizations(userID uint) ([]response.OrganizationResponse, error) {
	return s.orgService.GetUserOrganizations(userID)
}

func (s *AuthService) SwitchOrganization(userID, orgID uint) (string, error) {
	user, err := s.userService.GetUserByID(userID)
	if err != nil {
		return "", err
	}
	return auth.GenerateToken(user, orgID)
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

func (s *AuthService) createOrganizationWithTx(tx *gorm.DB, user *model.User, orgName string) error {
	org := &model.Organization{Name: orgName, OwnerID: user.ID}
	if err := s.orgService.CreateOrganizationWithTx(tx, org); err != nil {
		utils.LogError(s.log, "creating user organization", err)
		return err
	}

	user.DefaultOrganizationID = org.ID
	if err := s.userService.UpdateUserWithTx(tx, user); err != nil {
		utils.LogError(s.log, "updating user", err)
		return err
	}

	return nil
}
