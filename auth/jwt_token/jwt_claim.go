package jwttoken

import "github.com/dgrijalva/jwt-go"

type JWTClaims struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}
