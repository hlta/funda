package service

import (
	"funda/internal/auth"
	"funda/internal/model"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo model.UserRepository
}

func NewAuthService(userRepo model.UserRepository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (s *AuthService) Signup(user *model.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return s.userRepo.Create(user)
}

func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.userRepo.RetrieveByEmail(email)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}

	// Use the GenerateToken function from the auth package
	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
