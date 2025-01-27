package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/andycostintoma/http-servers/db/generated"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

type chirpResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	Body      string    `json:"body"`
}

func mapDbToResponse(chirp generated.Chirp) chirpResponse {
	return chirpResponse{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserID:    chirp.UserID,
	}
}

func (cfg *ApiConfig) HandleGetChirp(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		respondAndLog(w, http.StatusBadRequest, "Not a valid UUID", err)
		return
	}
	chirp, err := cfg.Db.GetChirpById(r.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondAndLog(w, http.StatusNotFound, "Chirp not found", err)
		} else {
			respondWithInternalServerError(w, err)
		}
		return
	}
	respondWithJSON(w, http.StatusOK, mapDbToResponse(chirp))
}

func (cfg *ApiConfig) HandlerGetChirps(w http.ResponseWriter, r *http.Request) {
	dbChirps, err := cfg.Db.GetChirps(r.Context())
	// TODO: maybe handle sql.ErrNoRows differently?
	if err != nil {
		respondWithInternalServerError(w, err)
		return
	}
	var response []chirpResponse
	for _, chirp := range dbChirps {
		response = append(response, mapDbToResponse(chirp))
	}
	respondWithJSON(w, http.StatusOK, response)
}

func (cfg *ApiConfig) HandlerChirpsCreate(w http.ResponseWriter, r *http.Request, userId uuid.UUID) {
	type parameters struct {
		Body string `json:"body"`
	}
	params := parameters{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respondAndLog(w, http.StatusBadRequest, "Invalid requestBody body format", err)
		return
	}

	if params.Body == "" {
		respondAndLog(w, http.StatusBadRequest, "Chirp body cannot be empty", nil)
		return
	}

	cleaned, err := validateChirp(params.Body)
	if err != nil {
		respondAndLog(w, http.StatusBadRequest, err.Error(), err)
		return
	}

	chirp, err := cfg.Db.CreateChirp(r.Context(), generated.CreateChirpParams{
		Body:   cleaned,
		UserID: userId,
	})
	if err != nil {
		respondWithInternalServerError(w, err)
		return
	}

	respondWithJSON(w, http.StatusCreated, mapDbToResponse(chirp))
}

func validateChirp(body string) (string, error) {
	const maxChirpLength = 140
	if len(body) > maxChirpLength {
		return "", errors.New("chirp is too long")
	}

	badWords := map[string]struct{}{
		"kerfuffle": {},
		"sharbert":  {},
		"fornax":    {},
	}
	cleaned := getCleanedBody(body, badWords)
	return cleaned, nil
}

func getCleanedBody(body string, badWords map[string]struct{}) string {
	words := strings.Split(body, " ")
	for i, word := range words {
		loweredWord := strings.ToLower(word)
		if _, ok := badWords[loweredWord]; ok {
			words[i] = "****"
		}
	}
	cleaned := strings.Join(words, " ")
	return cleaned
}

func (cfg *ApiConfig) HandlerDeleteChirp(w http.ResponseWriter, r *http.Request, userId uuid.UUID) {
	chirpId, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		respondAndLog(w, http.StatusBadRequest, "Not a valid UUID", err)
		return
	}
	chirp, err := cfg.Db.GetChirpById(r.Context(), chirpId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondAndLog(w, http.StatusNotFound, "Chirp not found", err)
		} else {
			respondWithInternalServerError(w, err)
		}
		return
	}
	if chirp.UserID != userId {
		respondAndLog(w, http.StatusForbidden, "Forbidden", nil)
		return
	}
	err = cfg.Db.DeleteChirp(r.Context(), chirpId)
	if err != nil {
		respondWithInternalServerError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
