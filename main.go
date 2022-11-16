package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aidenwang9867/ShorlServer/app"
	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Starting ShorlServer on port %s...\n", port)
	router := mux.NewRouter().StrictSlash(true)

	// API Index.
	router.HandleFunc("/", app.Index)

	// Generate the short link given the original.
	// GET query to generate the short link for an input long link.
	router.HandleFunc("/generate", app.GenerateGetHandler).Methods(http.MethodGet)
	// POST query to generate the short link for an input long link, bulk access is supported by POST.
	router.HandleFunc("/generate", app.GeneratePostHandler).Methods(http.MethodPost)

	// Redirect short to original.
	router.HandleFunc("/r/{short_link}", app.RedirectGetHandler).Methods(http.MethodGet)

	http.Handle("/", router)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal(err)
	}

}
