package util

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateToken 生成Token
func GenerateToken(userID uint) (string, error) {
	nowTime := time.Now()
	exp := nowTime.Add(3 * time.Hour)

	claims := Claims{
		userID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			Issuer:    "Todo-List",
		},
	}

	tokenClaims := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims)
	token, err := tokenClaims.SignedString([]byte("Todo-List"))

	return token, err
}

// ParseToken 解析Token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("Todo-List"), nil
	})

	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims)
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
