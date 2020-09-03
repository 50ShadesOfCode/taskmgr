package models

import (
	"github.com/dgrijalva/jwt-go"
)

//Credentials :
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//Claims :
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
