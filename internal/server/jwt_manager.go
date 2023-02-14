package server

import (
	"fmt"
	"time"

	db "github.com/Jay-T/go-devops-advanced-diploma/db/sqlc"
	"github.com/dgrijalva/jwt-go"
)

type JWTManager struct {
	secretKey     string
	tokenDuration time.Duration
}

func NewJWTManager(secretKey string, tokenDuration time.Duration) *JWTManager {
	return &JWTManager{
		secretKey:     secretKey,
		tokenDuration: tokenDuration,
	}
}

type Claims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

func (manager *JWTManager) GeneratetToken(acc *db.Account) (string, time.Time, error) {
	expirationTime := time.Now().Add(manager.tokenDuration)

	claims := &Claims{
		Username: acc.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(manager.secretKey))

	if err != nil {
		return "", time.Time{}, fmt.Errorf("could not generate token for user sign in request")
	}

	return tokenString, expirationTime, nil
}

func (manager *JWTManager) Verify(accessToken string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}

			return []byte(manager.secretKey), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
