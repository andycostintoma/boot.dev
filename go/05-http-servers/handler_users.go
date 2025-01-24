package main

import (
	"encoding/json"
	"github.com/andycostintoma/http-servers/internal/auth"
	"github.com/andycostintoma/http-servers/internal/database"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerUsersCreate(w http.ResponseWriter, r *http.Request) {
	type UserRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	requestBody := UserRequest{}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	hashedPassword, err := auth.HashPassword(requestBody.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error hashing password", err)
	}

	user, err := cfg.db.CreateUser(r.Context(), database.CreateUserParams{
		Email:          requestBody.Email,
		HashedPassword: hashedPassword,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create user", err)
		return
	}

	type UserResponse struct {
		Id        uuid.UUID `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Email     string    `json:"email"`
	}

	respondWithJSON(w, http.StatusCreated, UserResponse{
		Id:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
	})
}

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {

	type LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Expires  int    `json:"expires_in_seconds"`
	}
	requestBody := LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error decoding json", err)
		return
	}
	if requestBody.Expires == 0 {
		requestBody.Expires = 3600
	}

	user, err := cfg.db.GetUserByEmail(r.Context(), requestBody.Email)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "User not found", err)
		return
	}

	err = auth.CheckPasswordHash(requestBody.Password, user.HashedPassword)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Incorrect email or password"))
		return
	}

	token, err := auth.MakeJWT(user.ID, cfg.secret, time.Duration(requestBody.Expires)*time.Second)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error generating token", err)
		return
	}

	type UserResponse struct {
		Id        uuid.UUID `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Email     string    `json:"email"`
		Token     string    `json:"token"`
	}

	respondWithJSON(w, 200, UserResponse{
		Id:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
		Token:     token,
	})
}
