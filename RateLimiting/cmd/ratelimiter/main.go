package main

import (
	"fmt"
	"net/http"
	r "ratelimiter/internal/ratelimiter"
)

func main() {
	limiterStore := r.NewLimiterStore(5, 10) // 5 RPS, burst 10

	mux := http.NewServeMux()
	mux.Handle("/api", r.RateLimitMiddleware(limiterStore)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Request allowed")
	})))

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
