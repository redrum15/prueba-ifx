package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

func SendJSONError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	errorResponse := map[string]string{"error": message}
	json.NewEncoder(w).Encode(errorResponse)
}

func SendJSONSuccess(w http.ResponseWriter, response interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if response != nil {
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Error().Err(err).Msg("Error encoding response")
			http.Error(w, "Error generating response", http.StatusInternalServerError)
			return
		}
	}

}
