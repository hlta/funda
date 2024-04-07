package auth

import (
	"funda/configs"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Assuming you load the configuration at the application startup and pass it where needed
var jwtKey []byte

func SetupAuth(config configs.OAuthConfig) {
	jwtKey = []byte(config.JWTSecret)
}

type Claims struct {
	UserID uint
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
