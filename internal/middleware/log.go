package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		next.ServeHTTP(ww, r)

		duration := time.Since(start)
		log.Printf(
			"%-6s %-25s %-3d %-10s %-15s",
			r.Method,
			r.URL.Path,
			ww.Status(),
			duration.String(),
			r.RemoteAddr,
		)
	})
}
