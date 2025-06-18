package midleware

import (
	"context"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		xRequestID := r.Header.Get("X-Request-ID")
		if xRequestID == "" {
			xRequestID = uuid.NewV4().String()
		}

		ctx := context.WithValue(r.Context(), "X-Request-ID", xRequestID)
		r = r.WithContext(ctx)

		defer func(startedAt time.Time) {
			// Log the duration of the request and response time.
			// L.L.WithRequestID(r.Context()).Info("Outgoing Response", L.Int64("Duration", time.Since(startedAt).Milliseconds()), L.String("URL", r.URL.String()))
		}(time.Now())

		// L.L.WithRequestID(r.Context()).Info("Incoming Request", L.String("URL", r.URL.String()), L.String("Method", r.Method), L.String("Host", r.Host), L.String("Proto", r.Proto))

		// Log the request
		next.ServeHTTP(w, r)
	})
}
