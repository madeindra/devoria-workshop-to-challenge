package account

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AccountContext string

const EmailContex AccountContext = "email"

// properties of account
// json attributes will set the field name on json form
type Account struct {
	ID             int64      `json:"id"`
	Email          string     `json:"email"`
	Password       *string    `json:"password,omitempty"`
	FirstName      string     `json:"firstName"`
	LastName       string     `json:"lastName"`
	CreatedAt      time.Time  `json:"createdAt"`
	LastModifiedAt *time.Time `json:"lastModifiedAt"`
}

type AccountStandardJWTClaims struct {
	jwt.StandardClaims
	Email string `json:"email"`
}
