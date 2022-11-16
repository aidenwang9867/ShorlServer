package app

import (
	"encoding/json"
	"net/http"
	"path"

	"github.com/aidenwang9867/ShorlServer/utils"
)

func GetResultsHandler(w http.ResponseWriter, r *http.Request) {
	longLink := r.URL.Query().Get("long_link")
	link := utils.Link{
		Long:  longLink,
		Short: path.Join("sho.rl/l", utils.EncodeLink(longLink)),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(link)
}
