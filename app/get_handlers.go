package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/aidenwang9867/ShorlServer/utils"
)

func GenerateGetHandler(w http.ResponseWriter, r *http.Request) {
	longLink := r.URL.Query().Get("long_link")
	if longLink == "" {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		_, err := fmt.Fprint(w, "long link not specified")
		if err != nil {
			log.Printf("error during Write: %v", err)
		}
		return
	}
	link := utils.Link{
		Long:  longLink,
		Short: path.Join("sho.rl/r", utils.EncodeLink(longLink)),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(link)
}

func RedirectGetHandler(w http.ResponseWriter, r *http.Request) {
	// shortLink := r.URL.Query().Get("short_link")
	// w.Header().Set("Location", shortLink)
	// w.WriteHeader(http.StatusTemporaryRedirect)
	http.Redirect(w, r, "http://www.google.com", http.StatusTemporaryRedirect)

}
