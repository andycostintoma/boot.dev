package handlers

import (
	"encoding/json"
	"github.com/andycostintoma/http-servers/auth"
	"github.com/andycostintoma/http-servers/db/generated"
	"github.com/google/uuid"
	"net/http"
)

func (cfg *ApiConfig) HandlerEditUser(w http.ResponseWriter, r *http.Request, userId uuid.UUID) {

	type RequestBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	requestBody := RequestBody{}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		respondAndLog(w, http.StatusBadRequest, "Invalid requestBody body", err)
		return
	}

	hashedPassword, err := auth.HashPassword(requestBody.Password)
	if err != nil {
		respondWithInternalServerError(w, err)
		return
	}

	updatedUser, err := cfg.Db.UpdateUser(r.Context(), generated.UpdateUserParams{
		ID:             userId,
		Email:          requestBody.Email,
		HashedPassword: hashedPassword,
	})

	if err != nil {
		respondWithInternalServerError(w, err)
	}

	type UserResponse struct {
		Id    uuid.UUID `json:"id"`
		Email string    `json:"email"`
	}

	respondWithJSON(w, 200, UserResponse{
		Id:    updatedUser.ID,
		Email: updatedUser.Email,
	})
}
