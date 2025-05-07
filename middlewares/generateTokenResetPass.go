package middlewares

import (
	"time"
	"log"
	"github.com/golang-jwt/jwt/v5"
	config "github.com/RaihanMalay21/config-tb-berkah-jaya-development"
)

func GenerateResetToken(email string) (string, error) {

	expTime := time.Now().Add(5 * time.Minute)
	tokenBeforeSigned := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp": expTime.Unix(),
	})

	token, err := tokenBeforeSigned.SignedString(config.JWT_KEY)
	if err != nil {
		log.Println("Error cant signature token function Generate token for reset password:", err)
		return "", err
	}

	return token, nil
}