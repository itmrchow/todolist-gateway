package utils

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

func ResponseWriter(r *http.Request, w http.ResponseWriter, status int, resp any) {
	jsonData, _ := json.Marshal(resp)

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(jsonData)
	if err != nil {
		log.Error().Err(err).Msg("failed to write response")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
