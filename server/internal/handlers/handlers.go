package handlers

import (
	"net/http"
	database "server/video-streaming/internal/database/output"
)

// RegisterRoutes sets up all API routes with versioning
func RegisterRoutes(mux *http.ServeMux, q *database.Queries) {
	// Health check endpoint
	mux.HandleFunc("GET /health", healthHandler)

	// // API v1 routes
	// mux.Handle("/api/v1/", http.StripPrefix("/api/v1", v1.InitRouter(q, c)))
}

// healthHandler provides a simple health check endpoint
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok","message":"Server is healthy"}`))
}
