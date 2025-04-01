package auth

import (
	"encoding/json"
	"net/http"

	"github.com/redrum15/prueba/src/db/querys"
	"github.com/redrum15/prueba/src/middlewares"
	"github.com/redrum15/prueba/src/utils"
	"github.com/rs/zerolog/log"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var login utils.LoginBody

	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := login.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := querys.GetUserFromEmailNPassword(login.Email, login.Password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	token, err := middlewares.GenerateJWT(result.ID.String(), result.Type)
	if err != nil {
		http.Error(w, "Error generating authentication token", http.StatusInternalServerError)
		return
	}

	response := utils.LoginResponse{
		Token:  token,
		UserID: result.ID.String(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Error().Err(err).Msg("Error encoding response")
		http.Error(w, "Error generating response", http.StatusInternalServerError)
		return
	}
}
