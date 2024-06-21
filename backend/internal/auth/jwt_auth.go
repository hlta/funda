package auth

import (
	"funda/configs"
	"funda/internal/model"
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
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	OrgID  uint   `json:"org_id"`
	jwt.RegisteredClaims
}

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
	return token.SignedString([]byte(GetJWTKey()))
}
