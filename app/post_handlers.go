package app

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"path"

	"github.com/aidenwang9867/ShorlServer/utils"
)

func PostResultsHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		_, err := fmt.Fprintf(w, "error reading request body")
		if err != nil {
			log.Printf("error during Write: %v", err)
		}
		return
	}
	bulkLinks := []string{}
	err = json.Unmarshal(reqBody, &bulkLinks)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusBadRequest)
		_, err := fmt.Fprint(w, "error unmarshaling inputs")
		if err != nil {
			log.Printf("error during Write: %v", err)
		}
		return
	}
	ret := []utils.Link{}
	for _, l := range bulkLinks {
		link := utils.Link{
			Long:  l,
			Short: path.Join("sho.rl/l", utils.EncodeLink(l)),
		}
		ret = append(ret, link)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(ret)
}
