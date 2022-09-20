package jwt_auth

import (
	jwt "github.com/dgrijalva/jwt-go/v4"
)

type Claims struct {
	Username string `json:"string"`
	jwt.StandardClaims
}
