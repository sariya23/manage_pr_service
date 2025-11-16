package middleware

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
)

type ctxKey string

const RequestIDKey ctxKey = "request_id"

func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := fmt.Sprintf("%08x", rand.Uint32())
		ctx := context.WithValue(r.Context(), RequestIDKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
