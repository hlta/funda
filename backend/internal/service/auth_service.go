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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		s.log.WithField("action", "hashing password").Error(err.Error())
		return err
	}
	user.Password = string(hashedPassword)
	if err := s.userRepo.Create(user); err != nil {
		if errors.Is(err, model.ErrEmailExists) {
			return errors.New("email already exists")
		}
		s.log.WithField("action", "creating user").Error(err.Error())
		return errors.New("signup failure")
	}
	s.log.WithField("action", "user signed up").Info("User successfully registered")
	return nil
}

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

func (s *AuthService) VerifyToken(tokenString string) (*jwt.Token, error) {
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

	return token, nil
}
