package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime"
)

func logErrorWithStackTrace(err error) {
	stackTrace := make([]byte, 1024)
	n := runtime.Stack(stackTrace, false)
	log.Printf("Error: %s\nStack Trace: %s\n", err, stackTrace[:n])
}

func respondAndLog(w http.ResponseWriter, code int, msg string, err error) {
	if err != nil {
		logErrorWithStackTrace(err)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errorResponse{
		Error: msg,
	})
}

func respondWithInternalServerError(w http.ResponseWriter, err error) {
	respondAndLog(w, http.StatusInternalServerError, "Oops. Something went wrong.", err)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	w.Write(dat)
}