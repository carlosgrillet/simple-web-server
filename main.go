package main

import (
	"fmt"
	"net/http"
)

func root(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "ok\n")
}

func main() {
	http.HandleFunc("/", root)
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		fmt.Printf("Error ocurrend in the server: %s", err)
	}
}
