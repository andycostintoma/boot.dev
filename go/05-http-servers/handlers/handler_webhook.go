package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/andycostintoma/http-servers/auth"
	"github.com/google/uuid"
	"net/http"
)

func (cfg *ApiConfig) HandlerWebhook(w http.ResponseWriter, r *http.Request) {
	key, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondAndLog(w, http.StatusUnauthorized, "Couldn't get authentication header", err)
		return
	}

	if key != cfg.PolkaKey {
		respondAndLog(w, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	type requestBody struct {
		Event string `json:"event"`
		Data  struct {
			UserID uuid.UUID `json:"user_id"`
		} `json:"data"`
	}
	reqBody := requestBody{}
	err = json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		respondAndLog(w, http.StatusBadRequest, "Invalid requestBody body", err)
		return
	}

	if reqBody.Event != "user.upgraded" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	_, err = cfg.Db.UpgradeUser(r.Context(), reqBody.Data.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondAndLog(w, http.StatusNotFound, "User not found", err)
		} else {
			respondWithInternalServerError(w, err)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
