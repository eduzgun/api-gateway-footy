package models

import "github.com/dgrijalva/jwt-go"

//It will include the standard claims defined by the JWT specification, as well as a field for the user's role within the application.
type Claims struct {
	Role string `json:"role"`
	jwt.StandardClaims
}
