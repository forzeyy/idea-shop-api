package auth

import (
	"time"

	"github.com/forzeyy/idea-shop-api/middleware"
	"github.com/golang-jwt/jwt/v5"
)

var AccessTokenLifetime time.Time = time.Now().Add(time.Minute * 10)
var RefreshTokenLifetime time.Time = time.Now().Add(14 * 24 * time.Hour)

func GenerateAccessToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     AccessTokenLifetime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(middleware.AccessSecret)
}

func GenerateRefreshToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     RefreshTokenLifetime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(middleware.RefreshSecret)
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return middleware.AccessSecret, nil
	})
}
