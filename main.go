package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getRootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Root endpoint hit")
	fmt.Fprintln(w, "Root endpoint reached")
}

func getTestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Test endpoint reached")
}

func getHeadersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Headers endpoint reached\n")

	for key, values := range r.Header {
		fmt.Fprintf(w, "%s: %s\n", key, values)
	}
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UID: %d\n", os.Getuid())
	fmt.Fprintf(w, "GID: %d\n", os.Getgid())
}

func getStartJobHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://run-as-user:8000/api/v1/jobs")

	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
		return
	}

	fmt.Fprintf(w, "Body:\n%s", string(body))
}

func redirectMiddleware(next http.Handler) http.Handler {
	// return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	if r.Header.Get("X-Forwarded-Proto") != "https" || r.TLS != nil {
	// 		target := "https://" + r.Host + r.URL.RequestURI()
	// 		http.Redirect(w, r, target, http.StatusMovedPermanently)

	// 		return
	// 	}

	// 	next.ServeHTTP(w, r)
	// })

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", getRootHandler)
	mux.HandleFunc("/test", getTestHandler)
	mux.HandleFunc("/headers", getHeadersHandler)
	mux.HandleFunc("/runasuser", getUserHandler)
	mux.HandleFunc("/startjob", getStartJobHandler)

	fmt.Println("Server starting on :1234")
	err := http.ListenAndServe(":1234", redirectMiddleware(mux))

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server was closed")
	} else if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
