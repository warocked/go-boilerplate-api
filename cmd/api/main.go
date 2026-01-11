package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-boilerplate-api/internal/api/config"
	"go-boilerplate-api/internal/api/db"
	"go-boilerplate-api/internal/api/middlewares"
	"go-boilerplate-api/internal/api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load configuration
	err := config.LoadEnvFile()
	if err != nil {
		log.Fatalf("Failed to load env file: %v", err)
	}

	err = config.LoadAllConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	ctx := context.Background()

	if config.DATABASE_URL != "" {
		err = db.InitPostgres(ctx, config.DATABASE_URL)
		if err != nil {
			log.Fatalf("Failed to initialize PostgreSQL: %v", err)
		}
		defer db.ClosePostgres()

		if err = db.RunMigrations(ctx, config.DATABASE_URL); err != nil {
			log.Fatalf("Failed to run database migrations: %v", err)
		}
	}

	if config.REDIS_URL != "" {
		err = db.InitRedis(ctx, config.REDIS_URL)
		if err != nil {
			log.Fatalf("Failed to initialize Redis: %v", err)
		}
		defer db.CloseRedis()
	}

	if config.IS_PROD {
		if config.SECRET_KEY == "" || config.SECRET_KEY == "qweasd123" || len(config.SECRET_KEY) < 32 {
			log.Fatalf("SECRET_KEY must be set to a secure value (minimum 32 characters) in production")
		}
	}

	app := fiber.New(fiber.Config{
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      300 * time.Second,
		IdleTimeout:       120 * time.Second,
		ReadBufferSize:    4096,
		WriteBufferSize:   4096,
		EnablePrintRoutes: !config.IS_PROD,
		Prefork:           false,
		CaseSensitive:     false,
		StrictRouting:     false,
		ReduceMemoryUsage: true,
		Network:           fiber.NetworkTCP,
		BodyLimit:         int(config.REQUEST_BODY_LIMIT),
		ErrorHandler:      errorHandler,
	})

	if !config.IS_PROD {
		app.Use(func(c *fiber.Ctx) error {
			c.Set("Cache-Control", "no-store")
			return c.Next()
		})
	}

	middlewares.SetupMiddlewares(app)
	routes.SetupRoutes(app)

	port := config.PORT
	if port == "" {
		port = "3000"
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := app.Listen(":" + port); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	<-quit

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	app.ShutdownWithContext(shutdownCtx)
}
