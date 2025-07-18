package main

import (
	"fmt"
	"net/http"
)

var port = "8080"

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("{\"method\": \"%s\", \"path\": \"%s\", \"remoteAddr\": \"%s\"}\n", r.Method, r.URL.Path, r.RemoteAddr)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "ok\n")
}

func secure(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("{\"method\": \"%s\", \"path\": \"%s\", \"remoteAddr\": \"%s\"}\n", r.Method, r.URL.Path, r.RemoteAddr)
	w.WriteHeader(http.StatusUnauthorized)
	fmt.Fprintf(w, "Unauthorized\n")
}

func main() {
	fmt.Printf("Listening on port %s\n", port)
	http.HandleFunc("/", root)
	http.HandleFunc("/secure", secure)
	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), nil)
	if err != nil {
		fmt.Printf("Error ocurrend in the server: %s", err)
	}
}
