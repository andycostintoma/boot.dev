package main

import (
	"database/sql"
	"github.com/andycostintoma/http-servers/db/generated"
	handlers2 "github.com/andycostintoma/http-servers/handlers"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"sync/atomic"
)

func main() {
	const filepathRoot = "static"
	const port = "8080"

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Couldn't load environment variables")
	}
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL must be set")
	}
	platform := os.Getenv("PLATFORM")
	if platform == "" {
		log.Fatal("PLATFORM must be set")
	}
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable is not set")
	}
	polkaKey := os.Getenv("POLKA_KEY")
	if polkaKey == "" {
		log.Fatal("POLKA_KEY environment variable is not set")
	}

	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error opening database: %s", err)
	}
	dbQueries := generated.New(dbConn)

	apiCfg := handlers2.ApiConfig{
		FileserverHits: atomic.Int32{},
		Db:             dbQueries,
		Platform:       platform,
		JwtSecret:      jwtSecret,
		PolkaKey:       polkaKey,
	}

	mux := NewAuthenticatedServeMux(jwtSecret)
	fsHandler := apiCfg.MiddlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot))))
	mux.HandleFunc("/app/", fsHandler.ServeHTTP)

	mux.HandleFunc("GET /api/health", handlers2.HandlerReadiness)

	mux.HandleFunc("GET /api/chirps", apiCfg.HandlerGetChirps)
	mux.HandleFunc("GET /api/chirps/{id}", apiCfg.HandleGetChirp)
	mux.HandleAuthenticated("POST /api/chirps", apiCfg.HandlerChirpsCreate)
	mux.HandleAuthenticated("DELETE /api/chirps/{id}", apiCfg.HandlerDeleteChirp)

	mux.HandleFunc("POST /api/users", apiCfg.HandlerRegister)
	mux.HandleAuthenticated("PUT /api/users", apiCfg.HandlerEditUser)
	mux.HandleFunc("POST /api/login", apiCfg.HandlerLogin)
	mux.HandleFunc("POST /api/refresh", apiCfg.HandlerRefresh)
	mux.HandleFunc("POST /api/revoke", apiCfg.HandlerRevoke)

	mux.HandleFunc("POST /api/polka/webhooks", apiCfg.HandlerWebhook)

	mux.HandleFunc("POST /admin/reset", apiCfg.HandlerReset)
	mux.HandleFunc("GET /admin/metrics", apiCfg.HandlerMetrics)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}
