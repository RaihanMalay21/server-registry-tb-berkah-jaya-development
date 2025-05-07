package middlewares

import (
	"github.com/golang-jwt/jwt/v5"
	config "github.com/RaihanMalay21/config-tb-berkah-jaya-development"
)

func VerifyResetToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return config.JWT_KEY, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["email"].(string), nil
	}
	return "", err
}