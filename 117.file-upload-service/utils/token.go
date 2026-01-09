package utils

import (
	"file-upload-service/models"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("supersecretkey")

type Claims struct {
	UserId string `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateToken(user models.User) (string, error) {
	claims := &Claims{
		UserId: strconv.FormatUint(uint64(user.ID), 10),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(25 * 7 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return token.Claims.(*Claims), nil
}
