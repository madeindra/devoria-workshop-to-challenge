package jwt

import (
	"context"
	"crypto/rsa"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

// error messages
var (
	ErrInvalidToken      error = fmt.Errorf("invalid token")
	ErrExpiredOrNotReady error = fmt.Errorf("token is either expired or not ready to use")
)

// JWT interface
type JSONWebToken interface {
	Sign(ctx context.Context, claims jwt.Claims) (string, error)
	Parse(ctx context.Context, tokenString string, claims jwt.Claims) (*jwt.Token, error)
}

type jsonWebToken struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func NewJSONWebToken(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) JSONWebToken {
	return &jsonWebToken{privateKey, publicKey}
}

func (j *jsonWebToken) Sign(ctx context.Context, claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(j.privateKey)
}

func (j *jsonWebToken) Parse(ctx context.Context, tokenString string, claims jwt.Claims) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, claims, j.keyFunc)
	if err = j.checkError(err); err != nil {
		return token, nil
	}

	if !token.Valid {
		return token, ErrInvalidToken
	}

	return token, nil
}

func (j *jsonWebToken) keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
		return nil, ErrInvalidToken
	}
	return j.publicKey, nil
}

func (j *jsonWebToken) checkError(err error) error {
	if err == nil {
		return err
	}

	vld, ok := err.(*jwt.ValidationError)

	if !ok {
		return ErrInvalidToken
	}

	if vld.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
		return ErrExpiredOrNotReady
	}

	return ErrInvalidToken
}
