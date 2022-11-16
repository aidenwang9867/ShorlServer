package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	endpts := struct {
		GenerateShortLink  string `json:"generate_short_link"`
		RedirectToOriginal string `json:"redirect_to_original"`
	}{
		GenerateShortLink:  "/generate",
		RedirectToOriginal: "/r",
	}
	endptsBytes, err := json.MarshalIndent(endpts, "", " ")
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if _, err := fmt.Fprint(w, string(endptsBytes)); err != nil {
		log.Fatal(err)
	}
}
