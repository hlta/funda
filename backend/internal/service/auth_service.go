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
	userRepo model.UserRepository
	log      logger.Logger
}

func NewAuthService(userRepo model.UserRepository, log logger.Logger) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		log:      log,
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

	// Create the user in the repository
	if err := s.userRepo.Create(user); err != nil {
		s.log.WithField("action", "creating user").WithError(err).Error("Failed to create user")
		return err
	}

	// Log success
	s.log.WithField("action", "user signed up").WithField("userID", user.ID).Info("User successfully registered")
	return nil
}

func (s *AuthService) Login(email, password string) (*model.User, error) {
	user, err := s.userRepo.RetrieveByEmail(email)
	if err != nil {
		s.log.WithField("action", "retrieving user").Error(err.Error())
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		s.log.WithField("action", "password verification").Error("Invalid credentials")
		return nil, err
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		s.log.WithField("action", "generating token").Error(err.Error())
		return nil, err
	}

	user.Token = token
	s.log.WithField("action", "user logged in").Info("Token successfully generated")
	return user, nil
}

func (s *AuthService) VerifyToken(tokenString string) (*model.User, error) {
	token, err := jwt.ParseWithClaims(tokenString, &auth.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return auth.GetJWTKey(), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*auth.Claims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	user, err := s.userRepo.RetrieveByID(claims.UserID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
