package service

import (
	"funda/internal/auth"
	"funda/internal/logger"
	"funda/internal/model"

	"golang.org/x/crypto/bcrypt"
)

// AuthService provides authentication functionalities.
type AuthService struct {
	userRepo model.UserRepository
	log      logger.Logger
}

// NewAuthService creates a new instance of AuthService with dependency injection.
func NewAuthService(userRepo model.UserRepository, log logger.Logger) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		log:      log,
	}
}

// Signup handles user registration with password hashing.
func (s *AuthService) Signup(user *model.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		s.log.WithField("action", "hashing password").Error(err.Error())
		return err
	}
	user.Password = string(hashedPassword)
	if err := s.userRepo.Create(user); err != nil {
		s.log.WithField("action", "creating user").Error(err.Error())
		return err
	}
	s.log.WithField("action", "user signed up").Info("User successfully registered")
	return nil
}

// Login handles user authentication and token generation.
func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.userRepo.RetrieveByEmail(email)
	if err != nil {
		s.log.WithField("action", "retrieving user").Error(err.Error())
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		s.log.WithField("action", "password verification").Error("Invalid credentials")
		return "", err
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		s.log.WithField("action", "generating token").Error(err.Error())
		return "", err
	}

	s.log.WithField("action", "user logged in").Info("Token successfully generated")
	return token, nil
}
