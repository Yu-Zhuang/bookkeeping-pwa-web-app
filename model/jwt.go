package model

import (
	jwt "gopkg.in/dgrijalva/jwt-go.v3"
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
