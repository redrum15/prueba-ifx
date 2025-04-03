package auth

import (
	"encoding/json"
	"net/http"

	"github.com/redrum15/prueba/src/db/querys"
	"github.com/redrum15/prueba/src/handlers"
	"github.com/redrum15/prueba/src/middlewares"
	"github.com/redrum15/prueba/src/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var login utils.LoginBody

	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		handlers.SendJSONError(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := login.Validate(); err != nil {
		handlers.SendJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := querys.GetUserFromEmailNPassword(login.Email, login.Password)

	if err != nil {
		handlers.SendJSONError(w, err.Error(), http.StatusNotFound)
		return
	}

	token, err := middlewares.GenerateJWT(result.ID.String(), result.Type)
	if err != nil {
		handlers.SendJSONError(w, "Error generating authentication token", http.StatusInternalServerError)
		return
	}

	response := utils.LoginResponse{
		Token:    token,
		UserID:   result.ID.String(),
		UserType: result.Type,
	}

	handlers.SendJSONSuccess(w, response, http.StatusOK)
}
