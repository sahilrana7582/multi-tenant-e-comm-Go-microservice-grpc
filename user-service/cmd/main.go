package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"
	"github.com/sahilrana7582/multi-tenent-e-com-user-service/config"
	"github.com/sahilrana7582/multi-tenent-e-com-user-service/internal/db"
	"github.com/sahilrana7582/multi-tenent-e-com-user-service/internal/handler"
	"github.com/sahilrana7582/multi-tenent-e-com-user-service/internal/repository"
	"github.com/sahilrana7582/multi-tenent-e-com-user-service/internal/routes"
	"github.com/sahilrana7582/multi-tenent-e-com-user-service/internal/service"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using environment variables.")
	}

	cfg := config.Load()

	fmt.Printf("🚀 Starting %s on port %d", cfg.AppName, cfg.Port)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbConn := db.MustConnect(ctx, cfg.DatabaseURL)
	defer dbConn.Close()

	userRepo := repository.NewUserRepository(dbConn)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	router := routes.NewRouter(userHandler)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      router,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	fmt.Printf("✅ Listening on port %d", cfg.Port)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("❌ Could not listen: %v", err)
	}
}
