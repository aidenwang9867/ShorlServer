package main

import (
	"fmt"
	"github.com/aidenwang9867/ShorlServer/app"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Starting ShorlServer on port %s...\n", port)
	r := mux.NewRouter().StrictSlash(true)

	// API Index.
	r.HandleFunc("/", app.Index)

	// Generate the short link given the original.
	// GET query to generate the short link for an input long link.
	r.HandleFunc("/generate", app.GetResultsHandler).Methods(http.MethodGet)
	// POST query to generate the short link for an input long link, bulk access is supported by POST.
	r.HandleFunc("/generate", app.PostResultsHandler).Methods(http.MethodPost)

	// Redirect short to original.
	r.HandleFunc("/r/{short_link}", app.GetResultsHandler).Methods(http.MethodGet)

	http.Handle("/", r)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal(err)
	}

}
