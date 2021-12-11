package middleware

import "net/http"

type RouteMiddleware interface {
	Verify(next http.HandlerFunc) http.HandlerFunc
}

type RouteMiddlewareBearer interface {
	VerifyBearer(next http.HandlerFunc) http.HandlerFunc
}
