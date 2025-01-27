package main

import (
	"github.com/andycostintoma/http-servers/auth"
	"github.com/google/uuid"
	"net/http"
)

type AuthenticatedHandlerFunc func(http.ResponseWriter, *http.Request, uuid.UUID)

type AuthenticatedServeMux struct {
	mux       *http.ServeMux
	jwtSecret string
}

func NewAuthenticatedServeMux(jwtSecret string) *AuthenticatedServeMux {
	return &AuthenticatedServeMux{
		mux:       http.NewServeMux(),
		jwtSecret: jwtSecret,
	}
}

func (a *AuthenticatedServeMux) HandleAuthenticated(pattern string, handler AuthenticatedHandlerFunc) {
	a.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		token, err := auth.GetBearerToken(r.Header)
		if err != nil {
			http.Error(w, "Unauthorized: Missing or invalid token", http.StatusUnauthorized)
			return
		}

		userID, err := auth.ValidateJWT(token, a.jwtSecret)
		if err != nil {
			http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
			return
		}

		handler(w, r, userID)
	})
}

func (a *AuthenticatedServeMux) HandleFunc(pattern string, handler http.HandlerFunc) {
	a.mux.HandleFunc(pattern, handler)
}

func (a *AuthenticatedServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}
