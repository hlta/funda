package auth

import (
	"errors"
	"funda/configs"
	"funda/internal/model"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Claims struct defined to match the token's payload structure
type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	OrgID  uint   `json:"org_id"`
	jwt.RegisteredClaims
}

// SetupAuth initializes the jwtKey with the given config.
var jwtKey []byte

// SetupAuth initializes the jwtKey with the given config.
func SetupAuth(config configs.OAuthConfig) {
	jwtKey = []byte(config.JWTSecret)
}

// GenerateToken creates a JWT for a user.
func GenerateToken(user *model.User, orgID uint) (string, error) {
	claims := Claims{
		UserID: user.ID,
		Email:  user.Email,
		OrgID:  orgID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// ValidateToken verifies the token and returns the claims.
func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
