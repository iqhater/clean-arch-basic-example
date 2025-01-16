package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type RequestKey string

const RequestIDKey RequestKey = "request-id"

func RequestID(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		ctx := context.WithValue(req.Context(), RequestIDKey, uuid.New())
		next.ServeHTTP(w, req.WithContext(ctx))
	})
}
