package main

import (
	"fmt"
	"net/http"
	"time"
)

// LoggingMiddleware is a middleware that logs the request details
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Call the next handler
		next.ServeHTTP(w, r)

		// Log the request details
		duration := time.Since(start)
		fmt.Printf("Request: %s %s took %v\n", r.Method, r.URL, duration)
	})
}

// HelloHandler is a simple handler that responds with a message
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	// Create a new router
	mux := http.NewServeMux()

	// Add routes
	mux.Handle("/", LoggingMiddleware(http.HandlerFunc(HelloHandler)))

	// Start the server
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", mux)
}
