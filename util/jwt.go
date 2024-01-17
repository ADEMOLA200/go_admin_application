package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const secretKey = "secret"

func GenerateJwt(issuer string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: issuer,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	return claims.SignedString([]byte(secretKey))
}


func ParseJwt(cookie string) (string, error) {
	token, err := jwt.ParseWithClaims(cookie, jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil || !token.Valid {
		return "", err
	}

	claims := token.Claims.(*jwt.StandardClaims)	

	return claims.Issuer, nil
}