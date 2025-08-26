package usecase

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)
type Claims struct {
    UserID   uint   `json:"user_id"`
    Username string `json:"username"`
    jwt.RegisteredClaims
}

func ValidateToken(authHeader string) (*Claims, error) {
    tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
    claims := &Claims{}

    token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })
    if err != nil {
        return nil, err
    }
    if !token.Valid {
        return nil, errors.New("invalid token")
    }
    if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
        return nil, errors.New("token expired")
    }

    return claims, nil
}
