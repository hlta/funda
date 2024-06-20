package auth

import (
	"funda/configs"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey []byte

// SetupAuth initializes the jwtKey with the given config.
func SetupAuth(config configs.OAuthConfig) {
	jwtKey = []byte(config.JWTSecret)
}

// GetJWTKey returns the jwtKey.
func GetJWTKey() []byte {
	return jwtKey
}

type Claims struct {
	UserID      uint     `json:"user_id"`
	Email       string   `json:"email"`
	Roles       []string `json:"roles"`
	Permissions []string `json:"permissions"`
	jwt.RegisteredClaims
}

// GenerateToken creates a JWT for a given user ID.
func GenerateToken(userID uint) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, err
}
