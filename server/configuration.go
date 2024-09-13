package server

import (
	"net/http"
	"time"
)

// Create && return http.Server instance
//
// With Specific configuration,including ::
//
// Protection delays, CORS handler to manage cross-origin requests
func configurationServer(w http.Handler) *http.Server {
	return &http.Server{
		Addr:              "127.0.0.1:8080",
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
		MaxHeaderBytes:    1 << 20,
		Handler:           enableCors(w),
	}
}

// Checks the ORIGIN requests, and set appropriate CORS
func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Limit requests by this address => frontend
		path := r.Header.Get("Origin")
		//Imagine addr is
		if path == "http://localhost:3000" {
			// Set appropriate response
			w.Header().Set("Access-Control-Allow-Origin", path)
		}
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
