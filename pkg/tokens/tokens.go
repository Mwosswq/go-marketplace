package tokens

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type MyCustomClaims struct {
	UserID int `json:"id"`
	jwt.RegisteredClaims
}

func SignAccessToken(id int) (string, error) {
	return signToken(id, 15*time.Minute)
}

func SignRefreshToken(id int) (string, error) {
	return signToken(id, 30*24*time.Hour)
}

func signToken(id int, ttl time.Duration) (string, error) {
	sk := os.Getenv("SIGNING_KEY")
	if sk == "" {
		return "", errors.New("SIGNING_KEY is not set")
	}

	claims := MyCustomClaims{
		UserID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "go-marketplace",
			Subject:   "user_token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(sk))
}
