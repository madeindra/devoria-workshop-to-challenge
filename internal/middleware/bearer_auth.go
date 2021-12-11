package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/madeindra/devoria-workshop-to-challenge/internal/jwt"
)

type BearerAuth struct {
	jsonWebToken jwt.JSONWebToken
}

func NewBearerAuth(jsonWebToken jwt.JSONWebToken) RouteMiddlewareBearer {
	return &BearerAuth{jsonWebToken}
}

// Verify will verify the request to ensure it comes with an authorized bearer auth token.
func (be *BearerAuth) VerifyBearer(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		auth := r.Header.Get("Authorization")
		token := strings.Split(auth, " ")[1]
		res, err := be.jsonWebToken.Parse(ctx, token, &jwt.AccountStandardJWTClaims{})
		if err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		claims := (res.Claims).(*jwt.AccountStandardJWTClaims)
		ctx = context.WithValue(ctx, jwt.EmailContex, claims.Email)

		next(w, r.WithContext(ctx))
	})
}
