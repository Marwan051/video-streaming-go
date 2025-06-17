package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"server/video-streaming/internal/config"
	database "server/video-streaming/internal/database/output"
	"server/video-streaming/internal/handlers"
	"server/video-streaming/internal/middleware"
	"syscall"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	log.Print("Starting video streaming server...")
	// Load configuration
	cfg := config.Load()

	conn, err := sql.Open("sqlite3", cfg.SQLiteDBPath)
	if err != nil {
		log.Fatalf("failed to open sqlite db: %v", err)
	}
	defer conn.Close()

	conn.SetMaxOpenConns(1) // SQLite doesn't like high concurrency
	conn.SetConnMaxIdleTime(5 * time.Minute)

	if err := conn.Ping(); err != nil {
		log.Fatalf("failed to ping sqlite: %v", err)
	}

	queries := database.New(conn)

	// Setup routes with middleware
	mux := http.NewServeMux()

	// Register API routes
	handlers.RegisterRoutes(mux, queries)

	// Apply global middleware
	var middlewares []middleware.Middleware
	middlewares = append(middlewares, middleware.Recovery)

	if cfg.Environment == config.EnvDevelopment {
		middlewares = append(middlewares, middleware.GlobalCORS)
	} else {
		// TODO: add specific CORS for prod server
	}

	middlewares = append(
		middlewares,
		middleware.Logging,
	)

	handler := middleware.Chain(mux, middlewares...)

	// Create server
	server := &http.Server{
		Addr:         cfg.Port,
		Handler:      handler,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Server starting on %s", cfg.Port)
		log.Printf("Server starting in %s mode", cfg.Environment)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server failed to start: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Server shutting down...")

	// Graceful shutdown with timeout
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer shutdownCancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}
