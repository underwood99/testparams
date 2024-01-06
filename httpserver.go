package main

import (
	"fmt"
	"net/http"

	digcloudlog "github.com/digautos-library/digCloudLogGo"
)

func Home(w http.ResponseWriter, r *http.Request) {
	digcloudlog.DCL_Info("home got http request")
	fmt.Fprintf(w, "Homepage")
}
func Home1(w http.ResponseWriter, r *http.Request) {
	digcloudlog.DCL_Info("home1 got http request")
	fmt.Fprintf(w, "Homepage1")
}

func httpStart() error {
	http.HandleFunc("/", Home)
	http.HandleFunc("/vlhu/testservice/go-tcp-greeter-5c6/v1.0", Home1)

	// Spinning up the server.
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return err
	}

	return nil
}
