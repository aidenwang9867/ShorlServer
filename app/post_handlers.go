package app

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func PostResultsHandler(w http.ResponseWriter, r *http.Request) {
	_, err := io.ReadAll(r.Body)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		_, err := fmt.Fprintf(w, "error reading request body")
		if err != nil {
			log.Printf("error during Write: %v", err)
		}
		return
	}

}
