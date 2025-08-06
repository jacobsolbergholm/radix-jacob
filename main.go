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

func main() {
	http.HandleFunc("/", getRootHandler)
	http.HandleFunc("/test", getTestHandler)

	err := http.ListenAndServe(":1234", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server was closed")
	} else if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
