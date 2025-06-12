package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/GurramKarimunisa/go-user-api/config"
	"github.com/GurramKarimunisa/go-user-api/db/sqlc"
	"github.com/GurramKarimunisa/go-user-api/internal/handler"
	"github.com/GurramKarimunisa/go-user-api/internal/logger"
	"github.com/GurramKarimunisa/go-user-api/internal/repository"
	"github.com/GurramKarimunisa/go-user-api/internal/routes"
	"github.com/GurramKarimunisa/go-user-api/internal/service"
	"go.uber.org/zap"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize logger
	logger.InitLogger(cfg.Environment)
	defer logger.SyncLogger()

	logger.Log.Info("Starting application", zap.String("environment", cfg.Environment))

	// --- DIAGNOSTIC LINE ADDED HERE ---
	logger.Log.Info("Database URL being used", zap.String("url", cfg.DatabaseURL))
	// --- END DIAGNOSTIC LINE ---

	// Connect to database
	connPool, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		logger.Log.Fatal("Unable to connect to database", zap.Error(err))
	}
	defer connPool.Close()

	if err = connPool.Ping(context.Background()); err != nil {
		logger.Log.Fatal("Failed to ping database", zap.Error(err))
	}
	logger.Log.Info("Successfully connected to database")

	// Initialize sqlc queries
	queries := db.New(connPool)

	// Initialize repository, service, and handler
	userRepo := repository.NewUserRepository(queries)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Create Fiber app
	app := fiber.New()

	// Setup routes
	routes.SetupUserRoutes(app, userHandler)

	// Start server
	logger.Log.Info("Server is starting", zap.String("port", cfg.Port))
	log.Fatal(app.Listen(":" + cfg.Port))
}
