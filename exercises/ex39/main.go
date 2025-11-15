package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// loggingMiddleware is a middleware function that logs details of each incoming request.
// It takes an http.Handler as an argument (the next handler in the chain) and
// returns a new http.Handler.
func loggingMiddleware(next http.Handler) http.Handler {
	// http.HandlerFunc is an adapter that allows us to use ordinary functions
	// as HTTP handlers.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Log the request details.
		log.Printf(
			"Incoming Request -> Method: %s | URI: %s | Remote Addr: %s",
			r.Method,
			r.RequestURI,
			r.RemoteAddr,
		)

		// Call the next handler in the chain. This is the crucial part of middleware.
		// The request is passed down to the actual handler (e.g., rootHandler or helloHandler).
		next.ServeHTTP(w, r)

		// After the handler has finished, log the duration.
		log.Printf("Request completed in %v", time.Since(start))
	})
}

// rootHandler handles requests to the "/" path.
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the homepage!")
}

// helloHandler handles requests to the "/hello" path.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Gopher!")
}

func main() {
	// Create the handlers as http.HandlerFuncs.
	root := http.HandlerFunc(rootHandler)
	hello := http.HandlerFunc(helloHandler)

	// Register the handlers with the loggingMiddleware.
	// The middleware wraps the actual handler. When a request comes in,
	// the middleware's logic is executed first, and then it calls the
	// original handler's ServeHTTP method.
	http.Handle("/", loggingMiddleware(root))
	http.Handle("/hello", loggingMiddleware(hello))

	fmt.Println("Server is starting on port 8080...")
	fmt.Println("Try accessing http://localhost:8080/ or http://localhost:8080/hello")

	// Start the server.
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
