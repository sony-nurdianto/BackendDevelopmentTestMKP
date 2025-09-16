package encryption

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID int64, cardID string) (string, error) {
	secretKey := os.Getenv("JWTSECRET")
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(tokenStr string) (*jwt.Token, error) {
	secretKey := os.Getenv("JWTSECRET")
	return jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		return []byte(secretKey), nil
	})
}
