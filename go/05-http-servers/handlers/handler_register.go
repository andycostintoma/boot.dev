package handlers

import (
	"encoding/json"

	"github.com/andycostintoma/http-servers/auth"
	"github.com/andycostintoma/http-servers/db/generated"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (cfg *ApiConfig) HandlerRegister(w http.ResponseWriter, r *http.Request) {
	requestBody := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		respondAndLog(w, http.StatusBadRequest, "Invalid requestBody body.", err)
		return
	}

	hashedPassword, err := auth.HashPassword(requestBody.Password)
	if err != nil {
		respondWithInternalServerError(w, err)
		return
	}

	user, err := cfg.Db.CreateUser(r.Context(), generated.CreateUserParams{
		Email:          requestBody.Email,
		HashedPassword: hashedPassword,
	})
	if err != nil {
		respondWithInternalServerError(w, err)
		return
	}

	type responseBody struct {
		Id        uuid.UUID `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Email     string    `json:"email"`
	}

	respondWithJSON(w, http.StatusCreated, responseBody{
		Id:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
	})
}
