package handlers

import (
	"net/http"
)

func (cfg *ApiConfig) HandlerReset(w http.ResponseWriter, r *http.Request) {
	if cfg.Platform != "dev" {
		respondAndLog(w, http.StatusForbidden, "Reset is only allowed in dev environment.", nil)
	}

	cfg.FileserverHits.Store(0)
	err := cfg.Db.DeleteUsers(r.Context())
	if err != nil {
		respondWithInternalServerError(w, err)
		return
	}
	respondWithJSON(w, http.StatusOK, struct {
		Msg string `json:"msg"`
	}{Msg: "Hits reset to 0 and db reset to initial state."})
}
