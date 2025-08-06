package main

import (
	"errors"
	"fmt"
	"net/http"
)

func getRootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Root endpoint reached")
}

func getTestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Test endpoint reached")
}

func redirectMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Forwarded-Proto") != "https" || r.TLS != nil {
			target := "https://" + r.Host + r.URL.RequestURI()
			http.Redirect(w, r, target, http.StatusMovedPermanently)

			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", getRootHandler)
	mux.HandleFunc("/test", getTestHandler)

	err := http.ListenAndServe(":1234", redirectMiddleware(mux))

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server was closed")
	} else if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
