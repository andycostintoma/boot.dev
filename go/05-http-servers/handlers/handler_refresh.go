package handlers

import (
	"database/sql"
	"errors"
	"github.com/andycostintoma/http-servers/auth"
	"net/http"
	"time"
)

func (cfg *ApiConfig) HandlerRefresh(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondAndLog(w, http.StatusBadRequest, "Couldn't get token", err)
		return
	}

	dbToken, err := cfg.Db.GetToken(r.Context(), token)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondAndLog(w, http.StatusUnauthorized, "Unauthorized", err)
		} else {
			respondWithInternalServerError(w, err)
		}
		return
	}

	if dbToken.ExpiresAt.Before(time.Now()) {
		respondAndLog(w, http.StatusUnauthorized, "Unauthorized", err)
		return
	}

	sqlNullTime := sql.NullTime{}
	if dbToken.RevokedAt != sqlNullTime {
		respondAndLog(w, http.StatusUnauthorized, "Unauthorized", err)
		return
	}

	jwt, err := auth.MakeJWT(dbToken.UserID, cfg.JwtSecret, time.Hour)
	if err != nil {
		respondWithInternalServerError(w, err)
		return
	}

	respondWithJSON(w, http.StatusOK, struct {
		Token string `json:"token"`
	}{Token: jwt})
}
