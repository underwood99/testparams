package main

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage")
}

func httpStart() error {
	http.HandleFunc("/", Home)

	// Spinning up the server.
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return err
	}

	return nil
}
