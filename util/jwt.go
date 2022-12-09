package util

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"os"
)



func GenerateJwt(issuer string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: issuer, //strconv to convert integer to an int
		ExpiresAt: time.Now().Add(time.Hour * 48).Unix(), // expires in two days
	})

	return claims.SignedString([]byte(os.Getenv("SECRET_KEY"))) //secret key we want to use
}

func ParseJwt(cookie string) (string, error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil //to decode it, we need to encode the secret
	})

	if err != nil || !token.Valid {
		return "", err
	}

	claims := token.Claims.(*jwt.StandardClaims)

	return claims.Issuer, nil
}