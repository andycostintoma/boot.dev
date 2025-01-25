package handlers

import (
	"database/sql"
	"errors"
	"github.com/andycostintoma/http-servers/auth"
	"net/http"
)

func (cfg *ApiConfig) HandlerRevoke(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondAndLog(w, http.StatusBadRequest, "Couldn't get bearer token.", err)
		return
	}

	_, err = cfg.Db.RevokeToken(r.Context(), token)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondAndLog(w, http.StatusUnauthorized, "Unauthorized", err)
		} else {
			respondWithInternalServerError(w, err)
		}
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
