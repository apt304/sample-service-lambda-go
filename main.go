package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

// Universal Main function called by
// environment specific main functions
// See main_standalone.go or main_aws_lambda.go
func Main() http.Handler {
	fmt.Println("Shared bootstrapping code goes here")

	// Load configs

	// Setup DB connection

	// Setup services

	// Setup http handlers
	r := chi.NewRouter()

	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	// Return any http handler, use your favorite mux library.
	return r
}
