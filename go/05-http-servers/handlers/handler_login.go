package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/andycostintoma/http-servers/auth"
	"github.com/andycostintoma/http-servers/db/generated"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (cfg *ApiConfig) HandlerLogin(w http.ResponseWriter, r *http.Request) {

	type LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	requestBody := LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		respondAndLog(w, http.StatusBadRequest, "Invalid requestBody body", err)
		return
	}

	user, err := cfg.Db.GetUserByEmail(r.Context(), requestBody.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondAndLog(w, http.StatusUnauthorized, "Incorrect email or password", err)
		} else {
			respondWithInternalServerError(w, err)
		}
		return
	}

	err = auth.CheckPasswordHash(requestBody.Password, user.HashedPassword)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Incorrect email or password"))
		return
	}

	token, err := auth.MakeJWT(user.ID, cfg.JwtSecret, time.Hour)
	if err != nil {
		respondWithInternalServerError(w, err)
		return
	}

	refreshToken, err := auth.MakeRefreshToken()
	if err != nil {
		respondWithInternalServerError(w, err)
		return
	}

	dbRefreshToken, err := cfg.Db.CreateRefreshToken(r.Context(), generated.CreateRefreshTokenParams{
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(60 * 24 * time.Hour).UTC(),
		RevokedAt: sql.NullTime{},
		UserID:    user.ID,
	})

	if err != nil {
		respondWithInternalServerError(w, err)
	}

	type UserResponse struct {
		Id           uuid.UUID `json:"id"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
		Email        string    `json:"email"`
		IsChirpyRed  bool      `json:"is_chirpy_red"`
		Token        string    `json:"token"`
		RefreshToken string    `json:"refresh_token"`
	}

	respondWithJSON(w, 200, UserResponse{
		Id:           user.ID,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		Email:        user.Email,
		IsChirpyRed:  user.IsChirpyRed,
		Token:        token,
		RefreshToken: dbRefreshToken.Token,
	})
}
