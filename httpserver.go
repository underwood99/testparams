package main

import (
	"fmt"
	"net/http"

	digcloudlog "github.com/digautos-library/digCloudLogGo"
)

func Home(w http.ResponseWriter, r *http.Request) {
	digcloudlog.DCL_Info("got http request")
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
