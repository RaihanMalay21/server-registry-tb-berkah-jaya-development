package config

import "github.com/golang-jwt/jwt/v5"

var JWT_KEY = []byte("igvif4lcnde0942ufo02vow884050t05hgbejwdln")

type JWTClaim struct {
	UserName string
	jwt.RegisteredClaims
}