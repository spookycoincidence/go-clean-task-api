package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/evrone/go-clean-template/config"
	"github.com/evrone/go-clean-template/internal/app"
	"github.com/evrone/go-clean-template/internal/delivery/http" // Ajusta según tu path
	"github.com/evrone/go-clean-template/internal/repository"
	"github.com/evrone/go-clean-template/internal/usecase"
	_ "github.com/lib/pq"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Connect to DB (PostgreSQL)
	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Database connection error: %s", err)
	}
	defer db.Close()

	// Set DB connection settings (opcional)
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Test DB connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Database ping error: %s", err)
	}

	// Initialize repository and usecase
	taskRepo := repository.NewPostgresTaskRepository(db)
	taskUseCase := usecase.NewTaskUseCase(taskRepo)

	// Setup router
	r := http.SetupRouter(taskUseCase)

	// Run server
	if err := r.Run(cfg.ServerAddress); err != nil {
		log.Fatalf("Server run failed: %s", err)
	}
}
