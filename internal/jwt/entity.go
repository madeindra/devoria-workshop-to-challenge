package jwt

import "github.com/dgrijalva/jwt-go"

type AccountContext string

const EmailContex AccountContext = "email"

type AccountStandardJWTClaims struct {
	jwt.StandardClaims
	Email string `json:"email"`
}
